package main

import (
	"log"
	"os"

	"github.com/juddbaguio/top-ten-words-api/api"
	"github.com/juddbaguio/top-ten-words-api/service"
)

func main() {
	topTenWordsSrv := service.InitTopTenWordsService()
	srv := api.InitServer(topTenWordsSrv)
	if err := srv.Start(); err != nil {
		log.Println("server error: %w", err.Error())
		os.Exit(1)
	}
}
