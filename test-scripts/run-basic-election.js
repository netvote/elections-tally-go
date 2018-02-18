// ------------------------------------------------------------------------------
// This file is part of netvote.
//
// netvote is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// netvote is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with solidity.  If not, see <http://www.gnu.org/licenses/>
//
// (c) 2017 netvote contributors.
//------------------------------------------------------------------------------

const election = require("../elections-solidity/test/end-to-end/jslib/basic-election.js");

module.exports = async function(callback) {
    let accounts = web3.eth.accounts;

    let vote1Json = {
        encryptionSeed: 12345,
        ballotVotes: [
            {
                choices: [
                    {
                        selection: 2
                    },
                    {
                        selection: 1
                    },
                    {
                        selection: 0
                    }
                ]
            }
        ]
    };
    let vote2Json = {
        encryptionSeed: 54321,
        ballotVotes: [
            {
                choices: [
                    {
                        selection: 1
                    },
                    {
                        selection: 1
                    },
                    {
                        writeIn: "John Doe"
                    }
                ]
            }
        ]
    };


    let vote1 = await election.toEncryptedVote(vote1Json);
    let vote2 = await election.toEncryptedVote(vote2Json);

    let config = await election.doEndToEndElectionAutoActivate({
        account: {
            allowance: 3,
            owner: accounts[0]
        },
        netvote: accounts[1],
        admin: accounts[2],
        allowUpdates: false,
        autoActivate: true,
        skipGasMeasurment:  true,
        gateway: accounts[3],
        metadata: "Qmc9oXZnUtcHoa7GxE7Ujwq4zG9SqSqKk5w9Qjqbi1cEWB",
        provider: "http://localhost:8545",
        protoPath: "protocol/vote.proto",
        voters: {
            voter1: {
                voteId: "vote-id-1",
                vote: vote1
            },
            voter2: {
                voteId: "vote-id-2",
                vote: vote2
            }
        }
    });

    console.log(config.contract.address);
    callback();
};