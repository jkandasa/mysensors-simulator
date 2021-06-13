package start

import (
	cfg "github.com/jkandasa/mysensors-simulator/pkg/service/configuration"
	"github.com/jkandasa/mysensors-simulator/pkg/version"
	loggerUtils "github.com/mycontroller-org/backend/v2/pkg/utils/logger"
	"go.uber.org/zap"
)

// InitLogger func
func InitLogger() {
	logger := loggerUtils.GetLogger(cfg.CFG.Logger.Mode, cfg.CFG.Logger.Level, cfg.CFG.Logger.Encoding, false, 0)
	zap.ReplaceGlobals(logger)
	zap.L().Info("welcome to the mysensors-simulator :)")
	zap.L().Info("server detail", zap.Any("version", version.Get()), zap.Any("loggerConfig", cfg.CFG.Logger))
}
