# Apache kafka test repo
## docker
simple setup by [compose.yaml](docker/compose.yaml)
```
cd docker
docker-compose up -d
```
check 
```
rustam@rustam-zenbook:~/kafka_demo$ docker ps 
CONTAINER ID   IMAGE                              COMMAND                  CREATED         STATUS         PORTS                                                               NAMES
73160b57e69d   confluentinc/cp-kafka:latest       "/etc/confluent/dock…"   4 minutes ago   Up 4 minutes   9092/tcp, 0.0.0.0:29092->29092/tcp, [::]:29092->29092/tcp           kafka
7d83cbe670c3   confluentinc/cp-zookeeper:latest   "/etc/confluent/dock…"   4 minutes ago   Up 4 minutes   2888/tcp, 3888/tcp, 0.0.0.0:22181->2181/tcp, [::]:22181->2181/tcp   zookeeper

```
## python service
uv package manager should be installed
```
pip3 install uv
```
in py-kafka
```
cd py-kafka
```
generate msg 
```
uv run producer
```
get mgs
```
uv run consumer
```
example 
```
rustam@rustam-zenbook:~/kafka_demo/py-kafka$ uv run producer
2025-03-08 18:54:16.168 | INFO     | py_kafka.producer:main:20 - send i=0
2025-03-08 18:54:20.177 | INFO     | py_kafka.producer:main:20 - send i=1
```
```
rustam@rustam-zenbook:~/kafka_demo/py-kafka$ uv run consumer
2025-03-08 18:54:43.874 | INFO     | py_kafka.consumer:main:18 - {"message": "Hello, Kafka! - test 0"}
2025-03-08 18:54:43.874 | INFO     | py_kafka.consumer:main:18 - {"message": "Hello, Kafka! - test 1"}
```

## rust
create new topic
```
rustam@rustam-zenbook:~/kafka_demo$ docker exec kafka kafka-topics --bootstrap-server kafka:29092  --create --topic rustStories
Created topic rustStories.
```
check list 
```
rustam@rustam-zenbook:~/kafka_demo$ docker exec kafka kafka-topics --bootstrap-server kafka:29092  --list 
rustStories
test
```
write msg once
```
rustam@rustam-zenbook:~/kafka_demo/rust_kafka$ cargo run --bin producer
    Finished `dev` profile [unoptimized + debuginfo] target(s) in 0.03s
     Running `target/debug/producer`
About to publish a message at ["localhost:29092"] to: rustStories
```
get msg in loop 
```
rustam@rustam-zenbook:~/kafka_demo/rust_kafka$ cargo run --bin consumer
   Compiling rust_kafka v0.1.0 (/home/rustam/kafka_demo/rust_kafka)
    Finished `dev` profile [unoptimized + debuginfo] target(s) in 0.66s
     Running `target/debug/consumer`
rustStories:0@6: "rust, kafka"
rustStories:0@7: "rust, kafka"

```

