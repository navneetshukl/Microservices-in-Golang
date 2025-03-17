package event

import amqp "github.com/rabbitmq/amqp091-go"

type Consumer struct {
	conn *amqp.Connection
	queueName string

}

func NewConsumer(conn *amqp.Connection)(Consumer,error){
	consumer:=Consumer{
		conn: conn,
	}

	err:=consumer.setup()
	if err!=nil{
		return Consumer{},nil
	}

	return consumer,nil
}

func(consumer *Consumer) setup() error{
	channel,err:=consumer.conn.Channel()
	if err!=nil{
		return err
	}
	return declareExchange(channel)
}