// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package flowprobe

import (
	"context"
	"fmt"
	"io"

	api "git.fd.io/govpp.git/api"
	memclnt "github.com/achiarato/GoVPP/vppbinapi/memclnt"
)

// RPCService defines RPC service flowprobe.
type RPCService interface {
	FlowprobeGetParams(ctx context.Context, in *FlowprobeGetParams) (*FlowprobeGetParamsReply, error)
	FlowprobeInterfaceAddDel(ctx context.Context, in *FlowprobeInterfaceAddDel) (*FlowprobeInterfaceAddDelReply, error)
	FlowprobeInterfaceDump(ctx context.Context, in *FlowprobeInterfaceDump) (RPCService_FlowprobeInterfaceDumpClient, error)
	FlowprobeParams(ctx context.Context, in *FlowprobeParams) (*FlowprobeParamsReply, error)
	FlowprobeSetParams(ctx context.Context, in *FlowprobeSetParams) (*FlowprobeSetParamsReply, error)
	FlowprobeTxInterfaceAddDel(ctx context.Context, in *FlowprobeTxInterfaceAddDel) (*FlowprobeTxInterfaceAddDelReply, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) FlowprobeGetParams(ctx context.Context, in *FlowprobeGetParams) (*FlowprobeGetParamsReply, error) {
	out := new(FlowprobeGetParamsReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) FlowprobeInterfaceAddDel(ctx context.Context, in *FlowprobeInterfaceAddDel) (*FlowprobeInterfaceAddDelReply, error) {
	out := new(FlowprobeInterfaceAddDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) FlowprobeInterfaceDump(ctx context.Context, in *FlowprobeInterfaceDump) (RPCService_FlowprobeInterfaceDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_FlowprobeInterfaceDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_FlowprobeInterfaceDumpClient interface {
	Recv() (*FlowprobeInterfaceDetails, error)
	api.Stream
}

type serviceClient_FlowprobeInterfaceDumpClient struct {
	api.Stream
}

func (c *serviceClient_FlowprobeInterfaceDumpClient) Recv() (*FlowprobeInterfaceDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *FlowprobeInterfaceDetails:
		return m, nil
	case *memclnt.ControlPingReply:
		err = c.Stream.Close()
		if err != nil {
			return nil, err
		}
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) FlowprobeParams(ctx context.Context, in *FlowprobeParams) (*FlowprobeParamsReply, error) {
	out := new(FlowprobeParamsReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) FlowprobeSetParams(ctx context.Context, in *FlowprobeSetParams) (*FlowprobeSetParamsReply, error) {
	out := new(FlowprobeSetParamsReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) FlowprobeTxInterfaceAddDel(ctx context.Context, in *FlowprobeTxInterfaceAddDel) (*FlowprobeTxInterfaceAddDelReply, error) {
	out := new(FlowprobeTxInterfaceAddDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}
