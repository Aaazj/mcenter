package instance_test

import (
	"testing"

	"github.com/Aaazj/mcenter/apps/instance"
)

func TestParseStrLable(t *testing.T) {
	lables := instance.ParseStrLable("k1=v1,k2=v2")
	t.Log(lables)
}
