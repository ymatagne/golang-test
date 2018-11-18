package services

import (
	"strconv"
	"testing"
	"time"

	"../models"
)

func TestReturnTrueIfIsTheNewDay(t *testing.T) {
	yesterday := time.Date(2018, 11, 17, 12, 30, 0, 0, time.UTC)
	today := time.Date(2018, 11, 18, 12, 30, 0, 0, time.UTC)

	var isNewDay = IsNewDay(yesterday, today)
	if isNewDay == false {
		t.Errorf("Expected isNewDay of true, but it was %s instead.", strconv.FormatBool(isNewDay))
	}
}

func TestReturnFalseIfItsNotNewDay(t *testing.T) {
	yesterday := time.Date(2018, 11, 17, 12, 30, 0, 0, time.UTC)
	today := time.Date(2018, 11, 17, 12, 30, 0, 0, time.UTC)

	var isNewDay = IsNewDay(yesterday, today)
	if isNewDay == true {
		t.Errorf("Expected isNewDay of false, but it was %s instead.", strconv.FormatBool(isNewDay))
	}
}

func TestReturnTrueIfIsTheNeWeek(t *testing.T) {
	currentWeek := time.Date(2018, 11, 17, 12, 30, 0, 0, time.UTC)
	nextWeek := time.Date(2018, 11, 15, 12, 30, 0, 0, time.UTC)

	var isNewWeek = IsNewWeek(currentWeek, nextWeek)
	if isNewWeek == true {
		t.Errorf("Expected isNewWeek of false, but it was %s instead.", strconv.FormatBool(isNewWeek))
	}
}
func TestReturnFalseIfIsNotANeWeek(t *testing.T) {
	currentWeek := time.Date(2018, 11, 17, 12, 30, 0, 0, time.UTC)
	nextWeek := time.Date(2018, 11, 25, 12, 30, 0, 0, time.UTC)

	var isNewWeek = IsNewWeek(currentWeek, nextWeek)
	if isNewWeek == false {
		t.Errorf("Expected isNewWeek of true, but it was %s instead.", strconv.FormatBool(isNewWeek))
	}
}

func TestReturnTrueIfTheEventCanBeLoadOnTheAccount(t *testing.T) {

	var account = models.Account{LoadID: "1", MaxPerDay: 0, MaxPerWeek: 0, NumberPerDay: 0, History: make(map[string]bool)}
	var amount = 10.0

	var canBeLoaded = CanBeLoaded(account, amount)
	if canBeLoaded == false {
		t.Errorf("Expected isNewWeek of true, but it was %s instead.", strconv.FormatBool(canBeLoaded))
	}
}

func TestReturnFalseIfTheEventCannotBeLoadOnTheAccount(t *testing.T) {

	var account = models.Account{LoadID: "1", MaxPerDay: 0, MaxPerWeek: 0, NumberPerDay: 0, History: make(map[string]bool)}
	var amount = 10000.0

	var canBeLoaded = CanBeLoaded(account, amount)
	if canBeLoaded == true {
		t.Errorf("Expected isNewWeek of true, but it was %s instead.", strconv.FormatBool(canBeLoaded))
	}
}

func TestReturnFalseIfTheEventCannotBeLoadOnTheAccountBecauseMaxOfWeek(t *testing.T) {

	var account = models.Account{LoadID: "1", MaxPerDay: 0, MaxPerWeek: 20000, NumberPerDay: 0, History: make(map[string]bool)}
	var amount = 1000.0

	var canBeLoaded = CanBeLoaded(account, amount)
	if canBeLoaded == true {
		t.Errorf("Expected isNewWeek of true, but it was %s instead.", strconv.FormatBool(canBeLoaded))
	}
}

func TestReturnFalseIfTheEventCannotBeLoadOnTheAccountBecauseNumberOfLoaded(t *testing.T) {

	var account = models.Account{LoadID: "1", MaxPerDay: 0, MaxPerWeek: 10000, NumberPerDay: 4, History: make(map[string]bool)}
	var amount = 1000.0

	var canBeLoaded = CanBeLoaded(account, amount)
	if canBeLoaded == true {
		t.Errorf("Expected isNewWeek of true, but it was %s instead.", strconv.FormatBool(canBeLoaded))
	}
}
