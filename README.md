# General information about Kafka

### Installation

- [Installing Kafka with Homebrew](https://medium.com/@Ankitthakur/apache-kafka-installation-on-mac-using-homebrew-a367cdefd273)

```bash
brew cask install java
brew install kafka
```

### Start (local machine)

Start zookeeper

```bash
zookeeper-server-start /usr/local/etc/kafka/zookeeper.properties
```

Start Kafka

```bash
kafka-server-start /usr/local/etc/kafka/server.properties
```

### Start (Docker)

Build kafka container from [wurstmeister/kafka-docker](https://github.com/wurstmeister/kafka-docker)

```bash
./scripts/download-build.sh
```

Start kafka and zookeeper

```bash
./scripts/start-kafka.sh
```

### Use

Produce messages on topic `test`

```bash
kafka-console-producer --broker-list localhost:9092 --topic test
```

Consume messages on specific server

```bash
kafka-console-consumer --bootstrap-server localhost:9092 --topic test --from-beginning
```

### Resources

- [Kafka - Protobuf or Avro](https://medium.com/@felipedutratine/kafka-protobuf-or-avro-178c629b7327)
- [Avro explained](https://www.confluent.io/blog/avro-kafka-data/)
- [Schema evolution in Avro, Protocol Buffers and Thrift](http://martin.kleppmann.com/2012/12/05/schema-evolution-in-avro-protocol-buffers-thrift.html)
- [Understanding Kafka with Legos](https://youtu.be/Q5wOegcVa8E)
  - 10 minute, cool overview
- [Apache Kafka with Go](https://medium.com/@jawadahmadd/apache-kafka-with-go-f67986afb9a9)
  - Great!
  - Shows alternatives
  - Good code with great comments

```golang
// The tuple (topic, partition, offset) can be used as a unique identifier
// for a message in a Kafka cluster.
fmt.Println( "Your data is stored with unique identifier important/%d/%d", partition, offset)
```

- [Confluent 6: part series](https://www.confluent.io/blog/data-dichotomy-rethinking-the-way-we-treat-data-and-services/)
- [20 Kafka Best Practices](https://blog.newrelic.com/engineering/kafka-best-practices/)
- [Keynote - The Present and Future of the Streaming Platform](https://youtu.be/eublKlalobg)
  - Keynote Kafka Summit 2018 - London
  - Pretty good keynote on 5 stages of using streaming fully
    Some real life examples, Nordea, Norwegian government

#### Other

- [Apache Samza - Stream processing](http://samza.apache.org/)

##### Ongoing

- [Building Event Driven Services with Apache Kafka and Kafka Streams](https://www.youtube.com/watch?v=p9wcx3aTjuo)
  - ![image](image_ksql.png)

### Questions

- Are there unique message ids?
- What message format to use?
  - Protobuf
  - Avro
  - JSON
- Format important for KSQL? - [Yes](https://medium.com/@felipedutratine/kafka-protobuf-or-avro-178c629b7327)
- Partition per user?
- How does Kafka handle "pending" messages?
  - Poison queue?
- How do we handle multiple consumers in same group? For example when we deploy a new service version
