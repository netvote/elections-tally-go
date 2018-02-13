package tally

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/netvote/elections-tally-go/contracts"
	"github.com/netvote/elections-tally-go/decoder"
	"github.com/netvote/elections-tally-go/storage"
	"github.com/netvote/elections-tally-go/protocol"
	"math/big"
	"strings"
	"github.com/ethereum/go-ethereum/common"
	"bytes"
	"sync"
)

var (
	empty = TallyResults{}
)


func voteToResults(ballotAddress []string, ballotMetadata []*storage.Ballot, ballotGroups map[string][]string, vote protocol.Vote)(TallyResults){

	result := TallyResults{
		BallotResults:make(map[string]BallotResult),
	}
	for i:=0; i<len(vote.BallotVotes); i++ {
		address := ballotAddress[i]
		ballotVote := vote.BallotVotes[i]
		metadata := ballotMetadata[i]

		result.BallotResults[address] = BallotResult{
			TotalVotes: 1,
			GroupResults: make(map[string]GroupResult),
		}
		gr := GroupResult{
			TotalVotes: 1,
			Decisions: make([]map[string]int, 0),
		}

		for dIdx, choice := range ballotVote.Choices {
			if w := choice.GetWriteIn(); w == "" {
				//SELECTION
				selection := choice.GetSelection()
				d := metadata.Decisions[dIdx]
				id := d.Choices[selection].Id
				gr.Decisions = append(gr.Decisions, map[string]int{
					id: 1,
				})
			}else {
				// WRITEIN
				w = strings.ToUpper(w)
				w = strings.TrimSpace(w)
				gr.Decisions = append(gr.Decisions, map[string]int{
					"WRITEIN-"+w: 1,
				})
			}
		}

		if groups, ok := ballotGroups[address]; ok {
			for _,g := range groups {
				result.BallotResults[address].GroupResults[g] = gr
			}
		}

	}

	return result
}

func TallyPool(poolAddress string, decoder *decoder.Decoder, client *ethclient.Client) (TallyResults, error){
	pool, err := contracts.NewBasePool(common.HexToAddress(poolAddress), client)
	if err != nil{
		return empty, err
	}
	ballotCount, err := pool.GetBallotCount(nil)
	if err != nil{
		return empty, err
	}
	ballots := make([]string, 0)
	metadata := make([]*storage.Ballot, 0)
	groups := make(map[string][]string)

	result := TallyResults{
		BallotResults:make(map[string]BallotResult),
	}

	for i:=int64(0); i<ballotCount.Int64(); i++ {
		idx := big.NewInt(i)
		ballot, err := pool.GetBallot(nil, idx)
		if err != nil{
			return empty, err
		}


		ballotContract, err := contracts.NewTieredBallot(ballot, client)
		if err != nil{
			return empty, err
		}
		_, err = ballotContract.GetGroupCount(nil)
		groups[ballot.Hex()] = []string{"ALL"}
		if err == nil{
			groups[ballot.Hex()] = []string{}
			//TODO: get groups for this pool
			gCount, err := ballotContract.GetPoolGroupCount(nil, common.HexToAddress(poolAddress))
			if err != nil{
				return empty, err
			}
			for i:=int64(0); i<gCount.Int64(); i++ {
				group, err := ballotContract.GetPoolGroupAt(nil, common.HexToAddress(poolAddress), big.NewInt(i))
				if err != nil{
					return empty, err
				}
				groupSlice := group[:]
				n := bytes.Index(groupSlice, []byte{0})
				groupStr := string(group[:n])

				groups[ballot.Hex()] = append(groups[ballot.Hex()], groupStr)
			}
		}

		location, err := ballotContract.MetadataLocation(nil)
		if err != nil{
			return empty, err
		}
		ballotMetadata, err := storage.GetIpfsBallot(location)
		if err != nil{
			return empty, err
		}

		decisionResults := make([]map[string]int, 0)
		for _, d := range ballotMetadata.Decisions{
			choiceMap := make(map[string]int)
			for _, c := range d.Choices {
				choiceMap[c.Id] = 0
			}
			decisionResults = append(decisionResults, choiceMap)
		}

		br := BallotResult{
			Metadata: ballotMetadata,
			GroupResults: make(map[string]GroupResult),
		}
		for _, g := range groups[ballot.Hex()] {
			br.GroupResults[g] = GroupResult {
				Decisions: decisionResults,
			}
		}

		ballots = append(ballots, ballot.Hex())
		metadata = append(metadata, ballotMetadata)
		result.BallotResults[ballot.Hex()] = br
	}

	voteCount, err := pool.GetVoteCount(nil)
	if err != nil{
		return empty, err
	}

	tallyChan := make(chan TallyResults, 100)
	//defer close(tallyChan)
	var wg sync.WaitGroup
	wg.Add(int(voteCount.Int64()))
	for i := int64(0); i < voteCount.Int64(); i++ {
		idx := big.NewInt(i)
		go func() {
			payload, _ := pool.GetVoteAt(nil, idx)
			vote, _ := decoder.DecodeVote(payload)

			//TODO: validate errors
			//TODO: validate vote

			tallyChan <- voteToResults(ballots, metadata, groups, vote)
		}()
	}

	go func() {
		for vote := range tallyChan {
			go func() {
				defer wg.Done()
				result.MergeSum(vote)
			}()
		}
	}()


	wg.Wait()
	close(tallyChan)

	return result, nil
}
