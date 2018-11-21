package main

import (
	"encoding/csv"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)


type match struct {
	date      string `json:"Date"`
	homeTeam  string `json:"HomeTeam"`
	homeGoals string `json:"FTHG"`
	awayTeam  string `json:"AwayTeam"`
	awayGoals string `json:"FTAG"`
}

func LoadScores(s string) []match {
	csvFile, err := os.Open(s)
	if err != nil {
		log.Fatal("error opening file: "+err.Error())
	}
	reader := csv.NewReader(csvFile)
	var scores []match
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatalf("%v %s", line, error)
		}
		scores = append(scores, match{
			date:      line[1],
			homeTeam:  line[2],
			homeGoals: line[4],
			awayTeam:  line[3],
			awayGoals: line[5],
		})
	}

	return scores
}


func GetAllResults(w http.ResponseWriter, r *http.Request) {

	m:= LoadScores("E0.csv")
	var compilation []string
	for i := 1; i < len(m); i++ {
		l := fmt.Sprintf("%v %v %v - %v %v\n", m[i].date, m[i].homeTeam, m[i].homeGoals, m[i].awayGoals, m[i].awayTeam)
		fmt.Print(l)
		compilation = append(compilation, l)

	}

	all_scores := strings.Join(compilation, "")
	fmt.Fprintf(w, all_scores)

	return

}


func GetResultsByDate(w http.ResponseWriter, r *http.Request){

	m:= LoadScores("E0.csv")
	r_param:= mux.Vars(r)
	fmt.Println(r_param["d"])
	var compilation []string
	var l string

	for i := 1; i < len(m); i++ {

		if r_param["d"] == strings.Join(strings.Split(m[i].date, "/"), "") {
			l = fmt.Sprintf("%v %v %v - %v %v\n", m[i].date, m[i].homeTeam, m[i].homeGoals, m[i].awayGoals, m[i].awayTeam)
			compilation = append(compilation, l)

		}

	}

	all_scores := strings.Join(compilation, "")
	fmt.Fprintf(w, all_scores)


}

func GetResultsByTeam(w http.ResponseWriter, r *http.Request){

	m:= LoadScores("E0.csv")
	r_param:= mux.Vars(r)
	fmt.Println(r_param["d"])
	var compilation []string
	var l string

	for i := 1; i < len(m); i++ {

		if strings.ToLower(r_param["d"])  == strings.ToLower(strings.Join(strings.Split(m[i].homeTeam, "/"), ""))|| strings.ToLower(r_param["d"]) == strings.ToLower(strings.Join(strings.Split(m[i].awayTeam, "/"), "")) {
			l = fmt.Sprintf("%v %v %v - %v %v\n", m[i].date, m[i].homeTeam, m[i].homeGoals, m[i].awayGoals, m[i].awayTeam)
			compilation = append(compilation, l)

		}

	}

	all_scores := strings.Join(compilation, "")
	fmt.Fprintf(w, all_scores)


}

func GetResultsByHomeTeam(w http.ResponseWriter, r *http.Request){

	m:= LoadScores("E0.csv")
	r_param:= mux.Vars(r)
	fmt.Println(r_param["d"])
	var compilation []string
	var l string

	for i := 1; i < len(m); i++ {

		if strings.ToLower(r_param["d"])  == strings.ToLower(strings.Join(strings.Split(m[i].homeTeam, "/"), "")) {
			l = fmt.Sprintf("%v %v %v - %v %v\n", m[i].date, m[i].homeTeam, m[i].homeGoals, m[i].awayGoals, m[i].awayTeam)
			compilation = append(compilation, l)

		}

	}

	all_scores := strings.Join(compilation, "")
	fmt.Fprintf(w, all_scores)


}

func GetResultsByAwayTeam(w http.ResponseWriter, r *http.Request) {

	m := LoadScores("E0.csv")
	r_param := mux.Vars(r)
	fmt.Println(r_param["d"])
	var compilation []string
	var l string

	for i := 1; i < len(m); i++ {

		if strings.ToLower(r_param["d"]) == strings.ToLower(strings.Join(strings.Split(m[i].awayTeam, "/"), "")) {
			l = fmt.Sprintf("%v %v %v - %v %v\n", m[i].date, m[i].homeTeam, m[i].homeGoals, m[i].awayGoals, m[i].awayTeam)
			compilation = append(compilation, l)

		}

	}

	all_scores := strings.Join(compilation, "")
	fmt.Fprintf(w, all_scores)
}

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/results", GetAllResults).Methods("GET")
	r.HandleFunc("/bydate/{d}", GetResultsByDate).Methods("GET")
	r.HandleFunc("/byteam/{d}", GetResultsByTeam).Methods("GET")
	r.HandleFunc("/byhteam/{d}", GetResultsByHomeTeam).Methods("GET")
	r.HandleFunc("/byateam/{d}", GetResultsByAwayTeam).Methods("GET")


	log.Fatal(http.ListenAndServe(":8000", r))

}
