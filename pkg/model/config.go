package model

// Config
type Config struct {
	Logger LoggerConfig `yaml:"logger"`
	MQTT   MQTTConfig   `yaml:"mqtt"`
	Nodes  []NodeConfig `yaml:"nodes"`
}

// LoggerConfig struct
type LoggerConfig struct {
	Mode     string `yaml:"mode"`
	Encoding string `yaml:"encoding"`
	Level    string `yaml:"level"`
}

// MQTTConfig struct
type MQTTConfig struct {
	Broker             string `yaml:"broker"`
	InsecureSkipVerify bool   `yaml:"insecure_skip_verify"`
	Username           string `yaml:"username"`
	Password           string `yaml:"password" json:"-"`
	Subscribe          string `yaml:"subscribe"`
	Publish            string `yaml:"publish"`
	QoS                int    `yaml:"qos"`
	TransmitPreDelay   string `yaml:"transmit_pre_delay"`
	ReconnectDelay     string `yaml:"reconnect_delay"`
}

type NodeConfig struct {
	ID         string         `yaml:"id"`
	Name       string         `yaml:"name"`
	Version    string         `yaml:"version"`
	LibVersion string         `yaml:"lib_version"`
	IsRepeater bool           `yaml:"is_repeater"`
	Sensors    []SensorConfig `yaml:"sensors"`
}

type SensorConfig struct {
	ID     string   `yaml:"id"`
	Name   string   `yaml:"name"`
	Fields []string `yaml:"fields"`
}

type Message struct {
	Topic string
	Data  string
	QoS   byte
}

// Device interface
type Device interface {
	Close() error
	Write(message *Message) error
}
