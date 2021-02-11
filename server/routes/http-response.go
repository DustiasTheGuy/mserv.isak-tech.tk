package routes

// HTTPResponse struct is how data should be sent out
type HTTPResponse struct {
	Message interface{} `json:"message"` // An explanation of what went wrong
	Success bool        `json:"success"` // the request result
	Data    interface{} `json:"data"`    // any data associated with the response
}
