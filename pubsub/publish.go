package pubsub

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

var messageSize int
var maxIntarval int
var qos byte
var pid string
var baseTopic string
var trial int
var count int

func initPubOpts(opts PublishOptions2) {
	baseTopic = opts.Topic
	count = opts.Count
	messageSize = opts.MessageSize
	pid = strconv.FormatInt(int64(os.Getpid()), 16)
	maxIntarval = maxInterval
	trial = opts.TrialNum
	qos = opts.Qos
}

// "sync publish"
func spub(pOpts PublishOptions) PublishResult {
	var pResult PublishResult
	message := RandomMessage(messageSize)
	clientID := fmt.Sprintf("%s-%d", pid, pOpts.ID)
	topic := fmt.Sprintf(baseTopic+"%s"+"/"+"%d", clientID, pOpts.TrialNum)
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
func SyncPublish(opts PublishOptions2) {
	initPubOpts(opts)
	var pResults []PublishResult
	for index := 0; index < opts.Count; index++ {
		for id := 0; id < len(opts.Clients); id++ {
			var pOpts PublishOptions
			pOpts.Client = opts.Clients[id]
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

	message := RandomMessage(messageSize)
	clientID := fmt.Sprintf("%s-%d", pid, pOpts.ID)
	topic := fmt.Sprintf(baseTopic+"%s"+"/"+"%d", clientID, 0)
	if maxIntarval > 0 {
		waitTime = RandomInterval(maxIntarval)
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
		message = RandomMessage(messageSize)
		topic = fmt.Sprintf(baseTopic+"%s"+"/"+"%d", clientID, index)
		if maxIntarval > 0 {
			waitTime = RandomInterval(maxIntarval)
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
func AsyncPublish(opts PublishOptions2) {
	initPubOpts(opts)
	var pResults []PublishResult
	wg := &sync.WaitGroup{}
	freeze := &sync.WaitGroup{}
	freeze.Add(1)
	for id := 0; id < len(opts.Clients); id++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			var pOpts PublishOptions
			pOpts.Client = opts.Clients[id]
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
