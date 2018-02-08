package decoder

import (
	"crypto/cipher"
	"crypto/aes"
	"github.com/netvote/elections-tally-go/protocol"
	"github.com/sirupsen/logrus"
	"encoding/base64"
	"strings"
	"hash"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"crypto/md5"
)

type Decoder struct {
	key string
}

func addBase64Padding(value string) string {
	m := len(value) % 4
	if m != 0 {
		value += strings.Repeat("=", 4-m)
	}

	return value
}

func genIvAndKey(salt, data []byte, h hash.Hash, keyLen, blockLen int) (key []byte, iv []byte) {
	res := make([]byte, 0, keyLen+blockLen)
	p := append(data, salt...)
	var d_last []byte

	for ; len(res) < keyLen+blockLen; h.Reset() {
		h.Write(append(d_last, p...))
		resNew := h.Sum(res)
		d_last = resNew[len(res):]
		res = resNew
	}

	return res[:keyLen], res[keyLen:]
}

func DecodeVote(key string, payload string) (protocol.Vote, error){
	logrus.Debugf("encrypted size=%d", len(payload))
	decodedMsg, err := base64.StdEncoding.DecodeString(addBase64Padding(payload))

	if err != nil {
		return protocol.Vote{}, err
	}
	logrus.Debugf("size=%d, msg=%s", len(decodedMsg), decodedMsg)
	msgBytes := []byte(decodedMsg)

	hashKey, iv := genIvAndKey([]byte{}, []byte(key), md5.New(), 32, 1)
	block, err := aes.NewCipher([]byte(hashKey))
	if err != nil {
		return protocol.Vote{}, err
	}
	mode := cipher.NewCBCDecrypter(block, iv)

	mode.CryptBlocks(msgBytes, msgBytes)

	var vote protocol.Vote

	padding := len(msgBytes) - int(msgBytes[len(msgBytes)-1])

	err = proto.Unmarshal(msgBytes[:padding], &vote)
	if err != nil {
		return protocol.Vote{}, err
	}

	b, _ := json.Marshal(vote)
	logrus.Infof("VOTE=%s", string(b))

	return vote, nil
}
