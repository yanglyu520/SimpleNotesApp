package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	mux := http.NewServeMux()
	p := PlayerServer{}
	mux.HandleFunc("/players", p.GetAPlayerScore)
	fmt.Println("starting server")
	http.ListenAndServe(":8000", nil)
}

func (p *PlayerServer) GetAPlayerScore(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	fmt.Fprint(w, GetPlayerScore(player))
}

type PlayerServer struct {
	store PlayerStore
}

type PlayerStore interface {
	GetPlayerScore(player string) int
}

func GetPlayerScore(player string) int {
	switch player {
	case "A":
		return 3
	case "B":
		return 2
	default:
		return 1
	}
}
