package storage

import (
	ipfs "github.com/ipfs/go-ipfs-api"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"strings"
)

var (
	shell *ipfs.Shell
)

type Choice struct {
	Id string `json:"itemTitle"`
}

type Decision struct {
	Title string `json:"sectionTitle"`
	Choices []Choice `json:"ballotItems"`
}

type Ballot struct {
	Title string `json:"groupTitle"`
	Decisions []Decision `json:"ballotSections"`
}

type BallotMetadata struct {
	Title   string      `json:"ballotTitle"`
	BallotGroups []Ballot `json:"ballotGroups"`
}

func InitShell(){
	shell = ipfs.NewShell("https://gateway.ipfs.io")
}

func GetIpfsBallot(ref string)(*Ballot, error){
	obj, err := shell.ObjectGet(ref)

	if err != nil {
		return nil, err
	}

	var metadata BallotMetadata
	jsonStr := obj.Data[strings.Index(obj.Data, "{"):strings.LastIndex(obj.Data, "}")+1]

	if err := json.Unmarshal([]byte(jsonStr), &metadata); err != nil {
		return nil, err
	}

	logrus.Infof("group size: %d", len(metadata.BallotGroups))
	res := &Ballot{
		Title: metadata.Title,
		Decisions: []Decision{},
	}
	// squash into flat array per ballot
	for _, group := range metadata.BallotGroups {

		res.Decisions = append(res.Decisions, group.Decisions...)
	}
	return res, nil
}