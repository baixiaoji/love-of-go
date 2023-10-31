package main

import (
	_ "bookstore/internal/store"
	"bookstore/server"
	"bookstore/store/factory"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	s, err := factory.New("mem") // 创建图书数据存储实例
	if err != nil {
		panic(err)
	}

	srv := server.NewBookStoreServer(":8080", s) // 创建http服务

	errChan, err := srv.ListenAndServe() // 运行 http
	if err != nil {
		log.Println("Failed to start server", err)
		return
	}
	log.Println("Server is running on :8080")

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err = <-errChan:
		log.Println("web server failed: ", err)
		return
	case <-c:
		log.Println("Server is exiting...")
		ctx, cf := context.WithTimeout(context.Background(), time.Second)
		defer cf()
		err = srv.Shutdown(ctx) // 关闭http 服务
	}

	if err != nil {
		log.Println("bookstore server failed:", err)
		return
	}

	log.Println("bookstore server exited normally")
}
