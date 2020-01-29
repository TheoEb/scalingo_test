package main

import (
	"github.com/TheoEb/scalingo_test/backend/src/github"
	"github.com/TheoEb/scalingo_test/backend/src/server"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const port = ":8765"

func main() {
	apiKey, err := ioutil.ReadFile(os.Getenv("GITHUB_KEY_FILE"))
	if err != nil {
		log.Println(err)
		os.Exit(2)
	}
	ginEngine := gin.Default()
	g := github.NewClient(strings.TrimSpace(string(apiKey)))
	s := server.NewServer(ginEngine, g, port)
	s.Init()

	if err := s.Run(); err != nil {
		log.Println(err)
		os.Exit(2)
	}
}
