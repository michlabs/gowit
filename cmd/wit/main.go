package main 

import (
	"flag"
	"fmt"
	"os"
	"log"

	"github.com/michlabs/gowit"
)

var target string
var inputFP string
var token string

func main() {
	trainCmd := flag.NewFlagSet("train", flag.ExitOnError)
	trainCmd.StringVar(&target, "t", "", "required, intent or entity")
	trainCmd.StringVar(&inputFP, "i", "", "required, path to the input file")
	trainCmd.StringVar(&token, "token", "", "required, Wit.AI token")

	testCmd := flag.NewFlagSet("test", flag.ExitOnError)
	testCmd.StringVar(&target, "t", "", "required, intent or entity")
	testCmd.StringVar(&inputFP, "i", "", "required, path to the input file")
	testCmd.StringVar(&token, "token", "", "required, Wit.AI token")

	if len(os.Args) < 2 {
		fmt.Println("Error: Input is not enough")
		fmt.Println(helpMessage)
		os.Exit(1)
	}

	command := os.Args[1]
	switch command {
	case "train":
		trainCmd.Parse(os.Args[2:])
	case "test":
		testCmd.Parse(os.Args[2:])
	case "help":
		fmt.Println(helpMessage)
		os.Exit(0)
	default:
		fmt.Printf("Error: Do not have command %s \n", command)
		fmt.Println(helpMessage)
		os.Exit(1)
	}

	if target != "intent" && target != "entity" {
		fmt.Println("Error: You must choose intent or entity")
		fmt.Println(helpMessage)
		os.Exit(1)
	}

	if inputFP == "" {
		fmt.Println("Error: Input file is required but empty")
		fmt.Println(helpMessage)
		os.Exit(1)
	}

	if token == "" {
		fmt.Println("Error: Token is required")
		fmt.Println(helpMessage)
		os.Exit(1)
	}

	wit := gowit.NewClient(token)

	if trainCmd.Parsed() {
		if target == "intent" {
			err := TrainIntent(wit, inputFP)	
			if err != nil {
				log.Fatal(err)
			}
		} else {
			TrainEntity()
		}
	}

	if testCmd.Parsed() {
		if target == "intent" {
			if err := TestIntent(wit, inputFP); err != nil {
				log.Fatal(err)
			}
		} else {
			TestEntity()
		}
	}
}

const helpMessage string = `
wit is CLI tool that helps you train and test Wit.AI in terminal

Usage: wit <command> <option>
Available commands and corresponding options:
	train
	  -t string
	    	required, type of training (intent, entity)
	  -i string
	    	required, path to your input file
	  -token string
	  		required, Wit.AI token

	test
	  -t string
	    	required, type of training (intent, entity)
	  -i string
	    	required, path to your input file
	  -token string
	  		required, Wit.AI token

	help
`