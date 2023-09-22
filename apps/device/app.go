package device

import (
	"errors"
	"fmt"
	"hash/fnv"
	"net/http"
	"time"

	"github.com/Aaazj/mcenter/conf"
	"github.com/go-playground/validator/v10"
	request "github.com/infraboard/mcube/http/request"
	"go.mongodb.org/mongo-driver/bson"
)

func NewDefaultDevice() *Device {
	return &Device{}
}

// GenHashID hash id
func GenHashID(service, grpcPath string) string {
	hashedStr := fmt.Sprintf("%s-%s", service, grpcPath)
	h := fnv.New32a()
	h.Write([]byte(hashedStr))
	return fmt.Sprintf("%x", h.Sum32())
}

func (req *AllocationRequest) Devices() []*Device {

	eps := make([]*Device, 0, len(req.Entries))
	for i := range req.Entries {
		if req.Entries[i].AllocateDays == 0 {
			req.Entries[i].AllocateDays = 30
		}
		ep := &Device{
			Id:            GenHashID("device", req.Entries[i].Name),
			CreateAt:      time.Now().Unix(),
			Domain:        req.Domain,
			Namespace:     req.Namespace,
			CreateDate:    time.Now().Format("2006-01-02 15:04:05"),
			RemainingDays: req.Entries[i].AllocateDays,
			Entry:         req.Entries[i],
		}
		eps = append(eps, ep)
	}
	return eps
}

var (
	validate = validator.New()
)

// Validate 校验注册请求合法性
func (req *AllocationRequest) Validate() error {
	if len(req.Entries) == 0 {
		return fmt.Errorf("must require device")
	}

	return validate.Struct(req)
}

func NewAllocationResponse(message string) *AllocationResponse {
	return &AllocationResponse{Message: message}
}

func NewDefaultAllocationRequest() *AllocationRequest {
	return &AllocationRequest{
		Domain:    "default",
		Namespace: "default",
		Entries:   []*DeviceEntry{},
	}
}

func NewDescriptDeviceRequestWithName(name string) *DescribeDeviceRequest {
	return &DescribeDeviceRequest{Name: name}
}

// Validate 校验
func (req *DescribeDeviceRequest) Validate() error {
	if req.Name == "" {
		return fmt.Errorf("device name is required")
	}

	return nil
}

func (x *Device) ValidateDevice() bool {
	Expiredtime := x.CreateAt + x.Entry.AllocateDays*24*3600
	// fmt.Printf("Expiredtime: %v\n", Expiredtime)
	// expiredAt := time.Unix(x.CreateAt, 0).Add(time.Duration(x.Entry.AllocateDays) * time.Hour * 24)
	// fmt.Printf("expiredAt: %v\n", expiredAt)
	// t := time.Unix(Expiredtime, 0)
	// // 将Time转换成我们想要的日期格式。
	// s := t.Format("2006-01-02 15:04:05")
	// fmt.Println(s)

	now := time.Now()
	expiredAt := time.Unix(x.CreateAt, 0).Add(time.Duration(x.Entry.AllocateDays) * time.Hour * 24)

	ex := now.Sub(expiredAt).Hours() / 24
	fmt.Printf("ex: %v\n", ex)
	x.RemainingDays = int64(-ex)
	fmt.Printf("x: %v\n", x)
	return time.Unix(Expiredtime, 0).After(time.Now())

}

func NewReleaseDevicesRequestWithName() *ReleaseDevicesRequest {
	return &ReleaseDevicesRequest{
		Names: []string{},
	}
}

// Validate 校验参数
func (m *ValidateDeviceRequest) Validate() error {
	if err := validate.Struct(m); err != nil {
		return err
	}

	if m.Name == "" {
		return errors.New("device name required one")
	}

	return nil
}

func NewQueryDeviceRequestFromHTTP(r *http.Request) *QueryDeviceRequest {
	query := NewQueryDeviceRequest()

	query.Page = request.NewPageRequestFromHTTP(r)
	qs := r.URL.Query()
	namespace := qs.Get("namespace")
	if namespace != "" {
		query.Namespace = namespace
	}

	return query
}

func NewQueryDeviceRequest() *QueryDeviceRequest {
	return &QueryDeviceRequest{
		Page: request.NewPageRequest(conf.DEFAULT_PAGE_SIZE, 1),
	}
}

func NewQueryDeviceByNamespaceRequest(namespace string) *QueryDeviceByNamespaceRequest {
	return &QueryDeviceByNamespaceRequest{
		Namespace: namespace,
		Domain:    "default",
	}
}

func (r *QueryDeviceByNamespaceRequest) FindFilter() bson.M {
	filter := bson.M{}

	if r.Namespace != "" {
		filter["namespace"] = r.Namespace
	}

	if r.Domain != "" {
		filter["domain"] = r.Domain
	}
	return filter
}
