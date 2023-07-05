package impl

import (
	"context"

	"github.com/Aaazj/mcenter/apps/role"
	"github.com/infraboard/mcube/cache"
)

func WrapImpl(i *impl) *Decorator {
	return &Decorator{i}
}

type Decorator struct {
	*impl
}

func (d *Decorator) DescribeRole(ctx context.Context, req *role.DescribeRoleRequest) (
	*role.Role, error) {
	if req.Id != "" {
		ins := role.NewDefaultRole()
		if err := cache.C().Get(req.Id, ins); err != nil {
			d.log.Infof("get %s from cache error, %s", req.Id, err)
		} else {
			return ins, nil
		}
	}

	ins, err := d.impl.DescribeRole(ctx, req)
	if err != nil {
		return nil, err
	}

	if req.Id != "" {
		if err := cache.C().Put(req.Id, ins); err != nil {
			d.log.Infof("set %s to cache error, %s", req.Id, err)
		}
	}
	return ins, nil
}

func (d *Decorator) DeleteRole(ctx context.Context, req *role.DeleteRoleRequest) (
	*role.Role, error) {
	ins, err := d.impl.DeleteRole(ctx, req)
	if err != nil {
		return nil, err
	}

	if err := cache.C().Delete(req.Id); err != nil {
		d.log.Infof("delete %s to cache error, %s", req.Id, err)
	}
	return ins, nil
}

func (d *Decorator) AddPermissionToRole(ctx context.Context, req *role.AddPermissionToRoleRequest) (
	*role.PermissionSet, error) {
	ins, err := d.impl.AddPermissionToRole(ctx, req)
	if err != nil {
		return nil, err
	}

	if err := cache.C().Delete(req.RoleId); err != nil {
		d.log.Infof("delete %s to cache error, %s", req.RoleId, err)
	}
	return ins, nil
}

func (d *Decorator) RemovePermissionFromRole(ctx context.Context, req *role.RemovePermissionFromRoleRequest) (
	*role.PermissionSet, error) {
	ins, err := d.impl.RemovePermissionFromRole(ctx, req)
	if err != nil {
		return nil, err
	}

	if err := cache.C().Delete(req.RoleId); err != nil {
		d.log.Infof("delete %s to cache error, %s", req.RoleId, err)
	}
	return ins, nil
}

func (d *Decorator) UpdatePermission(ctx context.Context, req *role.UpdatePermissionRequest) (
	*role.Permission, error) {
	ins, err := d.impl.UpdatePermission(ctx, req)
	if err != nil {
		return nil, err
	}

	if err := cache.C().Delete(ins.RoleId); err != nil {
		d.log.Infof("delete %s to cache error, %s", req.Id, err)
	}
	return ins, err
}
