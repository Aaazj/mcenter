package impl

import (
	"github.com/Aaazj/mcenter/apps/audit"
	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func newQueryRequest(req *audit.QueryAuditRequest) *queryAuditRequest {
	return &queryAuditRequest{req}
}

type queryAuditRequest struct {
	*audit.QueryAuditRequest
}

func (r *queryAuditRequest) FindOptions() *options.FindOptions {
	pageSize := int64(r.Page.PageSize)
	skip := int64(r.Page.PageSize) * int64(r.Page.PageNumber-1)

	opt := &options.FindOptions{
		Sort:  bson.D{{Key: "create_at", Value: -1}},
		Limit: &pageSize,
		Skip:  &skip,
	}

	return opt
}

func (r *queryAuditRequest) FindFilter() bson.M {
	filter := bson.M{}

	if r.Username != "" {
		filter["username"] = r.Username
	}

	return filter
}

func newDescriptAuditRequest(req *audit.DescribeAuditRequest) (*describeAuditRequest, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	return &describeAuditRequest{req}, nil
}

type describeAuditRequest struct {
	*audit.DescribeAuditRequest
}

func (req *describeAuditRequest) FindFilter() bson.M {
	filter := bson.M{}
	if req.Userid != "" {
		filter["_id"] = req.Userid
	}

	return filter
}
