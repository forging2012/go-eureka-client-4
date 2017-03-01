package main


import "github.com/ibrokethecloud/go-eureka-client/eureka"
import "time"
import  log "github.com/Sirupsen/logrus"
import "sync"
//var EUREKA_ENDPOINT = os.Getenv("EUREKA_ENDPOINT")
//var APP_NAME = os.Getenv("APP_NAME")

func main() {

  var wg sync.WaitGroup
  wg.Add(1)
  // Function to call manage eureka //
  go func() {
    defer wg.Done()
    eureka.ManageEureka()
   }()

  // Main code comes in here //
  for {
    log.Error("Running the service in a loop")
    time.Sleep(time.Second * 30)
  }

  wg.Wait()
}
