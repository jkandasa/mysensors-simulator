package start

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	cfg "github.com/jkandasa/mysensors-simulator/pkg/service/configuration"

	mqttDriver "github.com/jkandasa/mysensors-simulator/pkg/client"
	"github.com/jkandasa/mysensors-simulator/pkg/model"
	"go.uber.org/zap"
)

var DEVICE model.Device

func Start() {

	start := time.Now()

	cfg.InitConfig()
	InitLogger()

	// start mqtt services
	mqttDevice, err := mqttDriver.New("simulator", &cfg.CFG.MQTT, onMessageReceive)
	if err != nil {
		zap.L().Fatal("error on connection to mqtt", zap.Error(err))
		return
	}
	DEVICE = mqttDevice

	zap.L().Info("services started", zap.String("timeTaken", time.Since(start).String()))

	handleShutdownSignal()
}

// handleShutdownSignal func
func handleShutdownSignal() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// waiting for signal
	sig := <-sigs
	close(sigs)

	zap.L().Info("shutdown initiated..", zap.Any("signal", sig))
	triggerShutdown()
}

func triggerShutdown() {
	start := time.Now()

	if DEVICE != nil {
		DEVICE.Close()
	}

	zap.L().Debug("closing services are done", zap.String("timeTaken", time.Since(start).String()))
	zap.L().Debug("bye, see you soon :)")

	os.Exit(0)
}
