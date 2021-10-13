package main

import (
	"context"
	"flag"
	pb "github.com/Kratos40-sba/bc-grpc/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

var client pb.BlockChainClient

func main() {
	start := flag.Bool("start", false, "Start sending concurrent blocks to the server ")
	stream := flag.Bool("stream", false, "Receive a stream of blocks")
	flag.Parse()
	conn, err := grpc.Dial(":8084", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Cannot dial server : %v", err)
	}
	client = pb.NewBlockChainClient(conn)
	if *start {
		startBlockChain()
	}
	if *stream {
		startStreaming()
	}
}
func startBlockChain() {
	for {
		block, err := client.AddBlock(context.Background(), &pb.BlockRequest{Data: time.Now().String()})
		if err != nil {
			log.Fatalf("Unable to add block : %v", err)
		}
		log.Printf("New block added -> %s\n", block.Hash)
		time.Sleep(1 * time.Second)
	}
}
func startStreaming() {
	req := &pb.ChainRequest{}
	stream, err := client.StreamGetBlocks(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling stream")
	}
	for {
		blockStream, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while reading from stream")
		}
		block := blockStream.Block
		log.Printf("Streaming => prevHash : %s | data : %s | hash : %s \n", block.PrevHash, block.Data, block.Hash)
	}

}
