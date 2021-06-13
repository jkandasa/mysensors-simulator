package start

import (
	"fmt"
	"strings"

	"github.com/jkandasa/mysensors-simulator/pkg/model"
	cfg "github.com/jkandasa/mysensors-simulator/pkg/service/configuration"
	"go.uber.org/zap"
)

func onMessageReceive(message *model.Message) {
	// node-id/child-sensor-id/command/ack/type  payload
	topicSlice := strings.Split(message.Topic, "/")
	if len(topicSlice) < 5 {
		zap.L().Warn("invalid topic", zap.Any("message", message))
		return
	}
	topicSlice = topicSlice[len(topicSlice)-5:]
	payload := message.Data

	// write acknowledgement
	updatedMessage := &model.Message{Data: payload, QoS: message.QoS, Topic: strings.Join(topicSlice, "/")}
	err := DEVICE.Write(updatedMessage)
	if err != nil {
		zap.L().Error("error on sending message", zap.Error(err), zap.Any("message", message))
	}

	if topicSlice[2] == CmdInternal {
		switch topicSlice[4] {
		case InternalPresentation, InternalReboot:
			sendNodeInfo(topicSlice[0])
			return
		case InternalHeartbeatRequest:
			topicSlice[3] = "0"
			topicSlice[4] = InternalHeartbeatResponse
			payload = "123456"
			updatedMessage := &model.Message{Data: payload, QoS: message.QoS, Topic: strings.Join(topicSlice, "/")}
			err := DEVICE.Write(updatedMessage)
			if err != nil {
				zap.L().Error("error on sending message", zap.Error(err), zap.Any("message", message))
			}

		case InternalDiscoverRequest:
			sendAllNodes()
			return

		default:
			// noop
		}
	}

}

func presentNodeInfo(node model.NodeConfig) {
	messages := make([]model.Message, 0)
	// node-id/child-sensor-id/command/ack/type  payload
	// 0;255;3;0;14;Gateway startup complete.
	topic := fmt.Sprintf("%s/255/3/0/14", node.ID)
	messages = append(messages, model.Message{Topic: topic, QoS: 0, Data: "Gateway startup complete"})

	// 0;255;0;0;17;2.3.2
	topic = fmt.Sprintf("%s/255/0/0/17", node.ID)
	messages = append(messages, model.Message{Topic: topic, QoS: 0, Data: node.LibVersion})

	// 0;255;3;0;11;Gateway
	topic = fmt.Sprintf("%s/255/3/0/11", node.ID)
	messages = append(messages, model.Message{Topic: topic, QoS: 0, Data: node.Name})

	// 0;255;3;0;12;1.0.0
	topic = fmt.Sprintf("%s/255/3/0/12", node.ID)
	messages = append(messages, model.Message{Topic: topic, QoS: 0, Data: node.Version})

	for _, sensor := range node.Sensors {
		// 0;0;0;0;3;Tube Light
		topic = fmt.Sprintf("%s/%s/3/0/12", node.ID, sensor.ID)
		messages = append(messages, model.Message{Topic: topic, QoS: 0, Data: sensor.Name})

		for _, field := range sensor.Fields {
			// 0;0;1;0;2;1
			if fieldID, ok := setReqFieldMapForRx[field]; ok {
				topic = fmt.Sprintf("%s/%s/1/0/%s", node.ID, sensor.ID, fieldID)
				messages = append(messages, model.Message{Topic: topic, QoS: 0, Data: ""})
			}
		}
	}

	for _, msg := range messages {
		err := DEVICE.Write(&msg)
		if err != nil {
			zap.L().Error("error on sending message", zap.Error(err), zap.Any("message", msg))
		}
	}

}

func sendNodeInfo(nodeId string) {
	for index := range cfg.CFG.Nodes {
		node := cfg.CFG.Nodes[index]
		if node.ID == nodeId {
			presentNodeInfo(node)
			return
		}
	}
}

func sendAllNodes() {
	for index := range cfg.CFG.Nodes {
		node := cfg.CFG.Nodes[index]
		presentNodeInfo(node)
	}
}
