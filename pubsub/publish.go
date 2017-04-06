package pubsub

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var messagesize int
var pubintarval int
var qos byte
var pid string
var basetopic string
var trial int
var count int

func initPubOpts(opts ExecOptions) {
	basetopic = opts.Topic
	count = opts.Count
	messagesize = opts.MessageSize
	pid = strconv.FormatInt(int64(os.Getpid()), 16)
	pubintarval = maxInterval
	trial = opts.TrialNum
	qos = opts.Qos
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
func aspub(pOpts PublishOptions, freeze *sync.WaitGroup) []PublishResult {
	var pResults []PublishResult
	var waitTime time.Duration

	message := RandomMessage(messagesize)
	clientID := fmt.Sprintf("%s-%d", pid, pOpts.ID)
	topic := fmt.Sprintf(basetopic+"%s"+"/"+"%d", clientID, 0)
	if pubintarval > 0 {
		waitTime = RandomInterval(pubintarval)
	}
	fmt.Print("ready!!")
	freeze.Wait()
	if waitTime > 0 {
		time.Sleep(waitTime)
	}
	fmt.Print("go!!")

	starttime := time.Now()
	token := pOpts.Client.Publish(topic, qos, false, message)
	token.Wait()
	endTime := time.Now()

	var vals PublishResult
	vals.StartTime = starttime
	vals.EndTime = endTime
	vals.DurTime = endTime.Sub(starttime)
	vals.Topic = topic
	vals.ClientID = clientID
	pResults = append(pResults, vals)

	for index := 1; index < count; index++ {
		message = RandomMessage(messagesize)
		topic = fmt.Sprintf(basetopic+"%s"+"/"+"%d", clientID, index)
		if pubintarval > 0 {
			waitTime = RandomInterval(pubintarval)
			time.Sleep(waitTime)
		}
		starttime := time.Now()
		token := pOpts.Client.Publish(topic, qos, false, message)
		token.Wait()
		endTime := time.Now()

		vals = PublishResult{}
		vals.StartTime = starttime
		vals.EndTime = endTime
		vals.DurTime = endTime.Sub(starttime)
		vals.Topic = topic
		vals.ClientID = clientID
		pResults = append(pResults, vals)
	}
	return pResults
}

// AsyncPublish is
func AsyncPublish(clients []MQTT.Client, opts ExecOptions) {
	initPubOpts(opts)
	var pResults []PublishResult
	wg := &sync.WaitGroup{}
	freeze := &sync.WaitGroup{}
	freeze.Add(1)
	for id := 0; id < len(clients); id++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			var pOpts PublishOptions
			pOpts.Client = clients[id]
			pOpts.ID = id
			re := aspub(pOpts, freeze)
			pResults = append(pResults, re...)
		}(id)
	}
	time.Sleep(3 * time.Second)
	freeze.Done()
	wg.Wait()

	for _, vals := range pResults {
		fmt.Printf("### dtime=%s, clientID=%s, topic=%s ###\n",
			vals.DurTime, vals.ClientID, vals.Topic)
	}
}

// LoadPublish is
func LoadPublish() {

}
