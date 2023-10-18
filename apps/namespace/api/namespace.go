package api

import (
	"github.com/emicklei/go-restful/v3"
	"k8s.io/klog/v2"

	"github.com/Aaazj/mcenter/apps/namespace"
	"github.com/Aaazj/mcenter/apps/token"
	"github.com/Aaazj/mcenter/conf"
)

func (h *handler) CreateNamespace(r *restful.Request, w *restful.Response) {

	res := conf.GeneralResponse{
		Errcode: 0,
		Errmsg:  "OK",
	}
	req := namespace.NewCreateNamespaceRequest()

	if err := r.ReadEntity(req); err != nil {
		//response.Failed(w, err)
		res.Errcode = 401001
		res.Errmsg = err.Error()
		klog.V(4).Info(res)
		if err := w.WriteAsJson(res); err != nil {
			klog.Error(err)
		}
		return
	}

	req.UpdateOwner(token.GetTokenFromRequest(r))

	set, err := h.service.CreateNamespace(r.Request.Context(), req)
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

func (h *handler) QueryNamespace(r *restful.Request, w *restful.Response) {
	res := conf.GeneralResponse{
		Errcode: 0,
		Errmsg:  "OK",
	}
	req := namespace.NewQueryNamespaceRequestFromHTTP(r.Request)

	set, err := h.service.QueryNamespace(r.Request.Context(), req)
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

func (h *handler) DeleteNamespace(r *restful.Request, w *restful.Response) {
	res := conf.GeneralResponse{
		Errcode: 0,
		Errmsg:  "OK",
	}
	domain := r.PathParameter("d_id")
	name := r.PathParameter("id")
	req := namespace.NewDeleteNamespaceRequest(domain, name)
	set, err := h.service.DeleteNamespace(r.Request.Context(), req)
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

func (u *handler) DescribeNamespace(r *restful.Request, w *restful.Response) {
	res := conf.GeneralResponse{
		Errcode: 0,
		Errmsg:  "OK",
	}
	req := namespace.NewDescriptNamespaceRequest(r.PathParameter("d_id"), r.PathParameter("id"))
	ins, err := h.service.DescribeNamespace(r.Request.Context(), req)
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

func (h *handler) PutNamespace(r *restful.Request, w *restful.Response) {

	res := conf.GeneralResponse{
		Errcode: 0,
		Errmsg:  "OK",
	}
	req := namespace.NewPutNamespaceRequest(r.PathParameter("d_id"), r.PathParameter("id"))

	if err := r.ReadEntity(req.Spec); err != nil {
		//response.Failed(w, err)
		res.Errcode = 401001
		res.Errmsg = err.Error()
		klog.V(4).Info(res)
		if err := w.WriteAsJson(res); err != nil {
			klog.Error(err)
		}
		return
	}

	set, err := h.service.UpdateNamespace(r.Request.Context(), req)
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
	//response.Success(w, set)
	res.Data = set
	if err = w.WriteAsJson(res); err != nil {
		klog.Error(err)
	}
}
