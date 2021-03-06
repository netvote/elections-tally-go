package main

import (
	"github.com/netvote/elections-tally-go/contracts"
	"github.com/netvote/elections-tally-go/decoder"
	"github.com/netvote/elections-tally-go/storage"
	"github.com/netvote/elections-tally-go/tally"
	"os"
	log "github.com/sirupsen/logrus"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"flag"
	"sync"
	"errors"
	"fmt"
)

func doTally(address string, provider string) (tally.TallyResults, error){
	// get encryption key
	conn, err := ethclient.Dial(provider)
	noError("error getting ethclient", err)

	election, err := contracts.NewBaseElection(common.HexToAddress(address), conn)
	noError("error getting election", err)

	key, err := election.PrivateKey(nil)
	log.Info("address="+address)
	noError("error getting decryption key", err)

	if key == "" {
		log.Info("This election is not ready to tally, because it has no decryption key set.")
		os.Exit(1)
	}

	noError("error creating decoder", err)

	electionType, err := election.ElectionType(nil)
	noError("error getting type", err)

	d, err := decoder.NewDecoder(key)
	noError("error getting decoder", err)

	storage.InitShell()

	switch electionType {
	case "BASIC":
		noError("error getting election", err)

		result, err := tally.TallyPool(address, d, conn)
		noError("error tallying pool", err)
		log.Info("done doing tally: "+result.Json())
		return result, nil

	case "TIERED":
		tiered, err := contracts.NewTieredElection(common.HexToAddress(address), conn)
		noError("error getting election", err)

		poolCnt, err := tiered.GetPoolCount(nil)
		noError("error getting pool count", err)

		var wg sync.WaitGroup
		wg.Add(int(poolCnt.Int64()))
		tallyChan := make(chan tally.TallyResults, 10)

		for i := int64(0); i<poolCnt.Int64(); i++ {
			go func() {
				result, err := tally.TallyPool(address, d, conn)
				noError("error tallying pool", err)
				tallyChan <- result
			}()
		}

		result := tally.NewTallyResults()
		go func() {
			for vote := range tallyChan {
				result.MergeSum(vote)
				wg.Done()
			}
		}()

		wg.Wait()
		close(tallyChan)
		return result, nil


	default:
		log.Errorf("Unrecognized election type: %s", electionType)
		os.Exit(1)
	}
	return tally.TallyResults{}, errors.New("invalid condition")
}

func main() {
	address := flag.String("address", "", "Ethereum address of election")
	testNet := flag.Bool("testnet", false, "Use ropsten network")
	debug := flag.Bool("debug", false, "Show more detail")

	flag.Parse()

	if *debug {
		log.SetLevel(log.DebugLevel)
	}

	network := "mainnet"
	if *testNet {
		network = "ropsten"
	}

	if *address == "" {
		log.Errorf("-address (of election) is required")
		os.Exit(1)
	}
	provider := "https://"+network+".infura.io"

	res, err := doTally(*address, provider)
	noError("error during tally", err)
	fmt.Println(res.Json())
}

func noError(message string, err error){
	if err != nil {
		log.Fatalf(message, err)
		os.Exit(1)
	}
}