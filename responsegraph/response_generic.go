package responsegraph

type ResponseGeneric struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseGenericPaginate struct {
	Code     int         `json:"code"`
	Message  string      `json:"message"`
	Data     interface{} `json:"data"`
	Paginate Paginate    `json:"paginate"`
}

type Paginate struct {
	Page       int `json:"page"`
	Page_Sieze int `json:"page_size"`
	Total      int `json:"total"`
}
