package main

import "os"
import "github.com/ibrokethecloud/go-eureka-client/eureka"
import "github.com/ibrokethecloud/go-utils"
import "encoding/json"
import "fmt"

var EUREKA_ENDPOINT = os.Getenv("EUREKA_ENDPOINT")
var APP_NAME = os.Getenv("APP_NAME")

func main() {
  var JsonRegistrationInfo,err = json.Marshal(eureka.DefaultRegisterationInfo)
  utils.CheckError(err)
  fmt.Println(string(JsonRegistrationInfo))

  // Register the client
  //eureka.RegisterClient(JsonRegistrationInfo)

  // Starts a go func to listen for signal
  eureka.HandleSigterm()

  // Sending heart beat signals //
  eureka.HealthCheck()

}
