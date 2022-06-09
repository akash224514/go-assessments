package main

import (
	"fmt"
	"log"
	"net/http"
)

//var wg sync.WaitGroup

func main() {
	fmt.Println("\n***Program for capturing the status codes for errors***")

	websitelist := []string{
		"https://google.com",
		"https://go.dev",
		"https://fb.com",
		"https://github.com",
		"https://youtube.com",
		"https://www.metamaxx.io",
		"https://stackoverflow.com",
		"https://akashshelkande.com", // Providing an Invalid URL to capture the error produced by HTTP request
	}

	for _, web := range websitelist {

		getStatusCode(web)
		//wg.Add(1)
	}
	//wg.Wait()
}

func getStatusCode(site string) {
	//defer wg.Done()
	res, err := http.Get(site) // Using Net/HTTPs Get method to recieve a status for the URL

	if err != nil { //handling an error for a visited URL (if any)
		log.Fatal(err)
		//fmt.Printf("\n%d status code for %s\n", res.StatusCode, site)
	} else {
		fmt.Printf("\nstatus code: %d status: %s for %s\n", res.StatusCode, res.Status, site) // Producing an output status code as well as status

	}

}

/*
APPROACH ->
-Here, we need to capture the status code for errors produced by an HTTP Request
- We are creating a slice of different URLs to which we are making a request using HTTPs get method.
- If URL is valid and server responds to the request then we will be getting status code 200 and status like OK
- If URL is invalid or server is not responding back then we can expect a detailed error description with exit status 1
*/
