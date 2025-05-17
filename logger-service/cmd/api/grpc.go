package main

import (
	"context"
	"fmt"
	"log"
	"log-service/data"
	"log-service/logs"
	"net"

	"google.golang.org/grpc"
)

type LogServer struct {
	logs.UnimplementedLogServiceServer
	Models data.Models
}

func (l *LogServer) WriteLog(ctx context.Context, req *logs.LogRequest) (*logs.LogResponse, error) {
	input := req.GetLogEntry()

	log := data.LogEntry{
		Name: input.Name,
		Data: input.Data,
	}

	err := l.Models.LogEntry.Insert(log)
	if err != nil {
		res := &logs.LogResponse{
			Result: "failed to insert log",
		}
		return res, err

	}

	res := &logs.LogResponse{
		Result: "logged",
	}

	return res, nil
}

func (app *Config) grpcListen() {
	listen, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", grpcPort))
	if err != nil {
		log.Fatalf("Failed to listen for grpc: %v", err)
	}
	defer listen.Close()

	grpcServer := grpc.NewServer()
	logs.RegisterLogServiceServer(grpcServer, &LogServer{Models: app.Models})

	log.Printf("Starting grpc server on port %s", grpcPort)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to start grpc server: %v", err)
	}
}
