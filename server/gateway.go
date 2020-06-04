package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"soulogic/pb"

	"google.golang.org/protobuf/proto"
)

func handleGateway(w http.ResponseWriter, r *http.Request) {

	limit := 100000 // POST 长度上限

	gw := &gateway{
		id:    1,
		limit: limit,
		w:     w,
		r:     r,
		rsp:   &pb.RspGateway{},
		req:   &pb.ReqGateway{},
	}

	gw.run()

	gw.output()
}

type gateway struct {
	id     int
	limit  int
	w      http.ResponseWriter
	r      *http.Request
	req    *pb.ReqGateway
	rsp    *pb.RspGateway
	reqLen int
	rspLen int
}

func (gw *gateway) run() {

	stop := gw.check()
	if stop {
		return
	}

	gw.dispatch()
}

func (gw *gateway) output() {

	gw.rsp.Ts = uint64(time.Now().UnixNano() / 1000000)

	b, err := proto.Marshal(gw.rsp)
	if err != nil {
		jw(`errMarshal`, err)
		return
	}
	gw.rspLen = len(b)
	gw.w.Write(b)
}

func (gw *gateway) check() (stop bool) {

	body, err := ioutil.ReadAll(io.LimitReader(gw.r.Body, int64(gw.limit)))
	if err != nil {
		return gw.error(1, err.Error())
	}

	gw.reqLen = len(body)
	if gw.reqLen == gw.limit {
		return gw.error(2, fmt.Sprintf(`body overflow (limit %d)`, gw.limit))
	}
	if gw.reqLen == 0 {
		return gw.error(3, `empty body`)
	}

	err = proto.Unmarshal(body, gw.req)
	if err != nil {
		return gw.error(4, err.Error())
	}

	return
}

func (gw *gateway) error(code uint32, msg string) (stop bool) {
	gw.rsp.Error = &pb.RspError{
		Code: code,
		Msg:  msg,
	}
	return true
}

func (gw *gateway) dispatch() {

	switch req := gw.req.One.(type) {

	case *pb.ReqGateway_Public:

		one := &pb.RspGateway_Public{}
		one.Public, gw.rsp.Error = gw.reqPublic(req.Public)
		if gw.rsp.Error != nil {
			return
		}
		gw.rsp.One = one

	case *pb.ReqGateway_Op:

		one := &pb.RspGateway_Op{}
		one.Op, gw.rsp.Error = gw.reqOp(req.Op)
		if gw.rsp.Error != nil {
			return
		}
		gw.rsp.One = one

	default:

		gw.rsp.Error = &pb.RspError{
			Code: 10,
			Msg:  `unknown type of req.one`,
		}
	}

}

func (gw *gateway) reqPublic(o *pb.ReqPublicList) (rl *pb.RspPublicList, rspErr *pb.RspError) {

	var list []*pb.RspPublic

	for _, v := range o.List {
		var r *pb.RspPublic
		r, rspErr = gw.eachPublic(v)
		if rspErr != nil {
			return
		}
		list = append(list, r)
	}

	rl = &pb.RspPublicList{
		List: list,
	}

	return
}

func (gw *gateway) auth(o *pb.Auth) (rspErr *pb.RspError) {
	return
}

func (gw *gateway) reqOp(o *pb.ReqOpList) (rl *pb.RspOpList, rspErr *pb.RspError) {

	rspErr = gw.auth(o.Auth)
	if rspErr != nil {
		return
	}

	var list []*pb.RspOp

	for _, v := range o.List {
		var r *pb.RspOp
		r, rspErr = gw.eachOp(v)
		if rspErr != nil {
			return
		}
		list = append(list, r)
	}

	rl = &pb.RspOpList{
		List: list,
	}

	return
}
