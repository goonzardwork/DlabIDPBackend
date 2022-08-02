package v3

import (
	"DlabIDPBackend/orm"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func IdpGetMacro(database *redis.Client, c *gin.Context) {
	var result []IdpDataMacro
	d, err := orm.JSONGet(database, orm.MACRO_KEY, orm.MACRO_PATH_ALL, &result)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusForbidden, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, d)
}
