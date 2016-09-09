package pump

import (
	"fmt"
	pb "github.com/pingcap/tidb-binlog/proto"
	"github.com/ngaut/log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"
)

type server struct{}

func (s *server) WriteBinlog(ctx context.Context, in *pb.WriteBinlogReq) (*pb.WriteBinlogResp, error) {
	return nil, nil
}

func (s *server) PullBinlog(ctx context.Context, in *pb.PullBinlogReq) (*pb.PullBinlogResp, error) {
	return nil, nil
}

func Run(cfg *Config) {
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", cfg.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterPumpServer(s, &server{})
	s.Serve(lis)
}