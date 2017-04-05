package pubsub

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func connect(id int, broker string) ConnectResult {
	var cRresult ConnectResult
	prosessID := strconv.FormatInt(int64(os.Getpid()), 16)
	//clientID := fmt.Sprintf("go-mqtt-bench%s-%d", prosessID, id)
	clientID := fmt.Sprintf("%s-%d", prosessID, id)
	opts := MQTT.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID(clientID)
	client := MQTT.NewClient(opts)

	startTime := time.Now()
	token := client.Connect()
	if token.Wait() && token.Error() != nil {
		fmt.Printf("Connected error: %s\n", token.Error())
		client = nil
	}
	endTime := time.Now()

	cRresult.StartTime = startTime
	cRresult.EndTime = endTime
	cRresult.DurTime = endTime.Sub(startTime)
	cRresult.Client = client
	cRresult.ClientID = clientID
	return cRresult
}

// MQTT.clinet=nilに対してdisconnect要求するとpanicに陥るので, 中身があるかどうかをチェックする必要がある.
func iscompleat(results []ConnectResult) ([]MQTT.Client, bool) {
	var clietns []MQTT.Client
	haserr := false
	for _, r := range results {
		if r.Client != nil {
			clietns = append(clietns, r.Client)
		} else {
			haserr = true
		}
	}
	return clietns, haserr
}

// SyncConnect is
func SyncConnect(execOpts ExecOptions) []MQTT.Client {
	var cResults []ConnectResult
	broker := execOpts.Broker
	for id := 0; id < execOpts.ClientNum; id++ {
		r := connect(id, broker)
		cResults = append(cResults, r)
	}
	clients, haserr := iscompleat(cResults)
	if haserr {
		SyncDisconnect(clients)
		os.Exit(0)
	}
	/*
		TODO
		   export ElasticSearch
	*/
	CDebug(cResults)
	return clients
}

// AsyncConnect is
func AsyncConnect(execOpts ExecOptions) []MQTT.Client {
	var cResults []ConnectResult
	wg := sync.WaitGroup{}
	broker := execOpts.Broker
	for id := 0; id < execOpts.ClientNum; id++ {
		wg.Add(1)
		go func(id int) {
			r := connect(id, broker)
			cResults = append(cResults, r)
			wg.Done()
		}(id)
	}
	wg.Wait()

	clients, haserr := iscompleat(cResults)
	if haserr {
		SyncDisconnect(clients)
		os.Exit(0)
	}

	/*
		TODO
			export ElasticSearch
	*/
	CDebug(cResults)
	return clients
}

// SyncDisconnect is
func SyncDisconnect(clinets []MQTT.Client) {
	for _, c := range clinets {
		c.Disconnect(250)
	}
}

// AsyncDisconnect is
func AsyncDisconnect(clients []MQTT.Client) {
	wg := sync.WaitGroup{}
	for _, c := range clients {
		wg.Add(1)
		go func(c MQTT.Client) {
			c.Disconnect(250)
			wg.Done()
		}(c)
	}
	wg.Wait()
}

// LoadConnect is
func LoadConnect() {

}
