package services

import "fmt"
import "../models"

func Velocity( attempts []*models.Attempt) {
	for _, attempt := range attempts {
        fmt.Println(attempt)
    }
}
