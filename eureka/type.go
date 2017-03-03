package eureka

import "github.com/ibrokethecloud/go-utils"
import "os"
import "strings"

var instanceId = utils.GetHostName()+":"+strings.ToLower(os.Getenv("APP_NAME"))+":"+os.Getenv("PORT")

type RegisterationInfo struct {

  Instance InstanceInfo `json:"instance"`
}

type InstanceInfo struct {
  Hostname string `json:"hostName"`
  App string `json:"app"`
  VipAddress  string `json:"vipAddress"`
  SecureVipAddress  string  `json:"secureVipAddress"`
  IpAddress string `json:"ipAddr"`
  Status string `json:"status"`
  Port  PortInfo  `json:"port"`
  SecurePort SecurePortInfo `json:"securePort"`
  HealthCheckUrl string `json:"healthCheckUrl"`
  StatusPageUrl string  `json:"statusPageUrl"`
  HomePageUrl  string `json:"homePageUrl"`
  DataCenter DataCenterInfo `json:"dataCenterInfo"`
  InstanceId string `json:"instanceId"`
}

type DataCenterInfo struct{
  Class string  `json:"@class"`
  Name string `json:"name"`
}

type PortInfo struct{
  PortNumber string `json:"$"`
  Status string `json:"@enabled"`
}

type SecurePortInfo struct{
  PortNumber string `json:"$"`
  Status string `json:"@enabled"`
}

var ContainerIp = utils.GetIP()

var DefaultDataCenterInfo = DataCenterInfo {
  Class: "com.netflix.appinfo.InstanceInfo$DefaultDataCenterInfo",
  Name: "MyOwn",
}

var DefaultPortInfo = PortInfo {
  PortNumber: PORT,
  Status: "true",
}

var DefaultSecurePortInfo = SecurePortInfo {
  PortNumber: "8443",
  Status: "false",
}

var DefaultInstanceInfo = InstanceInfo {
  Hostname: utils.GetHostName(),
  App:  os.Getenv("APP_NAME"),
  VipAddress: ContainerIp,
  SecureVipAddress: ContainerIp,
  IpAddress:  ContainerIp,
  Status: "UP",
  Port: DefaultPortInfo,
  SecurePort: DefaultSecurePortInfo,
  HealthCheckUrl: "http://"+ContainerIp+":"+os.Getenv("PORT")+"/health",
  StatusPageUrl: "http://"+ContainerIp+":"+os.Getenv("PORT")+"/status",
  HomePageUrl: "http://"+ContainerIp+":"+os.Getenv("PORT")+"/",
  DataCenter: DefaultDataCenterInfo,
  InstanceId: instanceId,
}

var DefaultRegisterationInfo = RegisterationInfo {
  Instance: DefaultInstanceInfo,
}


/*  {
      "instance": {
          "hostName": "WKS-SOF-L011",
          "app": "com.automationrhapsody.eureka.app",
          "vipAddress": "com.automationrhapsody.eureka.app",
          "secureVipAddress": "com.automationrhapsody.eureka.app"
          "ipAddr": "10.0.0.10",
          "status": "STARTING",
          "port": {"$": "8080", "@enabled": "true"},
          "securePort": {"$": "8443", "@enabled": "true"},
          "healthCheckUrl": "http://WKS-SOF-L011:8080/healthcheck",
          "statusPageUrl": "http://WKS-SOF-L011:8080/status",
          "homePageUrl": "http://WKS-SOF-L011:8080",
          "dataCenterInfo": {
              "@class": "com.netflix.appinfo.InstanceInfo$DefaultDataCenterInfo",
              "name": "MyOwn"
          },
      }
  } */
