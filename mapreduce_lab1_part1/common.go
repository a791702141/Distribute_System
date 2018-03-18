package mapreduce

import (
	"fmt" // fmt is a package used for I/O
	"strconv" // used for string operation 
)

// Debugging enabled?
const debugEnabled = false // constant variable

// debug() will only print if debugEnabled is true
func debug(format string, a ...interface{}) (n int, err error) { // first is parameter.second is return type
	if debugEnabled {
		n, err = fmt.Printf(format, a...)
	}
	return
}

// jobPhase indicates whether a task is scheduled as a map or reduce task.
type jobPhase string // use jobPhase as string

const (
	mapPhase    jobPhase = "mapPhase"
	reducePhase          = "reducePhase"
)

// KeyValue is a type used to hold the key/value pairs passed to the map and
// reduce functions.
type KeyValue struct {
	Key   string
	Value string
}

// reduceName constructs the name of the intermediate file which map task
// <mapTask> produces for reduce task <reduceTask>.
func reduceName(jobName string, mapTask int, reduceTask int) string {
	return "mrtmp." + jobName + "-" + strconv.Itoa(mapTask) + "-" + strconv.Itoa(reduceTask)
}

// mergeName constructs the name of the output file of reduce task <reduceTask>
func mergeName(jobName string, reduceTask int) string {
	return "mrtmp." + jobName + "-res-" + strconv.Itoa(reduceTask)
}
