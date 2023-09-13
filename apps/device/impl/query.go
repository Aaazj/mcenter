package impl

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/Aaazj/mcenter/apps/device"
	"github.com/infraboard/mcube/exception"
)

func newDescriptDeviceRequest(req *device.DescribeDeviceRequest) (*describeDeviceRequest, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	return &describeDeviceRequest{req}, nil
}

type describeDeviceRequest struct {
	*device.DescribeDeviceRequest
}

func (req *describeDeviceRequest) FindFilter() bson.M {
	filter := bson.M{}
	if req.Name != "" {
		filter["name"] = req.Name
	}

	return filter
}

func newQueryRequest(req *device.QueryDeviceRequest) *queryDeviceRequest {
	return &queryDeviceRequest{req}
}

type queryDeviceRequest struct {
	*device.QueryDeviceRequest
}

func (r *queryDeviceRequest) FindOptions() *options.FindOptions {
	pageSize := int64(r.Page.PageSize)
	skip := int64(r.Page.PageSize) * int64(r.Page.PageNumber-1)

	opt := &options.FindOptions{
		Sort:  bson.D{{Key: "create_at", Value: -1}},
		Limit: &pageSize,
		Skip:  &skip,
	}

	return opt
}

func (r *queryDeviceRequest) FindFilter() bson.M {
	filter := bson.M{}

	if r.Namespace != "" {
		filter["namespace"] = r.Namespace
	}

	if r.Domain != "" {
		filter["domain"] = r.Domain
	}
	return filter
}
