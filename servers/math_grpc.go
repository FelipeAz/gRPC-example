package servers

import (
	"context"
	"time"

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

func (m *Math) Fibonacci(in *pb.FibonacciRequest, stream pb.MathService_FibonacciServer) error {
	n := in.GetN()
	var i int32
	for i = 1; i <= n; i++ {
		res := &pb.FibonacciResponse{
			Result: FibonacciResolver(i),
		}
		err := stream.Send(res)
		if err != nil {
			return err
		}
		time.Sleep(time.Second * 2)
	}
	return nil
}

func FibonacciResolver(n int32) int32 {
	if n <= 1 {
		return n
	}
	return FibonacciResolver(n-1) + FibonacciResolver(n-2)
}
