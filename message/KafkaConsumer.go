package message

import (
	"context"
	"mojiayi-go-journey/setting"
	"sync"

	"github.com/Shopify/sarama"
)

type KafkaConsumer struct {
}

func Subscribe(groupName string, topic string) {
	version, err := sarama.ParseKafkaVersion(setting.KafkaSetting.Version)
	if err != nil {
		panic(err)
	}

	config := sarama.NewConfig()
	config.Version = version
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	consumer := Consumer{
		ready: make(chan bool),
	}

	ctx, _ := context.WithCancel(context.Background())
	client, err := sarama.NewConsumerGroup([]string{setting.KafkaSetting.Broker}, groupName, config)
	if err != nil {
		panic(err)
	}

	waitGroup := &sync.WaitGroup{}
	waitGroup.Add(1)

	go func() {
		defer waitGroup.Done()

		for {
			err := client.Consume(ctx, []string{topic}, &consumer)
			if err != nil {
				setting.MyLogger.Error("消费者发生错误,", err)
			}
			consumer.ready = make(chan bool)
		}
	}()

	<-consumer.ready
	setting.MyLogger.Info("启动sarama消费者")
}

type Consumer struct {
	ready chan bool
}

func (c *Consumer) Setup(sarama.ConsumerGroupSession) error {
	close(c.ready)
	return nil
}

func (c *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (c *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message := <-claim.Messages():
			{
				setting.MyLogger.Info("收到消息=", string(message.Value))
				session.MarkMessage(message, "")
			}

		case <-session.Context().Done():
			{
				return nil
			}
		}
	}
}
