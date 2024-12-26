package ventas

type Filtro struct {
	Atributo      string     `json:"atributo,omitempty"`
	Valor         string     `json:"valor,omitempty"`
	Valor2        string     `json:"valor2,omitempty"`
	TipoFiltro    TipoFiltro `json:"tipoFiltro,omitempty"`
	TipoDato      TipoDato   `json:"tipoDato,omitempty"`
	NombreColumna string     `json:"nombreColumna,omitempty"`
	Valores       []string   `json:"valores,omitempty"`
	Clase         string     `json:"clase,omitempty"`
	Operador      Operador   `json:"operador,omitempty"`
	SubGrupo      string     `json:"subGrupo,omitempty"`
}

type SearchRequest struct {
	Filtros            []Filtro `json:"filtros,omitempty"`
	ColumnaOrdenar     string   `json:"columnaOrdenar,omitempty"`
	Pagina             int      `json:"pagina,omitempty"`
	RegistrosPorPagina int      `json:"registrosPorPagina,omitempty"`
	Orden              Orden    `json:"orden,omitempty"`
	Canal              Canal    `json:"canal,omitempty"`
	RegistroInicial    int      `json:"registroInicial,omitempty"`
}
