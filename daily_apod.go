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
	copyright       string `json:"copyright"`       //copyright
	date            string `json:"date"`            // date
	explanation     string `json:"explanation"`     // explanation
	hdurl           string `json:"hdurl"`           // hdurl
	media_type      string `json:"media_type"`      // media_type
	service_version string `json:"service_version"` //service_version
	title           string `json:"title"`           // title
	url             string `json:"url"`             // url
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

	url := "https://api.nasa.gov/planetary/apod?api_key=lmkpX63zD6oG70WGFswSig6GY6pze0TtemdkciYz&date=2020-07-15"
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

	fmt.Println(apod.url)

	// get image from apod URL
	imageresp, err := http.Get(apod.url)
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
