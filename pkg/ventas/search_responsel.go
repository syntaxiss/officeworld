package ventas

import "time"

type SearchResponseLista struct {
	TempalteSearchResponse
	Data ResponseDataLista `json:"data"`
}

type ResponseDataLista struct {
	Content []ContentItemLista `json:"content"`
	TemplateResponseData
}

type ContentItemLista struct {
	ID                            int64     `json:"id"`
	Prefijo                       string    `json:"prefijo"`
	Motivo                        string    `json:"motivo"`
	IDPrefijo                     int64     `json:"idPrefijo"`
	Numero                        int       `json:"numero"`
	Fecha                         time.Time `json:"fecha"`
	Empresa                       string    `json:"empresa"`
	IDEmpresa                     int64     `json:"idEmpresa"`
	TerceroExterno                string    `json:"terceroExterno"`
	PrimerNombreTerceroExterno    string    `json:"primerNombreTerceroExterno"`
	SegundoNombreTerceroExterno   string    `json:"segundoNombreTerceroExterno"`
	PrimerApellidoTerceroExterno  string    `json:"primerApellidoTerceroExterno"`
	SegundoApellidoTerceroExterno string    `json:"segundoApellidoTerceroExterno"`
	IDTerceroExterno              int64     `json:"idTerceroExterno"`
	TerceroInterno                string    `json:"terceroInterno"`
	PrimerNombreTerceroInterno    string    `json:"primerNombreTerceroInterno"`
	SegundoNombreTerceroInterno   string    `json:"segundoNombreTerceroInterno"`
	PrimerApellidoTerceroInterno  string    `json:"primerApellidoTerceroInterno"`
	SegundoApellidoTerceroInterno string    `json:"segundoApellidoTerceroInterno"`
	IDTerceroInterno              int64     `json:"idTerceroInterno"`
	FormaPago                     string    `json:"formaPago"`
	IDFormaPago                   int64     `json:"idFormaPago"`
	Concepto                      string    `json:"concepto"`
	Responsable                   string    `json:"responsable"`
}
