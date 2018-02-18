package main

import (
	"net/http"
	"log"
	"time"
	"strconv"
)

func main() {
	timeout := 10
	website := "https://www.google.com.au"
	for {
		_, err := http.Get(website)
		if err != nil {
			log.Fatal("Internet ded", err)
		}
		log.Println("Found Google now I sleep for "+ strconv.Itoa(timeout) + " Minute(s). To exit press Ctrl-C")
		time.Sleep(time.Duration(timeout) * time.Minute)
	}
}
