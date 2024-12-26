package ventas

type BillElectronicResponse struct {
	TempalteSearchResponse
	Data ResponseDataBillElectronic `json:"data"`
}

type ResponseDataBillElectronic struct {
	Error       bool   `json:"error"`
	Mensaje     string `json:"mensaje"`
	Entero      int64  `json:"entero"`
	CodError    string `json:"codError"`
	CodigoError string `json:"codigoError"`
	Cufe        string `json:"cufe"`
}
