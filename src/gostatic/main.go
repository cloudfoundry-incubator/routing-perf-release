package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/valyala/fasthttp"
)

var data string

func main() {
	s := os.Getenv("RESPONSE_SIZE")
	if s == "" {
		s = "1"
	}
	responseSize, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Error parsing response size: %s", err)
	}
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalf("The PORT environment variable is empty")
	}
	data = strings.Repeat("Z", responseSize*1024)
	log.Fatal(fasthttp.ListenAndServe(":"+os.Getenv("PORT"), index))
}

func index(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Path()) {
	case "/slow":
		slowHandler(ctx)
	default:
		defaultHandler(ctx)
	}
}

func slowHandler(ctx *fasthttp.RequestCtx) {
	time.Sleep(5 * time.Second)
	fmt.Fprintf(ctx, data)
}

func defaultHandler(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, data)
}
