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

	// intent.DeleteAllValues()

	for _, iu := range ius {
		var value gowit.Value
		value.Name = iu.Intent
		value.Expressions = append(value.Expressions, iu.Utterance)
		intent.Values = append(intent.Values, value)
	}

	return wit.UpdateEntity(&intent)
}

func TestIntent(wit *gowit.Client, inputFP string) error {
	ius, err := ReadIntentsFromFile(inputFP)
	if err != nil {
		return err
	}

	total := len(ius)
	correct_counter := 0
	for i, iu := range ius {
		meaning, err := wit.Detect(iu.Utterance)
		if err != nil {
			fmt.Printf("%d\tError: %s", i, err.Error())
			continue
		}

		if iu.Intent == meaning.Intent() {
			correct_counter++
			fmt.Printf("%d/%d:\tcorrect\t%d/%d\n", i, total, correct_counter, total)
		} else {
			fmt.Printf("%d/%d:\tincorrect\n", i, total)
		}
	}
	fmt.Printf("Result: %f\n", float64(correct_counter)/float64(total))

	return nil
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
		text := strings.Replace(scanner.Text(), `"`, ``, -1) // Remove double quotes if have
		tokens := strings.SplitN(text, ",", 2)
		if len(tokens) < 2 {
			log.Println("Skip line %d: it is not complete: %s", counter, scanner.Text())
			continue
		}
		iu.Intent, iu.Utterance = strings.TrimSpace(tokens[0]), strings.TrimSpace(tokens[1])
		ius = append(ius, iu)
	}

	return ius, nil
}