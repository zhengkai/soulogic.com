package main

import (
	"soulogic/db"
	"soulogic/pb"
)

func (gw *gateway) eachPublic(req *pb.ReqPublic) (r *pb.RspPublic, rspErr *pb.RspError) {

	rsp := &pb.RspPublic{}

	switch one := req.One.(type) {
	case *pb.ReqPublic_BlogIndex:
		o := &pb.RspPublic_BlogIndex{}
		o.BlogIndex, rspErr = gw.publicBlogIndex(one.BlogIndex)
		if rspErr != nil {
			break
		}
		rsp.One = o
	}

	r = rsp

	return
}

func (gw *gateway) publicBlogIndex(_ bool) (re *pb.ItemList, rspErr *pb.RspError) {

	re = &pb.ItemList{
		List: db.ItemRecent(),
	}

	return
}
