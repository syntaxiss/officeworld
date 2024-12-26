package ventas

type TipoFiltro string
type TipoDato string
type Operador string
type Orden string
type Canal string
type Codigo string

const (
	IGUAL          TipoFiltro = "IGUAL"
	CONTIENE       TipoFiltro = "CONTIENE"
	MENOR_QUE      TipoFiltro = "MENOR_QUE"
	MAYOR_QUE      TipoFiltro = "MAYOR_QUE"
	EMPIEZA_CON    TipoFiltro = "EMPIEZA_CON"
	MAYOR_O_IGUAL  TipoFiltro = "MAYOR_O_IGUAL"
	MENOR_O_IGUAL  TipoFiltro = "MENOR_O_IGUAL"
	TERMINA_CON    TipoFiltro = "TERMINA_CON"
	ENTRE          TipoFiltro = "ENTRE"
	IS_NULL        TipoFiltro = "IS_NULL"
	LENGTH         TipoFiltro = "LENGTH"
	DIFERENTE      TipoFiltro = "DIFERENTE"
	IS_NOT_NULL    TipoFiltro = "IS_NOT_NULL"
	LENGTH_IGUAL   TipoFiltro = "LENGTH_IGUAL"
	NO_EMPIEZA_CON TipoFiltro = "NO_EMPIEZA_CON"

	STRING    TipoDato = "STRING"
	BOOLEAN   TipoDato = "BOOLEAN"
	NUMERIC   TipoDato = "NUMERIC"
	FECHA     TipoDato = "FECHA"
	LONG      TipoDato = "LONG"
	LISTA     TipoDato = "LISTA"
	IN        TipoDato = "IN"
	NOT_IN    TipoDato = "NOT_IN"
	ENUM      TipoDato = "ENUM"
	ENTIDAD   TipoDato = "ENTIDAD"
	ARRAY_INT TipoDato = "ARRAY_INT"

	AND Operador = "AND"
	OR  Operador = "OR"

	ASC  Orden = "ASC"
	DESC Orden = "DESC"

	WEB    Canal = "WEB"
	MOBILE Canal = "MOBILE"
	POS    Canal = "POS"
	API    Canal = "API"

	BLOQUEADO                   Codigo = "BLOQUEADO"
	ANULADO                     Codigo = "ANULADO"
	VERIFICADO                  Codigo = "VERIFICADO"
	EXPORTADO                   Codigo = "EXPORTADO"
	PRESENTADO_ELECTRONICAMENTE Codigo = "PRESENTADO_ELECTRONICAMENTE"
	AJUSTADO_AL_INVENTARIO      Codigo = "AJUSTADO_AL_INVENTARIO"
	EN_PROCESO_DE_CONTABILIZAR  Codigo = "EN_PROCESO_DE_CONTABILIZAR"
	TRASLADO_INVENTARIO         Codigo = "TRASLADO_INVENTARIO"
)
