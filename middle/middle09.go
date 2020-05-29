package middle09

import (
"strings"

"github.com/gin-gonic/gin"
)

func Default(resp_template map[string]interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if c.Writer.Written() {
			return
		}

		resp := make(map[string]interface{}, len(resp_template))
		for k, v := range resp_template {
			resp[k] = v
		}

		params := c.Keys
		for name, value := range params {
			if !strings.HasPrefix(name, "$.") {
				continue
			}
			data := resp
			name = strings.TrimLeft(name, "$.")
			name_parts := strings.Split(name, ".")
			for index, name_part := range strings.Split(name, ".") {
				if index == len(name_parts)-1 {
					data[name_part] = value

				} else if _, ok := data[name_part]; !ok {
					data[name_part] = make(map[string]interface{})
					data = data[name_part].(map[string]interface{})
				}
			}
		}

		c.JSON(c.Writer.Status(), resp)

	}
}
