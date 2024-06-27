package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sort"
)

// App struct
type App struct {
	ctx context.Context
}

type RandomImage struct {
	Message string
	Status  string
}

type AllBreeds struct {
	Message map[string][]string `json:"message"`
	Status  string              `json:"status"`
}

type ImageByBreed struct {
	Message []string `json:"message"`
	Status  string   `json:"status"`
}

func (a *App) GetRandomImageUrl() string {
	log.Println("GetRandomImageUrl")
	response, err := http.Get("https://dog.ceo/api/breeds/image/random")
	if err != nil {
		log.Fatal(err)
	}
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var data RandomImage
	err = json.Unmarshal(responseData, &data)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(data.Message)
	return data.Message
}

func (a *App) GetBreedList() []string {
	var breeds []string
	response, err := http.Get("https://dog.ceo/api/breeds/list/all")
	if err != nil {
		log.Fatal(err)
	}
	responseData, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}
	var data AllBreeds
	err = json.Unmarshal(responseData, &data)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(data.Message)
	for k := range data.Message {

		breeds = append(breeds, k)
	}
	sort.Strings(breeds)
	log.Println(breeds)
	return breeds
}

func (a *App) GetImageUrlByBreed(breed string) []string {
	url := fmt.Sprintf("%s%s%s%s", "https://dog.ceo/api/", "breed/", breed, "/images")
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var data ImageByBreed
	log.Println(string(responseData))
	err = json.Unmarshal(responseData, &data)
	if err != nil {
		log.Fatal(err)
	}
	return data.Message
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}
