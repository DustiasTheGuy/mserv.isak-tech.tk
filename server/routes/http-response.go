package routes

// HTTPResponse struct is how data should be sent out
type HTTPResponse struct {
	Message interface{} `json:"message"` // What's the status of the request?
	Success bool        `json:"success"` // Did the request succeed?
	Data    interface{} `json:"data"`    // any data associated with the response
}
