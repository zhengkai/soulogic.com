package main

import (
	"fmt"
	"soulogic/db"
	"soulogic/pb"
)

func (gw *gateway) eachOp(req *pb.ReqOp) (r *pb.RspOp, rspErr *pb.RspError) {

	rsp := &pb.RspOp{}

	switch one := req.One.(type) {

	case *pb.ReqOp_Echo:

		o := &pb.RspOp_Echo{}
		o.Echo, rspErr = gw.opEcho(one.Echo)
		if rspErr != nil {
			return
		}
		rsp.One = o

	case *pb.ReqOp_ItemEdit:

		o := &pb.RspOp_ItemEdit{}
		o.ItemEdit, rspErr = gw.opItemEdit(one.ItemEdit)
		if rspErr != nil {
			return
		}
		rsp.One = o

	default:

		rspErr = &pb.RspError{
			Code: 101,
			Msg:  fmt.Sprintf(`unknown type %T`, one),
		}
		return
	}

	r = rsp

	return
}

func (gw *gateway) opEcho(s string) (re string, rspErr *pb.RspError) {
	re = s
	return
}

func (gw *gateway) opItemEdit(ie *pb.ItemEdit) (id uint32, rspErr *pb.RspError) {

	var err error
	id, err = db.ItemSet(ie.ID, ie.Revision)
	if err != nil {
		rspErr = &pb.RspError{
			Code: 102,
			Msg:  err.Error(),
		}
		return
	}

	return
}
