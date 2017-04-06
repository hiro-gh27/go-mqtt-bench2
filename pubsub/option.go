package pubsub

import (
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

// ExecOptions is
type ExecOptions struct {
	Broker      string // Broker URI
	Qos         byte   // QoS(0|1|2)
	Retain      bool   // Retain
	Debug       bool   //デバック
	Topic       string // Topicのルート
	Method      string // 実行メソッド
	ClientNum   int    // クライアントの同時実行数
	Count       int    // 1クライアント当たりのメッセージ数
	MessageSize int    // 1メッセージのサイズ(byte)
	SleepTime   int    // 実行前の待機時間(ms)
	MaxInterval int    // メッセージ毎の実行間隔時間(ms)
	Test        bool   //テスト
	TrialNum    int    //試行回数
	SynBacklog  int    //net.ipv4.tcp_max_syn_backlog =
	AsyncFlag   bool   //ture mean asyncmode
}

// LoadOptions is
type LoadOptions struct {
}

// PublishOptions is
type PublishOptions struct {
	Client   MQTT.Client
	ID       int
	TrialNum int
	Count    int

	ProsessID          string
	MessageSize        int
	MaxPublishIntarval int
	Qos                byte
}
