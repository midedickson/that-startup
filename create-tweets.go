package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
)

func createTweet() {
	newTweet := Tweet{Text: "@victorobabatund Here's me tagging you with automated tweets to show progress on our bot 😌"}
	newTweetJson, err := json.Marshal(newTweet)
	if err != nil {
		log.Fatal("Error making new tweet.")
	}
	// Create an io.Reader from the JSON data
	reader := bytes.NewReader(newTweetJson)

	response, err := requester.Post("https://api.twitter.com/2/tweets", "application/json", reader)
	if err != nil {
		log.Fatal("Error making new tweet.")
	}
	defer response.Body.Close()
	// Process the response as needed
	fmt.Println("Response Status:", response.Status)
	// Read and print the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	fmt.Println("Response Body:", string(body))
}
