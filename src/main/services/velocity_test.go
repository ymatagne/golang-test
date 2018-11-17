package services

import (
	"strconv"
	"testing"
	"time"
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
