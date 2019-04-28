package mqtt

import (
	"fmt"
	"testing"
	"time"

	"github.com/eclipse/paho.mqtt.golang"
	logger "github.com/sirupsen/logrus"
)

func TestSub(t *testing.T) {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("%s://%s:%s", "tcp", "localhost", "1883"))
	//opts.SetClientID("test-clientID")
	opts.SetKeepAlive(time.Second * time.Duration(60))
	// If lost connection, reconnect again
	opts.SetConnectionLostHandler(func(client mqtt.Client, e error) {
		logger.Warn(fmt.Sprintf("Connection lost : %v", e))
		if client.IsConnected() {
			client.Disconnect(500)
		}
		client = mqtt.NewClient(opts)
		connect(client)
		subscribe(client)
	})

	client := mqtt.NewClient(opts)

	connect(client)
	subscribe(client)

	logger.Info("Start to subscribe....")

	defer func() {
		if r := recover(); r != nil {
			logger.Warn("unknown panic error, try to recover connection,", r)
			connect(client)
			subscribe(client)
		}
	}()

	select {}
}

func connect(client mqtt.Client) {
	for {
		// create connection
		token := client.Connect()
		if token.Wait() && token.Error() != nil {
			logger.Errorf("Fail to connect broker, %v", token.Error())
			time.Sleep(5 * time.Second)

			logger.Error("Retry the connection")
			continue
		} else {
			logger.Error("Reconnection successful!")
			break
		}
	}
}

func subscribe(client mqtt.Client) {
	for {
		// subscribe the topic, "#" means all topics
		token := client.Subscribe("#", byte(0), onIncomingDataReceived)
		if token.Wait() && token.Error() != nil {
			logger.Error("Fail to sub... ", token.Error())
			time.Sleep(5 * time.Second)

			logger.Errorf("Retry to subscribe")
			continue
		} else {
			logger.Info("Subscribe successful!")
			break
		}
	}
}

func onIncomingDataReceived(client mqtt.Client, message mqtt.Message) {
	logger.Info(message.Topic(), " ", string(message.Payload()))
}
