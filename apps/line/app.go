package line

import (
	"fmt"
	"hash/fnv"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
)

func NewDefaultLine() *Line {
	return &Line{}
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

// Validate 校验注册请求合法性
func (req *CreateLineRequest) Validate() error {
	if req.LineId == "" {
		return fmt.Errorf("must require line")
	}

	return validate.Struct(req)
}

func (r *QueryLineByNamespaceRequest) FindFilter() bson.M {
	filter := bson.M{}

	if r.Namespace != "" {
		filter["namespace"] = r.Namespace
	}

	if r.Domain != "" {
		filter["domain"] = r.Domain
	}
	return filter
}

func NewCreateLineRequest(namespace string, id string, name string, owner string) *CreateLineRequest {
	return &CreateLineRequest{
		Namespace: namespace,
		Domain:    "default",
		LineId:    id,
		Name:      name,
		Owner:     owner,
	}
}

func NewDeleteLinesRequestWithID() *DeleteLinesRequest {
	return &DeleteLinesRequest{
		Ids: []string{},
	}
}

func NewQueryLineByNamespaceRequest(namespace string) *QueryLineByNamespaceRequest {
	return &QueryLineByNamespaceRequest{
		Namespace: namespace,
		Domain:    "default",
	}
}
