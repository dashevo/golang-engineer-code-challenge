package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/shotonoff/golang-engineer-code-challenge/internal/app/config"
	"github.com/shotonoff/golang-engineer-code-challenge/internal/app/metric"
	"github.com/shotonoff/golang-engineer-code-challenge/internal/app/network"
	"github.com/shotonoff/golang-engineer-code-challenge/internal/app/usecase"
	"github.com/shotonoff/golang-engineer-code-challenge/internal/app/util"
)

func main() {
	conf, err := config.InitFromEnv()
	check(err)

	inputData, err := util.LoadSampleData(conf.TestSampleFile)
	check(err)

	metrics := metric.NewInMemory(nil)

	p2pClient, err := network.NewHTTPClient(metrics, network.P2PNetwork)
	check(err)

	hostedClient, err := network.NewHTTPClient(metrics, network.HostedNetwork)
	check(err)

	srv := usecase.NewService(p2pClient, hostedClient)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	err = srv.Store(ctx, inputData)
	check(err)

	outputData, err := srv.Fetch(ctx)
	check(err)

	if !util.Compare(inputData, outputData) {
		log.Fatalf("fetched data is not equal expected")
	}

	// aggregate stored metrics
	aggr := metrics.Aggregator()
	stats, err := aggr.SummaryStats()
	check(err)

	// write in output summary statistics
	err = metric.RenderSummaryStats(os.Stdout, stats)
	check(err)
}

func check(err error) {
	if err != nil {
		log.Fatalf(err.Error())
	}
}
