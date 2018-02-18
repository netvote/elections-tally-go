// +build system

package main

import (
	"testing"
	"io/ioutil"
	"github.com/stretchr/testify/assert"
	"github.com/ethereum/go-ethereum/common"
	"strings"
)

func TestTally(t *testing.T) {
	addrBytes, err := ioutil.ReadFile("../../elections-solidity/election.address")
	assert.NoError(t, err, "error while reading address")

	address := strings.TrimSpace(string(addrBytes))
	address = common.HexToAddress(address).String()

	t.Log("address: "+address)
	res, err := doTally(address, "http://localhost:8545")

	assert.NoError(t, err, "received error while tallying")

	br := res.BallotResults[address]

	assert.Equal(t, 2, br.TotalVotes, "expected 2 votes")
	assert.Equal(t, 2, br.GroupResults["ALL"].TotalVotes, "expected 2 votes")
	assert.Equal(t, 3, len(br.GroupResults["ALL"].Decisions), "expected 3 decisions")

	assert.Equal(t, 0, br.GroupResults["ALL"].Decisions[0]["John Smith"])
	assert.Equal(t, 1, br.GroupResults["ALL"].Decisions[0]["Sally Gutierrez"])
	assert.Equal(t, 1, br.GroupResults["ALL"].Decisions[0]["Tyrone Williams"])

	assert.Equal(t, 0, br.GroupResults["ALL"].Decisions[1]["Yes"])
	assert.Equal(t, 2, br.GroupResults["ALL"].Decisions[1]["No"])

	assert.Equal(t, 1, br.GroupResults["ALL"].Decisions[2]["Doug Hall"])
	assert.Equal(t, 0, br.GroupResults["ALL"].Decisions[2]["Emily Washington"])
	assert.Equal(t, 1, br.GroupResults["ALL"].Decisions[2]["WRITEIN-JOHN DOE"])
}