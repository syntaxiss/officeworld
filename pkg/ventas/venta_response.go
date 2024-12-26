package ventas

type DocumentVentaResponse struct {
	Prefijo                     Prefijo                      `json:"prefijo"`
	DocumentoTipo               DocumentoTipo                `json:"documentoTipo"`
	Empresa                     Empresa                      `json:"empresa"`
	TerceroExterno              TerceroExterno               `json:"terceroExterno"`
	TerceroInterno              TerceroInterno               `json:"terceroInterno"`
	DireccionTerceroExterno     DireccionTerceroExterno      `json:"direccionTerceroExterno"`
	FormaPago                   FormaPago                    `json:"formaPago"`
	Moneda                      Moneda                       `json:"moneda"`
	DocumentoEncabezadosEstados []DocumentoEncabezadosEstado `json:"documentoEncabezadosEstados"`
	DocAsociado                 DocAsociado                  `json:"docAsociado"`
	DocumentoMotivoGeneracion   DocumentoMotivoGeneracion    `json:"documentoMotivoGeneracion"`
	DocumentoTipoNotaCredito    DocumentoTipoNotaCredito     `json:"documentoTipoNotaCredito"`
	ElaboradoPor                ElaboradoPor                 `json:"elaboradoPor"`
	HistorialPrefijoPos         HistorialPrefijoPos          `json:"historialPrefijoPos"`

	ID                            int64  `json:"id"`
	Fecha                         string `json:"fecha"`
	Numero                        int64  `json:"numero"`
	Concepto                      string `json:"concepto"`
	Trm                           int    `json:"trm"`
	SenContabilizado              bool   `json:"senContabilizado"`
	NumeroExterno                 string `json:"numeroExterno"`
	PrefijoExterno                string `json:"prefijoExterno"`
	TieneRenglones                bool   `json:"tieneRenglones"`
	SenGenerado                   bool   `json:"senGenerado"`
	TieneResolucionElectronica    bool   `json:"tieneResolucionElectronica"`
	AtributosRequeridos           bool   `json:"atributosRequeridos"`
	AtributosEncabezado           bool   `json:"atributosEncabezado"`
	AtributosRenglones            bool   `json:"atributosRenglones"`
	SenExisteDocumentoAsociado    bool   `json:"senExisteDocumentoAsociado"`
	SenSectorTransporte           bool   `json:"senSectorTransporte"`
	ValidarContabilizacion        bool   `json:"validarContabilizacion"`
	ValidarDescuadrados           bool   `json:"validarDescuadrados"`
	SalidaAlmacenAsociada         int64  `json:"salidaAlmacenAsociada"`
	TieneSalidaAlmacenAsociada    bool   `json:"tieneSalidaAlmacenAsociada"`
	TieneProductosConFichaTecnica bool   `json:"tieneProductosConFichaTecnica"`
}

type Prefijo struct {
	ID          int64  `json:"id"`
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
}

type DocumentoTipo struct {
	ID              int64  `json:"id"`
	CodigoDocumento string `json:"codigoDocumento"`
	NombreDocumento string `json:"nombreDocumento"`
}

type Empresa struct {
	Tercero               Tercero                 `json:"tercero"`
	UbicacionCiudad       UbicacionCiudad         `json:"ubicacionCiudad"`
	ResponsabilidadFiscal []ResponsabilidadFiscal `json:"responsabilidadFiscal"`

	ID                         int64  `json:"id"`
	Nombre                     string `json:"nombre"`
	Prefijo                    string `json:"prefijo"`
	Identificacion             string `json:"identificacion"`
	DigitoVerificacion         int64  `json:"digitoVerificacion"`
	DigitoVerificacionGenerado bool   `json:"digitoVerificacionGenerado"`
	SenPrincipal               bool   `json:"senPrincipal"`
	NumeroEstablecimientos     int64  `json:"numeroEstablecimientos"`
	InfoTributariaAdicional    string `json:"infoTributariaAdicional"`
}
type Tercero struct {
	TerceroTipoIdentificacion TerceroTipoIdentificacion `json:"terceroTipoIdentificacion"`
	Ciudad                    Ciudad                    `json:"ciudad"`
	TerceroTipos              []TerceroTipo             `json:"terceroTipos"`

	ID                         int64  `json:"id"`
	Identificacion             string `json:"identificacion"`
	PrimerNombre               string `json:"primerNombre"`
	SegundoNombre              string `json:"segundoNombre"`
	PrimerApellido             string `json:"primerApellido"`
	SegundoApellido            string `json:"segundoApellido"`
	NombreCompleto             string `json:"nombreCompleto"`
	Codigo                     string `json:"codigo"`
	DigitoVerificacion         int64  `json:"digitoVerificacion"`
	TarifaICA                  int    `json:"tarifaICA"`
	DigitoVerificacionGenerado bool   `json:"digitoVerificacionGenerado"`
	SenActivo                  bool   `json:"senActivo"`
	SenManejaNomina            bool   `json:"senManejaNomina"`
	Atributos                  bool   `json:"atributos"`
	AplicaICAVentas            bool   `json:"aplicaICAVentas"`
}

