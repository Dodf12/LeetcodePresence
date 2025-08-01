package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/hugolgst/rich-go/client"
)

type leetcodeData struct {
	Title         string `json:"title"`
	Slug          string `json:"slug"`
	Difficulty    string `json:"difficulty"`
	Status        string `json:"status"`
	Language      string `json:"language"`
	ProblemNumber string `json:"problemNumber"`
}

func postsHandler(w http.ResponseWriter, r *http.Request) {
	// Allow CORS for all origins and methods
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		// Handle preflight request
		w.WriteHeader(http.StatusOK)
		return
	}

	log.Println("[server] postsHandler called with method:", r.Method)
	if r.Method == http.MethodPost {
		var dataclient leetcodeData
		err := json.NewDecoder(r.Body).Decode(&dataclient)
		if err != nil {
			log.Println("[server] Invalid request payload:", err)
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}
		log.Printf("[server] Received data: %+v\n", dataclient)
		updateDiscordPresence(dataclient)
		fmt.Fprintln(w, "Discord presence updated!")
	} else {
		log.Println("[server] Non-POST request received")
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
	}
}

func updateDiscordPresence(data leetcodeData) {
	log.Printf("[server] updateDiscordPresence called with: %+v\n", data)
	err := client.SetActivity(client.Activity{
		State:      fmt.Sprintf("Status: %s", data.Status),
		Details:    fmt.Sprintf("#%s: %s | Difficulty: %s", data.ProblemNumber, data.Title, data.Difficulty),
		LargeImage: "leetcodelogo", // Use the key from the Developer Portal
		LargeText:  fmt.Sprintf("Difficulty: %s", data.Difficulty),
		SmallImage: "leetcode_logo", // Use the key from the Developer Portal 
		Timestamps: &client.Timestamps{
			Start: func() *time.Time { t := time.Now(); return &t }(),
		},
	})
	if err != nil {
		log.Println("[server] Error updating Discord presence:", err)
	} else {
		log.Println("[server] Discord presence updated successfully.")
	}
}

func main() {
	fmt.Println("[server] Starting server...")
	err := client.Login("1400605323097669632")
	if err != nil {
		log.Println("[server] Error logging into Discord:", err)
		return
	}
	log.Println("[server] Logged into Discord successfully.")
	http.HandleFunc("/posts", postsHandler)
	log.Println("[server] Listening on :8080")
	fmt.Println("[server] Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
