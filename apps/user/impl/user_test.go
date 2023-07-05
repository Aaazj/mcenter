package impl_test

import (
	"testing"

	"github.com/Aaazj/mcenter/apps/domain"
	"github.com/Aaazj/mcenter/apps/user"
)

func TestCreateSupperUser(t *testing.T) {
	req := user.NewCreateUserRequest()
	// req.Domain = domain.DEFAULT_DOMAIN
	//req.Domain = "80"
	//req.Namespace = "80"
	req.Username = "80user999"
	req.Password = "123456"
	req.Type = user.TYPE_PRIMARY
	r, err := impl.CreateUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r.ToJson())
}

func TestCreateSubUser(t *testing.T) {
	req := user.NewCreateUserRequest()
	req.Domain = domain.DEFAULT_DOMAIN
	req.Domain = "wangheng"
	req.Username = "wangheng"
	req.Password = "123456"
	req.Type = user.TYPE_SUB
	r, err := impl.CreateUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r.ToJson())
}

func TestQueryUser(t *testing.T) {
	req := user.NewQueryUserRequest()
	//req.Domain = "wangheng"
	r, err := impl.QueryUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r.ToJson())
}

func TestDescribeUser(t *testing.T) {
	req := user.NewDescriptUserRequestById("kobe01")
	r, err := impl.DescribeUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r.ToJson())
}

// func TestDescribeUserByFeishuUserId(t *testing.T) {
// 	req := user.NewDescriptUserRequestByFeishuUserId(os.Getenv("FEISHU_USER_ID"))
// 	r, err := impl.DescribeUser(ctx, req)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	t.Log(r.ToJson())
// }

// func TestPatchUser(t *testing.T) {
// 	req := user.NewPatchUserRequest("kobe01")
// 	req.Profile.Phone = os.Getenv("TEST_CALL_NUMBER")
// 	req.Feishu.UserId = os.Getenv("FEISHU_USER_ID")
// 	r, err := impl.UpdateUser(ctx, req)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	t.Log(r.ToJson())
// }
