package eureka

import "github.com/ibrokethecloud/go-utils"
import "net/http"
import "time"
import  log "github.com/Sirupsen/logrus"


func RegisterClient(registerationInfo []byte){
  // Call Eureka end point and register //

  resp_code := 000

  for resp_code != 204 {

    req, err := http.NewRequest("POST", EUREKA_ENDPOINT, registerationInfo)
    req.Header.Set("Content-Type", "application/json")
    client := &http.Client{}
    resp, err := client.Do(req)

    utils.CheckError(err)

    defer resp.Body.Close()
    log.Infof("response Status: %v", resp.Status)
    log.Infof("response Headers: %v", resp.Header)

    body, _ := ioutil.ReadAll(resp.Body)
    log.Infof("response Body: %s", string(body))

    resp_code = resp.Status

     time.Sleep(time.Second * 10)

  }

}
