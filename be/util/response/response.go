package response

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}
