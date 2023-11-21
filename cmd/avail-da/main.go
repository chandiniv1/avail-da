package main

import (
	"context"
	"errors"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/rollkit/avail-da"
	"github.com/rollkit/go-da/proxy"
)

func main() {
	var appID uint32
	appID = 1
	ctx := context.Background()
	da := avail.NewAvailDA(appID, ctx)
	srv := proxy.NewServer(da, grpc.Creds(insecure.NewCredentials()))
	lis, err := net.Listen("tcp", "")
	if err != nil {
		log.Fatalln("failed to create network listener:", err)
	}
	log.Println("serving avail-da over gRPC on:", lis.Addr())
	err = srv.Serve(lis)
	if !errors.Is(err, grpc.ErrServerStopped) {
		log.Fatalln("gRPC server stopped with error:", err)
	}
}
