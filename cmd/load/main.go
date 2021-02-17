package main

import (
	"encoding/json"
	"fmt"
	vegeta "github.com/tsenart/vegeta/v12/lib"
	"github.com/tsenart/vegeta/v12/lib/plot"
	"log"
	"time"
)

var (
	endpoint = "http://127.0.0.1:8081"
	metrics  vegeta.Metrics
	p        = plot.New()
)

func main() {
	rate := vegeta.Rate{Freq: 5, Per: time.Second}
	targeter := NewTargeter(endpoint)
	attacker := vegeta.NewAttacker()
	duration := 60 * time.Second
	for res := range attacker.Attack(targeter, rate, duration, "Load Test") {
		if err := p.Add(res); err != nil {
			log.Fatalln(err)
		}
		fmt.Println(time.Now())
		metrics.Add(res)
	}
	metrics.Close()
	pretty, _ := json.MarshalIndent(metrics, "", "    ")
	fmt.Printf("%+v\n", string(pretty))
}

func NewTargeter(url string) vegeta.Targeter {
	return func(tgt *vegeta.Target) error {
		if tgt == nil {
			return vegeta.ErrNilTarget
		}
		tgt.Method = "GET"
		tgt.URL = url
		return nil
	}
}
