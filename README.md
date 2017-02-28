# gowit
Go SDK for Wit.AI 

## wit CLI
Installation and usage
```
$ export GOPATH=$(pwd)
$ export PATH=$PATH:$GOPATH/bin
$ go get github.com/michlabs/gowit/cmd/wit
$ wit help
$ wit train -t intent -i training.csv -token <your_wit_token>
```

Training file must be a CSV file and in following format:
```
intent_name1, intent utterance 1
intent_name1, intent utterance 2
intent_name2, intent utterance 3
intent_name1, intent utterance 4
...
```