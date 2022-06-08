package models

import (
	"fmt"
	"strconv"
)

type AfterEod struct {
	ID               int
	Nama             string
	Age              int
	Balanced         int
	No2BThreadNo     string
	No3ThreadNo      string
	PreviousBalanced int
	AverageBalanced  float32
	No1ThreadNo      string
	FreeTransfer     int
	No2AThreadNo     string
}

func NewAfterEod(data []string) (*AfterEod, error) {

	id, err := strconv.Atoi(data[0])
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	nama := data[1]
	age, err := strconv.Atoi(data[2])
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	balanced, err := strconv.Atoi(data[3])
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	previousBalanced, err := strconv.Atoi(data[4])
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	averageBalanced, err := strconv.ParseFloat(data[5], 32)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	freeTransfer, err := strconv.Atoi(data[6])
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &AfterEod{
		ID:               id,
		Nama:             nama,
		Age:              age,
		Balanced:         balanced,
		PreviousBalanced: previousBalanced,
		AverageBalanced:  float32(averageBalanced),
		FreeTransfer:     freeTransfer,
	}, nil
}
