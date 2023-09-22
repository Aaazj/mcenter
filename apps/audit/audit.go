package audit

import (
	"net/http"

	"github.com/Aaazj/mcenter/conf"
	request "github.com/infraboard/mcube/http/request"
)

func NewQueryAuditRequestFromHTTP(r *http.Request) *QueryAuditRequest {
	query := NewQueryAuditRequest()

	query.Page = request.NewPageRequestFromHTTP(r)
	qs := r.URL.Query()
	username := qs.Get("username")
	if username != "" {
		query.Username = username
	}

	return query
}

func NewQueryAuditRequest() *QueryAuditRequest {
	return &QueryAuditRequest{
		Page: request.NewPageRequest(conf.DEFAULT_PAGE_SIZE, 1),
	}
}

func NewAuditSet() *OperateLogSet {
	return &OperateLogSet{
		Items: []*OperateLog{},
	}
}

func (s *OperateLogSet) Add(item *OperateLog) {
	s.Items = append(s.Items, item)
	s.Total += 1
}
