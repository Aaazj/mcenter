package refresh_test

import (
	"context"
	"testing"

	"github.com/Aaazj/mcenter/apps/token"
	"github.com/Aaazj/mcenter/apps/token/provider"
	"github.com/Aaazj/mcenter/test/tools"
)

var (
	impl provider.TokenIssuer
	ctx  = context.Background()
)

func TestIssueToken(t *testing.T) {
	req := token.NewRefreshIssueTokenRequest("GlwoJphMPJjTajgSFM0oScFl", "rZbeCx21XsiS7O1BfCnqvO4W63GdgYT6")
	tk, err := impl.IssueToken(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tk.JsonFormat())
}

func init() {
	tools.DevelopmentSetup()
	impl = provider.GetTokenIssuer(token.GRANT_TYPE_REFRESH)
}
