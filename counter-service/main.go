package main

import (
  "github.com/micro/go-micro/v2"
  "context"
  "counter-service/proto" // replace counter-service if you didn't name counter-service when go mod init
  "log"
  "strconv"
)

// CounterService ...
type CounterService struct {}

var number int = 1

// Count ... Implement interface left in proto/counter.pb.micro.go server part
func (g *CounterService) Count(ctx context.Context, req *counter.Request, res *counter.Response) error {
  log.Println("Counter service handle Count")
  res.Number = strconv.Itoa(number)
  number++
  return nil
}

func main() {
  service := micro.NewService(
    micro.Name("micro.service.counter"), // The service name to register in the registry
  )

  service.Init()

  // The 'RegisterCounterHandler' is implement in the proto/counter.pb.micro.go file
  counter.RegisterCounterHandler(service.Server(), &CounterService{})
  
  if err := service.Run(); err != nil {
    log.Print(err.Error())
    return
  }
}