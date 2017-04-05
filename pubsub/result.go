package pubsub

import (
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

// ReturnConnect is
type ReturnConnect struct {
	StartTime time.Time   // when trial for connect start
	EndTime   time.Time   // when trial for connect end
	RTT       time.Time   // =(endtime - starttime)
	Client    MQTT.Client //client
	ClientID  string      // identification times
}

// ReturnPublish is
type ReturnPublish struct {
}

// ReturnSubscribe is
type ReturnSubscribe struct {
}
