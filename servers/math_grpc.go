package servers

import (
	"context"

	"github.com/FelipeAz/gRPC-example/pb/pb"
)

type Math struct {

}

func (m *Math) Sum(ctx context.Context, req *pb.NewSumRequest) (*pb.NewSumResponse, error) {
	res := req.Sum.A + req.Sum.B
	return &pb.NewSumResponse{
		Result: res,
	}, nil
}