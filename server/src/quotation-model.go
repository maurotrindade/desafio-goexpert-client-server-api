package src

type Quotation struct {
	Bid string `json:"bid"`
}

type QuotationRepository interface {
	Save(quotation *Quotation) error
}
