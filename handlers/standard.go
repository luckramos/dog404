package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"text/template"
	"time"

	"math/rand"
)

type Manifest map[string]struct {
	File string   `json:"file"`
	Css  []string `json:"css"`
}

type Dog struct {
	Name  string
	Image string
}

var dogList = []Dog{
	{Name: "Buddy", Image: "images/dog1.jpg"},
	{Name: "Charlie", Image: "images/dog2.jpg"},
	{Name: "Max", Image: "images/dog3.jpg"},
	{Name: "Rocky", Image: "images/dog4.jpg"},
}

type PageData struct {
	CSSFile string
	JSFile  string
	Dog     Dog
}

func getRandomDog() Dog {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randN := r.Intn(len(dogList))
	return dogList[randN]
}

func loadManifest() (Manifest, error) {
	file, err := os.Open("../static/dist/.vite/manifest.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var manifest Manifest
	if err := json.NewDecoder(file).Decode(&manifest); err != nil {
		return nil, err
	}
	return manifest, nil
}

func StandardHandler(w http.ResponseWriter, r *http.Request) {

	manifest, err := loadManifest()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("../static/dist/index.html"))

	entry, ok := manifest["index.html"]
	if !ok {
		http.Error(w, "Não foi possível encontrar a entrada 'index.html' no manifest", http.StatusInternalServerError)
		return
	}

	data := PageData{
		CSSFile: entry.Css[0],
		JSFile:  entry.File,
		Dog:     getRandomDog(),
	}

	tmpl.Execute(w, data)
}
