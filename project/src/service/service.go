package service

import (
	"external"
	"kafka"
	"time"
)

var (
	settlementService = external.Settlement
	adjustService     = external.AdjustPaymnet
)

type Request struct {
	Id                   int       `json:"id"`
	Type                 string    `json:"type"`
	Date                 time.Time `json:"date"`
	RejectReason         string    `json:"rejectReason"`
	SettlementAdjustment time.Time `json:"settlementAdjustment"`
	GrossValue           int       `json:"grossValue"`
	UserId               string    `json:"userId"`
	TechnologyType       int       `json:"tecnologyType"`
	StatusRequest        string    `json:"requestStatus"`
	Customer             string    `json:"customer"`
	ModCustomer          int       `json:"modCustomer"`
	BatchId              int       `json:"batchId"`
	MovementType         int       `json:"movementType"`
	RefundStatus         int       `json:"refundStatus"`
	AdjustmentReason     int       `json:"adjustmentReason"`
	AdjustmentComments   string    `json:"AdjustmentComments"`
	FileId               int       `json:"fieldId"`
	Product              int       `json:"product"`
	EntryType            string    `json:"entryType"`
	ProductCode          int       `json:"productCode"`

	CardAssociation     string `json:"cardAssociation"`
	FundingCurrencyCode string `json:"fundingCurrencyCode"`
}

func Service(request Request) {

	settlementService.GetCodCurrencyByCustomer(request.Customer)

	settlementService.GetContractedProduct(request.Customer)

	adjustService.GetAdjustmentProduct(request.Customer)

	association, currencyCode, _ := adjustService.SetCardAssociationAndPrincipalToMovement(request.Customer)

	request.CardAssociation = association

	request.FundingCurrencyCode = currencyCode

	//request.TypeIndicator = typeIndicator

	kafka.Send("persistence", request, request.Id)
}
