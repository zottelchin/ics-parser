package main

import (
	"time"
)

type Event struct {
	Name     string
	Start    time.Time
	End      time.Time
	Location string
}
