package kratos

import (
	"testing"
	"time"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

func TestApp(t *testing.T) {
	hs := http.NewServer()
	gs := grpc.NewServer()
	app := New(
		Name("kratos"),
		Version("v1.0.0"),
		Server(hs, gs),
	)
	ctx, done := OnInterrupt()
	time.AfterFunc(time.Second, func() {
		done()
	})
	if err := app.Run(ctx); err != nil {
		t.Fatal(err)
	}
}
