package main

import (
	"log"
	"net/http"
	"time"

	_ "gin_demo/docs/v1"
	_ "gin_demo/docs/v2"

	"golang.org/x/sync/errgroup"

	routers "gin_demo/routers"
)

func main() {
	server01 := &http.Server{
		Addr:         ":9090",
		Handler:      routers.Router01(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server02 := &http.Server{
		Addr:         ":9091",
		Handler:      routers.Router02(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	// 开启这两个服务,
	// 可以使用errgroup.Group或者开启两个goroutine
	var g errgroup.Group

	g.Go(func() error {
		return server01.ListenAndServe()
	})
	g.Go(func() error {
		return server02.ListenAndServe()
	})
	
	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
