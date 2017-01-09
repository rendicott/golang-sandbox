package main

import "fmt"
import "bufio"
import "os"
import "regexp"
import "time"

type liner struct {
	ip       []string
	time     string
	hostname []string
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
		}
		// fmt.Printf("Regex '%s' matched to '%s'\n", name, match)
		switch name {
		case "ipaddr":
			profile.ip = match
		case "hostname":
			profile.hostname = match
		}
	}
	return &profile, matched
}

func main() {
	filters := make(map[string]*regexp.Regexp)
	filters["ipaddr"] = regexp.MustCompile(`\b(?:[0-9]{1,3}\.){3}[0-9]{1,3}\b`)
	filters["hostname"] = regexp.MustCompile(`hostname: [a-z0-9\-]*`)

	var linecount int
	var matchcount int
	startTime := time.Now()
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		linecount++
		// fmt.Println(reIP.MatchString(line))
		_, matched := mapLine(line, filters)
		if matched {
			matchcount++
		}
		// fmt.Printf("IP: %s, hostname: %s\n\n", profile.ip, profile.hostname)
		// fmt.Printf("\t%s\n", line)
		if err := input.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "parse-text-stdin: %v\n", err)
		}
	}
	fmt.Printf("Lines: %d\n", linecount)
	fmt.Printf("matches: %d\n", matchcount)
	fmt.Printf("time taken: %s", time.Since(startTime))
}
