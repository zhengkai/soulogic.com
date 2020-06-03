package main

import "soulogic/pb"

func (gw *gateway) eachOp(req *pb.ReqOp) (r *pb.RspOp, rspErr *pb.RspError) {

	rsp := &pb.RspOp{}

	switch one := req.One.(type) {
	case *pb.ReqOp_Echo:
		o := &pb.RspOp_Echo{}
		o.Echo, rspErr = gw.opEcho(one.Echo)
		if rspErr != nil {
			break
		}
		rsp.One = o
	}

	r = rsp

	return
}

func (gw *gateway) opEcho(_ string) (re string, rspErr *pb.RspError) {
	return
}
