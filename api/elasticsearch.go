package api

import (
	"fmt"
	"time"

	"github.com/grafana/worldping-gw/elasticsearch"
)

func ElasticsearchProxy(c *Context) {
	proxyPath := c.Params("*")
	y, m, d := time.Now().Date()
	idxDate := fmt.Sprintf("%s-%d-%02d-%02d", elasticsearch.IndexName, y, m, d)
	if c.Req.Request.Method == "GET" && proxyPath == fmt.Sprintf("%s/_stats", idxDate) {
		c.JSON(200, "ok")
		return
	}
	if c.Req.Request.Method == "POST" && proxyPath == "_msearch" {
		elasticsearch.Proxy(c.OrgId, c.Context)
		return
	}
	c.JSON(404, "Not Found")
}
