package eureka

import (
        "os"
        "io/ioutil"
        "strings"
        "strconv"
        "fmt"
        "time"
        "goeureka/util"
)

var instanceId string

func Register() {
        instanceId = util.GetUUID();


        dir, _ := os.Getwd()
        data, _ := ioutil.ReadFile(dir + "/templates/regtpl.json")

        tpl := string(data)
        tpl = strings.Replace(tpl, "${ipAddress}", util.GetLocalIP(), -1)
        tpl = strings.Replace(tpl, "${port}", "1080", -1)
        tpl = strings.Replace(tpl, "${instanceId}", instanceId, -1)

        fmt.Println("going to send " + tpl)

        // Register.
        registerAction := HttpAction {
                Url : "http://" + os.Getenv("EURKEAIP") + ":8761/eureka/apps/fruits",
                Method: "POST",
                ContentType: "application/json",
                Body: tpl,
        }
        var result bool
        for {
                result = DoHttpRequest(registerAction)
                if result {
                        break
                } else {
                        time.Sleep(time.Second * 5)
                }
        }
        fmt.Println("Result from Eureka was " + strconv.FormatBool(result))
}

func StartHeartbeat() {
        for {
                fmt.Println("Sending heartbeat to Eureka")
                time.Sleep(time.Second * 30)
                heartbeat()
        }
}

func heartbeat() {
        heartbeatAction := HttpAction{
                Url : "http://" + os.Getenv("EURKEAIP") + ":8761/eureka/apps/fruits/fruits:" + instanceId,
                Method: "PUT",
        }
        DoHttpRequest(heartbeatAction)
}

func Deregister() {
        fmt.Println("Trying to deregister application...")
        // Deregister
        deregisterAction := HttpAction {
                Url : "http://" + os.Getenv("EURKEAIP") + ":8761/eureka/apps/fruits/fruits:" + instanceId,
                Method: "DELETE",
        }
        DoHttpRequest(deregisterAction)
        fmt.Println("Deregistered application, exiting. Check Eureka...")
}
