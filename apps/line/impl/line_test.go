package impl_test

import (
	"testing"

	"github.com/Aaazj/mcenter/apps/line"
)

func TestQueryDeviceByNamespace(t *testing.T) {
	// req := device.NewQueryDeviceByNamespaceRequest("80")

	// r, err := impl.QueryDeviceByNamespace(ctx, req)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	//t.Log(r)
}

func TestCreateLine(t *testing.T) {
	req := line.NewCreateLineRequest("100", "id", "name1", "kobe2")

	r, err := impl.CreateLine(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}

func TestDeleteLines(t *testing.T) {
	LineIds := []string{"id19"}
	req := line.NewDeleteLinesRequestWithID()
	req.Ids = LineIds
	r, err := impl.DeleteLines(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}

func TestQueryLineByNamespace(t *testing.T) {

	req := line.NewQueryLineByNamespaceRequest("100")

	r, err := impl.QueryLineByNamespace(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}
