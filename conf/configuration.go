package conf

import (
	"log"

	"go.uber.org/zap"
)

// Cfg configuration object
type Cfg struct {
	Logger *zap.Logger
}

var _Cfg Cfg

// _Logger logger object
var _Logger *zap.Logger

//	initialize configuration object
func init() {
	_Logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize logger:%v", err)
	}
	_Cfg.Logger = _Logger
}

// GetCfg get a logger
func GetCfg() (cfg *Cfg) {
	cfg = &_Cfg
	return
}
