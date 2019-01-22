package main

import (
	"fmt"
	"time"
)

type Event struct {
	Name        string
	Start       time.Time
	End         time.Time
	Location    string
	Description string
}

const timeLayout string = "20060102T150405"

func PrintEvent(e Event) {
	fmt.Printf("%s findet vom %s bis zum %s an diesem Ort %s statt. Beschreibung: %s\n", e.Name, e.Start.String(), e.End.String(), e.Location, e.Description)
}
