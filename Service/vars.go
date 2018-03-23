package service

// These are the arguments for the server
type Request struct {
	A, B, C float64
}

// This is the response from the server
type Response struct {
	Result float64
}
