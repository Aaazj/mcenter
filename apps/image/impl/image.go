package impl

import (
	"context"

	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Aaazj/mcenter/apps/image"
)

func (s *impl) CreateImage(ctx context.Context, req *image.CreateImageRequest) (*image.Image, error) {

	img, err := image.New(req)

	if err != nil {
		return nil, err
	}

	// if err := s.save(ctx, img); err != nil {
	// 	return nil, err
	// }

	if _, err := s.col.InsertOne(ctx, img); err != nil {
		return nil, exception.NewInternalServerError("inserted a imgage document error, %s", err)
	}

	return img, nil

}

func (s *impl) DeleteImages(ctx context.Context, req *image.DeleteImagesRequest) (*image.ImageSet, error) {
	// 专门优化单个删除

	var result *mongo.DeleteResult
	var err error
	if len(req.ImageIds) == 1 {
		result, err = s.col.DeleteMany(ctx, bson.M{"image_id": req.ImageIds[0]})
	} else {
		result, err = s.col.DeleteMany(ctx, bson.M{"image_id": bson.M{"$in": req.ImageIds}})
	}

	if err != nil {
		return nil, exception.NewInternalServerError("delete image(%s) error, %s", req.ImageIds[0], err)
	}
	if result.DeletedCount == 0 {
		return nil, exception.NewNotFound("image %s not found", req.ImageIds)
	}
	return nil, nil
}

func (s *impl) QueryImageByNamespace(ctx context.Context, req *image.QueryImageByNamespaceRequest) (*image.ImageByNamespaceSet, error) {
	if req.Domain == "" {
		req.Domain = "default"
	}
	resp, err := s.col.Find(ctx, req.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("find image error, error is %s", err)
	}
	set := image.NewImageByNamespaceSet()

	// 循环
	for resp.Next(ctx) {
		ins := image.NewDefaultImage()
		if err := resp.Decode(ins); err != nil {
			return nil, exception.NewInternalServerError("decode image error, error is %s", err)
		}

		set.Add(ins.ImageId)

	}
	return set, nil
}
