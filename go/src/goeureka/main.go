package main

import (
        "net/http"
        "sync"
        "log"
        "goeureka/service"
        "goeureka/eureka"
        "os"
        "os/signal"
        "syscall"
)

func main() {
        handleSigterm()                              // Graceful shutdown on Ctrl+C or kill

        go startWebServer()                          // Starts HTTP service  (async)

        eureka.Register()                            // Performs Eureka registration

        go eureka.StartHeartbeat()                   // Performs Eureka heartbeating (async)

        // Block...
        wg := sync.WaitGroup{}                       // Use a WaitGroup to block main() exit
        wg.Add(1)
        wg.Wait()
}

func handleSigterm() {
        c := make(chan os.Signal, 1)
        signal.Notify(c, os.Interrupt)
        signal.Notify(c, syscall.SIGTERM)
        go func() {
                <-c
                eureka.Deregister()
                os.Exit(1)
        }()
}

func startWebServer() {
        router := service.NewRouter()
        log.Println("Starting HTTP service at 1080")
        err := http.ListenAndServe(":1080", router)
        if err != nil {
                log.Println("An error occured starting HTTP listener at port 1080")
                log.Println("Error: " + err.Error())
        }
}

