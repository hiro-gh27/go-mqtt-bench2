package pubsub

import (
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

// ConnectResult is
type ConnectResult struct {
	StartTime time.Time     // when trial for connect start
	EndTime   time.Time     // when trial for connect end
	DurTime   time.Duration // =[endtime - starttime]
	Client    MQTT.Client   //client
	ClientID  string        // identification times
}

// PublishResult is
type PublishResult struct {
}

// SubscribeResult is
type SubscribeResult struct {
}
