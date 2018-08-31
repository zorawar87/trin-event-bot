package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

/*
<title>All Events</title>
<description>Trinity College Events</description>
<link>https://events.trincoll.edu/live/rss/events/header/All%20Events</link>
<category domain="https://events.trincoll.edu/">events</category>
<language>en-us</language>
<ttl>60</ttl>
<generator>LiveWhale 1.6.2</generator>
<atom:link href="https://events.trincoll.edu/live/rss/events/header/All%20Events" rel="self" type="application/rss+xml"/>
*/

// Represents part of the structure of the XML page representing events at Trincoll
type Result struct {
	Title       string  `xml:"channel>title"`
	Description string  `xml:"channel>description"`
	Link        string  `xml:"channel>link"`
	Category    string  `xml:"channel>category"`
	Events      []Event `xml:"channel>item"`
}

func (c Result) String() string {
	return fmt.Sprintf("%s [%s] (%s) Link:<%s> \nEvents:\n%s", c.Title, c.Category, c.Description, c.Link, c.Events)
}

/*
<item>
	<title>First-Year Shabbat</title>
	<link>https://events.trincoll.edu/#!view/event/event_id/61813</link>
	<pubDate>Fri, 31 Aug 2018 22:00:00 +0000</pubDate>
	<guid isPermaLink="true">https://events.trincoll.edu/#!view/event/event_id/61813#61813</guid>
	<livewhale:type>events</livewhale:type>
	<livewhale:id>61813</livewhale:id>
	<livewhale:timezone>America/New_York</livewhale:timezone>
	<livewhale:categories>Religious Service</livewhale:categories>
	<livewhale:ends>Fri, 31 Aug 2018 23:30:00 +0000</livewhale:ends>
	<source url="https://events.trincoll.edu/live/rss/events/header/All%20Events">All Events</source>
</item>
*/
// Represents a single event being held at trincoll
type Event struct {
	Title     string `xml:"title"`
	Link      string `xml:"link"`
	Type      string `xml:"http://www.livewhale.com/ type"`
	Category  string `xml:"http://www.livewhale.com/ categories"`
	StartsUtc string `xml:"pubDate"`
}

func (e Event) String() string {
	return fmt.Sprintf("[%s, %s] %-20s <%-10s> at %-15s\n", e.Category, e.Type, e.Title, e.Link, e.StartsUtc)
}

func main() {
	resp, err := http.Get("https://events.trincoll.edu/live/rss/events/header/All Events")
	panicOnErr(err)
	defer resp.Body.Close()
	bytes, _ := ioutil.ReadAll(resp.Body)

	var events Result
	xml.Unmarshal(bytes, &events)
	fmt.Println(events)

	t, err := time.Parse(time.RFC1123Z, "Tue, 26 Feb 2019 23:30:00 +0000")
	panicOnErr(err)
	fmt.Println(t.Local())
}

func panicOnErr(e error) {
	if e != nil {
		panic(e.Error())
	}
}
