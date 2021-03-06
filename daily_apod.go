package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

// Apod struct
type Apod struct {
	Copyright       string `json:"copyright"`       //copyright
	Date            string `json:"date"`            // date
	Explanation     string `json:"explanation"`     // explanation
	Hdurl           string `json:"hdurl"`           // hdurl
	Media_type      string `json:"media_type"`      // media_type
	Service_version string `json:"service_version"` //service_version
	Title           string `json:"title"`           // title
	Url             string `json:"url"`             // url
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJSON(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	fmt.Print(r.Body)

	return json.NewDecoder(r.Body).Decode(target)
}

func main() {

	url := "https://api.nasa.gov/planetary/apod?api_key=lmkpX63zD6oG70WGFswSig6GY6pze0TtemdkciYz"
	apod := new(Apod)

	var myClient = &http.Client{Timeout: 10 * time.Second}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := myClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	jsonErr := json.Unmarshal(body, &apod)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	// get image from apod URL
	imageresp, err := http.Get(apod.Hdurl)
	if err != nil {
		log.Fatalln(err)
	}

	// create empty image file
	file, err := os.Create("apod_image.jpg")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	//Write the bytes to the fiel
	_, err = io.Copy(file, imageresp.Body)
	if err != nil {
		log.Fatalln(err)
	}

}
