package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var CoffeeServed = 0
var Product = getEnv("PRODUCT", "K-Compact")
var AdmPassword = getEnv("ADMIN_PWD", "password123")
var ConfigFile = loadConfiguration("config.json")

type Config struct {
	MaxUsedPods      int      `json:"maxUsedPods"`
	AvailableFlavors []string `json:"availableFlavors"`
}

type Stats struct {
	MaxPods   int    `json:"maxPods"`
	Remaining int    `json:"remainingPods"`
	Served    int    `json:"totalServed"`
	Product   string `json:"product"`
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func loadConfiguration(file string) Config {
	var config Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("/health HIT")

	stats := getCurrentStats()
	if stats.Remaining <= 0 {
		w.WriteHeader(http.StatusTeapot)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{"alive":false}`)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{"alive":true}`)
	}
}

func readyCheckHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("/ready HIT")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"ready":true}`)
}

func coffeeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("/coffee HIT")

	stats := getCurrentStats()
	if stats.Remaining <= 0 {
		w.WriteHeader(http.StatusTeapot)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, "No more coffee for you!")
	} else {

		CoffeeServed = CoffeeServed + 1
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, "coffee!")
	}
}

func flavorsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("/flavors HIT")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, "flavors coming soon!")
}

func configHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("/config HIT")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(ConfigFile)
}

func getCurrentStats() Stats {
	stats := Stats{MaxPods: ConfigFile.MaxUsedPods, Remaining: (ConfigFile.MaxUsedPods - CoffeeServed), Served: CoffeeServed, Product: Product}
	return stats
}

func statsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("/stats HIT")

	stats := Stats{MaxPods: ConfigFile.MaxUsedPods, Remaining: (ConfigFile.MaxUsedPods - CoffeeServed), Served: CoffeeServed, Product: Product}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(stats)
}

func resetHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("/reset HIT")

	pwds, err := r.URL.Query()["password"]
	if !err || len(pwds) != 1 {
		log.Println("AUTHN: Incorrect number of passwords")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	pwd := pwds[0]
	if pwd != AdmPassword {
		log.Println("AUTHN: Incorrect password")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	log.Println("RESET: Keurig reset")
	CoffeeServed = 0
}

func main() {
	fmt.Println("Starting Keurig machine...")
	http.HandleFunc("/ready", readyCheckHandler)
	http.HandleFunc("/health", healthCheckHandler)
	http.HandleFunc("/coffee", coffeeHandler)
	http.HandleFunc("/stats", statsHandler)
	http.HandleFunc("/flavors", flavorsHandler)
	http.HandleFunc("/config", configHandler)
	http.HandleFunc("/reset", resetHandler)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
