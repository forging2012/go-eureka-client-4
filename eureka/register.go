package eureka

import "github.com/ibrokethecloud/go-utils"
import "net/http"
import "time"
import  log "github.com/Sirupsen/logrus"
import "io/ioutil"
import "bytes"


func RegisterClient(registerationInfo []byte){
  // Call Eureka end point and register //

  resp_code := "000"

  endpoint := "http://"+EUREKA_ENDPOINT+"/eureka/apps/"+APP_NAME

  log.Infof("About to hit endpoint %s",endpoint)

  for resp_code != "204 No Content" {

    req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(registerationInfo))
    req.Header.Set("Content-Type", "application/json")
    client := &http.Client{
      Timeout: 10 * time.Second,
    }
    resp, err := client.Do(req)

    utils.CheckError(err)

    defer resp.Body.Close()
    log.Debugf("response Status: %v", resp.Status)
    log.Debugf("response Headers: %v", resp.Header)

    body, _ := ioutil.ReadAll(resp.Body)
    log.Debugf("response Body: %s", string(body))

    resp_code = resp.Status

     time.Sleep(time.Second * 10)

  }

}
