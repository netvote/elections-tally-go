package tally

import (
	"github.com/netvote/elections-tally-go/storage"
	"encoding/json"
	"github.com/sirupsen/logrus"
)

const (
	GROUP_ALL = "ALL"
)

type GroupResult struct {
	TotalVotes int `json:"totalVotes"`
	Decisions []map[string]int `json:"decisions"`
}

type BallotResult struct {
	TotalVotes int `json:"totalVotes"`
	Metadata *storage.Ballot `json:"metadata,omitempty"`
	// group to BallotResult
	GroupResults map[string]GroupResult `json:"results"`
}

type TallyResults struct {
	// address to BallotResult
	BallotResults map[string]BallotResult `json:"ballotResults"`
}

func NewTallyResults() TallyResults{
	return TallyResults{
		BallotResults: make(map[string]BallotResult),
	}
}

// convert to channel approach?
func (r TallyResults) MergeSum(vote TallyResults){
	logrus.Debugf("merging %s to %s", vote.Json(), r.Json())
	for addr, br := range vote.BallotResults {
		ballotResult := r.BallotResults[addr]
		ballotResult.TotalVotes += 1

		for group, gr := range br.GroupResults {
			resultGroup := ballotResult.GroupResults[group]
			resultGroup.TotalVotes += 1
			for dIdx, d := range gr.Decisions {
				for choice, v := range d {
					resultGroup.Decisions[dIdx][choice] += v
				}
			}
			ballotResult.GroupResults[group] = resultGroup
		}
		r.BallotResults[addr] = ballotResult
	}
}

func (r TallyResults) Json()(string) {
	b, _ := json.Marshal(r)
	return string(b)
}
