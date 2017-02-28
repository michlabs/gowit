package main 

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"log"

	"github.com/michlabs/gowit"
)

type IntentUtterance struct {
	Intent string `json:"intent"`
	Utterance string `json:"utterance"`
}

func TrainIntent(wit *gowit.Client, inputFP string) error {
	ius, err := ReadIntentsFromFile(inputFP)
	if err != nil {
		return err
	}

	intent, err := wit.GetEntity("intent")
	if err != nil {
		return err
	}

	intent.DeleteAllValues()

	for _, iu := range ius {
		var value gowit.Value
		value.Name = iu.Intent
		value.Expressions = append(value.Expressions, iu.Utterance)
		intent.Values = append(intent.Values, value)
	}

	return wit.UpdateEntity(&intent)
}

func TestIntent() {
	fmt.Println("test intent")	
}

func ReadIntentsFromFile(inputFP string) ([]IntentUtterance, error) {
	input, err := os.Open(inputFP)
	if err != nil {
		return nil, err
	}
	defer input.Close()

	var ius []IntentUtterance
	scanner := bufio.NewScanner(input)
	counter := 0
	for scanner.Scan() {
		counter += 1
		var iu IntentUtterance
		tokens := strings.SplitN(scanner.Text(), ",", 2)
		if len(tokens) < 2 {
			log.Println("Skip line %d: it is not complete: %s", counter, scanner.Text())
			continue
		}
		iu.Intent, iu.Utterance = strings.TrimSpace(tokens[0]), strings.TrimSpace(tokens[1])
		ius = append(ius, iu)
	}

	return ius, nil
}