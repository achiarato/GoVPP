// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package vhost_user

import (
	"context"
	"fmt"
	"io"

	api "git.fd.io/govpp.git/api"
	memclnt "github.com/achiarato/GoVPP/vppbinapi/memclnt"
)

// RPCService defines RPC service vhost_user.
type RPCService interface {
	CreateVhostUserIf(ctx context.Context, in *CreateVhostUserIf) (*CreateVhostUserIfReply, error)
	CreateVhostUserIfV2(ctx context.Context, in *CreateVhostUserIfV2) (*CreateVhostUserIfV2Reply, error)
	DeleteVhostUserIf(ctx context.Context, in *DeleteVhostUserIf) (*DeleteVhostUserIfReply, error)
	ModifyVhostUserIf(ctx context.Context, in *ModifyVhostUserIf) (*ModifyVhostUserIfReply, error)
	ModifyVhostUserIfV2(ctx context.Context, in *ModifyVhostUserIfV2) (*ModifyVhostUserIfV2Reply, error)
	SwInterfaceVhostUserDump(ctx context.Context, in *SwInterfaceVhostUserDump) (RPCService_SwInterfaceVhostUserDumpClient, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) CreateVhostUserIf(ctx context.Context, in *CreateVhostUserIf) (*CreateVhostUserIfReply, error) {
	out := new(CreateVhostUserIfReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) CreateVhostUserIfV2(ctx context.Context, in *CreateVhostUserIfV2) (*CreateVhostUserIfV2Reply, error) {
	out := new(CreateVhostUserIfV2Reply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) DeleteVhostUserIf(ctx context.Context, in *DeleteVhostUserIf) (*DeleteVhostUserIfReply, error) {
	out := new(DeleteVhostUserIfReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) ModifyVhostUserIf(ctx context.Context, in *ModifyVhostUserIf) (*ModifyVhostUserIfReply, error) {
	out := new(ModifyVhostUserIfReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) ModifyVhostUserIfV2(ctx context.Context, in *ModifyVhostUserIfV2) (*ModifyVhostUserIfV2Reply, error) {
	out := new(ModifyVhostUserIfV2Reply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) SwInterfaceVhostUserDump(ctx context.Context, in *SwInterfaceVhostUserDump) (RPCService_SwInterfaceVhostUserDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_SwInterfaceVhostUserDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_SwInterfaceVhostUserDumpClient interface {
	Recv() (*SwInterfaceVhostUserDetails, error)
	api.Stream
}

type serviceClient_SwInterfaceVhostUserDumpClient struct {
	api.Stream
}

func (c *serviceClient_SwInterfaceVhostUserDumpClient) Recv() (*SwInterfaceVhostUserDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *SwInterfaceVhostUserDetails:
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
