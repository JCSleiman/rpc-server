package main

import (
        "log"
        "net/rcp"

        "github.com/jcsleiman/rcp-server/service"
)

type Areas struct {
  Client *rpc.Client
}

func main() {
  //Hora de conectarse al Servidor
  client, err := rpc.DialHTTP("tcp", "127.0.0.1:8081")
  if err != nil {
    log.Fatal("Dialing:", err)
  }

  a := &Areas{Client: client}

  result, err := a.triangle(5.3,2.2)
  if err!= nil {
          log.Println("Error de cálculo: " + err.Error())
  } else {
    log.Printf("Area del triángulo: %f", result)
  }

  result, err := a.circle(2, nil)
  if err!= nil {
          log.Println("Error de cálculo: " + err.Error())
  } else {
    log.Printf("Area del circulo: %f", result)
  }

  result, err := a.square(5, nil)
  if err!= nil {
          log.Println("Error de cálculo: " + err.Error())
  } else {
    log.Printf("Area del cuadrado: %f", result)
  }

  result, err := a.rectangle(10.5, 3.5)
  if err!= nil {
          log.Println("Error de cálculo: " + err.Error())
  } else {
    log.Printf("Area del rectángulo: %f", result)
  }

  result, err := a.cylinder(3.8, 4.5)
  if err!= nil {
          log.Println("Error de cálculo: " + err.Error())
  } else {
    log.Printf("Area del cilindro: %f", result)
  }

  func (a *Areas) triangle(a, b float64) (float64, error) {

    args := service.Request{A: a, B: b}
    var response service.Response
    err := a.Client.Call("Areas.Triangle", args, &response)
    return response.Result, err
  }

  func (a *Areas) circle(a, b float64) (float64, error) {

    args := service.Request{A: a, B: b}
    var response service.Response
    err := a.Client.Call("Areas.Circle", args, &response)
    return response.Result, err
  }

  func (a *Areas) square(a, b float64) (float64, error) {

    args := service.Request{A: a, B: b}
    var response service.Response
    err := a.Client.Call("Areas.Square", args, &response)
    return response.Result, err
  }

  func (a *Areas) rectangle(a, b float64) (float64, error) {

    args := service.Request{A: a, B: b}
    var response service.Response
    err := a.Client.Call("Areas.Rectangle", args, &response)
    return response.Result, err
  }

  func (a *Areas) Cylinder(a, b float64) (float64, error) {

    args := service.Request{A: a, B: b}
    var response service.Response
    err := a.Client.Call("Areas.Cylinder", args, &response)
    return response.Result, err
  }
