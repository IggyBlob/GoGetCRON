// GoGetCRON is a simple app that repeatedly executes HTTP requests
package main

import (
"time"
"net/http"
"log"
"errors"
)

const sleep = 2	// sleep times in hrs

var urls = []string {
	"http://www.paulhaunschmied.com/",
	// ...
}

func main() {
	log.Println("=== RadioCheckerCRON started ===")

	ticker := time.NewTicker(sleep * time.Hour)
	for {
		select {
		case <- ticker.C:
			for _, url := range urls {
				if err := getHTTP(url); err != nil {
					log.Printf("ERR %s\n", err.Error())
					continue
				}
				log.Println("OK  " + url)
			}

		}
	}
}

func getHTTP(url string) error {

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return errors.New("response code " + string(resp.StatusCode) + " for " + url)
	}
	return nil
}
