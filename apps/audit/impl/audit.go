package impl

import (
	"context"

	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Aaazj/mcenter/apps/audit"
)

func (s *service) AuditOperate(ctx context.Context, req *audit.OperateLog) (
	*audit.AuditOperateLogResponse, error) {
	if _, err := s.col.InsertOne(ctx, req); err != nil {
		return nil, exception.NewInternalServerError("inserted audit log(%s) document error, %s",
			req, err)
	}

	return &audit.AuditOperateLogResponse{}, nil
}

func (s *service) QueryAudit(ctx context.Context, req *audit.QueryAuditRequest) (
	*audit.OperateLogSet, error) {
	r := newQueryRequest(req)

	resp, err := s.col.Find(ctx, r.FindFilter(), r.FindOptions())
	if err != nil {
		return nil, exception.NewInternalServerError("find audit error, error is %s", err)
	}

	set := audit.NewAuditSet()
	// 循环
	for resp.Next(ctx) {
		ins := audit.NewDefaultAudit()
		if err := resp.Decode(ins); err != nil {
			return nil, exception.NewInternalServerError("decode audit error, error is %s", err)
		}

		set.Add(ins)
	}

	// count
	if len(r.Username) == 0 {
		count, err := s.col.CountDocuments(ctx, r.FindFilter())
		if err != nil {
			return nil, exception.NewInternalServerError("get audit count error, error is %s", err)
		}
		set.Total = count
	}

	return set, nil
}

func (s *service) DescribeAudit(ctx context.Context, req *audit.DescribeAuditRequest) (
	*audit.OperateLog, error) {
	r, err := newDescriptAuditRequest(req)
	if err != nil {
		return nil, err
	}

	ins := audit.NewDefaultAudit()
	if err := s.col.FindOne(ctx, r.FindFilter()).Decode(ins); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("Audit %s not found", req)
		}

		return nil, exception.NewInternalServerError("find Audit %s error, %s", req.Userid, err)
	}

	return ins, nil
}
