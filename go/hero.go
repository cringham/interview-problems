package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"sync"
)

// Globals
var mutex sync.Mutex
const maxExhaustion = 10
var heroes = make(map[string]Hero)
var fallenHeroes = make(map[string]Hero)

type Hero struct {
	PowerLevel int        `json:"PowerLevel"`
	Exhaustion int        `json:"Exhaustion"`
	Name       string     `json:"Name"`
}

type Calamity struct {
	PowerLevel int		 `json:"PowerLevel"`
	Heroes     []string `json:"Heroes"`
}

type CalamityReport struct {
	OutcomeMessage string   `json:"OutcomeMessage"`
	FallenHeroes   []string `json:"FallenHeroes"`
}

// A superhero has passed away, their name may not be taken up again
func killHeroByName(w http.ResponseWriter, r *http.Request) {
	var name string
	var ok bool
	if name, ok = mux.Vars(r)["name"]; !ok {
		log.Println("Error: There was an issue retrieving var 'name' from GET request.")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	mutex.Lock()
	defer mutex.Unlock()
	_, ok = heroes[name]
	if !ok {
		log.Println("Error: Hero", name, "was not found. Could not kill.")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	fallenHeroes[name] = heroes[name]
	delete(heroes, name)

	fmt.Println("Hero", name, "killed successfully.")
	w.WriteHeader(http.StatusOK)
	return
}

// Retire a superhero, someone may take up their name for the future passing on the title
func retireHeroByName(w http.ResponseWriter, r *http.Request) {
	var name string
	var ok bool
	if name, ok = mux.Vars(r)["name"]; !ok {
		log.Println("Error: There was an issue retrieving var 'name' from GET request.")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	mutex.Lock()
	defer mutex.Unlock()
	_, ok = heroes[name]
	if !ok {
		log.Println("Error: Hero", name, "was not found. Could not retire.")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	delete(heroes, name)

	fmt.Println("Hero", name, "retired successfully.")
	w.WriteHeader(http.StatusOK)
	return
}

// Recover 1 point of exhaustion
func restHeroByName(w http.ResponseWriter, r *http.Request) {
	var name string
	var ok bool
	if name, ok = mux.Vars(r)["name"]; !ok {
		log.Println("Error: There was an issue retrieving var 'name' from GET request.")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var hero Hero
	mutex.Lock()
	defer mutex.Unlock()
	hero, ok = heroes[name]
	if !ok {
		log.Println("Error: Hero", name, "was not found. Could not rest.")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if hero.Exhaustion == 0 {
		fmt.Println("Hero", hero.Name, "already fully rested.")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hero " + hero.Name + " already fully rested."))
		return
	}

	hero.Exhaustion -= 1
	heroes[name] = hero

	fmt.Println("Hero", name, "rested successfully.")
	w.WriteHeader(http.StatusOK)
	return
}

// A calamity of the supplied level requires heroes with an equivalent combined powerlevel to address it.
// Takes a calamity with powerlevel and at least 1 hero. On success return a 200 with json response indicating the calamity has been resolved.
// Otherwise return a response indicating that the heroes were not up to the task. Addressing a calamity adds 1 point of exhaustion.
func inciteCalamity(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error: Problem reading JSON body.")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var calamity Calamity
	err = json.Unmarshal(content, &calamity)
	if err != nil {
		log.Println("Error: JSON formatted incorrectly.")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var heroPowerSum = 0
	var hero Hero
	var ok bool
	mutex.Lock()
	defer mutex.Unlock()
	for _, name := range calamity.Heroes {
		hero, ok = heroes[name]
		if !ok {
			log.Println("Error: Hero", name, "was not found to fight in the calamity.")
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Hero " + name + " was not found to fight in the calamity."))
			return
		}
		heroPowerSum = heroPowerSum + hero.PowerLevel
	}

	var calamityReport CalamityReport
	for _, name := range calamity.Heroes {
		hero, _ = heroes[name]
		hero.Exhaustion += 1
		heroes[name] = hero

		if heroes[name].Exhaustion >= maxExhaustion {
			fallenHeroes[name] = heroes[name]
			delete(heroes, name)

			calamityReport.FallenHeroes = append(calamityReport.FallenHeroes, name)
		}
	}

	if heroPowerSum > calamity.PowerLevel {
		calamityReport.OutcomeMessage = "Heroes were victorious!"
	} else {
		calamityReport.OutcomeMessage = "Heroes could not defeat the calamity."
	}

	calamityReportMarshalled, _ := json.Marshal(calamityReport)
	calamityReportJson := string(calamityReportMarshalled)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(calamityReportJson))
	return
}

// Return a JSON representation of the hero with the name supplied
func getHeroJsonByName(w http.ResponseWriter, r *http.Request) {
	var name string
	var ok bool
	if name, ok = mux.Vars(r)["name"]; !ok {
		log.Println("Error: There was an issue retrieving var 'name' from GET request.")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	mutex.Lock()
	defer mutex.Unlock()
	var hero Hero
	hero, ok = heroes[name]
	if !ok {
		log.Println("Error: Hero", name, "was not found when getting hero.")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	heroMarshalled, _ := json.Marshal(hero)
	heroJson := string(heroMarshalled)

	fmt.Println("Hero", name, "successfully found.")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(heroJson))
	return
}

// Create a superhero according to the JSON body supplied
func makeHero(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error: Problem reading JSON body.")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var hero Hero
	err = json.Unmarshal(content, &hero)
	if err != nil {
		log.Println("Error: JSON formatted incorrectly.")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	mutex.Lock()
	defer mutex.Unlock()
	var ok bool
	if _, ok = fallenHeroes[hero.Name]; ok {
		log.Println("Error: Hero", hero.Name, "has fallen and their name cannot be used in making a new hero.")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Hero " + hero.Name + " has fallen and their name cannot be used in making a new hero."))
		return
	}

	if _, ok = heroes[hero.Name]; ok {
		fmt.Println("Error: Hero", hero.Name, "already exists.")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hero " + hero.Name + " already exists."))
		return
	}

	heroes[hero.Name] = hero

	fmt.Println("Hero", hero.Name, "created successfully.")
	w.WriteHeader(http.StatusOK)
	return
}

func linkRoutes(r *mux.Router) {
	r.HandleFunc("/hero", makeHero).Methods("POST")
	r.HandleFunc("/hero/{name}", getHeroJsonByName).Methods("GET")
	r.HandleFunc("/hero/rest/{name}", restHeroByName).Methods("GET")
	r.HandleFunc("/hero/retire/{name}", retireHeroByName).Methods("GET")
	r.HandleFunc("/hero/kill/{name}", killHeroByName).Methods("GET")
	r.HandleFunc("/calamity", inciteCalamity).Methods("POST")
}

func main() {
	// create a router
	router := mux.NewRouter()
	// and a server to listen on port 8081
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", 8081),
		Handler: router,
	}
	// link the supplied routes
	linkRoutes(router)
	// wait for requests
	log.Fatal(server.ListenAndServe())
}
