package eureka

import "github.com/ibrokethecloud/go-utils"
import "net/http"
import "time"
import  log "github.com/Sirupsen/logrus"
import "io/ioutil"
import "os"

var EUREKA_ENDPOINT = os.Getenv("EUREKA_ENDPOINT")
var APP_NAME = os.Getenv("APP_NAME")

// AppId comes from APP_NAME environment variable //

func HealthCheck(){
  // Call Eureka end point and register //

    endpoint := "http://"+EUREKA_ENDPOINT+"/eureka/v2/apps/"+APP_NAME+"/"+utils.GetHostName()

    for {
      req, err := http.NewRequest("PUT", endpoint, nil)
      req.Header.Set("Content-Type", "application/json")
      client := &http.Client{}
      resp, err := client.Do(req)

      utils.CheckError(err)

      defer resp.Body.Close()
      log.Infof("HealtCheck Status: %v", resp.Status)

      body, _ := ioutil.ReadAll(resp.Body)
      log.Infof("response Body: %s", string(body))

      resp.Status

       time.Sleep(time.Second * 30)
    }




}
