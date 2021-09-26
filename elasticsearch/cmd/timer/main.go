package main

import (
	"gitlab.com/Spide_IT/spide_it/elasticsearch/config"
	"gitlab.com/Spide_IT/spide_it/elasticsearch/internal/worker/server"
	"gitlab.com/Spide_IT/spide_it/elasticsearch/pkg/elasticsearch"
	"gopkg.in/robfig/cron.v2"
	"log"
	"time"
)

func main(){
	config, err := config.ReadConf("config/config-docker.yml")
	if err != nil {
		log.Fatal(err)
	}

	elasticsearchConn :=elasticsearch.NewGetElasticsearchConn(config.Elasticsearch.Dns)

	timer:=server.NewTimer(config,elasticsearchConn)

	c := cron.New()
	//Runs at 04:30 Bangkok time every day
	c.AddFunc("TZ=Asia/Bangkok 30 04 * * * *", timer.Run)
	c.Start()
	time.Sleep(10 * time.Second)
	c.Stop()
}