package v3

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

const (
	LANDING = `
	___ ___ ___ ___     _   ___ ___ 
   |_ _/ __|_ _/ __|   /_\ | _ \_ _|
	| | (_ || |\__ \  / _ \|  _/| | 
   |___\___|___|___/ /_/ \_\_| |___|
									`
)

type IdpWebServer struct {
	GinCon   *gin.Engine
	DataBase *redis.Client
}

type IdpWebServe interface {
	// landing page methods
	StartServer() IdpWebServer
	TestActive() IdpWebServer
	Version() IdpWebServer
}

type IdpDataMacro struct {
	FactSheet   string `json:"fsht"`
	Description string `json:"desc"`
	LastUpdate  string `json:"last"`
	Data        macros `json:"data"`
}

type macros struct {
	KR1Y      macroRows `json:"kr1y"`
	KR3Y      macroRows `json:"kr3y"`
	KR5Y      macroRows `json:"kr5y"`
	IFD1Y     macroRows `json:"ifd1y"`
	CD91D     macroRows `json:"cd91d"`
	CP91D     macroRows `json:"cp91d"`
	FB6M      macroRows `json:"fb6m"`
	FB1Y      macroRows `json:"fb1y"`
	FB3Y      macroRows `json:"fb3y"`
	KORIBOR3M macroRows `json:"koribor3m"`
	Feds      macroRows `json:"ffr"`
}

type macroRows []macroRow

type macroRow struct {
	Date  string  `json:"date"`
	Value float64 `json:"value"`
}
