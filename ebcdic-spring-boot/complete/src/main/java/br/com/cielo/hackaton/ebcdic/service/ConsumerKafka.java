package br.com.cielo.hackaton.ebcdic.service;

import java.io.IOException;
import java.math.BigDecimal;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.Date;
import java.util.List;
import java.util.Properties;
import java.util.UUID;
import java.util.concurrent.ArrayBlockingQueue;
import java.util.concurrent.ThreadPoolExecutor;
import java.util.concurrent.TimeUnit;

import org.apache.kafka.clients.consumer.ConsumerRecord;
import org.apache.kafka.clients.consumer.ConsumerRecords;
import org.apache.kafka.clients.consumer.KafkaConsumer;
import org.apache.kafka.clients.consumer.OffsetAndMetadata;
import org.apache.kafka.common.TopicPartition;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import com.fasterxml.jackson.core.JsonParseException;
import com.fasterxml.jackson.databind.JsonMappingException;
import com.fasterxml.jackson.databind.ObjectMapper;

import br.com.cielo.hackaton.ebcdic.domain.SettlementMessage;
import br.com.cielo.hackaton.ebcdic.domain.SettlementMovement;
import br.com.cielo.hackaton.ebcdic.persistence.SettlementMovementRepository;

@Component
public class ConsumerKafka {

	private KafkaConsumer<String, String> consumer;

	private ThreadPoolExecutor executor = new ThreadPoolExecutor(2, 2, 1, TimeUnit.SECONDS,
			new ArrayBlockingQueue<>(10), new ThreadPoolExecutor.CallerRunsPolicy());

	@Autowired
	SettlementMovementRepository settlementMovementRepository;

	public ConsumerKafka() throws Exception {
		Properties props = new Properties();
		props.put("bootstrap.servers", "localhost:9092");
		props.put("group.id", "settlement_movement_group");
		props.put("key.serializer", "org.apache.kafka.common.serialization.StringSerializer");
		props.put("value.serializer", "org.apache.kafka.common.serialization.StringSerializer");
		props.put("enable.auto.commit", "true");
		props.put("key.deserializer", "org.apache.kafka.common.serialization.StringDeserializer");
		props.put("value.deserializer", "org.apache.kafka.common.serialization.StringDeserializer");
		props.put("auto.commit.interval.ms", "1000");
		props.put("session.timeout.ms", "30000");

		consumer = new KafkaConsumer<>(props);
		consumer.subscribe(Arrays.asList("settlement_movement"));

	}

	public void consume() throws Exception {

		final int minBatchSize = 1000;
		List<ConsumerRecord<String, String>> buffer = new ArrayList<>();

		try {
			while (true) {
				ConsumerRecords<String, String> records = consumer.poll(Long.MAX_VALUE);
				for (TopicPartition partition : records.partitions()) {
					List<ConsumerRecord<String, String>> partitionRecords = records.records(partition);
					for (ConsumerRecord<String, String> record : records) {
						buffer.add(record);
					}

					if (buffer.size() >= minBatchSize) {
						executor.execute(() -> {
							try {
								insertIntoDb(buffer);
							} catch (Exception e) {
								e.printStackTrace();
							}
						});

						long lastOffset = partitionRecords.get(partitionRecords.size() - 1).offset();
						consumer.commitSync(Collections.singletonMap(partition, new OffsetAndMetadata(lastOffset + 1)));
						// consumer.commitSync();
						buffer.clear();

					}

				}
			}
		} finally {
			consumer.close();
		}
	}

	private void insertIntoDb(List<ConsumerRecord<String, String>> buffer) throws Exception {

		List<SettlementMovement> movements = new ArrayList<>();

		buffer.stream().forEach(record -> {
			try {
				movements.add(mapper(record));
			} catch (Exception e) {

				e.printStackTrace();
			}
		});

		for (ConsumerRecord<String, String> record : buffer) {
			System.out.println(record.offset() + ": " + record.value());

		}

		if (settlementMovementRepository != null) {
			settlementMovementRepository.insertMovementBatchCommand(movements);
		} else {
			System.out.println("nnnnnnnnnnnnnnnn NULL");
		}

	}

	private SettlementMovement mapper(ConsumerRecord<String, String> record) throws Exception {

		SettlementMessage msg = null;
		SettlementMovement sm = null;

		try {
			msg = new ObjectMapper().readValue(record.value(), SettlementMessage.class);
		} catch (JsonParseException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
			throw e;
		} catch (JsonMappingException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
			throw e;
		} catch (IOException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
			throw e;
		}

		if (msg != null) {
			sm = new SettlementMovement(
					// batchDate,
					msg.getDate(),
					// productCode,
					new Integer(msg.getProductCode()),
					// customerNumber,
					new Long(msg.getCustomer()),
					// settlementDate,
					msg.getSettlementAdjustment(),
					// netMovementValue,
					new BigDecimal(msg.getGrossValue()),
					// grossMovementValue,
					new BigDecimal(msg.getGrossValue()),
					// dailyDiscountValueAmt,
					new BigDecimal(20),
					// loadFileIdNumber,
					new Long(msg.getBatchId()),
					// fundingCurrencyCode,
					msg.getFundingCurrencyCode(),
					// traceNumberFinancialMovement,
					"",
					// traceSettlementDate,
					new Date(),
					// customerModNumber,
					msg.getModCustomer(),
					// movementTypeCode,
					msg.getMovementType(),
					// numberFinancialMovementOrig,
					UUID.randomUUID().toString(),
					// statusMovement
					msg.getRefundStatus());
		}

		return sm;
	}

}
