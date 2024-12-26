package ventas

type TempalteSearchResponse struct {
	Status           interface{} `json:"status"`
	UserMessage      string      `json:"userMessage"`
	DeveloperMessage interface{} `json:"developerMessage"`
	ErrorCode        string      `json:"errorCode"`
	MoreInfo         string      `json:"moreInfo"`
}

type SearchResponse struct {
	Status           Status      `json:"status"`
	UserMessage      string      `json:"userMessage"`
	DeveloperMessage interface{} `json:"developerMessage"`
	ErrorCode        string      `json:"errorCode"`
	MoreInfo         string      `json:"moreInfo"`
	Data             interface{} `json:"data"`
}

type VentaResponse struct {
	TempalteSearchResponse
	Data DocumentVentaResponse `json:"data"`
}

type TemplateResponseData struct {
	Sort             []SortItem `json:"sort"`
	TotalElements    int64      `json:"totalElements"`
	TotalPages       int        `json:"totalPages"`
	Size             int        `json:"size"`
	Number           int        `json:"number"`
	NumberOfElements int        `json:"numberOfElements"`
	Pageable         Pageable   `json:"pageable"`
	First            bool       `json:"first"`
	Last             bool       `json:"last"`
	Empty            bool       `json:"empty"`
}

type SortItem struct {
	Direction    string `json:"direction"`
	NullHandling string `json:"nullHandling"`
	Ascending    bool   `json:"ascending"`
	Property     string `json:"property"`
	IgnoreCase   bool   `json:"ignoreCase"`
}

type Pageable struct {
	Offset     int64      `json:"offset"`
	Sort       []SortItem `json:"sort"`
	PageSize   int        `json:"pageSize"`
	PageNumber int        `json:"pageNumber"`
	Paged      bool       `json:"paged"`
	Unpaged    bool       `json:"unpaged"`
}
