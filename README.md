# General information about Kafka

Produce messages on topic `test`

```bash
kafka-console-producer --broker-list localhost:9092 --topic test
```

Consume messages on specific server

```bash
kafka-console-consumer --bootstrap-server localhost:9092 --topic test --from-beginning
```


#### Resources

* [Kafka - Protobuf or Avro](https://medium.com/@felipedutratine/kafka-protobuf-or-avro-178c629b7327)
* [Avro explained](https://www.confluent.io/blog/avro-kafka-data/)
##### Other
* [Stream processing](http://samza.apache.org/)
