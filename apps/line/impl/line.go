package impl

import (
	"context"

	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Aaazj/mcenter/apps/line"
)

func (s *impl) CreateLine(ctx context.Context, req *line.CreateLineRequest) (*line.Line, error) {
	if req.Domain == "" {
		req.Domain = "default"
	}
	l, err := line.New(req)

	if err != nil {
		return nil, err
	}

	if err := s.save(ctx, l); err != nil {
		return nil, err
	}

	return l, nil

}

func (s *impl) DeleteLines(ctx context.Context, req *line.DeleteLinesRequest) (*line.LineSet, error) {
	// 专门优化单个删除

	var result *mongo.DeleteResult
	var err error
	if len(req.Ids) == 1 {
		result, err = s.col.DeleteMany(ctx, bson.M{"line_id": req.Ids[0]})
	} else {
		result, err = s.col.DeleteMany(ctx, bson.M{"line_id": bson.M{"$in": req.Ids}})
	}

	if err != nil {
		return nil, exception.NewInternalServerError("delete line(%s) error, %s", req.Ids[0], err)
	}
	if result.DeletedCount == 0 {
		return nil, exception.NewNotFound("line %s not found", req.Ids)
	}
	return nil, nil
}

func (s *impl) QueryLineByNamespace(ctx context.Context, req *line.QueryLineByNamespaceRequest) (*line.LineByNamespaceSet, error) {
	if req.Domain == "" {
		req.Domain = "default"
	}
	resp, err := s.col.Find(ctx, req.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("find line error, error is %s", err)
	}
	set := line.NewLineByNamespaceSet()

	// 循环
	for resp.Next(ctx) {
		ins := line.NewDefaultLine()
		if err := resp.Decode(ins); err != nil {
			return nil, exception.NewInternalServerError("decode line error, error is %s", err)
		}

		set.Add(ins.LineId)

	}
	return set, nil
}
