package main

import (
	"context"
	"log"
	"main/core"
	"main/ears"
	"main/storage"
	"os"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"gopkg.in/robfig/cron.v2"
)

var allEars = map[string]core.Ear{
	"NYTimes": &ears.NYTimesEar{},
	"Crypto": &ears.CryptoEar{
		APIKey: os.Getenv("ALPHAVANTAGE_API_KEY"),
	},
}

var store storage.Store

func makeRecord() {
	log.Println("Making new record")
	now := time.Now().UTC()
	record := core.Record{
		Date:     now,
		Executed: now,
	}
	for name, ear := range allEars {
		log.Printf("Fetching %s", name)
		err := ear.FetchAndPopulate(&record)
		if err != nil {
			log.Println("ERROR", err)
		}
	}
	log.Println("Saving record")
	err := store.SaveNewRecord(&record)
	if err != nil {
		log.Println("ERROR", err)
	}
}

func main() {
	pool, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}
	store = storage.Store{
		Pool: pool,
	}

	c := cron.New()
	c.AddFunc("@daily", makeRecord)
	c.Start()

	for {
	}
}
