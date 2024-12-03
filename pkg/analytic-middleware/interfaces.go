package analytic_middleware

import "github.com/gin-gonic/gin"

type Collector interface {
	GetList(api *gin.RouterGroup)
	Collect(c *gin.Context)
}
