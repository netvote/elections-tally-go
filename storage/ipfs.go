package storage

import (
	ipfs "github.com/ipfs/go-ipfs-api"
	"encoding/json"
	"strings"
)

var (
	shell *ipfs.Shell
	cache = make(map[string]*Ballot)
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
	if b, ok := cache[ref]; ok {
		return b, nil
	}

	obj, err := shell.ObjectGet(ref)

	if err != nil {
		return nil, err
	}

	var metadata BallotMetadata
	jsonStr := obj.Data[strings.Index(obj.Data, "{"):strings.LastIndex(obj.Data, "}")+1]

	if err := json.Unmarshal([]byte(jsonStr), &metadata); err != nil {
		return nil, err
	}

	res := &Ballot{
		Title: metadata.Title,
		Decisions: []Decision{},
	}
	// squash into flat array per ballot
	for _, group := range metadata.BallotGroups {
		res.Decisions = append(res.Decisions, group.Decisions...)
	}
	cache[ref] = res
	return res, nil
}