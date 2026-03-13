package echo

import (
	"fmt"
	"testing"
	"time"

	"github.com/omeid/conex"
	"github.com/omeid/echo"
	echoHttp "github.com/omeid/echo/http"
)

var (
	// Image to use for the box.
	Image = "omeid/echo:http"
	// Port used for connecting to the echo service.
	Port = "3000"

	// EchoUpWaitTime dictates how long we should wait for Echo to accept connections.
	EchoUpWaitTime = 10 * time.Second
)

func init() {
	conex.Require(func() string { return Image })
}

// Box returns an echo client connect to an echo container based on
// your provided tags.
func Box(t testing.TB, reverse bool) (echo.Echo, conex.Container) {
	params := []string{}

	if reverse {
		params = append(params, "-reverse")
	}

	c := conex.Box(t, &conex.Config{
		Image:  Image,
		Cmd:    params,
		Expose: []string{Port},
	})

	t.Log("Waiting for Echo to accept connections")

	err := c.Wait(Port, EchoUpWaitTime)
	if err != nil {
		c.Drop()
		t.Fatal("Echo failed to start:", err)
	}

	t.Log("Echo is now accepting connections")

	addr := fmt.Sprintf("http://%s:%s", c.Address(), Port)

	e, err := echoHttp.NewClient(addr)
	if err != nil {
		c.Drop()
		t.Fatal(err)
	}

	return e, c
}
