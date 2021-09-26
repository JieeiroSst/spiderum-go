package rabbitmq

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"github.com/streadway/amqp"
	"gitlab.com/Spide_IT/spide_it/email/config"
	email "gitlab.com/Spide_IT/spide_it/email/internal"
	"log"
)

type EmailsConsumer struct {
	amqpConn *amqp.Connection
	emailUC  email.EmailsUseCase
	cfg  *config.Config
}

const (
	exchangeKind       = "direct"
	exchangeDurable    = true
	exchangeAutoDelete = false
	exchangeInternal   = false
	exchangeNoWait     = false

	queueDurable    = true
	queueAutoDelete = false
	queueExclusive  = false
	queueNoWait     = false

	publishMandatory = false
	publishImmediate = false

	prefetchCount  = 1
	prefetchSize   = 0
	prefetchGlobal = false

	consumeAutoAck   = false
	consumeExclusive = false
	consumeNoLocal   = false
	consumeNoWait    = false
)

func NewImagesConsumer(amqpConn *amqp.Connection, emailUC email.EmailsUseCase,cfg  *config.Config) *EmailsConsumer {
	return &EmailsConsumer{amqpConn: amqpConn, emailUC: emailUC,cfg:cfg}
}

func (c *EmailsConsumer) CreateChannel(exchangeName, queueName, bindingKey, consumerTag string) (*amqp.Channel, error) {
	ch, err := c.amqpConn.Channel()
	if err != nil {
		return nil, errors.Wrap(err, "Error amqpConn.Channel")
	}
	err = ch.ExchangeDeclare(
		exchangeName,
		exchangeKind,
		exchangeDurable,
		exchangeAutoDelete,
		exchangeInternal,
		exchangeNoWait,
		nil,
	)
	if err != nil {
		return nil, errors.Wrap(err, "Error ch.ExchangeDeclare")
	}
	queue, err := ch.QueueDeclare(
		queueName,
		queueDurable,
		queueAutoDelete,
		queueExclusive,
		queueNoWait,
		nil,
	)
	if err != nil {
		return nil, errors.Wrap(err, "Error ch.QueueDeclare")
	}
	err = ch.QueueBind(
		queue.Name,
		bindingKey,
		exchangeName,
		queueNoWait,
		nil,
	)
	if err != nil {
		return nil, errors.Wrap(err, "Error ch.QueueBind")
	}
	err = ch.Qos(
		prefetchCount,  // prefetch count
		prefetchSize,   // prefetch size
		prefetchGlobal, // global
	)
	if err != nil {
		return nil, errors.Wrap(err, "Error  ch.Qos")
	}
	return ch, nil
}

func (c *EmailsConsumer) worker(ctx context.Context, messages <-chan amqp.Delivery) {
	for delivery := range messages {
		span, ctx := opentracing.StartSpanFromContext(ctx, "EmailsConsumer.worker")
		err := c.emailUC.SendEmail(ctx, delivery.Body)
		if err != nil {
			if err := delivery.Reject(false); err != nil {
				log.Println("Err delivery.Reject")
			}
		} else {
			err = delivery.Ack(false)
			if err != nil {
				log.Println("Failed to acknowledge delivery")
			}
		}
		span.Finish()
	}
}

func (c *EmailsConsumer) StartConsumer(workerPoolSize int, exchange, queueName, bindingKey, consumerTag string) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ch, err := c.CreateChannel(exchange, queueName, bindingKey, consumerTag)
	if err != nil {
		return errors.Wrap(err, "CreateChannel")
	}
	defer ch.Close()
	deliveries, err := ch.Consume(
		queueName,
		consumerTag,
		consumeAutoAck,
		consumeExclusive,
		consumeNoLocal,
		consumeNoWait,
		nil,
	)
	if err != nil {
		return errors.Wrap(err, "Consume")
	}

	for i := 0; i < workerPoolSize; i++ {
		go c.worker(ctx, deliveries)
	}
	chanErr := <-ch.NotifyClose(make(chan *amqp.Error))
	return chanErr
}