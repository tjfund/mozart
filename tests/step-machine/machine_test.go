package step_machine

import (
	"encoding/json"
	"github.com/coinbase/step/machine"
	"io/ioutil"
	"testing"

	"github.com/coinbase/step/utils/to"
	"github.com/stretchr/testify/assert"
)

func loadFixture(file string, t *testing.T) *machine.StateMachine {
	exampleMachine, err := machine.ParseFile(file)
	assert.NoError(t, err)
	return exampleMachine
}

func execute(json []byte, input interface{}, t *testing.T) (map[string]interface{}, error) {
	exampleMachine, err := machine.FromJSON(json)
	assert.NoError(t, err)
	exampleMachine.SetDefaultHandler()

	exec, err := exampleMachine.Execute(input)

	return exec.Output, err
}

func executeFixture(file string, input map[string]interface{}, t *testing.T) map[string]interface{} {
	exampleMachine := loadFixture(file, t)

	exec, err := exampleMachine.Execute(input)

	assert.NoError(t, err)

	return exec.Output
}

//////
// TESTS
//////

func Test_Machine_EmptyStateMachinePassExample(t *testing.T) {
	_, err := execute([]byte(machine.EmptyStateMachine), make(map[string]interface{}), t)
	assert.NoError(t, err)
}

func Test_Machine_SimplePassExample_With_Execute(t *testing.T) {
	json := []byte(`
  {
      "StartAt": "start",
      "States": {
        "start": {
          "Type": "Pass",
          "Result": "b",
          "ResultPath": "$.a",
          "End": true
        }
    }
  }`)

	output, err := execute(json, make(map[string]interface{}), t)
	assert.NoError(t, err)
	assert.Equal(t, output["a"], "b")

	output, err = execute(json, "{}", t)
	assert.NoError(t, err)
	assert.Equal(t, output["a"], "b")

	output, err = execute(json, to.Strp("{}"), t)
	assert.NoError(t, err)
	assert.Equal(t, output["a"], "b")
}

func Test_Machine_ErrorUnknownState(t *testing.T) {
	exampleMachine := loadFixture("../asl/bad_unknown_state.json", t)
	_, err := exampleMachine.Execute(make(map[string]interface{}))

	assert.Error(t, err)
	assert.Regexp(t, "Unknown State", err.Error())
}

func Test_Machine_MarshallAllTypes(t *testing.T) {
	file := "../asl/all_types.json"
	sm, err := machine.ParseFile(file)
	assert.NoError(t, err)

	sm.SetDefaultHandler()
	assert.NoError(t, sm.Validate())

	marshalledJson, err := json.Marshal(sm)
	assert.NoError(t, err)

	rawJson, err := ioutil.ReadFile(file)
	assert.NoError(t, err)

	assert.JSONEq(t, string(rawJson), string(marshalledJson))
}
