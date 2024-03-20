package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-co-op/gocron"
	"gopkg.in/yaml.v3"
)

type Monitor struct {
	Target string `yaml:"target"`
	Ping   string `yaml:"ping"`
	Period string `yaml:"period"`
	Status int    `yaml:"status"`
}

func main() {
	scheduler := gocron.NewScheduler(time.Local)

	monitors := readConfig()
	for name, monitor := range monitors {
		job, err := createJob(name, &monitor)
		if err == nil {
			fmt.Println("Schedule created for ", name, " every ", monitor.Period)
			scheduler.Every(monitor.Period).Do(job)
		} else {
			fmt.Println(err.Error())
		}
	}
}

func readConfig() map[string]Monitor {
	file, err := os.ReadFile(os.Getenv("configFile"))
	if err != nil {
		fmt.Printf("yamlFile.Get err #%v ", err)
	}

	obj := make(map[string]Monitor)
	err = yaml.Unmarshal(file, obj)
	if err != nil {
		fmt.Printf("Unmarshal: %v", err)
	}

	return obj
}

func checkStatusCode(target string, expected int) (bool, int) {
	resp, err := http.Get(target)
	if err != nil {
		return false, 0
	}

	return resp.StatusCode == expected, resp.StatusCode
}

func createJob(name string, monitor *Monitor) (func(), error) {
	if monitor.Target == "" {
		return nil, errors.New("Invalid config: No 'target' defined for monitor " + name)
	}

	if monitor.Ping == "" {
		return nil, errors.New("Invalid config: No 'ping' defined for monitor " + name)
	}

	if monitor.Status == 0 {
		monitor.Status = 200
	}

	if monitor.Period == "" {
		monitor.Period = "10m"
	}

	return func() {
		success, code := checkStatusCode(monitor.Target, monitor.Status)
		if success {
			http.Get(monitor.Ping)
		} else {
			dt := time.Now()
			timestamp := dt.Format("01-02-2006 15:04:05")
			fmt.Printf("%s: %s (%s) received http status code %d. Expected %d\n", timestamp, name, monitor.Target, code, monitor.Status)
		}
	}, nil
}
