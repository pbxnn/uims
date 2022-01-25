package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/spf13/cobra"
	"log"
)

var kafkaAddress = "152.136.131.96:9092"

func CreateTopic(cmd *cobra.Command, args []string) {
	var topicName string
	if len(args) < 1 {
		log.Panicln("args not enough")
	}
	topicName = args[0]

	log.Println("start craete topic...")
	topicDetail := sarama.TopicDetail{NumPartitions: 1, ReplicationFactor: 1}
	kafkaAdmin := initKafkaAdmin()

	defer func() {
		if err := kafkaAdmin.Close(); err != nil {
			log.Panicln("close admin err:", err)
		}
	}()

	err := kafkaAdmin.CreateTopic(topicName, &topicDetail, false)
	if err != nil {
		log.Panicln("create topic err:", err)
	}
}

func ListTopic(cmd *cobra.Command, args []string) {

	log.Println("start list topic...")

	kafkaAdmin := initKafkaAdmin()
	defer func() {
		if err := kafkaAdmin.Close(); err != nil {
			log.Panicln("close admin err:", err)
		}
	}()

	topics, err := kafkaAdmin.ListTopics()
	if err != nil {
		log.Panicln("create topic err:", err)
	}

	for idx, topic := range topics {
		fmt.Println(idx, topic)
	}
}

func initKafkaAdmin() sarama.ClusterAdmin {
	broker := sarama.NewBroker(kafkaAddress)
	err := broker.Open(nil)
	if err != nil {
		log.Panicln("open broker err:", err)
	}

	config := sarama.NewConfig()
	kafkaAdmin, err := sarama.NewClusterAdmin([]string{kafkaAddress}, config)
	if err != nil {
		log.Panicln("create new cluster err:", err)
	}

	return kafkaAdmin
}
