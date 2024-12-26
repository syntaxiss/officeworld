package ventas

type Renglon struct {
	ID            int64  `json:"id,omitempty"`
	IdInventario  int64  `json:"idInventario,omitempty"`
	UnidadMedida  string `json:"unidadMedida,omitempty"`
	Cantidad      int64  `json:"cantidad,omitempty"`
	ValorUnitario int64  `json:"valorUnitario,omitempty"`
	IdBodega      int64  `json:"idBodega,omitempty"`
	PorDescuento  int64  `json:"porDescuento,omitempty"`
	Concepto      string `json:"concepto,omitempty"`
}

type DocumentVentaRequest struct {
	Reglones                 []*Renglon `json:"reglones,omitempty"`
	ID                       int64      `json:"id,omitempty"`
	Fecha                    string     `json:"fecha,omitempty"`
	Prefijo                  int64      `json:"prefijo,omitempty"`
	DocumentoTipo            string     `json:"documentoTipo,omitempty"`
	IdEmpresa                int64      `json:"idEmpresa,omitempty"`
	IdTerceroExterno         int64      `json:"idTerceroExterno,omitempty"`
	IdTerceroInterno         int64      `json:"idTerceroInterno,omitempty"`
	IdFormaPago              int64      `json:"idFormaPago,omitempty"`
	IdMoneda                 int64      `json:"idMoneda,omitempty"`
	Trm                      int64      `json:"trm,omitempty"`
	PorcentajeDescuento      bool       `json:"porcentajeDescuento,omitempty"`
	PorcentajeTodosRenglones bool       `json:"porcentajeTodosRenglones,omitempty"`
	ValDescuento             int64      `json:"valDescuento,omitempty"`
	IdFactura                int64      `json:"idFactura,omitempty"`
	IdDetalles               []int64    `json:"idDetalles,omitempty"`
}
