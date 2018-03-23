package service

import (
	"log"
)

// This represents the available service over the network
type Areas int

// Lets begin with Triangle that give us the Area of this shape
func (a *Areas) Triangle(r *Request, rp *Response) error {

	log.Printf("Executing Triangle with args: %+v", rq)
	rp.Result = (rq.A * rq.B) / 2
	return nil
}

func (a *Areas) Circle(r *Request, rp *Response) error {

	log.Printf("Executing Circle with args: %+v", rq)
	rp.Result = 3.1415 * (A * A)
	return nil
}

func (a *Areas) Square(r *Request, rp *Response) error {

	log.Printf("Executing Square with args: %+v", rq)
	rp.Result = A * A
	return nil
}

func (a *Areas) Rectangle(r *Request, rp *Response) error {

	log.Printf("Executing Rectangle with args: %+v", rq)
	rp.Result = A * B
	return nil
}

func (a *Areas) Cylinder(r *Request, rp *Response) error {

	log.Printf("Executing Cylinder with args: %+v", rq)
	rp.Result = 2*3.1415*A*B + 2*3.1415*(A*A)
	return nil
}
