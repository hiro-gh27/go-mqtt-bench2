package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	pubsub "github.com/hiro-gh27/go-mqtt-bench2/pubsub"
)

const base = "go-mqtt-bench/"

func main() {
	rand.Seed(time.Now().UnixNano())

	cpus := runtime.NumCPU()
	runtime.GOMAXPROCS(cpus)

	opts := initOption()

	//clients := pubsub.SyncConnect(opts)

	var clients []MQTT.Client
	if opts.AsyncFlag {
		clients = pubsub.AsyncConnect(opts)
		switch opts.Method {
		case "pub":
			pubsub.AsyncPublish(clients, opts)
		case "sub":
		case "RTTpubsub":
		case "RTTconect":
		}
	} else {
		clients = pubsub.SyncConnect(opts)
		switch opts.Method {
		case "pub":
			pubsub.SyncPublish(clients, opts)
		case "sub":
		case "RTTpubsub":
		case "RTTconect":
		}
	}

	pubsub.SyncDisconnect(clients)
	fmt.Println("program is success!!")
}

func initOption() pubsub.ExecOptions {
	broker := flag.String("broker", "tcp://{host}:{port}", "URI of MQTT broker (required)")
	action := flag.String("action", "p|pub or s|sub", "Publish or Subscribe or Subscribe(with publishing) (required)")
	qos := flag.Int("qos", 0, "MQTT QoS(0|1|2)")
	retain := flag.Bool("retain", false, "MQTT Retain")
	topic := flag.String("topic", base, "Base topic")
	clients := flag.Int("clients", 10, "Number of clients")
	count := flag.Int("count", 10, "Number of loops per client")
	size := flag.Int("size", 1024, "Message size per publish (byte)")
	sleepTime := flag.Int("sleep", 3000, "sleep wait time (ms)")
	intervalTime := flag.Int("interval", 0, "Interval time per message (ms)")
	trial := flag.Int("trial", 1, "trial is number of how many loops are")
	//synBacklog := flag.Int("syn", 128, "net.ipv4.tcp_max_syn_backlog = ")
	asyncmode := flag.Bool("async", false, "ture mean asyncmode")
	debug := flag.Bool("x", false, "Debug mode")

	flag.Parse()
	if len(os.Args) < 1 {
		fmt.Println("call here")
		flag.Usage()
		os.Exit(0)
	}

	if broker == nil || *broker == "" || *broker == "tcp://{host}:{port}" {
		fmt.Println("Use Default Broker= tcp://10.0.0.4:1883")
		*broker = "tcp://10.0.0.10:1883"
		//*broker = "tcp://localhost:1883"
	}

	method := ""
	if *action == "p" || *action == "pub" {
		method = "pub"
	} else if *action == "s" || *action == "sub" {
		method = "sub"
	} else if *action == "ps" || *action == "pubsub" {
		method = "singlePubSub"
	}
	if method == "" {
		fmt.Printf("Invalid argument : -action -> %s\n", *action)
		os.Exit(0)
	}

	execOpts := pubsub.ExecOptions{}
	execOpts.Broker = *broker
	execOpts.Qos = byte(*qos)
	execOpts.Retain = *retain
	execOpts.Topic = *topic
	execOpts.MessageSize = *size
	execOpts.ClientNum = *clients
	execOpts.Count = *count
	execOpts.MaxInterval = *intervalTime
	execOpts.SleepTime = *sleepTime
	execOpts.TrialNum = *trial
	//execOpts.SynBacklog = *synBacklog
	execOpts.Method = method
	execOpts.AsyncFlag = *asyncmode
	//fmt.Println(execOpts.Topic)

	execOpts.Test = false
	execOpts.Debug = *debug

	return execOpts
}
