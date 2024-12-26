package ventas

type SearchResponseConsulta struct {
	TempalteSearchResponse
	Data ResponseDataDetalles `json:"data"`
}

type ResponseDataConsulta struct {
	Content []ContentItemConsulta `json:"content"`
	TemplateResponseData
}

type ContentItemConsulta struct {
	ID     int64  `json:"id"`
	Codigo string `json:"codigo"`
	Motivo string `json:"motivo"`
}
