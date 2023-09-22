package api

import (
	"fmt"

	"github.com/emicklei/go-restful/v3"

	"github.com/Aaazj/mcenter/apps/audit"
	"github.com/Aaazj/mcenter/conf"
	"k8s.io/klog/v2"
)

func (h *handler) QueryAudit(r *restful.Request, w *restful.Response) {
	res := conf.GeneralResponse{
		Errcode: 0,
		Errmsg:  "OK",
	}

	req := audit.NewQueryAuditRequestFromHTTP(r.Request)

	ins, err := h.service.QueryAudit(r.Request.Context(), req)
	if err != nil {
		//response.Failed(w, err)
		res.Errcode = 401002
		res.Errmsg = err.Error()
		klog.V(4).Info(res)
		if err := w.WriteAsJson(res); err != nil {
			klog.Error(err)
		}
		return
	}

	res.Data = ins
	if err = w.WriteAsJson(res); err != nil {
		klog.Error(err)
	}
	//response.Success(w, ins)
}

func (h *handler) DescribeAudit(r *restful.Request, w *restful.Response) {

	res := conf.GeneralResponse{
		Errcode: 0,
		Errmsg:  "OK",
	}
	req := audit.NewDescriptAuditRequestWithId(r.PathParameter("id"))
	fmt.Printf("req: %v\n", req)
	ins, err := h.service.DescribeAudit(r.Request.Context(), req)
	if err != nil {
		//response.Failed(w, err)
		res.Errcode = 401002
		res.Errmsg = err.Error()
		klog.V(4).Info(res)
		if err := w.WriteAsJson(res); err != nil {
			klog.Error(err)
		}
		return
	}

	res.Data = ins
	if err = w.WriteAsJson(res); err != nil {
		klog.Error(err)
	}
}
