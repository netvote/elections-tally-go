package decoder

import (
	"crypto/cipher"
	"crypto/aes"
	"github.com/netvote/elections-tally-go/protocol"
	"encoding/base64"
	"strings"
	"hash"
	"github.com/golang/protobuf/proto"
	"crypto/md5"
)

type Decoder struct {
	block cipher.Block
	iv []byte
}

func NewDecoder(key string) (*Decoder, error) {
	hashKey, iv := genIvAndKey([]byte{}, []byte(key), md5.New(), 32, 1)
	block, err := aes.NewCipher([]byte(hashKey))
	if err != nil {
		return &Decoder{}, err
	}

	return &Decoder{
		block: block,
		iv: iv,
	}, nil
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

func (d *Decoder) DecodeVote(payload string) (protocol.Vote, error){
	var vote protocol.Vote

	//base64 -> utf8
	decodedMsg, err := base64.StdEncoding.DecodeString(addBase64Padding(payload))
	if err != nil {
		return protocol.Vote{}, err
	}

	//decrypt
	msgBytes := []byte(decodedMsg)
	mode := cipher.NewCBCDecrypter(d.block, d.iv)
	mode.CryptBlocks(msgBytes, msgBytes)

	// decode proto
	protoBytes := trimPadding(msgBytes)
	err = proto.Unmarshal(protoBytes, &vote)
	if err != nil {
		return protocol.Vote{}, err
	}

	return vote, nil
}

func trimPadding(msg []byte)([]byte){
	padding := len(msg) - int(msg[len(msg)-1])
	return msg[:padding]
}