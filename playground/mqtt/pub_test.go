package mqtt

import (
	"fmt"
	"testing"
	"time"

	"github.com/eclipse/paho.mqtt.golang"
	logger "github.com/sirupsen/logrus"
)

func TestPub(t *testing.T) {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("%s://%s:%s", "tcp", "localhost", "1883"))
	//opts.SetClientID("test-clientID")
	opts.SetKeepAlive(time.Second * time.Duration(60))
	// If lost connection, reconnect again
	opts.SetConnectionLostHandler(func(client mqtt.Client, e error) {
		logger.Warn(fmt.Sprintf("Connection lost : %v", e))
	})

	// connect to broker
	client := mqtt.NewClient(opts)
	token := client.Connect()
	if token.Wait() && token.Error() != nil {
		logger.Fatalf("Fail to connect broker, %v", token.Error())
	}

	// publish to topic
	token = client.Publish("TestTopic", byte(0), false, "123")
	if token.Wait() && token.Error() != nil {
		logger.Errorf("Fail to publish, %v", token.Error())
	}
}
