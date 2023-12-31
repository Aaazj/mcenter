package impl

import (
	"context"
	"fmt"
	"time"

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
			return nil, exception.NewNotFound("Device %s not found", req.Id)
		}

		return nil, exception.NewInternalServerError("find Device %s error, %s", req.Id, err)
	}

	if !ins.ValidateDevice() { //过期  释放资源
		//_, err := s.col.DeleteOne(ctx, r.FindFilter())
		_, err := s.col.DeleteOne(ctx, bson.M{"id": ins.Entry.Id})

		if err != nil {
			return nil, exception.NewInternalServerError("device expired, delete device (%s)  error, %s", req.Id, err)
		}
		return nil, exception.NewInternalServerError("device expired, delete device (%s) ", req.Id)
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
			//_, err := s.col.DeleteOne(ctx, r.FindFilter())
			_, err := s.col.DeleteOne(ctx, bson.M{"id": ins.Entry.Id})
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
		if err := s.col.FindOneAndReplace(ctx, bson.M{"id": devices[i].Entry.Id}, devices[i]).Err(); err != nil {
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

	var result *mongo.DeleteResult
	var err error
	if len(req.Ids) == 1 {
		fmt.Printf("req.Ids[0]: %v\n", req.Ids[0])
		result, err = s.col.DeleteMany(ctx, bson.M{"id": req.Ids[0]})
	} else {
		result, err = s.col.DeleteMany(ctx, bson.M{"id": bson.M{"$in": req.Ids}})
	}
	fmt.Printf("err: %v\n", err)
	if err != nil {
		return nil, exception.NewInternalServerError("delete device(%s) error, %s", req.Ids[0], err)
	}
	if result.DeletedCount == 0 {

		return nil, nil
	}
	return nil, nil
}

func (s *impl) ValidateDevice(ctx context.Context, req *device.ValidateDeviceRequest) (*device.Device, error) {

	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	device, err := s.get(ctx, req.Id)
	if err != nil {
		return nil, exception.NewUnauthorized(err.Error())
	}

	if !device.ValidateDevice() { //过期  释放资源
		_, err := s.col.DeleteOne(ctx, bson.M{"id": device.Entry.Id})
		if err != nil {
			return nil, exception.NewInternalServerError("device expired, delete device (%s)  error, %s", device.Entry.Id, err)
		}
		return nil, exception.NewInternalServerError("device expired, delete device (%s) ", device.Entry.Id)
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

			_, err := s.col.DeleteOne(ctx, bson.M{"id": ins.Entry.Id})
			if err != nil {

				klog.Error(exception.NewInternalServerError("device expired, delete device (%s)  error, %s", ins.Entry.Id, err))

			}
			klog.Error(exception.NewInternalServerError("device expired, delete device (%s) ", ins.Entry.Id))

		} else {
			if err := s.col.FindOneAndReplace(ctx, bson.M{"id": ins.Entry.Id}, ins).Err(); err != nil {
				klog.Error("device (%s)剩余时间更新失败", ins.Entry.Id)
			}
			set.Add(ins.Entry.Id)
		}
	}
	return set, nil
}

func (s *impl) RenewalDevice(ctx context.Context, req *device.DeviceRenewalRequest) (*device.Device, error) {

	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	ins := device.NewDefaultDevice()
	if err := s.col.FindOne(ctx, bson.M{"id": req.Id}).Decode(ins); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("Device %s not found or Device expired", req.Id)
		}

		return nil, exception.NewInternalServerError("find Device %s error, %s", req.Id, err)
	}
	ins.Entry.AllocateDays += req.RenewalTime
	expiredAt := time.Unix(ins.CreateAt, 0).Add(time.Duration(ins.Entry.AllocateDays) * time.Hour * 24)

	ins.ExpiredDate = expiredAt.Format("2006-01-02 15:04:05")

	if !ins.ValidateDevice() { //过期  释放资源
		_, err := s.col.DeleteOne(ctx, bson.M{"id": req.Id})
		if err != nil {
			return nil, exception.NewInternalServerError("device expired, delete device (%s)  error, %s", req.Id, err)
		}
		return nil, exception.NewInternalServerError("device expired, delete device (%s) ", req.Id)
	} else {
		if err := s.col.FindOneAndReplace(ctx, bson.M{"_id": ins.Id}, ins).Err(); err != nil {
			klog.Error("device (%s)剩余时间更新失败", ins.Entry.Id)
		}
	}

	return ins, nil

}
