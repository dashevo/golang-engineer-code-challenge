package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"text/tabwriter"
	"time"

	"github.com/shotonoff/golang-engineer-code-challenge/internal/config"
	"github.com/shotonoff/golang-engineer-code-challenge/internal/httpclient"
	"github.com/shotonoff/golang-engineer-code-challenge/internal/metric"
	"github.com/shotonoff/golang-engineer-code-challenge/internal/usecase"
	"github.com/shotonoff/golang-engineer-code-challenge/internal/util"
)

func main() {
	conf, err := config.InitFromEnv()
	check(err)

	inputData, err := util.LoadSampleData(conf.TestSampleFile)
	check(err)

	metrics := metric.NewInMemory(nil)

	client := httpclient.New()
	httpclient.WithHeaders(client, config.DefaultHTTPHeaders)
	httpclient.WithMetricsMiddleware(client, conf.Network, metrics)

	srv := usecase.NewService(client, conf.FetchURL, conf.StoreURL)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	err = srv.Store(ctx, inputData)
	check(err)

	outputData, err := srv.Fetch(ctx)
	check(err)

	if !util.Compare(inputData, outputData) {
		log.Fatalf("fetched data is not equal expected")
	}

	stats := &metric.SummaryStats{Request: make(map[string]*metric.RequestStats)}
	err = metrics.Reduce(metric.SummaryStatsReduce(), stats)
	check(err)

	render(os.Stdout, conf.Network, stats)
	check(err)
}

func check(err error) {
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func render(w io.Writer, network string, stats *metric.SummaryStats) {
	_, _ = fmt.Fprintf(w, "Your total expenses: %.4f DASH\n\n", stats.TotalCost)
	tw := tabwriter.NewWriter(w, 0, 8, 1, '\t', tabwriter.AlignRight)
	_, _ = fmt.Fprintf(w, "Requests to the %s service\n", network)
	_, _ = fmt.Fprintf(tw, "Method\tUrl\tSize/b\tElapsed/ms\n")
	for _, req := range stats.Request {
		_, _ = fmt.Fprintf(tw, "%s\t%s\t%d\t%d\n", req.Method, req.URL, req.Size, time.Duration(req.Elapsed).Milliseconds())
	}
	tw.Flush()
}
