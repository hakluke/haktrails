package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sync"

	"gopkg.in/yaml.v2"
)

//Config struct holds all configuration data that comes from config.yml or environment variables
type Config struct {
	SecurityTrails struct {
		Key string `yaml:"key" envconfig:"SECURITYTRAILS_KEY"`
	} `yaml:"securitytrails"`
}

var config *Config
var apiEndpoint string

func main() {
	defaultConfigFile := os.Getenv("HOME") + "/.config/haktools/haktrails-config.yml"
	// parse the command line options
	concurrencyPtr := flag.Int("t", 4, "Number of threads to utilise. Keep in mind that the API has rate limits.")
	configFile := flag.String("c", defaultConfigFile, "Config file location")
	flag.Parse()

	// load config file
	f, err := os.Open(*configFile)
	if err != nil {
		fmt.Println("Error opening config file:", err)
		return
	}
	defer f.Close()

	// parse the config file
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println("Error decoding config file:", err)
		return
	}

	apiEndpoint = "https://api.securitytrails.com/v1/" // global

	numWorkers := *concurrencyPtr
	work := make(chan string)
	go func() {
		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {
			work <- s.Text()
		}
		close(work)
	}()

	wg := &sync.WaitGroup{}

	switch os.Args[1] {
	case "subdomains":
		for i := 0; i < numWorkers; i++ {
			wg.Add(1)
			go subdomains(work, wg)
		}
		wg.Wait()
	case "associated":
		for i := 0; i < numWorkers; i++ {
			wg.Add(1)
			go associatedDomains(work, wg)
		}
		wg.Wait()
	// no valid subcommand found - default to showing a message and exiting
	default:
		fmt.Println("Subcommand missing or incorrect. Hint: haktrails {subdomains|associateddomains}")
		os.Exit(1)
	}
}
