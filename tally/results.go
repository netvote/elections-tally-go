package tally

import (
	"github.com/netvote/elections-tally-go/storage"
	"encoding/json"
	"sync"
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
	GroupResults map[string]GroupResult `json:"groupVotes"`
}

type TallyResults struct {
	// address to BallotResult
	mux *sync.Mutex
	BallotResults map[string]BallotResult `json:"ballotResults"`
}

func NewTallyResults() TallyResults{
	return TallyResults{
		mux: &sync.Mutex{},
	}
}

// convert to channel approach?
func (r TallyResults) MergeSum(vote TallyResults){
	r.mux.Lock()
	defer r.mux.Unlock()
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