type TerceroTipoIdentificacion struct {
	UbicacionPais UbicacionPais `json:"ubicacionPais"`

	ID           int64  `json:"id"`
	Abreviatura  string `json:"abreviatura"`
	Nombre       string `json:"nombre"`
	Codigo       string `json:"codigo"`
	SenEsEmpresa bool   `json:"senEsEmpresa"`
	SenManejaDV  bool   `json:"senManejaDV"`
}

type UbicacionPais struct {
	ID         int64  `json:"id"`
	Nombre     string `json:"nombre"`
	Codigo     string `json:"codigo"`
	CodAlterno string `json:"codAlterno"`
}

type Ciudad struct {
	UbicacionDepartamento UbicacionDepartamento `json:"ubicacionDepartamento"`

	ID         int64  `json:"id"`
	Nombre     string `json:"nombre"`
	Codigo     string `json:"codigo"`
	SenSistema bool   `json:"senSistema"`
}

type UbicacionDepartamento struct {
	UbicacionPais UbicacionPais `json:"ubicacionPais"`

	ID         int64  `json:"id"`
	Nombre     string `json:"nombre"`
	Codigo     string `json:"codigo"`
	SenSistema bool   `json:"senSistema"`
}

type TerceroTipo struct {
	ID     int64  `json:"id"`
	Codigo string `json:"codigo"`
	Nombre string `json:"nombre"`
}

type UbicacionCiudad struct {
	UbicacionDepartamento UbicacionDepartamento `json:"ubicacionDepartamento"`

	ID         int64  `json:"id"`
	Nombre     string `json:"nombre"`
	Codigo     string `json:"codigo"`
	SenSistema bool   `json:"senSistema"`
}

type ResponsabilidadFiscal struct {
	ID          int64  `json:"id"`
	Codigo      string `json:"codigo"`
	Significado string `json:"significado"`
}

type TerceroExterno struct {
	TerceroTipos                  []TerceroTipo                 `json:"terceroTipos"`
	TerceroVendedorPredeterminado TerceroVendedorPredeterminado `json:"terceroVendedorPredeterminado"`
	FormaPagoPredeterminada       FormaPago                     `json:"formaPagoPredeterminada"`
	ListaPrecioPredeterminada     ListaPrecio                   `json:"listaPrecioPredeterminada"`

	ID                                        int64  `json:"id"`
	NombreCompleto                            string `json:"nombreCompleto"`
	PrimerNombre                              string `json:"primerNombre"`
	SegundoNombre                             string `json:"segundoNombre"`
	PrimerApellido                            string `json:"primerApellido"`
	SegundoApellido                           string `json:"secondApellido"`
	TerceroTipoIdentificacion                 string `json:"terceroTipoIdentificacion"`
	Identificacion                            string `json:"identificacion"`
	TerceroTipoContribuyente                  string `json:"terceroTipoContribuyente"`
	TerceroClasificacionAdministradorImpuesto string `json:"terceroClasificacionAdministradorImpuesto"`
	SenCupoCredito                            bool   `json:"senCupoCredito"`
}

