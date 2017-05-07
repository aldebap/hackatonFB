package br.com.cielo.hackaton.ebcdic.persistence;

import java.sql.Connection;
import java.sql.PreparedStatement;
import java.sql.SQLException;
import java.util.List;

import javax.sql.DataSource;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Bean;
import org.springframework.stereotype.Repository;

import com.zaxxer.hikari.HikariConfig;
import com.zaxxer.hikari.HikariDataSource;

import br.com.cielo.hackaton.ebcdic.domain.SettlementMovement;

@Repository
public class SettlementMovementRepositoryImpl implements SettlementMovementRepository {

	@Autowired
	DataSource dataSource;

	@Override
	public void insertMovement(SettlementMovement movement) throws SQLException {

		Connection connection = this.dataSource.getConnection();
		PreparedStatement ps = connection.prepareStatement(SettlementMovementRepository.INSERT_SETTLEMENT_MOVEMENT);

		int index = 0;

		ps.setDate(++index, new java.sql.Date(movement.getBatchDate().getTime()));
		ps.setInt(++index, movement.getProductCode());
		ps.setLong(++index, movement.getCustomerNumber());
		ps.setDate(++index, new java.sql.Date(movement.getSettlementDate().getTime()));
		ps.setBigDecimal(++index, movement.getNetMovementValue());
		ps.setBigDecimal(++index, movement.getGrossMovementValue());
		ps.setBigDecimal(++index, movement.getDailyDiscountValueAmt());
		ps.setLong(++index, movement.getLoadFileIdNumber());
		ps.setString(++index, movement.getFundingCurrencyCode());
		ps.setInt(++index, movement.getCustomerModNumber());
		ps.setInt(++index, movement.getStatusMovement());
		ps.setInt(++index, movement.getMovementTypeCode());
		ps.setString(++index, movement.getNumberFinancialMovementOrig());
		ps.addBatch();

		ps.executeBatch();
		ps.close();
		connection.close();
	}

	@Override
	public void insertMovementBatchCommand(List<SettlementMovement> movements) throws SQLException {

		System.out.println(" >>>>>Tamanho da Lista de Movimentos Financeiros:" + movements.size());

		Connection connection = this.dataSource.getConnection();

		final int batchSize = 10;
		int count = 0;
		int index = 0;
		int[] result;

		PreparedStatement ps = connection.prepareStatement(SettlementMovementRepository.INSERT_SETTLEMENT_MOVEMENT);

		long start = System.currentTimeMillis();

		for (SettlementMovement movement : movements) {
			ps.setDate(++index, new java.sql.Date(movement.getBatchDate().getTime()));
			ps.setInt(++index, movement.getProductCode());
			ps.setLong(++index, movement.getCustomerNumber());
			ps.setDate(++index, new java.sql.Date(movement.getSettlementDate().getTime()));
			ps.setBigDecimal(++index, movement.getNetMovementValue());
			ps.setBigDecimal(++index, movement.getGrossMovementValue());
			ps.setBigDecimal(++index, movement.getDailyDiscountValueAmt());
			ps.setLong(++index, movement.getLoadFileIdNumber());
			ps.setString(++index, movement.getFundingCurrencyCode());
			ps.setInt(++index, movement.getCustomerModNumber());
			ps.setInt(++index, movement.getStatusMovement());
			ps.setInt(++index, movement.getMovementTypeCode());
			ps.setString(++index, movement.getNumberFinancialMovementOrig());
			ps.addBatch();

			if (++count % batchSize == 0) {
				result = ps.executeBatch();
				System.out.println("Qtde Inserts: " + result.length);

			}

			index = 0;

		}
		int[] result2 = ps.executeBatch();
		System.out.println("Qtde Inserts: " + result2.length);
		connection.close();

		long end = System.currentTimeMillis();
		System.out.println("Tempo total = " + (end - start) + " ms");
		System.out.println("Mediaa = " + (end - start) / movements.size() + " ms");

		System.out.println(" Fim insertMovementBatchCommand");
	}

	@Bean(destroyMethod = "close")
	public DataSource getDataSouce() throws SQLException {
		HikariConfig config = new HikariConfig("/hikari.properties");
		HikariDataSource dataSource = new HikariDataSource(config);

		return dataSource;
	}
}
