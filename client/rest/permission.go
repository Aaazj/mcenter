package rest

import (
	//"github.com/Aaazj/mcenter/apps/policy"
	//"github.com/Aaazj/mcenter/apps/role"
	"github.com/infraboard/mcube/client/rest"
)

type PolicyService interface {
	// 权限校验
	//CheckPermission(context.Context, *policy.CheckPermissionRequest) (*role.Permission, error)
}

type policyImpl struct {
	client *rest.RESTClient
}

// func (i *policyImpl) CheckPermission(ctx context.Context, req *policy.CheckPermissionRequest) (*role.Permission, error) {
// 	// ins := token.NewDefaultToken()
// 	// resp := (ins)

// 	// fmt.Println("bearer " + req.AccessToken)
// 	// err := i.client.
// 	// 	Get("token").
// 	// 	Header(token.VALIDATE_TOKEN_HEADER_KEY, req.AccessToken).
// 	// 	Do(ctx).
// 	// 	Into(resp)
// 	// if err != nil {
// 	// 	return nil, err
// 	// }

// 	// if resp.Error() != nil {
// 	// 	return nil, err
// 	// }

// 	return nil, nil
// }
