package structs

type UpdateOrderLoan struct {
	IdOrder     int `json:"id_order"`
	Plafond     int `json:"plafond"`
	DownPayment int `json:"down_payment"`
	Installment int `json:"installment"`
	Tenor       int `json:"tenor"`
	UpdatedBy   int `json:"updated_by"`
}
