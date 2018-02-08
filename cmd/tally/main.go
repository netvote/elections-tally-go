package main

import (
	"github.com/netvote/elections-tally-go/contracts"
	"github.com/netvote/elections-tally-go/decoder"
	"os"
	log "github.com/sirupsen/logrus"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"flag"
	"math/big"
)

func main() {

	address := flag.String("address", "", "Ethereum address of election")
	testNet := flag.Bool("testnet", false, "Use ropsten network")

	flag.Parse()


	network := "mainnet"
	if *testNet {
		network = "ropsten"
	}

	if *address == "" {
		log.Errorf("-address (of election) is required")
		os.Exit(1)
	}

	log.Infof("Tallying election=%s", *address)

	// get encryption key
	conn, err := ethclient.Dial("https://"+network+".infura.io")
	noError("error getting ethclient", err)

	election, err := contracts.NewBasicElection(common.HexToAddress(*address), conn)
	noError("error getting election", err)

	key, err := election.PrivateKey(nil)
	noError("error getting decryption key", err)

	if key == "" {
		log.Info("This election is not ready to tally, because it has no decryption key set.")
		os.Exit(1)
	}


	noError("error creating decoder", err)

	electionType, err := election.ElectionType(nil)
	noError("error getting type", err)

	switch electionType {
	case "BASIC":

		voterCount, err := election.GetVoteCount(nil)
		noError("error getting count", err)

		if voterCount.Uint64() == 0 {
			log.Info("There are no votes for this election")
			os.Exit(0)
		}

		for i := int64(0); i < voterCount.Int64(); i++ {
			votePayload, err := election.GetVoteAt(nil, big.NewInt(i))
			noError("error getting vote", err)
			log.Infof("vote=%s", votePayload)
			_, err = decoder.DecodeVote(key, votePayload)
			noError("error getting vote", err)
		}

		log.Info("Basic")
	case "TIERED":
		log.Info("Tiered")
	default:
		log.Errorf("Unrecognized election type: %s", electionType)
		os.Exit(1)
	}


}

func noError(message string, err error){
	if err != nil {
		log.Fatalf(message, err)
		os.Exit(1)
	}
}