package main

import (
	"context"
	"log"
	"net"
	"os"

	pb_svc_firebase "github.com/byeol-i/firebase-auth-module/pb/svc/firebase"
	pb_svc_stream "github.com/byeol-i/firebase-auth-module/pb/svc/stream"
	auth "github.com/byeol-i/firebase-auth-module/pkg/authentication/firebase"
	"github.com/byeol-i/firebase-auth-module/pkg/config"
	"github.com/byeol-i/firebase-auth-module/pkg/logger"
	server "github.com/byeol-i/firebase-auth-module/pkg/svc/firebase"
	streamServer "github.com/byeol-i/firebase-auth-module/pkg/svc/stream"

	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

func main() {
	if err := realMain(); err != nil {
		log.Printf("err :%s", err)
		os.Exit(1)
	}
}

func realMain() error {
	configManager := config.GetInstance()
	gRPCL, err := net.Listen("tcp", configManager.GrpcConfig.GetAuthAddr())
	if err != nil {
		return err
	}
	defer gRPCL.Close()

	firebaseApp, err := auth.NewFirebaseApp(configManager.FirebaseConfig.GetFirebaseCredFilePath(), configManager.FirebaseConfig.GetFirebaseProjectID())
	if err != nil {
		return err
	}

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)

	streamSrv := streamServer.NewStreamServiceServer(firebaseApp)
	authSrv := server.NewAuthServiceServer(firebaseApp)

	pb_svc_firebase.RegisterFirebaseServer(grpcServer, authSrv)
	pb_svc_stream.RegisterStreamServer(grpcServer, streamSrv)

	wg, _ := errgroup.WithContext(context.Background())

	wg.Go(func() error {
		logger.Info("Starting grpc server..." + configManager.GrpcConfig.GetAuthAddr())
		err := grpcServer.Serve(gRPCL)
		if err != nil {
			log.Fatalf("failed to serve: %v", err)
			return err
		}

		return nil
	})

	return wg.Wait()
}
