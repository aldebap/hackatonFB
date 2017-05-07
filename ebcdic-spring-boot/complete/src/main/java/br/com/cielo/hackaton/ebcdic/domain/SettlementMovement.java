package br.com.cielo.hackaton.ebcdic.domain;

import java.io.Serializable;
import java.math.BigDecimal;
import java.util.Date;

public class SettlementMovement implements Serializable {
	
	
	/**
	 * Contructor
	 */

	public SettlementMovement(Date batchDate, Integer productCode, Long customerNumber, Date settlementDate,
			BigDecimal netMovementValue, BigDecimal grossMovementValue, BigDecimal dailyDiscountValueAmt,
			Long loadFileIdNumber, String fundingCurrencyCode, String traceNumberFinancialMovement,
			Date traceSettlementDate, Integer customerModNumber, Integer movementTypeCode,
			String numberFinancialMovementOrig, Integer statusMovement) {
		super();
		this.batchDate = batchDate;
		this.productCode = productCode;
		this.customerNumber = customerNumber;
		this.settlementDate = settlementDate;
		this.netMovementValue = netMovementValue;
		this.grossMovementValue = grossMovementValue;
		this.dailyDiscountValueAmt = dailyDiscountValueAmt;
		this.loadFileIdNumber = loadFileIdNumber;
		this.fundingCurrencyCode = fundingCurrencyCode;
		this.traceNumberFinancialMovement = traceNumberFinancialMovement;
		this.traceSettlementDate = traceSettlementDate;
		this.customerModNumber = customerModNumber;
		this.movementTypeCode = movementTypeCode;
		this.numberFinancialMovementOrig = numberFinancialMovementOrig;
		this.statusMovement = statusMovement;
	}

	/**
	 * serialVersionUID
	 */
	private static final long serialVersionUID = 1L;

	/**
	 * Campo Date batchDate. DT_BATCH DATE
	 */
	private Date batchDate;

	/**
	 * Campo Integer productCode. CD_PRODUCT NUMBER(4, 0) NOT NULL
	 */
	private Integer productCode;

	/**
	 * Campo Long customerNumber. NU_CUSTOMER NUMBER(10, 0) NOT NULL
	 */
	private Long customerNumber;

	/**
	 * Campo Date settlementDate. DT_SETTLEMENT DATE
	 */
	private Date settlementDate;

	/**
	 * Campo BigDecimal netMovementValue. VL_MOVEMENT_NET NUMBER(15,2)
	 */
	private BigDecimal netMovementValue;

	/**
	 * Campo BigDecimal grossMovementValue. VL_MOVEMENT_GROSS NUMBER(15,2)
	 */
	private BigDecimal grossMovementValue;

	/**
	 * Campo BigDecimal dailyDiscountValueAmt. VL_DAILY_DISCOUNT_AMOUNT
	 * NUMBER(15,2)
	 */
	private BigDecimal dailyDiscountValueAmt;

	/**
	 * Campo Long loadFileIdNumber. NU_LOAD_FILE_ID NUMBER(20, 0)
	 */
	private Long loadFileIdNumber;

	/**
	 * Campo String fundingCurrencyCode. CD_FUNDING_CURRENCY VARCHAR2(3)
	 */
	private String fundingCurrencyCode;

	/**
	 * Campo String numberFinancialMovement
	 */
	private String traceNumberFinancialMovement;

	/**
	 * Campo Date traceSettlementBaseDate. DT_SETTLEMENT_TRACE DATE
	 */
	private Date traceSettlementDate;

	/**
	 * Campo Integer customerModNumber. NU_MOD_CUSTOMER NUMBER
	 */
	private Integer customerModNumber;

	/**
	 * Campo Integer movementTypeCode. CD_MOVEMENT_TYPE NUMBER(3)
	 */
	private Integer movementTypeCode;

	/**
	 * Campo String numberFinancialMovementOrig
	 */
	private String numberFinancialMovementOrig;
	