type TerceroVendedorPredeterminado struct {
	ID                        int64  `json:"id"`
	NombreCompleto            string `json:"nombreCompleto"`
	PrimerNombre              string `json:"primerNombre"`
	SegundoNombre             string `json:"segundoNombre"`
	PrimerApellido            string `json:"primerApellido"`
	SegundoApellido           string `json:"segundoApellido"`
	TerceroTipoIdentificacion string `json:"terceroTipoIdentificacion"`
	Identificacion            string `json:"identificacion"`
}

type FormaPago struct {
	ID                   int64  `json:"id"`
	Codigo               string `json:"codigo"`
	Nombre               string `json:"nombre"`
	SenManejaCupoCredito bool   `json:"senManejaCupoCredito"`
	IdGrupo              int64  `json:"idGrupo"`
}

type ListaPrecio struct {
	ID     int64  `json:"id"`
	Nombre string `json:"nombre"`
}

type TerceroInterno struct {
	ID                        int64  `json:"id"`
	NombreCompleto            string `json:"nombreCompleto"`
	PrimerNombre              string `json:"primerNombre"`
	SegundoNombre             string `json:"segundoNombre"`
	PrimerApellido            string `json:"primerApellido"`
	SegundoApellido           string `json:"segundoApellido"`
	TerceroTipoIdentificacion string `json:"terceroTipoIdentificacion"`
	Identificacion            string `json:"identificacion"`
}

type DireccionTerceroExterno struct {
	ID                int64  `json:"id"`
	Direccion         string `json:"direccion"`
	Ciudad            string `json:"ciudad"`
	TelefonoPrincipal string `json:"telefonoPrincipal"`
	Sucursal          string `json:"sucursal"`
	SenPrincipal      bool   `json:"senPrincipal"`
}

type Moneda struct {
	ID                 int64  `json:"id"`
	Codigo             string `json:"codigo"`
	Nombre             string `json:"nombre"`
	SeparadorDecimales string `json:"separadorDecimales"`
	SeparadorMiles     string `json:"separadorMiles"`
	CantidadDecimales  int64  `json:"cantidadDecimales"`
	Simbolo            string `json:"simbolo"`
	TieneTRM           bool   `json:"tieneTRM"`
}

type DocumentoEncabezadosEstado struct {
	ID     int64  `json:"id"`
	Codigo Codigo `json:"codigo"`
	Nombre string `json:"nombre"`
}

type DocAsociado struct {
	Prefijo       Prefijo       `json:"prefijo"`
	Moneda        Moneda        `json:"moneda"`
	DocumentoTipo DocumentoTipo `json:"documentoTipo"`

	ID        int64  `json:"id"`
	Numero    int    `json:"numero"`
	Fecha     string `json:"fecha"`
	Trm       int    `json:"trm"`
	VersionFe int    `json:"versionFe"`
	Cufe      string `json:"cufe"`
}

type DocumentoMotivoGeneracion struct {
	ID     int64  `json:"id"`
	Codigo string `json:"codigo"`
	Motivo string `json:"motivo"`
}

type DocumentoTipoNotaCredito struct {
	ID                 int64  `json:"id"`
	Codigo             string `json:"codigo"`          //maxLength: 5 minLength: 0
	TipoNotaCredito    string `json:"tipoNotaCredito"` //maxLength: 15 minLength: 0
	EfectoEnInventario int    `json:"efectoEnInventario"`
}

type ElaboradoPor struct {
	ID            int64          `json:"id"`
	Nombres       string         `json:"nombres"`
	Activo        bool           `json:"activo"`
	Email         string         `json:"email"`
	Canal         string         `json:"canal"`
	UsuarioTokens []UsuarioToken `json:"usuarioTokens"`
}

type UsuarioToken struct {
	ID               int64  `json:"id"`
	JWT              string `json:"jwt"` //maxLength: 500 minLength: 0
	FechaRegistro    string `json:"fechaRegistro"`
	FechaVencimiento string `json:"fechaVencimiento"`
	SesTokenActivo   bool   `json:"sesTokenActivo"`
}

type HistorialPrefijoPos struct {
	Numero        int64  `json:"numero"`
	NombrePrefijo string `json:"nombrePrefijo"`
}
