package impl

import (
	"context"

	"github.com/Aaazj/mcenter/apps/line"
	"github.com/infraboard/mcube/exception"
)

func (s *impl) save(ctx context.Context, line *line.Line) error {
	if _, err := s.col.InsertOne(ctx, line); err != nil {
		return exception.NewInternalServerError("inserted line(%s) document error, %s",
			line.LineId, err)
	}

	return nil
}
