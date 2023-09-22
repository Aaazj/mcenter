package impl

import (
	"context"
	"fmt"

	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"k8s.io/klog/v2"

	"github.com/Aaazj/mcenter/apps/device"
)

func (s *impl) DescribeDevice(ctx context.Context, req *device.DescribeDeviceRequest) (
	*device.Device, error) {

	r, err := newDescriptDeviceRequest(req)
	if err != nil {
		return nil, err
	}

	ins := device.NewDefaultDevice()
	if err := s.col.FindOne(ctx, r.FindFilter()).Decode(ins); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("Device %s not found", req.Name)
		}

		return nil, exception.NewInternalServerError("find Device %s error, %s", req.Name, err)
	}

	if !ins.ValidateDevice() { //过期  释放资源
		_, err := s.col.DeleteOne(ctx, r.FindFilter())
		if err != nil {
			return nil, exception.NewInternalServerError("device expired, delete device (%s)  error, %s", req.Name, err)
		}
		return nil, exception.NewInternalServerError("device expired, delete device (%s) ", req.Name)
	} else {
		if err := s.col.FindOneAndReplace(ctx, bson.M{"_id": ins.Id}, ins).Err(); err != nil {
			klog.Error("device (%s)剩余时间更新失败", ins.Entry.Name)
		}
	}

	return ins, nil
}

func (s *impl) QueryDevice(ctx context.Context, req *device.QueryDeviceRequest) (
	*device.DeviceSet, error) {
	r := newQueryRequest(req)

	resp, err := s.col.Find(ctx, r.FindFilter(), r.FindOptions())
	if err != nil {
		return nil, exception.NewInternalServerError("find device error, error is %s", err)
	}

	set := device.NewDeviceSet()

	// 循环
	for resp.Next(ctx) {
		ins := device.NewDefaultDevice()
		if err := resp.Decode(ins); err != nil {
			return nil, exception.NewInternalServerError("decode device error, error is %s", err)
		}

		if !ins.ValidateDevice() { //过期  释放资源
			_, err := s.col.DeleteOne(ctx, r.FindFilter())
			if err != nil {

				klog.Error(exception.NewInternalServerError("device expired, delete device (%s)  error, %s", ins.Entry.Name, err))

			}
			klog.Error(exception.NewInternalServerError("device expired, delete device (%s) ", ins.Entry.Name))

		} else {
			if err := s.col.FindOneAndReplace(ctx, bson.M{"_id": ins.Id}, ins).Err(); err != nil {
				klog.Error("device (%s)剩余时间更新失败", ins.Entry.Name)
			}
			set.Add(ins)
		}
	}

	// count
	count, err := s.col.CountDocuments(ctx, r.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get device count error, error is %s", err)
	}
	set.Total = count
	return set, nil
}

func (s *impl) AllocationDevice(ctx context.Context, req *device.AllocationRequest) (*device.DeviceSet, error) {

	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}
	set := device.NewDeviceSet()
	devices := req.Devices()

	// 更新已有的记录
	news := make([]interface{}, 0, len(devices))
	for i := range devices {
		set.Add(devices[i])
		set.Total += 1
		if err := s.col.FindOneAndReplace(ctx, bson.M{"_id": devices[i].Id}, devices[i]).Err(); err != nil {
			if err == mongo.ErrNoDocuments {
				news = append(news, devices[i])
			} else {
				return nil, err
			}
		}
	}

	// 插入新增记录
	if len(news) > 0 {
		if _, err := s.col.InsertMany(ctx, news); err != nil {
			return nil, exception.NewInternalServerError("inserted a service document error, %s", err)
		}
	}

	return set, nil

}

func (s *impl) ReleaseDevices(ctx context.Context, req *device.ReleaseDevicesRequest) (*device.DeviceSet, error) {
	// 专门优化单个删除
	fmt.Printf("req.Names: %v\n", req.Names)
	var result *mongo.DeleteResult
	var err error
	if len(req.Names) == 1 {
		result, err = s.col.DeleteMany(ctx, bson.M{"name": req.Names[0]})
	} else {
		result, err = s.col.DeleteMany(ctx, bson.M{"name": bson.M{"$in": req.Names}})
	}

	// fmt.Printf("req.Names[0]: %v\n", req.Names[0])
	// //_, err := s.col.DeleteMany(ctx, bson.M{"name": req.Names[0]})
	// _, err := s.col.DeleteOne(ctx, bson.M{"name": req.Names[0]})
	fmt.Printf("err: %v\n", err)
	if err != nil {
		return nil, exception.NewInternalServerError("delete device(%s) error, %s", req.Names[0], err)
	}
	if result.DeletedCount == 0 {
		fmt.Printf("\"return exception.NewNotFound\": %v\n", "return exception.NewNotFound")
	}
	return nil, nil
}

func (s *impl) ValidateDevice(ctx context.Context, req *device.ValidateDeviceRequest) (*device.Device, error) {

	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	device, err := s.get(ctx, req.Name)
	if err != nil {
		return nil, exception.NewUnauthorized(err.Error())
	}

	if !device.ValidateDevice() { //过期  释放资源
		_, err := s.col.DeleteOne(ctx, bson.M{"name": device.Entry.Name})
		if err != nil {
			return nil, exception.NewInternalServerError("device expired, delete device (%s)  error, %s", device.Entry.Name, err)
		}
		return nil, exception.NewInternalServerError("device expired, delete device (%s) ", device.Entry.Name)
	}
	return device, nil
}

func (s *impl) QueryDeviceByNamespace(ctx context.Context, req *device.QueryDeviceByNamespaceRequest) (*device.DeviceByNamespaceSet, error) {

	resp, err := s.col.Find(ctx, req.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("find device error, error is %s", err)
	}
	set := device.NewDeviceByNamespaceSet()

	// 循环
	for resp.Next(ctx) {
		ins := device.NewDefaultDevice()
		if err := resp.Decode(ins); err != nil {
			return nil, exception.NewInternalServerError("decode device error, error is %s", err)
		}

		if !ins.ValidateDevice() { //过期  释放资源
			_, err := s.col.DeleteOne(ctx, req.FindFilter())
			if err != nil {

				klog.Error(exception.NewInternalServerError("device expired, delete device (%s)  error, %s", ins.Entry.Name, err))

			}
			klog.Error(exception.NewInternalServerError("device expired, delete device (%s) ", ins.Entry.Name))

		} else {
			if err := s.col.FindOneAndReplace(ctx, bson.M{"_id": ins.Id}, ins).Err(); err != nil {
				klog.Error("device (%s)剩余时间更新失败", ins.Entry.Name)
			}
			set.Add(ins.Entry.Name)
		}
	}
	return set, nil
}
