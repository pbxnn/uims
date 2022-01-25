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

	var listKafkaTopic = &cobra.Command{
		Use:   "listKafkaTopic",
		Short: "打印kafka topic列表",
		Run:   kafka.ListTopic,
	}

	rootCmd.AddCommand(createKafkaTopic, listKafkaTopic)
}

func main() {
	rootCmd := &cobra.Command{}
	initCmdList(rootCmd)
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
