package main

import (
	"io"
	"net/http"
	"os"

	"github.com/curtisnewbie/miso/miso"
	"github.com/gin-gonic/gin"
)

func main() {
	miso.SetDefProp(miso.PropAppName, "echo-server")
	miso.SetDefProp(miso.PropServerPort, 8080)
	miso.ManualBootstrapPrometheus()

	miso.RawAny("/*proxy", func(c *gin.Context, rail miso.Rail) {
		rail.Infof("Receive '%v %v' request from %v", c.Request.Method, c.Request.RequestURI, c.Request.RemoteAddr)
		rail.Infof("Content-Length: %v", c.Request.ContentLength)
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
	})
	miso.BootstrapServer(os.Args)
}
