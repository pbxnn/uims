package main

import (
	"github.com/spf13/cobra"
	"uims/cmd/kafka"
)

func initCmdList(rootCmd *cobra.Command) {

	var createKafkaTopic = &cobra.Command{
		Use:   "createKafkaTopic",
		Short: "创建kafka topic",
		Run:   kafka.CreateTopic,
	}

	var deleteKafkaTopic = &cobra.Command{
		Use:   "deleteKafkaTopic",
		Short: "打印consumer group列表",
		Run:   kafka.DeleteTopic,
	}

	var listKafkaTopic = &cobra.Command{
		Use:   "listKafkaTopic",
		Short: "打印kafka topic列表",
		Run:   kafka.ListTopic,
	}

	var listConsumerGroup = &cobra.Command{
		Use:   "listConsumerGroup",
		Short: "打印consumer group列表",
		Run:   kafka.ListConsumerGroups,
	}

	var describeConsumerGroup = &cobra.Command{
		Use:   "getConsumerGroup",
		Short: "打印consumer group列表",
		Run:   kafka.DescribeConsumerGroup,
	}

	var describeTopic = &cobra.Command{
		Use:   "describeTopic",
		Short: "打印consumer group列表",
		Run:   kafka.DescribeTopic,
	}

	rootCmd.AddCommand(
		createKafkaTopic,
		listKafkaTopic,
		listConsumerGroup,
		deleteKafkaTopic,
		describeConsumerGroup,
		describeTopic,
	)
}

func main() {
	rootCmd := &cobra.Command{}
	initCmdList(rootCmd)
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
