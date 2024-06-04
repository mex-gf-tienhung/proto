package grpc_client

import (
	"sync"

	"github.com/mex-gf-tienhung/proto/golang/staff"
	"google.golang.org/grpc"
)

var (
	_staffClient        *staffClientStruct
	loadstaffClientOnce sync.Once
)

func ConnectTostaffServer(addr string, options ...grpc.DialOption) error {
	var err error
	loadstaffClientOnce.Do(func() {
		_staffClient = new(staffClientStruct)
		err = _staffClient.Connect(addr, options...)
	})

	return err
}

func StaffClient() *staffClientStruct {
	if _staffClient == nil {
		panic("grpc staff client: like client is not initiated")
	}

	return _staffClient
}

type staffClientStruct struct {
	staff.StaffServiceClient
	clientConn *grpc.ClientConn
}

func (c *staffClientStruct) Connect(addr string, options ...grpc.DialOption) error {
	authConn, err := grpc.Dial(addr, options...)
	if err != nil {
		return err
	}

	c.StaffServiceClient = staff.NewStaffServiceClient(authConn)
	c.clientConn = new(grpc.ClientConn)
	c.clientConn = authConn
	return nil
}

func (c *staffClientStruct) Close() {
	if c.clientConn == nil {
		return
	}

	defer c.clientConn.Close()
}
