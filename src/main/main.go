package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"./models"
	"./services"
)

// main is the main function
// Just parsing the file on list of struct.
func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) == 1 {
		data, err := os.Open(argsWithoutProg[0])
		if err != nil {
			fmt.Println(err)
		}
		defer data.Close()
		scanner := bufio.NewScanner(data)

		events := []*models.Event{}

		for scanner.Scan() {
			event := &models.Event{}
			err := json.Unmarshal([]byte(scanner.Text()), event)
			if err != nil {
				fmt.Println(err)
			}
			events = append(events, event)
		}

		services.VelocityAnalyze(events)

	} else {
		fmt.Println(" I need arguments !")
	}
}
