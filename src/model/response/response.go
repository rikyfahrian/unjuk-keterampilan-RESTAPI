package response

type Response struct {
	Status  int         `json:"status"`
	Message interface{} `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
