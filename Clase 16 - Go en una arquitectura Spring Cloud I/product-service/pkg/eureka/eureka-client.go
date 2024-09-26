package eureka

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/procyon-projects/chrono"
)

type AppRegistrationBody struct {
	Instance InstanceDetails `json:"instance"`
}

type InstanceDetails struct {
	InstanceId       string         `json:"instanceId"`
	HostName         string         `json:"hostName"`
	App              string         `json:"app"`
	VipAddress       string         `json:"vipAddress"`
	SecureVipAddress string         `json:"secureVipAddress"`
	IpAddr           string         `json:"ipAddr"`
	Status           string         `json:"status"`
	Port             Port           `json:"port"`
	SecurePort       Port           `json:"securePort"`
	HealthCheckUrl   string         `json:"healthCheckUrl"`
	StatusPageUrl    string         `json:"statusPageUrl"`
	HomePageUrl      string         `json:"homePageUrl"`
	DataCenterInfo   DataCenterInfo `json:"dataCenterInfo"`
}

type DataCenterInfo struct {
	Class string `json:"@class"`
	Name  string `json:"name"`
}

type Port struct {
	Number  int    `json:"$"`
	Enabled string `json:"@enabled"`
}

func ScheduleHeartbeat(appName string, appId string) chrono.ScheduledTask {

	taskScheduler := chrono.NewDefaultTaskScheduler()
	task, err := taskScheduler.ScheduleWithFixedDelay(func(ctx context.Context) {
		fmt.Println("Fixed Delay Task")
		sendHeartbeat(appName, appId)

	}, 25*time.Second)

	if err != nil {
		log.Fatal(err)
	}
	return task
}
func RegisterApp(appId string, appName string) {
	fmt.Println("register app")
	body := buildBody(appId, appName, "STARTING")

	var buf bytes.Buffer
	error := json.NewEncoder(&buf).Encode(body)
	if error != nil {
		log.Fatal(error)
	}
	resp, err := http.Post("http://localhost:8761/eureka/apps/"+appName, "application/json", &buf)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	responseBody, parseErr := io.ReadAll(resp.Body)

	if parseErr != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseBody))

}

func UpdateAppStatus(appId string, appName string, status string) {

	fmt.Println("update app")
	body := buildBody(appId, appName, status)

	var buf bytes.Buffer
	error := json.NewEncoder(&buf).Encode(body)
	if error != nil {
		log.Fatal(error)
	}
	resp, err := http.Post("http://localhost:8761/eureka/apps/"+appName, "application/json", &buf)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	responseBody, parseErr := io.ReadAll(resp.Body)

	if parseErr != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseBody))

}

func DeleteApp(appName string, appId string) {
	fmt.Println("sending hearbeat")
	req, err := http.NewRequest("DELETE", "http://localhost:8761/eureka/apps/"+appName+"/"+appId+"", nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.Body)

}
func sendHeartbeat(appName string, appId string) {
	fmt.Println("sending hearbeat")
	req, err := http.NewRequest("PUT", "http://localhost:8761/eureka/apps/"+appName+"/"+appId+"", nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.Body)
}

func buildBody(appId string, appName string, status string) *AppRegistrationBody {

	hostname := "192.168.1.37"
	httpport := "8083"

	homePageUrl := "http://" + hostname + ":" + httpport
	statusPageUrl := "http://" + hostname + ":" + httpport + "/status"
	healthCheckUrl := "http://" + hostname + ":" + httpport + "/healthcheck"
	dataCenterInfo := DataCenterInfo{Class: "com.netflix.appinfo.InstanceInfo$DefaultDataCenterInfo", Name: "MyOwn"}
	port := Port{Number: 8083, Enabled: "false"}
	securePort := Port{Number: 8083, Enabled: "false"}
	instance := InstanceDetails{InstanceId: appId, HostName: hostname, App: appName, VipAddress: appName, SecureVipAddress: appName, IpAddr: "192.168.1.37", Status: status, Port: port, SecurePort: securePort, HealthCheckUrl: healthCheckUrl, StatusPageUrl: statusPageUrl, HomePageUrl: homePageUrl, DataCenterInfo: dataCenterInfo}

	requestBody := &AppRegistrationBody{Instance: instance}

	return requestBody
}
