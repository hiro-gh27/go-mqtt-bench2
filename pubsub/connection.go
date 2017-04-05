package pubsub

import (
	"fmt"
	"os"
	"strconv"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func connect(id int, broker string) MQTT.Client {
	prosessID := strconv.FormatInt(int64(os.Getpid()), 16)
	clientID := fmt.Sprintf("go-mqtt-bench%s-%d", prosessID, id)
	opts := MQTT.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID(clientID)
	client := MQTT.NewClient(opts)
	token := client.Connect()
	if token.Wait() && token.Error() != nil {
		fmt.Printf("Connected error: %s\n", token.Error())
		client = nil
	}
	return client
}

func iscompleat(clients []MQTT.Client) bool {
	for _, client := range clients {
		if client == nil {
			return false
		}
	}
	return true
}

// SyncConnect is
func SyncConnect(execOpts ExecOptions) []MQTT.Client {
	var clients []MQTT.Client
	broker := execOpts.Broker
	for id := 0; id < execOpts.ClientNum; id++ {
		client := connect(id, broker)
		clients = append(clients, client)
	}
	if iscompleat(clients) == false {
		var rClient []MQTT.Client
		for _, client := range clients {
			if client != nil {
				rClient = append(rClient, client)
			}
		}
		SyncDisconnect(rClient)
		os.Exit(0)
	}
	return clients
}

// AsyscConnect is
/*
func AsyscConnect(execOpts ExecOptions) []MQTT.Client {
	var clients []MQTT.Client
	broker := execOpts.Broker
	for id := 0; id < execOpts.ClientNum; id++ {

	}
}
*/

// SyncDisconnect is
func SyncDisconnect(clinets []MQTT.Client) {
	for _, clinet := range clinets {
		clinet.Disconnect(250)
	}
}

// AsyncDisconnect is
func AsyncDisconnect() {

}

// LoadConnect is
func LoadConnect() {

}
