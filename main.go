package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type DateTime struct {
	DateTime string `json:"datetime"`
	TimeZone string `json:"timeZone"`
}

type Event struct {
	Summary     string   `json:"summary"`
	Description string   `json:"description"`
	Location    string   `json:"location"`
	Start       DateTime `json:"start"`
	end         DateTime `json:"end"`
}

func Crawl(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var data map[string][]Event
	json.Unmarshal(body, &data)

	fmt.Printf("\nData \n")
	for _, item := range data["items"] {
		fmt.Printf("Summary:\n")
		fmt.Println(item.Summary)
		fmt.Printf("Description:\n")
		fmt.Println(item.Description)
	}
	return
}

func main() {
	urls := []string{
		fmt.Sprintf("https://www.googleapis.com/calendar/v3/calendars/santaclara.acm@gmail.com/events?key=%s&timeMin=2018-02-24T00:00:00Z&timeMax=2018-04-09T00:00:00Z&singmaxResults=9999&_=1520708172234", os.Getenv("ACM_KEY")),
		fmt.Sprintf("https://clients6.google.com/calendar/v3/calendars/csl@scu.edu/events?calendarId=csl@scu.edu&singleEvents=true&timeZone=America/Los_Angeles&maxAttendees=1&maxResults=250&sanitizeHtml=true&timeMin=2018-02-26T00:00:00-08:00&timeMax=2018-04-02T00:00:00-08:00&key=%s", os.Getenv("CSL_KEY")),
	}
	for _, url := range urls {
		Crawl(url)
	}
}
