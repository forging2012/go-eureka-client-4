package eureka

import "github.com/ibrokethecloud/go-utils"
import "encoding/json"
import "fmt"
import "os"

// Global variable definitions based on env variable declarations //
var EUREKA_ENDPOINT = os.Getenv("EUREKA_ENDPOINT")
var APP_NAME = os.Getenv("APP_NAME")
var PORT = os.Getenv("PORT")

// Wrapper function to manage the eureka registry lifecycle
func ManageEureka() {

  var JsonRegistrationInfo,err = json.Marshal(DefaultRegisterationInfo)
  utils.CheckError(err)
  fmt.Println(string(JsonRegistrationInfo))

  // Register the client
  RegisterClient(JsonRegistrationInfo)

  // Starts a go func to listen for signal
  HandleSigterm()

  // Sending heart beat signals //
  HealthCheck()
}
