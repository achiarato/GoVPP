// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package maps

import (
	"context"
	"fmt"
	"io"

	api "git.fd.io/govpp.git/api"
	memclnt "github.com/achiarato/GoVPP/vppbinapi/memclnt"
)

// RPCService defines RPC service map.
type RPCService interface {
	MapAddDelRule(ctx context.Context, in *MapAddDelRule) (*MapAddDelRuleReply, error)
	MapAddDomain(ctx context.Context, in *MapAddDomain) (*MapAddDomainReply, error)
	MapDelDomain(ctx context.Context, in *MapDelDomain) (*MapDelDomainReply, error)
	MapDomainDump(ctx context.Context, in *MapDomainDump) (RPCService_MapDomainDumpClient, error)
	MapDomainsGet(ctx context.Context, in *MapDomainsGet) (RPCService_MapDomainsGetClient, error)
	MapIfEnableDisable(ctx context.Context, in *MapIfEnableDisable) (*MapIfEnableDisableReply, error)
	MapParamAddDelPreResolve(ctx context.Context, in *MapParamAddDelPreResolve) (*MapParamAddDelPreResolveReply, error)
	MapParamGet(ctx context.Context, in *MapParamGet) (*MapParamGetReply, error)
	MapParamSetFragmentation(ctx context.Context, in *MapParamSetFragmentation) (*MapParamSetFragmentationReply, error)
	MapParamSetICMP(ctx context.Context, in *MapParamSetICMP) (*MapParamSetICMPReply, error)
	MapParamSetICMP6(ctx context.Context, in *MapParamSetICMP6) (*MapParamSetICMP6Reply, error)
	MapParamSetSecurityCheck(ctx context.Context, in *MapParamSetSecurityCheck) (*MapParamSetSecurityCheckReply, error)
	MapParamSetTCP(ctx context.Context, in *MapParamSetTCP) (*MapParamSetTCPReply, error)
	MapParamSetTrafficClass(ctx context.Context, in *MapParamSetTrafficClass) (*MapParamSetTrafficClassReply, error)
	MapRuleDump(ctx context.Context, in *MapRuleDump) (RPCService_MapRuleDumpClient, error)
	MapSummaryStats(ctx context.Context, in *MapSummaryStats) (*MapSummaryStatsReply, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) MapAddDelRule(ctx context.Context, in *MapAddDelRule) (*MapAddDelRuleReply, error) {
	out := new(MapAddDelRuleReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) MapAddDomain(ctx context.Context, in *MapAddDomain) (*MapAddDomainReply, error) {
	out := new(MapAddDomainReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) MapDelDomain(ctx context.Context, in *MapDelDomain) (*MapDelDomainReply, error) {
	out := new(MapDelDomainReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) MapDomainDump(ctx context.Context, in *MapDomainDump) (RPCService_MapDomainDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_MapDomainDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_MapDomainDumpClient interface {
	Recv() (*MapDomainDetails, error)
	api.Stream
}

type serviceClient_MapDomainDumpClient struct {
	api.Stream
}

func (c *serviceClient_MapDomainDumpClient) Recv() (*MapDomainDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *MapDomainDetails:
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

func (c *serviceClient) MapDomainsGet(ctx context.Context, in *MapDomainsGet) (RPCService_MapDomainsGetClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_MapDomainsGetClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_MapDomainsGetClient interface {
	Recv() (*MapDomainDetails, error)
	api.Stream
}

type serviceClient_MapDomainsGetClient struct {
	api.Stream
}

func (c *serviceClient_MapDomainsGetClient) Recv() (*MapDomainDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *MapDomainDetails:
		return m, nil
	case *MapDomainsGetReply:
		err = c.Stream.Close()
		if err != nil {
			return nil, err
		}
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) MapIfEnableDisable(ctx context.Context, in *MapIfEnableDisable) (*MapIfEnableDisableReply, error) {
	out := new(MapIfEnableDisableReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) MapParamAddDelPreResolve(ctx context.Context, in *MapParamAddDelPreResolve) (*MapParamAddDelPreResolveReply, error) {
	out := new(MapParamAddDelPreResolveReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) MapParamGet(ctx context.Context, in *MapParamGet) (*MapParamGetReply, error) {
	out := new(MapParamGetReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) MapParamSetFragmentation(ctx context.Context, in *MapParamSetFragmentation) (*MapParamSetFragmentationReply, error) {
	out := new(MapParamSetFragmentationReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) MapParamSetICMP(ctx context.Context, in *MapParamSetICMP) (*MapParamSetICMPReply, error) {
	out := new(MapParamSetICMPReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) MapParamSetICMP6(ctx context.Context, in *MapParamSetICMP6) (*MapParamSetICMP6Reply, error) {
	out := new(MapParamSetICMP6Reply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) MapParamSetSecurityCheck(ctx context.Context, in *MapParamSetSecurityCheck) (*MapParamSetSecurityCheckReply, error) {
	out := new(MapParamSetSecurityCheckReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) MapParamSetTCP(ctx context.Context, in *MapParamSetTCP) (*MapParamSetTCPReply, error) {
	out := new(MapParamSetTCPReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) MapParamSetTrafficClass(ctx context.Context, in *MapParamSetTrafficClass) (*MapParamSetTrafficClassReply, error) {
	out := new(MapParamSetTrafficClassReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) MapRuleDump(ctx context.Context, in *MapRuleDump) (RPCService_MapRuleDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_MapRuleDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_MapRuleDumpClient interface {
	Recv() (*MapRuleDetails, error)
	api.Stream
}

type serviceClient_MapRuleDumpClient struct {
	api.Stream
}

func (c *serviceClient_MapRuleDumpClient) Recv() (*MapRuleDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *MapRuleDetails:
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

func (c *serviceClient) MapSummaryStats(ctx context.Context, in *MapSummaryStats) (*MapSummaryStatsReply, error) {
	out := new(MapSummaryStatsReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}
