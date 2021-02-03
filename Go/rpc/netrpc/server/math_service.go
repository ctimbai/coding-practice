// Package server the rpc server
package server

// MathService 远程服务对象
type MathService struct{}

// Args 待计算的参数
type Args struct {
	A, B int
}

// Add 计算两个数相加的服务
func (m *MathService) Add(args Args, reply *int) error {
	*reply = args.A + args.B
	return nil
}
