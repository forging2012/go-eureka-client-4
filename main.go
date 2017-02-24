package main

import "fmt"
import "./type"
import "github.com/ibrokethecloud/go-utils"
import "encoding/json"

func main() {
  var JsonRegistrationInfo,err = json.Marshal(eureka.DefaultRegisterationInfo)
  utils.CheckError(err)
  fmt.Println(string(JsonRegistrationInfo))
}
