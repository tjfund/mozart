package step_machine

import (
	"encoding/json"
	"github.com/coinbase/step/machine"
	"testing"

	"github.com/coinbase/step/utils/to"
	"github.com/stretchr/testify/assert"
)

type stateTestData struct {
	Input  map[string]interface{}
	Output map[string]interface{}
	Error  *string
	Next   *string
}

func testState(state machine.State, std stateTestData, t *testing.T) {
	// Make sure the execution is on Valid State
	err := state.Validate()
	assert.NoError(t, err)

	// default empty input
	if std.Input == nil {
		std.Input = map[string]interface{}{}
	}

	output, next, err := state.Execute(nil, std.Input)

	// expecting error?
	if std.Error != nil {
		assert.Error(t, err)
		assert.Regexp(t, *std.Error, err.Error())
	} else if err != nil {
		assert.NoError(t, err)
	}

	if std.Output != nil {
		assert.Equal(t, std.Output, output)
	}

	if std.Next != nil {
		assert.Equal(t, *std.Next, *next)
	}
}

func parseChoiceState(b []byte, t *testing.T) *machine.ChoiceState {
	var p machine.ChoiceState
	err := json.Unmarshal(b, &p)
	assert.NoError(t, err)
	p.SetName(to.Strp("TestState"))
	p.SetType(to.Strp("Choice"))
	return &p
}

func parsePassState(b []byte, t *testing.T) *machine.PassState {
	var p machine.PassState
	err := json.Unmarshal(b, &p)
	assert.NoError(t, err)
	p.SetName(to.Strp("TestState"))
	p.SetType(to.Strp("Pass"))
	return &p
}

func parseWaitState(b []byte, t *testing.T) *machine.WaitState {
	var p machine.WaitState
	err := json.Unmarshal(b, &p)
	assert.NoError(t, err)
	p.SetName(to.Strp("TestState"))
	p.SetType(to.Strp("Wait"))
	return &p
}

func parseTaskState(b []byte, t *testing.T) *machine.TaskState {
	var p machine.TaskState
	err := json.Unmarshal(b, &p)
	assert.NoError(t, err)
	p.SetName(to.Strp("TestState"))
	p.SetType(to.Strp("Task"))
	return &p
}

func parseValidTaskState(b []byte, handler interface{}, t *testing.T) *machine.TaskState {
	state := parseTaskState(b, t)
	state.SetTaskHandler(handler)
	assert.NoError(t, state.Validate())
	return state
}

func parseMapState(b []byte, t *testing.T) *machine.MapState {
	var p machine.MapState
	err := json.Unmarshal(b, &p)
	assert.NoError(t, err)
	p.SetName(to.Strp("TestState"))
	p.SetType(to.Strp("Map"))
	return &p
}
