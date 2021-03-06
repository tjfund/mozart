// Copyright 2017 Netflix, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package task

import (
	"encoding/json"
)

type ResultStatus string

type Result struct {
	Status                ResultStatus           `json:"status"`
	WorkflowInstanceId    string                 `json:"workflowInstanceId"`
	TaskId                string                 `json:"taskId"`
	ReasonForIncompletion string                 `json:"reasonForIncompletion"`
	CallbackAfterSeconds  int64                  `json:"callbackAfterSeconds"`
	WorkerId              string                 `json:"workerId"`
	OutputData            map[string]interface{} `json:"outputData"`
	Logs                  []LogMessage           `json:"logs"`
}

// LogMessage used to sent logs to conductor server
type LogMessage struct {
	Log    	       string    `json:"log"`
	TaskID         string    `json:"taskId"`
	CreatedTime    int       `json:"createdTime"`
}

func NewEmptyTaskResult() *Result {
	taskResult := new(Result)
	taskResult.OutputData = make(map[string]interface{})
	return taskResult
}

func NewTaskResult(t *Task) *Result {
	taskResult := new(Result)
	taskResult.CallbackAfterSeconds = t.CallbackAfterSeconds
	taskResult.WorkflowInstanceId = t.WorkflowInstanceId
	taskResult.TaskId = t.TaskId
	taskResult.ReasonForIncompletion = t.ReasonForIncompletion
	taskResult.Status = ResultStatus(t.Status)
	taskResult.WorkerId = t.WorkerId
	taskResult.OutputData = t.OutputData
	return taskResult
}

func (t *Result) ToJSONString() (string, error) {
	var jsonString string
	b, err := json.Marshal(t)
	if err == nil {
		jsonString = string(b)
	}
	return jsonString, err
}

func ParseTaskResult(inputJSON string) (*Result, error) {
	t := NewEmptyTaskResult()
	err := json.Unmarshal([]byte(inputJSON), t)
	return t, err
}
