package eureka

import "github.com/ibrokethecloud/go-utils"
import "net/http"
import  log "github.com/Sirupsen/logrus"
import "os"
import "os/signal"
import "syscall"
import "time"

func UnregisterClient(){
  // Call Eureka end point and register //

  endpoint := "http://"+EUREKA_ENDPOINT+"/eureka/apps/"+APP_NAME+"/"+instanceId
  req, err := http.NewRequest("DELETE", endpoint, nil)
  req.Header.Set("Content-Type", "application/json")
  client := &http.Client{
    Timeout: 10 * time.Second,
  }
  resp, err := client.Do(req)

  utils.CheckError(err)

  defer resp.Body.Close()
  log.Debugf("HealtCheck Status: %v", resp.Status)


}


func HandleSigterm() {

// Got from here https://golang.org/pkg/os/signal/#Notify //

  c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
  signal.Notify(c, syscall.SIGTERM)

           go func() {                           // Start an anonymous func running in a goroutine
                   <-c                           // that will block until a message is recieved on
                   log.Infof("Received SIGTERM. Shutting down now")
                   UnregisterClient()           // the channel. When that happens, perform Eureka
                   os.Exit(1)                    // deregistration and exit program.
           }()

}
