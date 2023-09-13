package impl

import (
	"context"

	"github.com/Aaazj/mcenter/apps/device"
	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *impl) get(ctx context.Context, name string) (*device.Device, error) {
	filter := bson.M{"name": name}

	ins := device.NewDefaultDevice()
	if err := s.col.FindOne(ctx, filter).Decode(ins); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("device %s not found", name)
		}

		return nil, exception.NewInternalServerError("find device %s error, %s", name, err)
	}

	return ins, nil
}
