package external

var (
	Settlement    = SettlementImpl{}
	AdjustPaymnet = AdjustPaymentImpl{}
)

type SettlementService interface {
	GetContractedProduct(customer string) int
	GetCodCurrencyByCustomer(customer string) int
}

type AdjustPayment interface {
	GetAdjustmentProduct(customer string) int
	SetCardAssociationAndPrincipalToMovement(customer string) (string, string, string)
}

type SettlementImpl struct {
}

type AdjustPaymentImpl struct {
}

func (*SettlementImpl) GetContractedProduct(customer string) int {
	return 40
}

func (*SettlementImpl) GetCodCurrencyByCustomer(customer string) int {
	return 986
}

func (*AdjustPaymentImpl) GetAdjustmentProduct(customer string) int {
	return 40
}

func (*AdjustPaymentImpl) SetCardAssociationAndPrincipalToMovement(customer string) (string, string, int) {
	return "VISA", "986", 1
}
