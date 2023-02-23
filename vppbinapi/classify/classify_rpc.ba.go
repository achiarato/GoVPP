// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package classify

import (
	"context"
	"fmt"
	"io"

	api "git.fd.io/govpp.git/api"
	memclnt "github.com/achiarato/GoVPP/vppbinapi/memclnt"
)

// RPCService defines RPC service classify.
type RPCService interface {
	ClassifyAddDelSession(ctx context.Context, in *ClassifyAddDelSession) (*ClassifyAddDelSessionReply, error)
	ClassifyAddDelTable(ctx context.Context, in *ClassifyAddDelTable) (*ClassifyAddDelTableReply, error)
	ClassifyPcapGetTables(ctx context.Context, in *ClassifyPcapGetTables) (*ClassifyPcapGetTablesReply, error)
	ClassifyPcapLookupTable(ctx context.Context, in *ClassifyPcapLookupTable) (*ClassifyPcapLookupTableReply, error)
	ClassifyPcapSetTable(ctx context.Context, in *ClassifyPcapSetTable) (*ClassifyPcapSetTableReply, error)
	ClassifySessionDump(ctx context.Context, in *ClassifySessionDump) (RPCService_ClassifySessionDumpClient, error)
	ClassifySetInterfaceIPTable(ctx context.Context, in *ClassifySetInterfaceIPTable) (*ClassifySetInterfaceIPTableReply, error)
	ClassifySetInterfaceL2Tables(ctx context.Context, in *ClassifySetInterfaceL2Tables) (*ClassifySetInterfaceL2TablesReply, error)
	ClassifyTableByInterface(ctx context.Context, in *ClassifyTableByInterface) (*ClassifyTableByInterfaceReply, error)
	ClassifyTableIds(ctx context.Context, in *ClassifyTableIds) (*ClassifyTableIdsReply, error)
	ClassifyTableInfo(ctx context.Context, in *ClassifyTableInfo) (*ClassifyTableInfoReply, error)
	ClassifyTraceGetTables(ctx context.Context, in *ClassifyTraceGetTables) (*ClassifyTraceGetTablesReply, error)
	ClassifyTraceLookupTable(ctx context.Context, in *ClassifyTraceLookupTable) (*ClassifyTraceLookupTableReply, error)
	ClassifyTraceSetTable(ctx context.Context, in *ClassifyTraceSetTable) (*ClassifyTraceSetTableReply, error)
	FlowClassifyDump(ctx context.Context, in *FlowClassifyDump) (RPCService_FlowClassifyDumpClient, error)
	FlowClassifySetInterface(ctx context.Context, in *FlowClassifySetInterface) (*FlowClassifySetInterfaceReply, error)
	InputACLSetInterface(ctx context.Context, in *InputACLSetInterface) (*InputACLSetInterfaceReply, error)
	OutputACLSetInterface(ctx context.Context, in *OutputACLSetInterface) (*OutputACLSetInterfaceReply, error)
	PolicerClassifyDump(ctx context.Context, in *PolicerClassifyDump) (RPCService_PolicerClassifyDumpClient, error)
	PolicerClassifySetInterface(ctx context.Context, in *PolicerClassifySetInterface) (*PolicerClassifySetInterfaceReply, error)
	PuntACLAddDel(ctx context.Context, in *PuntACLAddDel) (*PuntACLAddDelReply, error)
	PuntACLGet(ctx context.Context, in *PuntACLGet) (*PuntACLGetReply, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) ClassifyAddDelSession(ctx context.Context, in *ClassifyAddDelSession) (*ClassifyAddDelSessionReply, error) {
	out := new(ClassifyAddDelSessionReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) ClassifyAddDelTable(ctx context.Context, in *ClassifyAddDelTable) (*ClassifyAddDelTableReply, error) {
	out := new(ClassifyAddDelTableReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) ClassifyPcapGetTables(ctx context.Context, in *ClassifyPcapGetTables) (*ClassifyPcapGetTablesReply, error) {
	out := new(ClassifyPcapGetTablesReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) ClassifyPcapLookupTable(ctx context.Context, in *ClassifyPcapLookupTable) (*ClassifyPcapLookupTableReply, error) {
	out := new(ClassifyPcapLookupTableReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) ClassifyPcapSetTable(ctx context.Context, in *ClassifyPcapSetTable) (*ClassifyPcapSetTableReply, error) {
	out := new(ClassifyPcapSetTableReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) ClassifySessionDump(ctx context.Context, in *ClassifySessionDump) (RPCService_ClassifySessionDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_ClassifySessionDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_ClassifySessionDumpClient interface {
	Recv() (*ClassifySessionDetails, error)
	api.Stream
}

type serviceClient_ClassifySessionDumpClient struct {
	api.Stream
}

func (c *serviceClient_ClassifySessionDumpClient) Recv() (*ClassifySessionDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *ClassifySessionDetails:
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

func (c *serviceClient) ClassifySetInterfaceIPTable(ctx context.Context, in *ClassifySetInterfaceIPTable) (*ClassifySetInterfaceIPTableReply, error) {
	out := new(ClassifySetInterfaceIPTableReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) ClassifySetInterfaceL2Tables(ctx context.Context, in *ClassifySetInterfaceL2Tables) (*ClassifySetInterfaceL2TablesReply, error) {
	out := new(ClassifySetInterfaceL2TablesReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) ClassifyTableByInterface(ctx context.Context, in *ClassifyTableByInterface) (*ClassifyTableByInterfaceReply, error) {
	out := new(ClassifyTableByInterfaceReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) ClassifyTableIds(ctx context.Context, in *ClassifyTableIds) (*ClassifyTableIdsReply, error) {
	out := new(ClassifyTableIdsReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) ClassifyTableInfo(ctx context.Context, in *ClassifyTableInfo) (*ClassifyTableInfoReply, error) {
	out := new(ClassifyTableInfoReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) ClassifyTraceGetTables(ctx context.Context, in *ClassifyTraceGetTables) (*ClassifyTraceGetTablesReply, error) {
	out := new(ClassifyTraceGetTablesReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) ClassifyTraceLookupTable(ctx context.Context, in *ClassifyTraceLookupTable) (*ClassifyTraceLookupTableReply, error) {
	out := new(ClassifyTraceLookupTableReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) ClassifyTraceSetTable(ctx context.Context, in *ClassifyTraceSetTable) (*ClassifyTraceSetTableReply, error) {
	out := new(ClassifyTraceSetTableReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) FlowClassifyDump(ctx context.Context, in *FlowClassifyDump) (RPCService_FlowClassifyDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_FlowClassifyDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_FlowClassifyDumpClient interface {
	Recv() (*FlowClassifyDetails, error)
	api.Stream
}

type serviceClient_FlowClassifyDumpClient struct {
	api.Stream
}

func (c *serviceClient_FlowClassifyDumpClient) Recv() (*FlowClassifyDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *FlowClassifyDetails:
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

func (c *serviceClient) FlowClassifySetInterface(ctx context.Context, in *FlowClassifySetInterface) (*FlowClassifySetInterfaceReply, error) {
	out := new(FlowClassifySetInterfaceReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) InputACLSetInterface(ctx context.Context, in *InputACLSetInterface) (*InputACLSetInterfaceReply, error) {
	out := new(InputACLSetInterfaceReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) OutputACLSetInterface(ctx context.Context, in *OutputACLSetInterface) (*OutputACLSetInterfaceReply, error) {
	out := new(OutputACLSetInterfaceReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) PolicerClassifyDump(ctx context.Context, in *PolicerClassifyDump) (RPCService_PolicerClassifyDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_PolicerClassifyDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_PolicerClassifyDumpClient interface {
	Recv() (*PolicerClassifyDetails, error)
	api.Stream
}

type serviceClient_PolicerClassifyDumpClient struct {
	api.Stream
}

func (c *serviceClient_PolicerClassifyDumpClient) Recv() (*PolicerClassifyDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *PolicerClassifyDetails:
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

func (c *serviceClient) PolicerClassifySetInterface(ctx context.Context, in *PolicerClassifySetInterface) (*PolicerClassifySetInterfaceReply, error) {
	out := new(PolicerClassifySetInterfaceReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) PuntACLAddDel(ctx context.Context, in *PuntACLAddDel) (*PuntACLAddDelReply, error) {
	out := new(PuntACLAddDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) PuntACLGet(ctx context.Context, in *PuntACLGet) (*PuntACLGetReply, error) {
	out := new(PuntACLGetReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}
