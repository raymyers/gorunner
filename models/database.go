package models

import (
	"io/ioutil"
	"os"
)

const (
	jobsFile     = "jobs.json"
	runsFile     = "runs.json"
	tasksFile    = "tasks.json"
	triggersFile = "triggers.json"
)

var (
	jobList     JobList
	taskList    TaskList
	runList     RunList
	triggerList TriggerList
)

type ListWriter func([]byte, string)
type ListReader func(string) []byte

func InitDatabase() {
	jobList = JobList{list{elements: make([]elementer, 10), fileName: jobsFile}}
	taskList = TaskList{list{elements: make([]elementer, 10), fileName: tasksFile}}
	triggerList = TriggerList{list{elements: make([]elementer, 10), fileName: triggersFile}}
	runList = RunList{list{elements: make([]elementer, 10), fileName: runsFile}}

	jobList.Load()
	taskList.Load(readFile)
	triggerList.Load(readFile)
	runList.Load()
}

func GetJobList() *JobList {
	return &jobList
}

func GetRunList() *RunList {
	return &runList
}

func GetTaskList() *TaskList {
	return &taskList
}

func GetTriggerList() *TriggerList {
	return &triggerList
}

func writeFile(bytes []byte, filePath string) {
	err := ioutil.WriteFile(filePath, bytes, 0644)
	if err != nil {
		panic(err)
	}
}

func readFile(filePath string) []byte {
	_, err := os.Stat(filePath)
	if err != nil {
		println("Couldn't read file, creating fresh:", filePath)
		err = ioutil.WriteFile(filePath, []byte("[]"), 0644)
		if err != nil {
			panic(err)
		}
	}

	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return bytes
}
