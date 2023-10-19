package image

import (
	"time"

	"github.com/infraboard/mcube/exception"
	"github.com/rs/xid"
)

// New 实例
func New(req *CreateImageRequest) (*Image, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	now := time.Now()
	img := &Image{
		Id:         GenHashID("image", xid.New().String()),
		CreateAt:   now.Unix(),
		Domain:     req.Domain,
		Namespace:  req.Namespace,
		CreateDate: now.Format("2006-01-02 15:04:05"),
		ImageId:    req.ImageId,
		Name:       req.Name,
		Owner:      req.Owner,
	}

	return img, nil
}

func NewImageByNamespaceSet() *ImageByNamespaceSet {
	return &ImageByNamespaceSet{
		Items: []string{},
	}
}

func (s *ImageByNamespaceSet) Add(item string) {
	s.Items = append(s.Items, item)
}
