package impl_test

import (
	"testing"

	"github.com/Aaazj/mcenter/apps/device"
)

func TestQueryDeviceByNamespace(t *testing.T) {
	req := device.NewQueryDeviceByNamespaceRequest("80")

	r, err := impl.QueryDeviceByNamespace(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}
