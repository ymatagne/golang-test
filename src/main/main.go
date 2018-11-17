package main

import "fmt"
import "os"
import "./models"
import "./services"
import "bufio"
import "encoding/json"

func main() {
	argsWithoutProg := os.Args[1:]
	if (len(argsWithoutProg) == 1){
		data, err := os.Open(argsWithoutProg[0])
		if err != nil {
			fmt.Println(err)
		}
		defer data.Close()
		scanner := bufio.NewScanner(data)

		attempts:=[]*models.Attempt{}

		for scanner.Scan() {
			attempt := &models.Attempt{}
			err := json.Unmarshal([]byte(scanner.Text()), attempt)
			if err != nil {
				fmt.Println(err)
			}
			attempts=append(attempts,attempt)
		}

		services.Velocity(attempts)
		

	}else{		
		fmt.Println(" I need arguments !")
	}
}