	private int statusMovement;
	
	
	/**
	 * Campo Int statusMovement
	 */
	public Integer getStatusMovement() {
		return statusMovement;
	}

	public void setStatusMovement(Integer statusMovement) {
		this.statusMovement = statusMovement;
	}

	public Date getBatchDate() {
		return batchDate;
	}

	public void setBatchDate(Date batchDate) {
		this.batchDate = batchDate;
	}

	public Integer getProductCode() {
		return productCode;
	}

	public void setProductCode(Integer productCode) {
		this.productCode = productCode;
	}

	public Long getCustomerNumber() {
		return customerNumber;
	}

	public void setCustomerNumber(Long customerNumber) {
		this.customerNumber = customerNumber;
	}

	public Date getSettlementDate() {
		return settlementDate;
	}

	public void setSettlementDate(Date settlementDate) {
		this.settlementDate = settlementDate;
	}

	public BigDecimal getNetMovementValue() {
		return netMovementValue;
	}

	public void setNetMovementValue(BigDecimal netMovementValue) {
		this.netMovementValue = netMovementValue;
	}

	public BigDecimal getGrossMovementValue() {
		return grossMovementValue;
	}

	public void setGrossMovementValue(BigDecimal grossMovementValue) {
		this.grossMovementValue = grossMovementValue;
	}

	public BigDecimal getDailyDiscountValueAmt() {
		return dailyDiscountValueAmt;
	}

	public void setDailyDiscountValueAmt(BigDecimal dailyDiscountValueAmt) {
		this.dailyDiscountValueAmt = dailyDiscountValueAmt;
	}

	public Long getLoadFileIdNumber() {
		return loadFileIdNumber;
	}

	public void setLoadFileIdNumber(Long loadFileIdNumber) {
		this.loadFileIdNumber = loadFileIdNumber;
	}

	public String getFundingCurrencyCode() {
		return fundingCurrencyCode;
	}

	public void setFundingCurrencyCode(String fundingCurrencyCode) {
		this.fundingCurrencyCode = fundingCurrencyCode;
	}

	public String getTraceNumberFinancialMovement() {
		return traceNumberFinancialMovement;
	}

	public void setTraceNumberFinancialMovement(String traceNumberFinancialMovement) {
		this.traceNumberFinancialMovement = traceNumberFinancialMovement;
	}

	public Date getTraceSettlementDate() {
		return traceSettlementDate;
	}

	public void setTraceSettlementDate(Date traceSettlementDate) {
		this.traceSettlementDate = traceSettlementDate;
	}

	public Integer getCustomerModNumber() {
		return customerModNumber;
	}

	public void setCustomerModNumber(Integer customerModNumber) {
		this.customerModNumber = customerModNumber;
	}

	public Integer getMovementTypeCode() {
		return movementTypeCode;
	}

	public void setMovementTypeCode(Integer movementTypeCode) {
		this.movementTypeCode = movementTypeCode;
	}

	public String getNumberFinancialMovementOrig() {
		return numberFinancialMovementOrig;
	}

	public void setNumberFinancialMovementOrig(String numberFinancialMovementOrig) {
		this.numberFinancialMovementOrig = numberFinancialMovementOrig;
	}

	@Override
	public String toString() {
		return "SettlementMovement [batchDate=" + batchDate + ", productCode=" + productCode + ", customerNumber="
				+ customerNumber + ", settlementDate=" + settlementDate + ", netMovementValue=" + netMovementValue
				+ ", grossMovementValue=" + grossMovementValue + ", dailyDiscountValueAmt=" + dailyDiscountValueAmt
				+ ", loadFileIdNumber=" + loadFileIdNumber + ", fundingCurrencyCode=" + fundingCurrencyCode
				+ ", traceNumberFinancialMovement=" + traceNumberFinancialMovement + ", traceSettlementDate="
				+ traceSettlementDate + ", customerModNumber=" + customerModNumber + ", movementTypeCode="
				+ movementTypeCode + ", numberFinancialMovementOrig=" + numberFinancialMovementOrig + "]";
	}
	
}
