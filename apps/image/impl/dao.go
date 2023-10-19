package impl

import (
	"context"

	"github.com/Aaazj/mcenter/apps/image"
	"github.com/infraboard/mcube/exception"
)

func (s *impl) save(ctx context.Context, img *image.Image) error {
	if _, err := s.col.InsertOne(ctx, img); err != nil {
		return exception.NewInternalServerError("inserted image(%s) document error, %s",
			img.ImageId, err)
	}

	return nil
}
