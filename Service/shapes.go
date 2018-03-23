package service

// These are all the operations that are implemented on the server
type Operations interface {
	Triangle(*Request, *Response) error
	Circle(*Request, *Response) error
	Square(*Request, *Response) error
	Rectangle(*Request, *Response) error
	Cylinder(*Request, *Response) error
}
