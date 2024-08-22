package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Arjun-Ingole/cache-proxy/internal/proxy"
)

func main() {
	// Parse Arguments
	PORT := flag.Int("port", 0, "Define PORT on which the caching proxy server will run")
	ORIGIN := flag.String("origin", "", "Define the URL of the server to which the requests will be forwarded")
	CLEAR_CACHE := flag.Bool("clear-cache", false, "Clear the Cache")
	flag.Parse()

	proxy := proxy.NewProxy("http://example.com")

	// Validate Arguments
	if *CLEAR_CACHE {
		fmt.Println("Clearing Cache....")
		proxy.ClearCache()
		os.Exit(0)
	}

	if *ORIGIN != "" || *PORT != 0 {
		if *ORIGIN == "" {
			log.Fatal("Origin server URL is required when starting the server")
		}
		proxy.Origin = *ORIGIN
		http.Handle("/", proxy)
		log.Printf("Starting proxy server on port %d", *PORT)
		log.Printf("Forwarding requests to %s", *ORIGIN)
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *PORT), nil))
	} else {
		fmt.Println("No action specified.\nUse --clear-cache to clear the cache\nProvide --port and --origin to start the server")
		flag.Usage()
	}
}
