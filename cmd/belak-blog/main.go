package main

import (
	"flag"
	"log"
	"net/http"

	blog "github.com/belak/blog"
)

func main() {
	addr := flag.String("addr", ":8080", "listen address")
	flag.Parse()

	site, err := blog.LoadSite(blog.BlogFS, blog.PagesFS, blog.AssetsFS)
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	site.RegisterRoutes(mux)

	log.Printf("listening on %s", *addr)
	log.Fatal(http.ListenAndServe(*addr, mux))
}
