package api

import (
	"fmt"

	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/restful/response"

	"github.com/Aaazj/mcenter/apps/token"
	"github.com/Aaazj/mcenter/conf"
	"k8s.io/klog/v2"
)

func (h *handler) IssueToken(r *restful.Request, w *restful.Response) {

	res := conf.GeneralResponse{
		Errcode: 0,
		Errmsg:  "OK",
	}

	req := token.NewIssueTokenRequest()
	fmt.Printf("\"tokennnnnnnnnnnnnnnnnnnnnn\": %v\n", "tokennnnnnnnnnnnnnnnnnnnnn")
	if err := r.ReadEntity(req); err != nil {

		//response.Failed(w, err)

		res.Errcode = 401001
		res.Errmsg = "读取账户密码信息失败:" + err.Error()
		klog.V(4).Info(res)
		if err := w.WriteAsJson(res); err != nil {
			klog.Error(err)
		}
		return

	}

	// 补充用户的登录时的位置信息
	req.Location = token.NewNewLocationFromHttp(r.Request)

	tk, err := h.service.IssueToken(r.Request.Context(), req)
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
	res.Data = tk
	if err = w.WriteAsJson(res); err != nil {
		klog.Error(err)
	}
	//tk.SetCookie(w)
	//response.Success(w, tk)
}

func (u *handler) RevolkToken(r *restful.Request, w *restful.Response) {
	res := conf.GeneralResponse{
		Errcode: 0,
		Errmsg:  "OK",
	}
	qs := r.Request.URL.Query()

	req := token.NewRevolkTokenRequest("", "")
	req.AccessToken = token.GetAccessTokenFromHTTP(r.Request)

	req.RefreshToken = qs.Get("refresh_token")

	tk := r.Attribute(token.TOKEN_ATTRIBUTE_NAME).(*token.Token)
	req.RefreshToken = tk.RefreshToken

	ins, err := h.service.RevolkToken(r.Request.Context(), req)
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

func (u *handler) ValidateToken(r *restful.Request, w *restful.Response) {
	tk := r.Request.Header.Get(token.VALIDATE_TOKEN_HEADER_KEY)
	req := token.NewValidateTokenRequest(tk)

	resp, err := h.service.ValidateToken(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, resp)
}

func (h *handler) QueryToken(r *restful.Request, w *restful.Response) {
	req := token.NewQueryTokenRequestFromHttp(r)

	resp, err := h.service.QueryToken(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, resp)
}
