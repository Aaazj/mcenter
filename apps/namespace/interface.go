package namespace

import context "context"

const (
	AppName = "namespace"
)

type Service interface {
	CreateNamespace(context.Context, *CreateNamespaceRequest) (*Namespace, error)
	DeleteNamespace(context.Context, *DeleteNamespaceRequest) (*Namespace, error)
	UpdateNamespace(context.Context, *UpdateNamespaceRequest) (*Namespace, error)
	RPCServer
}
