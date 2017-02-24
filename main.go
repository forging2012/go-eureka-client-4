package main

import "os"
import "github.com/ibrokethecloud/go-eureka-client/eureka"
import "github.com/ibrokethecloud/go-utils"
import "encoding/json"

var EUREKA_ENDPOINT = os.Getenv("EUREKA_ENDPOINT")

func main() {
  var JsonRegistrationInfo,err = json.Marshal(eureka.DefaultRegisterationInfo)
  utils.CheckError(err)

  eureka.RegisterClient(JsonRegistrationInfo)
}
