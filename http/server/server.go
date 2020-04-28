package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var f = func() {
	log.Println("sleep a few seconds")
	time.Sleep(20 * time.Second)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", ping)
	mux.HandleFunc("/sleep", sleep)
	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	// When listeners closed, executes given functions
	srv.RegisterOnShutdown(f)

	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Println("failed to serve: ", err)
		}
	}()

	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGTERM, os.Interrupt)
	<-s
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Println("failed to shutdown gracefully: ", err)
	}
	log.Println("Server shutdown")
}

/*

## [net/http server.Shutdown](https://github.com/golang/go/blob/master/src/net/http/server.go#L2669)の実装

* シャットダウン処理中にステートを更新する
* Serverオブジェクトをロックする
* リスナーをクローズする
* ListenAndServe()にErrServerClosedを返却する
* 事前に登録されたシャットダウン時のコールされる関数の実行
* Serverオブジェクトをアンロックする
* 一定間隔でIdle Connectionのクローズ処理、およびコンテキストクローズの監視などの処理をループして待機

pingとsleepハンドラは独立しており、sleepハンドラで待機していてもpingハンドラは正常に処理される。
*/

func ping(w http.ResponseWriter, r *http.Request) {
	log.Println("Received Request: ping")
	w.Header().Set("Content-Type", "text/plain")
	_, _ = w.Write([]byte("ping\n"))
}

func sleep(w http.ResponseWriter, r *http.Request) {
	log.Println("Received Request: sleep")
	time.Sleep(10 * time.Second)
	w.Header().Set("Content-Type", "text/plain")
	_, _ = w.Write([]byte("wake up\n"))
}
