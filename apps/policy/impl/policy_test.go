package impl_test

import (
	"testing"

	"github.com/Aaazj/mcenter/apps/domain"
	"github.com/Aaazj/mcenter/apps/label"
	"github.com/Aaazj/mcenter/apps/namespace"
	"github.com/Aaazj/mcenter/apps/policy"
)

func TestCreatePolicy(t *testing.T) {
	req := policy.NewCreatePolicyRequest()
	req.Username = "80user2"
	req.RoleId = "2be09843"
	// req.Domain = domain.DEFAULT_DOMAIN
	req.Domain = "80"
	req.Namespace = namespace.DEFAULT_NAMESPACE
	req.AddScope(label.NewLabelRequirement("env", "test", "prod"))
	req.CreateBy = "admin"
	r, err := impl.CreatePolicy(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r.ToJson())
}

func TestQueryPolicy(t *testing.T) {
	req := policy.NewQueryPolicyRequest()
	req.WithRole = true
	// 查询test用户在默认空间的策略
	req.Username = "kobe01"
	req.Domain = domain.DEFAULT_DOMAIN
	req.Namespace = namespace.DEFAULT_NAMESPACE

	r, err := impl.QueryPolicy(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r.ToJson())
}

func TestCheckPermissionOk(t *testing.T) {
	req := policy.NewCheckPermissionRequest()
	req.Domain = "80"
	req.Namespace = namespace.DEFAULT_NAMESPACE
	req.Username = "80user2"

	// 检查test用户在默认空间下是否有访问mpaas服务的构建配置功能
	req.ServiceId = "1430bd4e"
	req.Path = "POST./cmdb/api/v1/secrets/"
	r, err := impl.CheckPermission(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r.ToJson())
}

func TestCheckPermissionDeny(t *testing.T) {
	req := policy.NewCheckPermissionRequest()
	req.Domain = domain.DEFAULT_DOMAIN
	req.Namespace = namespace.DEFAULT_NAMESPACE
	req.Username = "test"

	// 检查是否有创建Pipeline权限
	req.ServiceId = "cd08fc9c"
	req.Path = "POST./mpaas/api/v1/pipelines"
	_, err := impl.CheckPermission(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
}
