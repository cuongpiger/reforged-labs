package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

// AdRequest represents the request payload
type AdRequest struct {
	Title          string   `json:"title"`
	Description    string   `json:"description"`
	Genre          string   `json:"genre"`
	TargetAudience []string `json:"targetAudience"`
	VisualElements []string `json:"visualElements"`
	CallToAction   string   `json:"callToAction"`
	Duration       int      `json:"duration"`
	Priority       int      `json:"priority"`
}

var titles = []string{"Dragon Kingdom", "Warriors of Light", "Cyber Future", "Mystic Legends"}
var descriptions = []string{"Build your empire!", "Fight in epic battles!", "Explore the unknown!", "Become the hero!"}
var genres = []string{"Strategy", "Action", "Adventure", "RPG"}
var audiences = [][]string{{"18-34", "Strategy Gamers"}, {"25-40", "Hardcore Gamers"}, {"13-25", "Casual Gamers"}}
var visuals = [][]string{{"Dragons", "Castle", "Battle Scenes"}, {"Cyber Warriors", "Future City", "Drones"}, {"Magic Spells", "Mystical Creatures", "Dark Forest"}}
var actions = []string{"Download Now!", "Play for Free!", "Join the Adventure!"}

// generateRandomAd creates a random AdRequest
func generateRandomAd() AdRequest {
	return AdRequest{
		Title:          titles[rand.Intn(len(titles))],
		Description:    descriptions[rand.Intn(len(descriptions))],
		Genre:          genres[rand.Intn(len(genres))],
		TargetAudience: audiences[rand.Intn(len(audiences))],
		VisualElements: visuals[rand.Intn(len(visuals))],
		CallToAction:   actions[rand.Intn(len(actions))],
		Duration:       rand.Intn(60) + 10, // Random duration between 10 and 70 seconds
		Priority:       rand.Intn(100) + 1, // Random priority between 1 and 100
	}
}

// sendRequest sends an HTTP POST request with a random ad payload
func sendRequest(wg *sync.WaitGroup) {
	defer wg.Done()

	ad := generateRandomAd()
	jsonData, _ := json.Marshal(ad)

	req, err := http.NewRequest("POST", "http://127.0.0.1:8000/api/v1/ads", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Request creation failed:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Request failed:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("Response Status: %s\n", resp.Status)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var wg sync.WaitGroup
	numRequests := 100

	for i := 0; i < numRequests; i++ {
		wg.Add(1)
		go sendRequest(&wg)
	}

	wg.Wait()
	fmt.Println("All requests completed.")
}
