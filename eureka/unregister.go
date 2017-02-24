package eureka

import "github.com/ibrokethecloud/go-utils"
import "net/http"
import "time"
import  log "github.com/Sirupsen/logrus"
import "io/ioutil"
import "bytes"

func UnregisterClient(registerationInfo []byte, endpoint string){
  // Call Eureka end point and register //

  endpoint := "http://"+EUREKA_ENDPOINT+"/eureka/v2/apps/"+APP_NAME+"/"+utils.GetHostName()
  req, err := http.NewRequest("DELETE", endpoint, nil)
  req.Header.Set("Content-Type", "application/json")
  client := &http.Client{}
  resp, err := client.Do(req)

  utils.CheckError(err)

  defer resp.Body.Close()
  log.Infof("HealtCheck Status: %v", resp.Status)


}
