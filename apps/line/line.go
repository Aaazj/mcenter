package line

import (
	"time"

	"github.com/infraboard/mcube/exception"
	"github.com/rs/xid"
)

// New 实例
func New(req *CreateLineRequest) (*Line, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	now := time.Now()
	img := &Line{
		Id:         GenHashID("line", xid.New().String()),
		CreateAt:   now.Unix(),
		Domain:     req.Domain,
		Namespace:  req.Namespace,
		CreateDate: now.Format("2006-01-02 15:04:05"),
		LineId:     req.LineId,
		Name:       req.Name,
		Owner:      req.Owner,
	}

	return img, nil
}

func NewLineByNamespaceSet() *LineByNamespaceSet {
	return &LineByNamespaceSet{
		Items: []string{},
	}
}

func (s *LineByNamespaceSet) Add(item string) {
	s.Items = append(s.Items, item)
}
