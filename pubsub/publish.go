package pubsub

import (
	"fmt"
	"os"
	"strconv"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var messagesize int
var pubintarval int
var qos byte
var pid string
var basetopic string

func initPubOpts(opts ExecOptions) {
	pid = strconv.FormatInt(int64(os.Getpid()), 16)
	messagesize = opts.MessageSize
	qos = opts.Qos
	pubintarval = maxInterval
	basetopic = opts.Topic
}

// "sync publish"
func spub(pOpts PublishOptions) PublishResult {
	var pResult PublishResult
	message := RandomMessage(messagesize)
	clientID := fmt.Sprintf("%s-%d", pid, pOpts.ID)
	topic := fmt.Sprintf(basetopic+"%s"+"/"+"%d", clientID, pOpts.TrialNum)
	startTime := time.Now()
	token := pOpts.Client.Publish(topic, qos, false, message)
	token.Wait()
	endTime := time.Now()
	//fmt.Println(basetopic)

	pResult.StartTime = startTime
	pResult.EndTime = endTime
	pResult.DurTime = endTime.Sub(startTime)
	pResult.Topic = topic
	pResult.ClientID = clientID
	fmt.Printf("### dtime=%s, clientID=%s, topic=%s ###\n",
		pResult.DurTime, pResult.ClientID, pResult.Topic)
	return pResult
}

// SyncPublish is
func SyncPublish(clients []MQTT.Client, opts ExecOptions) {
	initPubOpts(opts)
	var pResults []PublishResult
	for index := 0; index < opts.Count; index++ {
		for id := 0; id < len(clients); id++ {
			var pOpts PublishOptions
			pOpts.Client = clients[id]
			pOpts.ID = id
			pOpts.TrialNum = index
			pr := spub(pOpts)
			pResults = append(pResults, pr)
		}
	}
}

// "async publish""
func aspub() []PublishResult {
	var pResults []PublishResult
	return pResults
}

// AsyncPublish is
func AsyncPublish(clients []MQTT.Client, opts ExecOptions) {
	//prosessID := strconv.FormatInt(int64(os.Getpid()), 16)
	//clientID := fmt.Sprintf("%s-%d", prosessID, id)
	//fmt.Printf("In asyncpublish prosessID is: %s\n", prosessID)
	//wg := &sync.WaitGroup{}
	for index := 0; index < opts.TrialNum; index++ {
		for id := 0; id < opts.ClientNum; id++ {

		}
	}
}

// LoadPublish is
func LoadPublish() {

}
