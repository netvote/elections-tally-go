# elections-tally-go
This tallies elections using goroutines.  This is faster than the node library (for large elections). 

## Build
```
make contracts
make build
```

## Run
```
./main -address=0x6007cb06f4e91dd5fad6d639ff2a3caaf12dd4b3 -testnet
```

## Parameters
```
Usage of ./main:
  -address string
    	Ethereum address of election
  -debug
    	Show detailed logs   	
  -testnet
    	Use ropsten network
```

## Ouput
```
{
  "ballotResults": {
    "0x6007CB06f4e91dd5fad6d639ff2a3caAF12dd4B3": {
      "totalVotes": 4,
      "groupVotes": {
        "ALL": {
          "totalVotes": 4,
          "decisions": [
            {
              "John Smith": 4,
              "Sally Gutierrez": 0,
              "Tyrone Williams": 0
            },
            {
              "No": 4,
              "Yes": 0
            },
            {
              "Doug Hall": 0,
              "Emily Washington": 0,
              "WRITEIN-JOHN DOE": 4
            }
          ]
        }
      }
    }
  }
}
```

Contributing
-------------------

### Contribution Process
1. Fork repo
2. Make desired changes
3. Submit PR (Reference Issue #)
4. Reviewer will review
5. Reviewer Squash + Merges PR
