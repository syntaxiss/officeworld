package ventas

import "time"

type SearchResponseDetalles struct {
	TempalteSearchResponse
	Data ResponseDataDetalles `json:"data"`
}

type ResponseDataDetalles struct {
	Content []ContentItemDetalles `json:"content"`
	TemplateResponseData
}

type ContentItemDetalles struct {
	ID                  int64            `json:"id"`
	Inventario          Inventario       `json:"inventario"`
	InventarioBodega    InventarioBodega `json:"inventarioBodega"`
	UnidadMedida        UnidadMedida     `json:"unidadMedida"`
	Cantidad            int              `json:"cantidad"`
	ValorUnitario       float64          `json:"valorUnitario"`
	ValorTotalRenglon   float64          `json:"valorTotalRenglon"`
	PorcentajeDescuento float64          `json:"porcentajeDescuento"`
	CentroCosto         CentroCosto      `json:"centroCosto"`
	Tercero             TerceroD         `json:"tercero"`
	Concepto            string           `json:"concepto"`
	EstaCruzado         bool             `json:"estaCruzado"`
	SenObsequio         bool             `json:"senObsequio"`
	Lotes               []Lote           `json:"lotes"`
	Precio              float64          `json:"precio"`
	PrecioRenglon       float64          `json:"precioRenglon"`
}

type Inventario struct {
	ID             int64  `json:"id"`
	Descripcion    string `json:"descripcion"`
	Codigo         string `json:"codigo"`
	Clasificacion  string `json:"clasificacion"`
	Meses          string `json:"meses"`
	SenManejaLotes bool   `json:"senManejaLotes"`
}

type InventarioBodega struct {
	ID                int64  `json:"id"`
	Nombre            string `json:"nombre"`
	SenActiva         bool   `json:"senActiva"`
	SenPredeterminada bool   `json:"senPredeterminada"`
	Codigo            string `json:"codigo"`
}

type UnidadMedida struct {
	ID                                 int64                              `json:"id"`
	Nombre                             string                             `json:"nombre"`
	Codigo                             string                             `json:"codigo"`
	SenDeSistema                       bool                               `json:"senDeSistema"`
	UnidadMedidaTipo                   UnidadMedidaTipo                   `json:"unidadMedidaTipo"`
	UnidadMedidaAdministradorImpuestos UnidadMedidaAdministradorImpuestos `json:"unidadMedidaAdministradorImpuestos"`
	UnidadMedidaBaseID                 int                                `json:"unidadMedidaBaseId"`
	Factor                             float64                            `json:"factor"`
}

type UnidadMedidaTipo struct {
	ID     int64  `json:"id"`
	Nombre string `json:"nombre"`
}

type UnidadMedidaAdministradorImpuestos struct {
	ID            int64          `json:"id"`
	Codigo        string         `json:"codigo"`
	Descripcion   string         `json:"descripcion"`
	UbicacionPais UbicacionPaisD `json:"ubicacionPais"`
}

type UbicacionPaisD struct {
	ID         int64  `json:"id"`
	Nombre     string `json:"nombre"`
	Codigo     string `json:"codigo"`
	CodAlterno string `json:"codAlterno"`
}

type CentroCosto struct {
	ID     int64  `json:"id"`
	Nombre string `json:"nombre"`
	Codigo string `json:"codigo"`
}

type TerceroD struct {
	ID              int64  `json:"id"`
	NombreCompleto  string `json:"nombreCompleto"`
	PrimerNombre    string `json:"primerNombre"`
	SegundoNombre   string `json:"segundoNombre"`
	PrimerApellido  string `json:"primerApellido"`
	SegundoApellido string `json:"segundoApellido"`
}

type Lote struct {
	ID                                        int64       `json:"id"`
	Numero                                    string      `json:"numero"`
	Descripcion                               string      `json:"descripcion"`
	SenActivo                                 bool        `json:"senActivo"`
	SenVence                                  bool        `json:"senVence"`
	FechaVencimiento                          time.Time   `json:"fechaVencimiento"`
	Observacion                               string      `json:"observacion"`
	TerceroPojo                               interface{} `json:"terceroPojo"`
	IDInventarioLoteDocumentoMovimientoActual int         `json:"idInventarioLoteDocumentoMovimientoActual"`
	Cantidad                                  float64     `json:"cantidad"`
}
