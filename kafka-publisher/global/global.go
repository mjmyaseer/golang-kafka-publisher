package global

import (
	"github.com/Shopify/sarama"

 "Company import path goes here"

)

var (
	Subjects  = make(map[string]string)
	Versions  = make(map[string]int)
	Topics    = make(map[string][]string)
	Registry  *schema_registry.SchemaRegistry
	Publisher sarama.SyncProducer
)

func init() {
	config.DefaultConfigurator = config.NewConfigurator(config.NewZookeeperLoader())
}

func LoadConfigurations() {
	config.LoadConfiguration(&config.Configurations{
		new(config.AppConfig),
		new(configurations.ProducerConf),
		new(configurations.Map),
		new(configurations.UrlConf),
	})
}

func Init() {
	LoadConfigurations()
	CreateProducer()
	Registry = schema_registry.NewSchemaRegistry(configurations.UrlConfig.SchemaPath)
	for event := range configurations.Maps {
		temp := configurations.Maps[event]
		Subjects[temp.Event] = temp.Subject
		Versions[temp.Event] = temp.Version
		Topics[temp.Event] = temp.Topics
		Registry.Register(temp.Subject, temp.Version, nil)
	}
	log.Info("Event Map and Subject Map are loaded")
}

func CreateProducer() {

	conf := sarama.NewConfig()
	conf.Producer.RequiredAcks = configurations.ProducerConfig.Producer.RequiredAcks
	conf.Producer.Retry.Max = configurations.ProducerConfig.Producer.RetryMax
	conf.Producer.Return.Successes = configurations.ProducerConfig.Producer.ReturnSuccess
	conf.Producer.Return.Errors = configurations.ProducerConfig.Producer.ReturnErrors
	conf.Version = sarama.V1_1_0_0
	conf.Producer.Partitioner = sarama.NewHashPartitioner
	producer, err := sarama.NewSyncProducer(configurations.ProducerConfig.Brokers, conf)
	if err != nil {
		panic(err)
	}
	Publisher = producer
	log.Info("Producer is inited")
}
