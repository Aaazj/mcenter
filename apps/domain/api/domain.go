package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/restful/response"
	"k8s.io/klog/v2"

	"github.com/Aaazj/mcenter/apps/domain"
	"github.com/Aaazj/mcenter/conf"
)

func (h *handler) CreateDomain(r *restful.Request, w *restful.Response) {

	res := conf.GeneralResponse{
		Errcode: 0,
		Errmsg:  "OK",
	}

	req := domain.NewCreateDomainRequest()

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

	set, err := h.service.CreateDomain(r.Request.Context(), req)
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

func (u *handler) DescribeDomain(r *restful.Request, w *restful.Response) {
	req := domain.NewDescribeDomainRequestById(r.PathParameter("id"))
	ins, err := h.service.DescribeDomain(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (u *handler) PutDomain(r *restful.Request, w *restful.Response) {
	req := domain.NewPutDomainRequest(r.PathParameter("id"))

	if err := r.ReadEntity(req.Spec); err != nil {
		response.Failed(w, err)
		return
	}

	set, err := h.service.UpdateDomain(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (u *handler) PatchDomain(r *restful.Request, w *restful.Response) {
	req := domain.NewPatchDomainRequestById(r.PathParameter("id"))

	if err := r.ReadEntity(req.Spec); err != nil {
		response.Failed(w, err)
		return
	}

	set, err := h.service.UpdateDomain(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}
