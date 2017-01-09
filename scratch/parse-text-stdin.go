package main

import "fmt"
import "bufio"
import "os"
import "regexp"
import "time"

const dateFormat = "2006-01-02T15:04:05-07:00"

// define the regexes
const reTimestamp = `[0-9]{4}-[0-9]{2}-[0-9]{2}T[0-9]{2}:[0-9]{2}:[0-9]{2}-[0-9]{2}:[0-9]{2}`
const reHostname = `hostname: [a-z0-9\-]*`
const reIPAddr = `\b(?:[0-9]{1,3}\.){3}[0-9]{1,3}\b`

type liner struct {
	ip        []string
	timestamp time.Time
	hostname  []string
}

func (l *liner) display() string {
	return fmt.Sprintf("IP: %s, hostname: %s, timestamp: %s\n\n", l.ip, l.hostname, l.timestamp.Format(time.ANSIC))
}

func mapLine(s string, filters map[string]*regexp.Regexp) (*liner, bool) {
	var profile liner
	var matched bool
	for name, regex := range filters {
		match := regex.FindAllString(s, -1)
		if match == nil {
			matched = false
		} else {
			matched = true

			switch name {
			case "ipaddr":
				profile.ip = match
			case "hostname":
				profile.hostname = match
			case "timestamp":
				profile.timestamp, _ = time.Parse(dateFormat, match[0])
			}
		}

	}
	return &profile, matched
}

func mapLineSlim(s string, filters map[string]*regexp.Regexp) bool {
	var matched bool
	for _, regex := range filters {
		match := regex.FindAllString(s, -1)
		if match == nil {
			matched = false
		} else {
			matched = true
		}
	}
	return matched
}

func main() {
	filters := make(map[string]*regexp.Regexp)
	filters["ipaddr"] = regexp.MustCompile(reIPAddr)
	filters["hostname"] = regexp.MustCompile(reHostname)
	filters["timestamp"] = regexp.MustCompile(reTimestamp)

	var linecount int
	var matchcount int
	startTime := time.Now()
	input := bufio.NewScanner(os.Stdin)
	var profiles []string
	for input.Scan() {
		line := input.Text()
		linecount++
		// fmt.Println(reIP.MatchString(line))
		profile, matched := mapLine(line, filters)
		// matched := mapLineSlim(line, filters)
		if matched {
			matchcount++
		}
		profiles = append(profiles, profile.display())
		// fmt.Printf("\t%s\n", line)
		if err := input.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "parse-text-stdin: %v\n", err)
		}
	}
	fmt.Printf("Lines: %d\n", linecount)
	fmt.Printf("matches: %d\n", matchcount)
	fmt.Printf("time taken: %s\n", time.Since(startTime))
	fmt.Printf("Length of profiles: %d", len(profiles))
}
