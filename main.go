package main

import (
	"bufio"
	"os"
	"fmt"
	"net/http"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		subdomain := scanner.Text()
		url := "http://" + subdomain + ".s3.amazonaws.com/"

		req, _ := http.NewRequest("GET", url, nil)
		client := new(http.Client)
		res, _ := client.Do(req)

		if res.StatusCode == 200 {
			fmt.Println("\x1b[32m[+] " + url + " [200] [bucketname: " + subdomain+ "]\x1b[0m")
		} else if res.StatusCode == 403 {
			fmt.Println("\x1b[33m[+] " + url + " [403] [bucketname: " + subdomain+ "]\x1b[0m")
		}

		time.Sleep(time.Second * time.Duration(1))
	}
}
