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
	// 为防止内存溢出，此处调用返回的第二个参数*context.CancelFunc*不能忽略
	ctx, cancel := context.WithCancel(context.Background())

	go executeSubscribe(groupName, topic, ctx, cancel)
}

func executeSubscribe(groupName string, topic string, ctx context.Context, cancel context.CancelFunc) {
	version, err := sarama.ParseKafkaVersion(setting.KafkaSetting.Version)
	if err != nil {
		panic(err)
	}

	config := sarama.NewConfig()
	config.Version = version
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	client, err := sarama.NewConsumerGroup([]string{setting.KafkaSetting.Broker}, groupName, config)
	if err != nil {
		panic(err)
	}

	waitGroup := &sync.WaitGroup{}
	waitGroup.Add(1)

	consumer := Consumer{
		ready: make(chan bool),
	}

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
