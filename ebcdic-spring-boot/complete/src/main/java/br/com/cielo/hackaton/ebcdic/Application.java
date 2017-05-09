package br.com.cielo.hackaton.ebcdic;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.boot.CommandLineRunner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.ApplicationContext;
import org.springframework.context.annotation.Bean;

import br.com.cielo.hackaton.ebcdic.service.ConsumerKafka;


@SpringBootApplication
public class Application {
	

	private static final Logger log = LoggerFactory.getLogger(Application.class);

    public static void main(String[] args) {
        SpringApplication.run(Application.class, args);
    }

    @Bean 
    public CommandLineRunner commandLineRunner(ApplicationContext ctx) {
        return args -> {

            log.info("Iniciando servico Spring Boot");
           // getConsumer().consume();           
            
        };
    }
    
    @Bean 
    public static  ConsumerKafka getConsumer() throws Exception {
    	return new ConsumerKafka();
    }
}
