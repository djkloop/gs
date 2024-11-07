package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"

	flow "github.com/s8sg/goflow/flow/v1"
	goflow "github.com/s8sg/goflow/v1"
)

func Input(data []byte, option map[string][]string) ([]byte, error) {
	var input map[string]int
	if err := json.Unmarshal(data, &input); err != nil {
		return nil, err
	}

	outputInt := input["input"]

	return []byte(strconv.Itoa(outputInt)), nil
}

func AddOne(data []byte, option map[string][]string) ([]byte, error) {
	num, _ := strconv.Atoi(string(data))
	outputInt := num + rand.Intn(10) + 1
	fmt.Println("AddOne = ", outputInt)
	return []byte(strconv.Itoa(outputInt)), nil
}

func AddTwo(data []byte, option map[string][]string) ([]byte, error) {
	num, _ := strconv.Atoi(string(data))
	outputInt := num + rand.Intn(101) + 100
	fmt.Println("AddTwo = ", outputInt)
	return []byte(strconv.Itoa(outputInt)), nil
}

func Output(data []byte, option map[string][]string) ([]byte, error) {
	fmt.Println("data = ", string(data))
	return []byte("ok"), nil
}

func MyFlow(flow *flow.Workflow, context *flow.Context) error {
	dag := flow.Dag()

	//
	dag.Node("input", Input)
	dag.Node("add-one", AddOne)
	dag.Node("add-two", AddTwo)
	dag.Node("output", Output)

	//
	dag.Edge("input", "add-one")
	dag.Edge("add-one", "add-two")
	dag.Edge("add-two", "output")

	return nil
}

func main() {
	gfs := goflow.FlowService{
		Port:              8999,
		RedisURL:          "43.143.243.166:6379",
		RedisPassword:     "bS9@xG2?",
		WorkerConcurrency: 5,
	}

	err := gfs.Register("add-flow", MyFlow)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	err = gfs.Start()
	if err != nil {
		fmt.Println("err!2")
		panic(err)
	}
}
