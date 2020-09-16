package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

// Trend API response JSON
type Trend struct {
	Author   string `json:"author,omitempty"`
	Name     string `json:"name,omitempty"`
	URL      string `json:"url,omitempty"`
	Language string `json:"language,omitempty"`
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/repo", getRandomTrendHandler).Methods("GET")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	log.Printf("Listening http on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func getRandomTrendHandler(outResp http.ResponseWriter, inReq *http.Request) {
	apiOutReq, err := http.NewRequest("GET", "https://ghapi.huchen.dev/repositories", nil)
	if err != nil {
		outResp.WriteHeader(http.StatusInternalServerError)
		outResp.Write([]byte(err.Error()))
		return
	}
	apiOutReq.Header.Add("Accept", "application/json")
	outQuery := apiOutReq.URL.Query()
	inReqQueryValues := inReq.URL.Query()
	for paramName, paramValue := range inReqQueryValues {
		if paramName != "redirect" {
			outQuery.Add(paramName, paramValue[0])
		}
	}
	apiOutReq.URL.RawQuery = outQuery.Encode()
	client := &http.Client{}
	apiResp, err := client.Do(apiOutReq)
	if err != nil {
		outResp.WriteHeader(http.StatusInternalServerError)
		outResp.Write([]byte(err.Error()))
		return
	}
	defer apiResp.Body.Close()
	apiRespBody, err := ioutil.ReadAll(apiResp.Body)
	if err != nil {
		outResp.WriteHeader(http.StatusInternalServerError)
		outResp.Write([]byte(err.Error()))
		return
	}
	var trendsArray []Trend
	err = json.Unmarshal(apiRespBody, &trendsArray)
	if err != nil {
		outResp.WriteHeader(http.StatusInternalServerError)
		outResp.Write([]byte(err.Error()))
		return
	}
	// TODO random seed
	rndIndex := rand.Intn(len(trendsArray))
	rndURL := trendsArray[rndIndex].URL
	redirect := false
	if redirectParam, ok := inReqQueryValues["redirect"]; ok {
		redirect = strings.ToLower(redirectParam[0]) == "true"
	}
	outResp.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, post-check=0, pre-check=0")
	outResp.Header().Set("Pragma", "no-cache")
	outResp.Header().Set("Expires", "0")
	if redirect {
		http.Redirect(outResp, inReq, rndURL, 301)
	} else {
		outResp.WriteHeader(http.StatusOK)
		outResp.Write([]byte(rndURL))
	}
}
