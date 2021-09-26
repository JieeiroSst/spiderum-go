package rabbitmq

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/streadway/amqp"
	"gitlab.com/Spide_IT/spide_it/email/config"
	"gitlab.com/Spide_IT/spide_it/email/pkg/rabbitmq"
	"log"
	"time"
)

type EmailsPublisher struct {
	amqpChan *amqp.Channel
	cfg      *config.Config
}

func NewEmailsPublisher(cfg *config.Config) (*EmailsPublisher, error) {
	mqConn, err := rabbitmq.NewRabbitMQConn(cfg)
	if err != nil {
		return nil, err
	}
	amqpChan, err := mqConn.Channel()
	if err != nil {
		return nil, errors.Wrap(err, "p.amqpConn.Channel")
	}

	return &EmailsPublisher{cfg: cfg, amqpChan: amqpChan}, nil
}

func (p *EmailsPublisher) SetupExchangeAndQueue(exchange, queueName, bindingKey, consumerTag string) error {
	err := p.amqpChan.ExchangeDeclare(
		exchange,
		exchangeKind,
		exchangeDurable,
		exchangeAutoDelete,
		exchangeInternal,
		exchangeNoWait,
		nil,
	)
	if err != nil {
		return errors.Wrap(err, "Error ch.ExchangeDeclare")
	}

	queue, err := p.amqpChan.QueueDeclare(
		queueName,
		queueDurable,
		queueAutoDelete,
		queueExclusive,
		queueNoWait,
		nil,
	)
	if err != nil {
		return errors.Wrap(err, "Error ch.QueueDeclare")
	}

	err = p.amqpChan.QueueBind(
		queue.Name,
		bindingKey,
		exchange,
		queueNoWait,
		nil,
	)
	if err != nil {
		return errors.Wrap(err, "Error ch.QueueBind")
	}
	return nil
}

func (p *EmailsPublisher) CloseChan() {
	if err := p.amqpChan.Close(); err != nil {
		log.Printf("EmailsPublisher CloseChan: %v", err)
	}
}

func (p *EmailsPublisher) Publish(body []byte, contentType string) error {
	if err := p.amqpChan.Publish(
		p.cfg.RabbitMQ.Exchange,
		p.cfg.RabbitMQ.RoutingKey,
		publishMandatory,
		publishImmediate,
		amqp.Publishing{
			ContentType:  contentType,
			DeliveryMode: amqp.Persistent,
			MessageId:    uuid.New().String(),
			Timestamp:    time.Now(),
			Body:         body,
		},
	); err != nil {
		return errors.Wrap(err, "ch.Publish")
	}
	return nil
}