package main

import (
	"bufio"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func main() {
	fmt.Println("hello")
	Parse(3, "https://calendar.google.com/calendar/ical/rrlj3mbt8ohkiq9l10j597bavg@group.calendar.google.com/private-1e93e3ef97ce9779732182fe54c7b173/basic.ics", "test2")
}

func Parse(num int, urls ...string) {
	var a []Event
	for _, url := range urls {
		resp, err := http.Get(url)
		if err != nil || resp.StatusCode != 200 {
			continue
		}
		defer resp.Body.Close()
		scanner := bufio.NewScanner(resp.Body)
		for scanner.Scan() {
			if strings.Contains(scanner.Text(), "BEGIN:VEVENT") {
				event, _, _ := parseEvent(scanner)
				a = append(a, event)

			}
		}
	}
}

func parseEvent(scanner *bufio.Scanner) (Event, string, []string) {
	var a Event
	var rrule string
	var exclude []string
	for !strings.Contains(scanner.Text(), "END:VEVENT") {
		line := scanner.Text()
		switch {
		case strings.HasPrefix(line, "DTSTART"): //StartZeit
			if strings.Contains(line, "TZID") {
				location := strings.SplitAfterN(strings.TrimPrefix(line, "DTSTART;TZID="), ":", 2)[0]
				loc, _ := time.LoadLocation(location)
				a.Start, _ = time.ParseInLocation(timeLayout, strings.SplitAfterN(line, ":", 2)[1], loc)
			} else {
				a.Start, _ = time.Parse(timeLayout, strings.SplitAfterN(line, ":", 2)[1])
			}
		case strings.HasPrefix(line, "DTEND"): //EndZeit
			if strings.Contains(line, "TZID") {
				location := strings.SplitAfterN(strings.TrimPrefix(line, "DTEND;TZID="), ":", 2)[0]
				loc, _ := time.LoadLocation(location)
				a.End, _ = time.ParseInLocation(timeLayout, strings.SplitAfterN(line, ":", 2)[1], loc)
			} else {
				a.End, _ = time.Parse(timeLayout, strings.SplitAfterN(line, ":", 2)[1])
			}
		case strings.HasPrefix(line, "LOCATION"): //Ort
			a.Location = strings.SplitAfterN(line, ":", 2)[1]
		case strings.HasPrefix(line, "DESCRIPTION"): //Beschreibung
			a.Description = strings.SplitAfterN(line, ":", 2)[1]
		case strings.HasPrefix(line, "SUMMARY"): //Titel
			a.Name = strings.SplitAfterN(line, ":", 2)[1]
		case strings.HasPrefix(line, "RRULE"): //Wiederholungsregel
			rrule = strings.SplitAfterN(line, ":", 2)[1]
		case strings.HasPrefix(line, "EXDATE"): //Ausnahme zu Wiederholungsregel
			exclude = append(exclude, line)
		}
		scanner.Scan()
	}
	return a, rrule, exclude
}

/*
func parseRRule(e Event, rrule string, ex []string) []Event {
	var all []Event
	not := parseExclude(ex)
	var params map[string]string
	params = make(map[string]string)

	ss := strings.Split(rrule, ";")
	for _, pair := range ss {
		z := strings.Split(pair, "=")
		params[z[0]] = z[1]
	}

	switch params["FREQ"] {
	case "YEARLY": //Jährlich
		if params["BYDAY"] == "" { //Datum und Monat. 12 Januar

		}
		if params["BYDAY"] != "" { //Wochentag und Anzahl der Vorkommnisse: 1. Sonntag  oder so
			params["BYDAY"] //Tag der wiederholung
			params["BYSETPOS"] //vorkommnis des Tags im Monat

		}
	case "MONTHLY": //Monatlich
	case "WEEKLY": //Wöchentlich
	case "DAILY": //Täglich
	case "HOURLY": //Stündlich

	}
}

func parseExclude(ex []string) []time.Time {
	var ret []time.Time
	for _, line := range ex {
		var exdate time.Time
		if strings.Contains(line, "TZID") {
			location := strings.SplitAfterN(strings.TrimPrefix(line, "EXDATE;TZID="), ":", 2)[0]
			loc, _ := time.LoadLocation(location)
			exdate, _ = time.ParseInLocation(timeLayout, strings.SplitAfterN(line, ":", 2)[1], loc)
		} else {
			exdate, _ = time.Parse(timeLayout, strings.SplitAfterN(line, ":", 2)[1])
		}
		ret = append(ret, exdate)
	}
	return ret
}
*/
