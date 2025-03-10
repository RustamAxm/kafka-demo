use kafka::consumer::{Consumer, FetchOffset, GroupOffsetStorage};
use kafka::error::Error as KafkaError;

/// This program demonstrates consuming messages through a `Consumer`.
/// This is a convenient client that will fit most use cases.  Note
/// that messages must be marked and committed as consumed to ensure
/// only once delivery.
fn main() {
    tracing_subscriber::fmt::init();

    let broker = "localhost:29092".to_owned();
    let topic = "rustStories".to_owned();
    let group = "my-group".to_owned();

    if let Err(e) = consume_messages(group, topic, vec![broker]) {
        println!("Failed consuming messages: {}", e);
    }
}

fn consume_messages(group: String, topic: String, brokers: Vec<String>) -> Result<(), KafkaError> {
    let mut con = Consumer::from_hosts(brokers)
        .with_topic(topic)
        .with_group(group)
        .with_fallback_offset(FetchOffset::Earliest)
        .with_offset_storage(Some(GroupOffsetStorage::Kafka))
        .create()?;

    loop {
        for ms in con.poll().unwrap().iter() {
            for m in ms.messages() {
                println!(
                    "{}:{}@{}: {:?}",
                    ms.topic(),
                    ms.partition(),
                    m.offset,
                    String::from_utf8_lossy(m.value)
                );
            }
            let _ = con.consume_messageset(ms);
        }
        con.commit_consumed()?;
    }
}
