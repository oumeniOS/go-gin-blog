package main

import (
	"log"
	"github.com/robfig/cron"
	"github.com/oumeniOS/go-gin-blog/models"
	"time"
)

func main1() {
	log.Printf("starting...")

	c := cron.New()
	c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanAllTag...")
		models.ClearAllTag()
	})
	c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanAllArticle...")
		models.ClearAllArtical()
	})

	c.Start()

	t1 := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)
		}
	}
}
