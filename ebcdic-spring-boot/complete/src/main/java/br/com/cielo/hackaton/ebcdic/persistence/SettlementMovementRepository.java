package br.com.cielo.hackaton.ebcdic.persistence;

import java.sql.SQLException;
import java.util.List;

import br.com.cielo.hackaton.ebcdic.domain.SettlementMovement;

public interface SettlementMovementRepository {
	
	
	/**
     * SQL Inserção de lançamento financeiro na agenda.
     */
    String INSERT_SETTLEMENT_MOVEMENT = "INSERT INTO "
                    + " TBSETR_SETTLEMENT_MOVEMENT "
                    + " (DT_BATCH, "
                    + " CD_PRODUCT, "
                    + " NU_CUSTOMER, "
                    + " DT_SETTLEMENT, "
                    + " VL_MOVEMENT_NET, "
                    + " VL_MOVEMENT_GROSS, "
                    + " VL_DAILY_DISCOUNT_AMOUNT, "
                    + " NU_LOAD_FILE_ID, "
                    + " CD_FUNDING_CURRENCY, "
                    // + " NU_FINANCIAL_MOVEMENT_TRACE, "
                    // + " DT_SETTLEMENT_TRACE, "
                    // + " NU_MOD_CUSTOMER_TRACE, "
                    + " NU_MOD_CUSTOMER, " 
                    + " CD_MOVEMENT_STATUS, " 
                    + " CD_MOVEMENT_TYPE, "
                    + " NU_FINANCIAL_MOVEMENT " 
                    + " ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? )";
    
    
    /**
     * Criar lancamentos financeiros na agenda
     * @throws SQLException 
     */
    void insertMovement(SettlementMovement movement) throws SQLException;


	void insertMovementBatchCommand(List<SettlementMovement> movements) throws SQLException;

}
