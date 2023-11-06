package impl

import (
	"context"

	"github.com/Aaazj/mcenter/apps/device"
	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *impl) get(ctx context.Context, id string) (*device.Device, error) {
	filter := bson.M{"id": id}

	ins := device.NewDefaultDevice()
	if err := s.col.FindOne(ctx, filter).Decode(ins); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("device %s not found", id)
		}

		return nil, exception.NewInternalServerError("find device %s error, %s", id, err)
	}

	return ins, nil
}
