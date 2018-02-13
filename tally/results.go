package tally

import (
	"github.com/netvote/elections-tally-go/storage"
	"encoding/json"
)

const (
	GROUP_ALL = "ALL"
)

type GroupResult struct {
	TotalVotes int
	Decisions []map[string]int
}

type BallotResult struct {
	TotalVotes int
	Metadata *storage.Ballot
	// group to BallotResult
	GroupResults map[string]GroupResult
}

type TallyResults struct {
	// address to BallotResult
	BallotResults map[string]BallotResult
}

func (r TallyResults) MergeSum(vote TallyResults){
	for addr, br := range vote.BallotResults {
		for group, gr := range br.GroupResults {
			for dIdx, d := range gr.Decisions {
				for choice, v := range d {
					r.BallotResults[addr].GroupResults[group].Decisions[dIdx][choice] += v
				}
			}
		}
	}
}


func (r TallyResults) Json()(string) {
	b, _ := json.Marshal(r)
	return string(b)
}
