package main

type Map struct {
	Count    int `json:"count"`
	NextMap *string `json:"next"`
	PreviousMap *string `json:"previous"`
	Locations []Location `json:"results"`
}

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type CurrentLocation struct {
	Next      *string
	Previous *string
}