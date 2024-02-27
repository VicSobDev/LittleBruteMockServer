package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/vicsobdev/LittleBruteMockServer/server"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	log.Println("Initializing Server")

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Couldn't load .env file: %v", err)
	}

	listenAddr := os.Getenv("LISTEN_ADDR")

	api := server.New(listenAddr)

	panic(api.Start())

}
