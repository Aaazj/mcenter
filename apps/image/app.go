package image

import (
	"fmt"
	"hash/fnv"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
)

func NewDefaultImage() *Image {
	return &Image{}
}

// GenHashID hash id
func GenHashID(service, grpcPath string) string {
	hashedStr := fmt.Sprintf("%s-%s", service, grpcPath)
	h := fnv.New32a()
	h.Write([]byte(hashedStr))
	return fmt.Sprintf("%x", h.Sum32())
}

var (
	validate = validator.New()
)

func NewCreateImageRequest(namespace string, id string, name string, owner string) *CreateImageRequest {
	return &CreateImageRequest{
		Namespace: namespace,
		Domain:    "default",
		ImageId:   id,
		Name:      name,
		Owner:     owner,
	}
}

// Validate 校验注册请求合法性
func (req *CreateImageRequest) Validate() error {
	if req.ImageId == "" {
		return fmt.Errorf("must require image")
	}

	return validate.Struct(req)
}

func NewDeleteImagesRequestWithID() *DeleteImagesRequest {
	return &DeleteImagesRequest{
		ImageIds: []string{},
	}
}

func NewQueryImageByNamespaceRequest(namespace string) *QueryImageByNamespaceRequest {
	return &QueryImageByNamespaceRequest{
		Namespace: namespace,
		Domain:    "default",
	}
}

func (r *QueryImageByNamespaceRequest) FindFilter() bson.M {
	filter := bson.M{}

	if r.Namespace != "" {
		filter["namespace"] = r.Namespace
	}

	if r.Domain != "" {
		filter["domain"] = r.Domain
	}
	return filter
}
