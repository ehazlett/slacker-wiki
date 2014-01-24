package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	wikiURL  string
	wikiAuth string
)

type (
	SearchResult struct {
		Group   []string `json:"group"`
		Results []Page   `json:"result"`
	}

	Page struct {
		Title  string `json:"title"`
		Type   string `json:"type"`
		Author string `json:"username"`
		Links  []Link `json:"link"`
	}
	Link struct {
		HREF string `json:"href"`
		Type string `json:"type"`
	}
)

func handler(w http.ResponseWriter, r *http.Request) {
	query := r.FormValue("text")
	client := &http.Client{}
	path := fmt.Sprintf("%s/wiki/rest/prototype/1/search?query=%s&type=search&os_authType=basic", wikiURL, query)
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, fmt.Sprintf("%s", err))
		return
	}
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", wikiAuth))
	req.Header.Set("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, fmt.Sprintf("%s", err))
		return
	}
	defer resp.Body.Close()
	d := json.NewDecoder(resp.Body)
	var result SearchResult
	if err = d.Decode(&result); err != nil {
		log.Println(err)
		fmt.Fprintf(w, fmt.Sprintf("%s", err))
		return
	}
	for _, page := range result.Results {
		var link string
		if len(page.Links) > 0 {
			link = page.Links[0].HREF
		} else {
			link = ""
		}
		if page.Type == "page" {
			s := fmt.Sprintf("%s: <%s>\n", page.Title, link)
			fmt.Fprintf(w, s)
		}
	}
}

func init() {
	flag.StringVar(&wikiURL, "url", "", "Base URL to Confluence")
	flag.StringVar(&wikiAuth, "auth", "", "Basic auth for Confluence (base64 encoded 'username:password')")
	flag.Parse()
	if wikiURL == "" {
		log.Fatal("You must specify a url")
	}
	if wikiAuth == "" {
		log.Fatal("You must specify auth")
	}
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Running on :8080...")
	http.ListenAndServe(":8080", nil)
}
