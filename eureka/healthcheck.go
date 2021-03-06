package eureka

import "github.com/ibrokethecloud/go-utils"
import "net/http"
import "time"
import  log "github.com/Sirupsen/logrus"
import "io/ioutil"


// AppId comes from APP_NAME environment variable //

func HealthCheck(){
  // Call Eureka end point and register //

    endpoint := "http://"+EUREKA_ENDPOINT+"/eureka/apps/"+APP_NAME+"/"+instanceId

    for {
      req, err := http.NewRequest("PUT", endpoint, nil)
      req.Header.Set("Content-Type", "application/json")
      client := &http.Client{
        Timeout: 10 * time.Second,
      }
      resp, err := client.Do(req)

      utils.CheckError(err)

      defer resp.Body.Close()
      log.Debugf("HealtCheck Status: %v", resp.Status)

      body, _ := ioutil.ReadAll(resp.Body)
      log.Debugf("response Body: %s", string(body))

      time.Sleep(time.Second * 30)
    }




}
