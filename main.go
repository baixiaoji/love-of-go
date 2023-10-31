package main

import (
	"github.com/google/uuid"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)
	

var logger *zap.Logger

func init() {
	logger, _ = zap.NewProduction()
}

func fastHTTPHandler(ctx *fasthttp.RequestCtx) {
	logger.Info("hello, go module", zap.ByteString("uri", ctx.RequestURI()))
}

func main() {
	logger.Info(uuid.NewString())
	fasthttp.ListenAndServe(":8081", fastHTTPHandler)
}