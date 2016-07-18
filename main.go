package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Data struct {
	Frutas Frutas
}

type Frutas map[string]string

func serverRest(write http.ResponseWriter, request *http.Request) {
	response, err := getResponseJson()
	if err != nil {
		panic(err)
	}

	fmt.Fprintln(write, string(response))
}

func main() {
	http.HandleFunc("/usuario", serverRest)
	http.ListenAndServe("localhost:6565", nil)
}

func getResponseJson() ([]byte, error) {
	frutas := make(map[string]string)
	frutas["1"] = "Maca"
	frutas["2"] = "Laranja"
	frutas["3"] = "Manga"
	frutas["4"] = "Abacate"
	frutas["5"] = "Tangerina"
	frutas["6"] = "Fruta do Conde"

	d := Data{frutas}

	return json.MarshalIndent(d, "", " ")

}