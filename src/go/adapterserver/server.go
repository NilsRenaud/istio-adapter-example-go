package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/testdata"
)

var (
	tls      = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile = flag.String("cert_file", "", "The TLS cert file")
	keyFile  = flag.String("key_file", "", "The TLS key file")
	port     = flag.Int("port", 9070, "The adapterserver port")
)

type myProtoServiceServer struct {
	UnimplementedHandleMyprotoServiceServer
}

// GetFeature returns the feature at the given point.
func (s *myProtoServiceServer) HandleMyproto(ctx context.Context, template *Template) (*OutputTemplate, error) {
	log.Print("Serving a request !!!!")
	return &OutputTemplate{User: []byte("Nils"), ValidUseCount: 1}, nil
}

func newServer() *myProtoServiceServer {
	s := &myProtoServiceServer{}
	return s
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	if *tls {
		if *certFile == "" {
			*certFile = testdata.Path("server1.pem")
		}
		if *keyFile == "" {
			*keyFile = testdata.Path("server1.key")
		}
		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		if err != nil {
			log.Fatalf("Failed to generate credentials %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}
	grpcServer := grpc.NewServer(opts...)
	RegisterHandleMyprotoServiceServer(grpcServer, newServer())
	log.Print("Begin to serve")
	grpcServer.Serve(lis)
	log.Print("Serving")
}
