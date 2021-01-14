package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/urfave/cli/v2"
)

// Version is overwritten at build time.
var version = "0.0.0"

const doesNotExistStatus int = 000

func main() {
	app := &cli.App{
		Name:    "Probe",
		Version: version,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "output",
				Usage: "Where do you want to results written to? If left blank they will only be printed.",
			},
			&cli.StringFlag{
				Name:  "user-agent",
				Usage: "Custom user agent. Default 'probe/{version}'.",
				Value: fmt.Sprintf("probe/%s", version),
			},
			&cli.BoolFlag{
				Name:  "filtered",
				Usage: "When turned off all domains checked will be in the output/logs. If turned on only the domains that return with a successful status.",
				Value: true,
			},
			&cli.BoolFlag{
				Name:  "verbose",
				Usage: "Turns on verbose looking.",
			},
		},
		Action: func(c *cli.Context) error {
			run(c.String("output"), c.String("user-agent"), c.Bool("filtered"), c.Bool("verbose"))
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func run(output string, userAgent string, filtered, verbose bool) {
	started := time.Now()
	statuses := []status{}
	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		status, err := checkStatus(httpClient, userAgent, scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		statuses = append(statuses, status)
	}
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go stdout(wg, filtered, verbose, statuses)
	if len(output) > 0 {
		wg.Add(1)
		go writeFile(wg, output, statuses)

	}
	wg.Wait()
	if verbose {
		fmt.Printf("Completed in: %f2 secs\n", time.Since(started).Seconds())
	}
}

type status struct {
	domain string
	status int
}
