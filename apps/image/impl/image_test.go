package impl_test

import (
	"testing"

	"github.com/Aaazj/mcenter/apps/image"
)

func TestQueryDeviceByNamespace(t *testing.T) {
	// req := device.NewQueryDeviceByNamespaceRequest("80")

	// r, err := impl.QueryDeviceByNamespace(ctx, req)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	//t.Log(r)
}

func TestCreateImagee(t *testing.T) {
	req := image.NewCreateImageRequest("100", "id6", "name", "kobe2")

	r, err := impl.CreateImage(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}

func TestDeleteImages(t *testing.T) {
	ImagesIds := []string{"id"}
	req := image.NewDeleteImagesRequestWithID()
	req.ImageIds = ImagesIds
	r, err := impl.DeleteImages(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}

func TestQueryImageByNamespace(t *testing.T) {

	req := image.NewQueryImageByNamespaceRequest("100")

	r, err := impl.QueryImageByNamespace(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}
