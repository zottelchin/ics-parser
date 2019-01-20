# ICS-Parser
This a a simple ics-parser.
## Motivation
I wrote this because i was tried a few parser for my [E-Ink Server](https://github.com/zottelchin/E-Ink-Kalender-Servern). The first (and its fork) i tried had a memory leak i suspect or i wasn't using it correctly. In the end it made my server behave strange because the process used 1,5 gigs memory. So i tried other ones, some worked only for some of my google calendars other created segfaults. 
So this is my try to do it better. :sweat_smile:
## Usage 
Install it using go:
```bash
go get github.com/zottelchin/ics-parser
```
Because I'm devoloping this for the other project and I'm only interested in future events, there is no way to get older events.
```golang
...
var events []ics-parser.Event{}
events, err := ics-parser.Parse(7, url1, url2, ...)
fmt.Println("The next 7 events (local time) are:")
for _, e := range events {
    fmt.Printf("%d.%s %s takes place at %s.\n", e.StartDay, e.StartMonth, e.Name, e.location)
}
```
The first param is the number of events you want, after that you can put urls to ics files (tested only with google).

License MIT