package api

import (
	"fmt"

	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/restful/response"

	"github.com/Aaazj/mcenter/apps/namespace"
	"github.com/Aaazj/mcenter/apps/token"
)

func (h *handler) CreateNamespace(r *restful.Request, w *restful.Response) {
	req := namespace.NewCreateNamespaceRequest()

	if err := r.ReadEntity(req); err != nil {
		response.Failed(w, err)
		return
	}
	tk := r.Attribute(token.TOKEN_ATTRIBUTE_NAME).(*token.Token)
	fmt.Printf("tk: %v\n", tk)
	req.UpdateOwner(token.GetTokenFromRequest(r))

	set, err := h.service.CreateNamespace(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set)
}

func (h *handler) QueryNamespace(r *restful.Request, w *restful.Response) {
	req := namespace.NewQueryNamespaceRequestFromHTTP(r.Request)
	set, err := h.service.QueryNamespace(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set)
}

func (h *handler) DeleteNamespace(r *restful.Request, w *restful.Response) {
	domain := r.PathParameter("d_id")
	name := r.PathParameter("id")
	req := namespace.NewDeleteNamespaceRequest(domain, name)
	set, err := h.service.DeleteNamespace(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set)
}

func (u *handler) DescribeNamespace(r *restful.Request, w *restful.Response) {
	req := namespace.NewDescriptNamespaceRequest(r.PathParameter("d_id"), r.PathParameter("id"))
	ins, err := h.service.DescribeNamespace(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}
