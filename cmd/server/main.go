package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/ketsuna-org/sovrabase/internal/config"
	pb "github.com/ketsuna-org/sovrabase/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type server struct {
	pb.UnimplementedForwardCommandServiceServer
	nodeClients map[string]pb.ForwardCommandServiceClient
}

func (s *server) ForwardCommand(ctx context.Context, req *pb.ForwardCommandRequest) (*pb.ForwardCommandResponse, error) {
	log.Printf("Received command: %s for node: %s", req.Command, req.TargetNode)

	// Si c'est pour ce node, traiter localement
	if req.TargetNode == "" || req.TargetNode == "self" {
		return &pb.ForwardCommandResponse{
			Success: true,
			Result:  fmt.Sprintf("Command '%s' executed locally", req.Command),
		}, nil
	}

	// Sinon, forwarder au node cible
	if client, exists := s.nodeClients[req.TargetNode]; exists {
		resp, err := client.ForwardCommand(ctx, req)
		if err != nil {
			return &pb.ForwardCommandResponse{
				Success:      false,
				ErrorMessage: fmt.Sprintf("Failed to forward to %s: %v", req.TargetNode, err),
			}, nil
		}
		return resp, nil
	}

	// Node inconnu
	return &pb.ForwardCommandResponse{
		Success:      false,
		ErrorMessage: fmt.Sprintf("Unknown target node: %s", req.TargetNode),
	}, nil
}

func main() {
	cfg, err := config.LoadConfig("config.toml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// Initialiser les connexions aux autres nodes
	nodeClients := make(map[string]pb.ForwardCommandServiceClient)
	for _, addr := range cfg.Cluster.RPCServers {
		if addr != cfg.RPC.RPCAddr { // Ne pas se connecter à soi-même
			conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				log.Printf("Failed to connect to node %s: %v", addr, err)
				continue
			}
			nodeClients[addr] = pb.NewForwardCommandServiceClient(conn)
			log.Printf("Connected to node: %s", addr)
		}
	}

	lis, err := net.Listen("tcp", cfg.RPC.RPCAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterForwardCommandServiceServer(s, &server{nodeClients: nodeClients})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
