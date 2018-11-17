package services

import "fmt"
import "time"
import "../models"
import "strconv"
import "strings"
import "encoding/json"

func Velocity(attempts []*models.Attempt) {
	var isNewDay bool = false
	var isNewWeek bool = false
	var attemptPerCustomerId map[string]models.Account = make(map[string]models.Account)
	var actualAccount models.Account

	for _, attempt := range attempts {

		if val, ok := attemptPerCustomerId[attempt.CustomerId]; ok {
			actualAccount = val
			isNewDay = IsNewDay(actualAccount.LatestDate, attempt.Time)
			isNewWeek = IsNewWeek(actualAccount.LatestDate, attempt.Time)
			//if attempt.CustomerId == "137" {
			//	fmt.Println("**************")
			//	fmt.Println("               ")
			//	fmt.Println("               ")
			//	fmt.Println(isNewWeek)
			//	fmt.Println(actualAccount.LatestDate)
			//	fmt.Println(attempt.Time)
			//}
			if isNewWeek == true {
				actualAccount.MaxPerWeek = 0
			}
			if isNewDay == true {
				actualAccount.MaxPerDay = 0
				actualAccount.NumberPerDay = 0
			}
		} else {
			actualAccount = models.Account{Id: attempt.Id, MaxPerDay: 0, MaxPerWeek: 0, NumberPerDay: 0}
		}

		amount, err := strconv.ParseFloat(strings.TrimPrefix(attempt.LoadAmount, "$"), 64)
		if err != nil {
			fmt.Println(err)
		}

		//if attempt.CustomerId == "137" {
		//	fmt.Println("**************")
		//	fmt.Println("               ")
		//	fmt.Println("               ")
		//	fmt.Println("Want to add : " + string(attempt.LoadAmount))
		//	fmt.Println(attempt.Time)
		//	fmt.Println(isNewDay)
		//	fmt.Println(isNewWeek)
		//	fmt.Println(actualAccount)
		//}
		if GetAmount(actualAccount, amount) == true {
			//if attempt.CustomerId == "137" {
			PrintOutput(models.Activity{Id: attempt.Id, CustomerId: attempt.CustomerId, Accepted: true})
			//}
			actualAccount.MaxPerDay = actualAccount.MaxPerDay + amount
			actualAccount.MaxPerWeek = actualAccount.MaxPerWeek + amount
			actualAccount.NumberPerDay = actualAccount.NumberPerDay + 1
			actualAccount.LatestDate = attempt.Time
		} else {
			//if attempt.CustomerId == "137" {
			PrintOutput(models.Activity{Id: attempt.Id, CustomerId: attempt.CustomerId, Accepted: false})
			//}
		}

		attemptPerCustomerId[attempt.CustomerId] = actualAccount
	}
}

func PrintOutput(activity models.Activity) {
	response, err := json.Marshal(activity)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(response))
}

func GetAmount(actualAccount models.Account, amount float64) bool {
	if (amount + actualAccount.MaxPerDay) >= 5000 {
		return false
	}
	if (amount + actualAccount.MaxPerWeek) >= 20000 {
		return false
	}
	if actualAccount.NumberPerDay > 3 {
		return false
	}
	return true
}

func IsNewDay(actualDate time.Time, newDate time.Time) bool {
	return actualDate.Day() != newDate.Day()
}

func IsNewWeek(latestDate time.Time, newDate time.Time) bool {
	_, latestIsoWeek := latestDate.ISOWeek()
	_, newIsoWeek := newDate.ISOWeek()
	return latestIsoWeek != newIsoWeek
}
