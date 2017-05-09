package br.com.cielo.hackaton.ebcdic.domain;

import java.util.Date;

public class SettlementMessage {
	
	private int id;
	private String type;
	private Date date; 
	private String rejectReason;
	private Date settlementAdjustment;
	private int grossValue;
	private String userId;
	private int tecnologyType;
	private String requestStatus;
	private String customer;
	private int modCustomer;
	private int batchId;
	private int movementType;
	private int refundStatus;
	private int adjustmentReason;
	private String AdjustmentComments;
	private int fieldId;
	private int product;
	private String entryType;
	private int productCode;
	private String cardAssociation;
	private String fundingCurrencyCode;
	/**
	 * @return the id
	 */
	public int getId() {
		return id;
	}
	/**
	 * @param id the id to set
	 */
	public void setId(int id) {
		this.id = id;
	}
	/**
	 * @return the type
	 */
	public String getType() {
		return type;
	}
	/**
	 * @param type the type to set
	 */
	public void setType(String type) {
		this.type = type;
	}
	/**
	 * @return the date
	 */
	public Date getDate() {
		return date;
	}
	/**
	 * @param date the date to set
	 */
	public void setDate(Date date) {
		this.date = date;
	}
	/**
	 * @return the rejectReason
	 */
	public String getRejectReason() {
		return rejectReason;
	}
	/**
	 * @param rejectReason the rejectReason to set
	 */
	public void setRejectReason(String rejectReason) {
		this.rejectReason = rejectReason;
	}
	/**
	 * @return the settlementAdjustment
	 */
	public Date getSettlementAdjustment() {
		return settlementAdjustment;
	}
	/**
	 * @param settlementAdjustment the settlementAdjustment to set
	 */
	public void setSettlementAdjustment(Date settlementAdjustment) {
		this.settlementAdjustment = settlementAdjustment;
	}
	/**
	 * @return the grossValue
	 */
	public int getGrossValue() {
		return grossValue;
	}
	/**
	 * @param grossValue the grossValue to set
	 */
	public void setGrossValue(int grossValue) {
		this.grossValue = grossValue;
	}
	/**
	 * @return the userId
	 */
	public String getUserId() {
		return userId;
	}
	/**
	 * @param userId the userId to set
	 */
	public void setUserId(String userId) {
		this.userId = userId;
	}
	/**
	 * @return the tecnologyType
	 */
	public int getTecnologyType() {
		return tecnologyType;
	}
	/**
	 * @param tecnologyType the tecnologyType to set
	 */
	public void setTecnologyType(int tecnologyType) {
		this.tecnologyType = tecnologyType;
	}
	/**
	 * @return the requestStatus
	 */
	public String getRequestStatus() {
		return requestStatus;
	}
	/**
	 * @param requestStatus the requestStatus to set
	 */
	public void setRequestStatus(String requestStatus) {
		this.requestStatus = requestStatus;
	}
	/**
	 * @return the customer
	 */
	public String getCustomer() {
		return customer;
	}
	/**
	 * @param customer the customer to set
	 */
	public void setCustomer(String customer) {
		this.customer = customer;
	}
	/**
	 * @return the modCustomer
	 */
	public int getModCustomer() {
		return modCustomer;
	}
	/**
	 * @param modCustomer the modCustomer to set
	 */
	public void setModCustomer(int modCustomer) {
		this.modCustomer = modCustomer;
	}
	/**
	 * @return the batchId
	 */
	public int getBatchId() {
		return batchId;
	}
	/**
	 * @param batchId the batchId to set
	 */
	public void setBatchId(int batchId) {
		this.batchId = batchId;
	}
	/**
	 * @return the movementType
	 */
	public int getMovementType() {
		return movementType;
	}
	/**
	 * @param movementType the movementType to set
	 */
	public void setMovementType(int movementType) {
		this.movementType = movementType;
	}
	/**
	 * @return the refundStatus
	 */
	public int getRefundStatus() {
		return refundStatus;
	}
	/**
	 * @param refundStatus the refundStatus to set
	 */
	public void setRefundStatus(int refundStatus) {
		this.refundStatus = refundStatus;
	}
	/**
	 * @return the adjustmentReason
	 */
	public int getAdjustmentReason() {
		return adjustmentReason;
	}
	/**
	 * @param adjustmentReason the adjustmentReason to set
	 */
	public void setAdjustmentReason(int adjustmentReason) {
		this.adjustmentReason = adjustmentReason;
	}
	/**
	 * @return the adjustmentComments
	 */
	public String getAdjustmentComments() {
		return AdjustmentComments;
	}
	/**
	 * @param adjustmentComments the adjustmentComments to set
	 */
	public void setAdjustmentComments(String adjustmentComments) {
		AdjustmentComments = adjustmentComments;
	}
	/**
	 * @return the fieldId
	 */
	public int getFieldId() {
		return fieldId;
	}
	/**
	 * @param fieldId the fieldId to set
	 */
	public void setFieldId(int fieldId) {
		this.fieldId = fieldId;
	}
	/**
	 * @return the product
	 */
	public int getProduct() {
		return product;
	}
	/**
	 * @param product the product to set
	 */
	public void setProduct(int product) {
		this.product = product;
	}
	/**
	 * @return the entryType
	 */
	public String getEntryType() {
		return entryType;
	}
	/**
	 * @param entryType the entryType to set
	 */
	public void setEntryType(String entryType) {
		this.entryType = entryType;
	}
	/**
	 * @return the productCode
	 */
	public int getProductCode() {
		return productCode;
	}
	/**
	 * @param productCode the productCode to set
	 */
	public void setProductCode(int productCode) {
		this.productCode = productCode;
	}
	/**
	 * @return the cardAssociation
	 */
	public String getCardAssociation() {
		return cardAssociation;
	}
	/**
	 * @param cardAssociation the cardAssociation to set
	 */
	public void setCardAssociation(String cardAssociation) {
		this.cardAssociation = cardAssociation;
	}
	/**
	 * @return the fundingCurrencyCode
	 */
	public String getFundingCurrencyCode() {
		return fundingCurrencyCode;
	}
	/**
	 * @param fundingCurrencyCode the fundingCurrencyCode to set
	 */
	public void setFundingCurrencyCode(String fundingCurrencyCode) {
		this.fundingCurrencyCode = fundingCurrencyCode;
	}


}
