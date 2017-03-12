// GoGetCRON is a simple app that repeatedly executes HTTP requests
package main

import (
"time"
"net/http"
"log"
"errors"
)

const (
	sleep = 2	// sleep times in hrs
)

var remotes = [][3]string {
	{"http://www.paulhaunschmied.com/"},
	// {"url", "username", "password"}
}

func main() {
	log.Println("=== RadioCheckerCRON started ===")

	ticker := time.NewTicker(sleep * time.Hour)
	for {
		select {
		case <- ticker.C:
			for _, r := range remotes {
				if err := getHTTP(r[0], r[1], r[2]); err != nil {
					log.Printf("ERR %s\n", err.Error())
					continue
				}
				log.Println("OK  " + r[0])
			}

		}
	}
}

func getHTTP(url, username, password string) error {

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.SetBasicAuth(username, password)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return errors.New("response code " + string(resp.StatusCode) + " for " + url)
	}
	return nil
}
