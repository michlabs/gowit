package main 

import (
	"fmt"

	"github.com/michlabs/gowit"
	"github.com/michlabs/nlu"
)

func TrainIntent(wit *gowit.Client, inputFP string) error {
	ius, err := nlu.ReadIntentsFromFile(inputFP)
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
	ius, err := nlu.ReadIntentsFromFile(inputFP)
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