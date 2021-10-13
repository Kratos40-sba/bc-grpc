package main

import (
	"context"
	"github.com/Kratos40-sba/bc-grpc/chain"
	pb "github.com/Kratos40-sba/bc-grpc/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type Server struct {
	chain *chain.BlockChain
}

func (s Server) AddBlock(ctx context.Context, request *pb.BlockRequest) (*pb.BlockResponse, error) {
	b := s.chain.AppendBlocks(request.Data)
	return &pb.BlockResponse{
		Hash: b.Hash,
	}, nil
}

func (s Server) StreamGetBlocks(request *pb.ChainRequest, server pb.BlockChain_StreamGetBlocksServer) error {
	for _, block := range s.chain.Blocks {
		res := &pb.ChainStreamResponse{Block: &pb.Block{
			PrevHash: block.PrevHash,
			Data:     block.Data,
			Hash:     block.Hash,
		}}
		err := server.Send(res)
		if err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
	}
	return nil
}

func (s Server) GetChain(ctx context.Context, request *pb.ChainRequest) (*pb.ChainResponse, error) {
	panic("not implemented")
}

func main() {
	listener, err := net.Listen("tcp", ":8084")
	if err != nil {
		log.Fatalf("Unable to listen on port : %s", err)
	}
	srv := grpc.NewServer()

	server := &Server{chain.MakeBlockChain()}
	pb.RegisterBlockChainServer(srv, server)
	err = srv.Serve(listener)
	if err != nil {
		return
	}
}
