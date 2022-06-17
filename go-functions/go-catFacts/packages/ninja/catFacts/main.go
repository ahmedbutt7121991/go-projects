package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var client *http.Client

type CatFact struct {
	Fact string `json:"fact"`
	Length int `json:"length"`
}

func GetJson(url string, target interface{}) error {
	resp, err := client.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}


func GetCatFact() {
	url := "https://catfact.ninja/fact"

	var catFact CatFact

	err := GetJson(url, &catFact)
	if err != nil {
		fmt.Printf("Error getting cat facts: %s\n", err.Error())
	} else {
		fmt.Printf("A Super interesting Cat Fact: %s\n", catFact.Fact)
	}

}

func Main() {

	client = &http.Client{Timeout: 10 * time.Second}

	GetCatFact()

	// catFact := CatFact {
	// 	Fact: "A random cat fact",
	// 	Length: 17,
	// }

	// jsonStr, err := json.Marshal(catFact)
	// if err != nil {
	// 	fmt.Printf("error marshaling: %s\n", err.Error())
	// } else {
	// 	fmt.Printf("Test JSON: %s\n", string(jsonStr))
	// }

}