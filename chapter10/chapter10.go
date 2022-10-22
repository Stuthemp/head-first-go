package main

import (
	"awesomeProject/chapter10/calendar"
	"fmt"
	"log"
)

func main() {
	event := calendar.Event{}
	err := event.SetYear(2019)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetMonth(12)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetDay(26)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetTitle("Joe Mama")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(event)
}
