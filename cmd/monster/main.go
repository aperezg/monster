package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/aperezg/monster"

	"github.com/aperezg/monster/server"
	"github.com/aperezg/monster/storage"
)

func main() {

	host := flag.String("host", "localhost", "api host")
	port := flag.Int("port", 3000, "api port")
	withData := flag.Bool("withData", false, "initialize the api with some monsters")
	flag.Parse()
	var monsters map[string]*monster.Monster
	if *withData {
		monsters = monster.Monsters
	}

	repo := storage.NewMonsterRepository(monsters)

	log.Printf("Server running on: http://%s:%d\n", *host, *port)

	err := http.ListenAndServe(fmt.Sprintf("%s:%d", *host, *port), server.NewApi(repo))
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
