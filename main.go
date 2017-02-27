package main

import "os"
import "github.com/ibrokethecloud/go-eureka-client/eureka"

import "fmt"

var EUREKA_ENDPOINT = os.Getenv("EUREKA_ENDPOINT")
var APP_NAME = os.Getenv("APP_NAME")

func main() {

  // Function to call manage eureka //
  eureka.ManageEureka()

}
