package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"googlemaps.github.io/maps"
)

var apiKey = "YOUR_API_KEY_DUDE"
var stationsFile = "transum.json"

type Station struct {
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Type      string  `json:"type"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/stations", getStations).Methods("GET")
	r.HandleFunc("/check-location", checkLocation).Methods("POST")
	http.Handle("/", r)
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getStations(w http.ResponseWriter, r *http.Request) {
	stations, err := readStationsFromFile(stationsFile)
	if err != nil {
		http.Error(w, "Failed to read stations data", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(stations); err != nil {
		http.Error(w, "Failed to encode stations data", http.StatusInternalServerError)
	}
}

func readStationsFromFile(filename string) ([]Station, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var stations []Station
	if err := json.Unmarshal(byteValue, &stations); err != nil {
		return nil, err
	}

	return stations, nil
}

func checkLocation(w http.ResponseWriter, r *http.Request) {
	var loc Station
	err := json.NewDecoder(r.Body).Decode(&loc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("Received location: %+v\n", loc)

	c, err := maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	dest := &maps.GeocodingRequest{
		Address: "Stasiun Tujuan",
	}

	res, err := c.Geocode(context.Background(), dest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(res) == 0 {
		http.Error(w, "destination not found", http.StatusNotFound)
		return
	}

	destLoc := res[0].Geometry.Location
	log.Printf("Destination location: %+v\n", destLoc)

	userLoc := maps.LatLng{Lat: loc.Latitude, Lng: loc.Longitude}
	if isWithinRadius(userLoc, destLoc, 0.1) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("User has arrived at the destination"))
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User is not at the destination"))
	}
}

func isWithinRadius(loc1, loc2 maps.LatLng, radiusKm float64) bool {
	const earthRadiusKm = 6371.0
	dLat := (loc2.Lat - loc1.Lat) * (math.Pi / 180)
	dLon := (loc2.Lng - loc1.Lng) * (math.Pi / 180)

	lat1 := loc1.Lat * (math.Pi / 180)
	lat2 := loc2.Lat * (math.Pi / 180)

	a := (0.5 - (0.5 * math.Cos(dLat))) + (math.Cos(lat1) * math.Cos(lat2) * (1 - math.Cos(dLon)) / 2)
	return (earthRadiusKm * 2 * math.Asin(math.Sqrt(a))) <= radiusKm
}
