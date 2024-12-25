package client

type Renglon struct {
	IdInventario  int     `json:"idInventario,omitempty"`
	UnidadMedida  string  `json:"unidadMedida,omitempty"`
	Cantidad      float64 `json:"cantidad,omitempty"`
	ValorUnitario float64 `json:"valorUnitario,omitempty"`
	IdBodega      int     `json:"idBodega,omitempty"`
	PorDescuento  float64 `json:"porDescuento,omitempty"`
	Concepto      string  `json:"concepto,omitempty"`
}

type Request struct {
	Reglones                 []*Renglon `json:"reglones,omitempty"`
	Fecha                    string     `json:"fecha,omitempty"`
	Prefijo                  int        `json:"prefijo,omitempty"`
	DocumentoTipo            string     `json:"documentoTipo,omitempty"`
	IdEmpresa                int        `json:"idEmpresa,omitempty"`
	IdTerceroExterno         int        `json:"idTerceroExterno,omitempty"`
	IdTerceroInterno         int        `json:"idTerceroInterno,omitempty"`
	IdFormaPago              int        `json:"idFormaPago,omitempty"`
	IdMoneda                 int        `json:"idMoneda,omitempty"`
	Trm                      int        `json:"trm,omitempty"`
	PorcentajeDescuento      bool       `json:"porcentajeDescuento,omitempty"`
	PorcentajeTodosRenglones bool       `json:"porcentajeTodosRenglones,omitempty"`
	ValDescuento             int        `json:"valDescuento,omitempty"`
	IdFactura                int        `json:"idFactura,omitempty"`
	IdDetalles               []int      `json:"idDetalles,omitempty"`
}
