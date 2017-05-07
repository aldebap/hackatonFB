package br.com.cielo.hackaton.ebcdic.config;

import java.sql.SQLException;

import javax.sql.DataSource;

import org.springframework.context.annotation.Bean;

import com.zaxxer.hikari.HikariConfig;
import com.zaxxer.hikari.HikariDataSource;

//@Configuration
class DataSourceConfig {
  
	@Bean(destroyMethod = "close")
	public DataSource dataSource() throws SQLException {
	    HikariConfig config = new HikariConfig("/hikari.properties");
	    HikariDataSource dataSource = new HikariDataSource(config);

	    return dataSource;
	}
}
