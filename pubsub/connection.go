package pubsub

import (
	"fmt"
	"os"
	"strconv"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

// SyncConnect is
func SyncConnect(execOpts ExecOptions) []MQTT.Client {
	var clients []MQTT.Client
	for id := 0; id < execOpts.ClientNum; id++ {
		prosessID := strconv.FormatInt(int64(os.Getpid()), 16)
		clientID := fmt.Sprintf("go-mqtt-bench%s-%d", prosessID, id)
		opts := MQTT.NewClientOptions()
		opts.AddBroker(execOpts.Broker)
		opts.SetClientID(clientID)
		client := MQTT.NewClient(opts)
		token := client.Connect()
		if token.Wait() && token.Error() != nil {
			fmt.Printf("Connected error: %s\n", token.Error())
			return clients
		}
		clients = append(clients, client)
	}
	return clients
}

func asyscConnect() {

}

// SyncDisconnect is
func SyncDisconnect(clinets []MQTT.Client) {
	for _, clinet := range clinets {
		clinet.Disconnect(250)
	}
}
