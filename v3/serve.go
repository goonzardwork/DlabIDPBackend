package v3

import (
	"DlabIDPBackend/orm"
	"log"

	"github.com/gin-gonic/gin"
)

func (w IdpWebServer) StartServer() IdpWebServer {
	w.GinCon = gin.Default()
	db, err := orm.Conn("./token.json")
	if err != nil {
		log.Panicln(err)
	} else {
		w.DataBase = db
	}
	w.GinCon.GET("/", IdpLanding)
	return w
}

func (w IdpWebServer) TestActive() IdpWebServer {
	w.GinCon.GET("/ping", IdpPingPong)
	return w
}

func (w IdpWebServer) Version() IdpWebServer {
	w.GinCon.GET("/version", IdpVersion)
	return w
}

func (w IdpWebServer) Macro() IdpWebServer {
	w.GinCon.GET("/api/v3/macro/dataTable", func(c *gin.Context) {
		IdpGetMacro(w.DataBase, c)
	})
	return w
}
