package main

import (
	"io"
	"net/http"
	"os"

	"github.com/curtisnewbie/miso/miso"
	"github.com/gin-gonic/gin"
)

func init() {
	miso.SetDefProp(miso.PropAppName, "echo-server")
}

func main() {
	miso.RawAny("/echo", func(c *gin.Context, rail miso.Rail) {
		rail.Infof("Receive request from %v", c.Request.RemoteAddr)
		rail.Infof("Method: %v, Content-Length: %v", c.Request.Method, c.Request.ContentLength)

		body, e := io.ReadAll(c.Request.Body)
		if e != nil {
			rail.Errorf("Failed to read request body, %v", e)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		rail.Info("Headers: ")
		for k, v := range c.Request.Header {
			rail.Infof("  %-30s: %v", k, v)
		}

		rail.Info("")
		rail.Info("Body: ")
		rail.Infof("  %s", string(body))
		rail.Info("")
		rail.Info("Requested processed\n")
	})
	miso.BootstrapServer(os.Args)
}
