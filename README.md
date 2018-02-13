# elections-tally-go
This tallies elections using goroutines in an effort to be reasonably quick

## Build
```
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
  -testnet
    	Use ropsten network
```

## Ouput
```
INFO[0001] {"BallotResults":{"0x6007CB06f4e91dd5fad6d639ff2a3caAF12dd4B3":{"TotalVotes":0,"Metadata":{"groupTitle":"2020 Election","ballotSections":[{"sectionTitle":"Governor","ballotItems":[{"itemTitle":"John Smith"},{"itemTitle":"Sally Gutierrez"},{"itemTitle":"Tyrone Williams"}]},{"sectionTitle":"Proposition 33","ballotItems":[{"itemTitle":"Yes"},{"itemTitle":"No"}]},{"sectionTitle":"Tax Commissioner","ballotItems":[{"itemTitle":"Doug Hall"},{"itemTitle":"Emily Washington"}]}]},"GroupResults":{"ALL":{"TotalVotes":0,"Decisions":[{"John Smith":3,"Sally Gutierrez":0,"Tyrone Williams":0},{"No":3,"Yes":0},{"Doug Hall":0,"Emily Washington":0,"WRITEIN-JOHN DOE":3}]}}}}}
```
