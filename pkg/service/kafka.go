package service

import (
	"time"

	"github.com/NetEase-Media/ngo/adapter/log"
	"github.com/NetEase-Media/ngo/client/kafka"
)

// Product 将数据放入kafka
func Product(message, sync string) error {
	begin := time.Now()
	if sync != "" {
		err := producer.SyncSend("test", message)
		log.Infof("get result cost %v", time.Since(begin))
		if err != nil {
			log.Errorf("failed to send message, error: %v", err)
		}
	} else {
		producer.Send("test", message, func(meta *kafka.RecordMetadata, err error) {
			if err != nil {
				log.Errorf("failed to send message, error: %v", err)
			}
			log.Infof("get result cost %v", time.Since(begin))
		})
	}
	log.Infof("send cost %v", time.Since(begin))
	return nil
}

// Consume 消费者消费数据
func Consume() {
	consumer.AddListener("test", &listener{handle})
	consumer.Start()
}

func handle(message kafka.ConsumerMessage, acknowledgment *kafka.Acknowledgment) {
	log.Infof("消费topic: [test], message: %v", message)
	acknowledgment.Acknowledge()
}

type listener struct {
	listenFn func(message kafka.ConsumerMessage, ack *kafka.Acknowledgment)
}

func (l *listener) Listen(message kafka.ConsumerMessage, ack *kafka.Acknowledgment) {
	l.listenFn(message, ack)
}
