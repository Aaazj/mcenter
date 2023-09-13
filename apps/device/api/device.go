package api

import (
	"fmt"

	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/request"
	"k8s.io/klog/v2"

	"github.com/Aaazj/mcenter/apps/device"
	"github.com/Aaazj/mcenter/conf"
)

func (h *handler) AllocationDevice(r *restful.Request, w *restful.Response) {

	res := conf.GeneralResponse{
		Errcode: 0,
		Errmsg:  "OK",
	}

	req := device.NewDefaultAllocationRequest()

	if err := request.GetDataFromRequest(r.Request, req); err != nil {

		res.Errcode = 401001
		res.Errmsg = err.Error()
		klog.V(4).Info(res)
		if err := w.WriteAsJson(res); err != nil {
			klog.Error(err)
		}
		return
	}

	set, err := h.service.AllocationDevice(r.Request.Context(), req)
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

	res.Data = set
	if err = w.WriteAsJson(res); err != nil {
		klog.Error(err)
	}
	//response.Success(w, set)
}

func (h *handler) ReleaseDevices(r *restful.Request, w *restful.Response) {
	res := conf.GeneralResponse{
		Errcode: 0,
		Errmsg:  "OK",
	}

	req := device.NewReleaseDevicesRequestWithName()

	if err := request.GetDataFromRequest(r.Request, req); err != nil {

		res.Errcode = 401001
		res.Errmsg = err.Error()
		klog.V(4).Info(res)
		if err := w.WriteAsJson(res); err != nil {
			klog.Error(err)
		}
		return
	}
	fmt.Printf("req: %v\n", req)
	set, err := h.service.ReleaseDevices(r.Request.Context(), req)
	if err != nil {
		//response.Failed(w, err)
		res.Errcode = 401001
		res.Errmsg = err.Error()
		klog.V(4).Info(res)
		if err := w.WriteAsJson(res); err != nil {
			klog.Error(err)
		}
		return
	}

	res.Data = set
	if err = w.WriteAsJson(res); err != nil {
		klog.Error(err)
	}
	//response.Success(w, set)
}

func (h *handler) QueryDevice(r *restful.Request, w *restful.Response) {
	res := conf.GeneralResponse{
		Errcode: 0,
		Errmsg:  "OK",
	}
	req := device.NewQueryDeviceRequestFromHTTP(r.Request)

	set, err := h.service.QueryDevice(r.Request.Context(), req)
	if err != nil {
		//response.Failed(w, err)
		res.Errcode = 401001
		res.Errmsg = err.Error()
		klog.V(4).Info(res)
		if err := w.WriteAsJson(res); err != nil {
			klog.Error(err)
		}
		return
	}

	res.Data = set
	if err = w.WriteAsJson(res); err != nil {
		klog.Error(err)
	}
	//response.Success(w, set)
}

func (u *handler) DescribeDevice(r *restful.Request, w *restful.Response) {
	res := conf.GeneralResponse{
		Errcode: 0,
		Errmsg:  "OK",
	}
	req := device.NewDescriptDeviceRequestWithName(r.PathParameter("name"))
	ins, err := h.service.DescribeDevice(r.Request.Context(), req)
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

// func (h *handler) PutNamespace(r *restful.Request, w *restful.Response) {

// 	res := conf.GeneralResponse{
// 		Errcode: 0,
// 		Errmsg:  "OK",
// 	}
// 	req := namespace.NewPutNamespaceRequest(r.PathParameter("d_id"), r.PathParameter("id"))

// 	if err := r.ReadEntity(req.Spec); err != nil {
// 		//response.Failed(w, err)
// 		res.Errcode = 401001
// 		res.Errmsg = err.Error()
// 		klog.V(4).Info(res)
// 		if err := w.WriteAsJson(res); err != nil {
// 			klog.Error(err)
// 		}
// 		return
// 	}

// 	set, err := h.service.UpdateNamespace(r.Request.Context(), req)
// 	if err != nil {
// 		//response.Failed(w, err)
// 		res.Errcode = 401002
// 		res.Errmsg = err.Error()
// 		klog.V(4).Info(res)
// 		if err := w.WriteAsJson(res); err != nil {
// 			klog.Error(err)
// 		}
// 		return
// 	}
// 	//response.Success(w, set)
// 	res.Data = set
// 	if err = w.WriteAsJson(res); err != nil {
// 		klog.Error(err)
// 	}
// }
