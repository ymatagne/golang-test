package services

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"../models"
)

// VelocityAnalyze analyse all the "velocity limits" rules
func VelocityAnalyze(events []*models.Event) {
	var accounts map[string]models.Account = make(map[string]models.Account)
	var account models.Account

	for _, event := range events {

		if val, ok := accounts[event.CustomerID]; ok {
			if VerifyIfEventHasBeenObserved(val, event) {
				continue
			}
			account = UpdateAccountToHostNewEvent(val, event)
		} else {
			account = models.Account{LoadID: event.LoadID, MaxPerDay: 0, MaxPerWeek: 0, NumberPerDay: 0, History: make(map[string]bool)}
		}

		amount, err := strconv.ParseFloat(strings.TrimPrefix(event.LoadAmount, "$"), 64)
		if err != nil {
			fmt.Println(err)
		}

		var accepted = CanBeLoaded(account, amount)

		PrintOutput(models.Activity{LoadID: event.LoadID, CustomerID: event.CustomerID, Accepted: accepted})

		accounts[event.CustomerID] = UpdateAccount(account, event, accepted)
	}
}

//UpdateAccount : Update all fields of the account in function of the rules
func UpdateAccount(account models.Account, event *models.Event, accepted bool) models.Account {
	if accepted {
		account.MaxPerDay = account.MaxPerDay + GetAmount(event)
		account.MaxPerWeek = account.MaxPerWeek + GetAmount(event)
		account.NumberPerDay = account.NumberPerDay + 1
		account.LatestDate = event.Time
	}
	account.History[event.LoadID] = accepted
	return account
}

// UpdateAccountToHostNewEvent : Update the total per day or per Week in function of the day
func UpdateAccountToHostNewEvent(account models.Account, event *models.Event) models.Account {
	var isNewDay = IsNewDay(account.LatestDate, event.Time)
	var isNewWeek = IsNewWeek(account.LatestDate, event.Time)
	if isNewWeek == true {
		account.MaxPerWeek = 0
	}
	if isNewDay == true {
		account.MaxPerDay = 0
		account.NumberPerDay = 0
	}
	return account
}

// VerifyIfEventHasBeenObserved : Verify if the event has been already observed, if yes we keep the first and ignore the next.
//	if a load ID is observed more than once for a particular user, all but the first instance can be ignored
func VerifyIfEventHasBeenObserved(account models.Account, event *models.Event) bool {
	if _, exists := account.History[event.LoadID]; exists {
		return true
	}
	return false
}

// PrintOutput : Convert the Activity struct in Json string
func PrintOutput(activity models.Activity) {
	response, err := json.Marshal(activity)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(response))
}

// CanBeLoaded : Verify all the Load rules.
// 	Each customer is subject to three limits:
//
// 		- A maximum of $5,000 can be loaded per day
// 		- A maximum of $20,000 can be loaded per week
// 		- A maximum of 3 loads can be performed per day, regardless of amount
func CanBeLoaded(account models.Account, amount float64) bool {
	if (amount + account.MaxPerDay) >= 5000 {
		return false
	}
	if (amount + account.MaxPerWeek) >= 20000 {
		return false
	}
	if account.NumberPerDay > 3 {
		return false
	}
	return true
}

// IsNewDay : Return True if the day of latestDate is different of the day of the new date
func IsNewDay(latestDate time.Time, newDate time.Time) bool {
	return latestDate.Day() != newDate.Day()
}

// IsNewWeek : Return True if the week of the new Date is different of the week of the latest date
func IsNewWeek(latestDate time.Time, newDate time.Time) bool {
	_, latestIsoWeek := latestDate.ISOWeek()
	_, newIsoWeek := newDate.ISOWeek()
	return latestIsoWeek != newIsoWeek
}

// GetAmount : Get the amount from the Evemt and convert the amount in float64
func GetAmount(event *models.Event) float64 {
	amount, err := strconv.ParseFloat(strings.TrimPrefix(event.LoadAmount, "$"), 64)
	if err != nil {
		fmt.Println(err)
	}
	return amount
}
