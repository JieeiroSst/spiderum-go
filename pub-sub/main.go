package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type DataEvent struct {
	Data interface{}
	Topic string
}

type DataChannel chan DataEvent
type DataChannelSlice []DataChannel

type EvenBus struct {
	subscriber map[string]DataChannelSlice
	rm         sync.RWMutex
}

func (eb *EvenBus) Publish(topic string,data interface{}){
	eb.rm.RLock()
	defer eb.rm.RUnlock()
	if chans,found:=eb.subscriber[topic];found{
		channels:=append(DataChannelSlice{},chans...)
		go func(data DataEvent,dataChannelSlice DataChannelSlice) {
			for _,ch:=range dataChannelSlice{
				ch<- data
			}
		}(DataEvent{Data:data,Topic:topic},channels)
	}
}

func (eb *EvenBus) Subscribe(topic string,ch DataChannel){
	eb.rm.RLock()
	defer eb.rm.RUnlock()
	if prev,found := eb.subscriber[topic];found{
		eb.subscriber[topic] =append(prev,ch)
	}else {
		eb.subscriber[topic] = append([]DataChannel{},ch)
	}
}

var eb = &EvenBus{
	subscriber: map[string]DataChannelSlice{},
}


func printDataEvent(ch string, data DataEvent) {
	fmt.Printf("Channel: %s; Topic: %s; DataEvent: %v\n", ch, data.Topic, data.Data)
}

func publisTo(topic string, data string) {
	for {
		eb.Publish(topic, data)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}

func main(){
	ch1 := make(chan DataEvent,2)
	ch2 := make(chan DataEvent,2)
	ch3 := make(chan DataEvent,2)

	eb.Subscribe("topic1", ch1)
	eb.Subscribe("topic2", ch2)
	eb.Subscribe("topic3", ch3)

	go publisTo("topic1", "Hi topic 1")
	go publisTo("topic2", "Welcome to topic 2")
	go publisTo("topic3", "Welcome to topic 3")

	for {
		select {
		case d := <-ch1:
			go printDataEvent("ch1", d)
		case d := <-ch2:
			go printDataEvent("ch2", d)
		case d := <-ch3:
			go printDataEvent("ch3", d)
		}
	}
}

