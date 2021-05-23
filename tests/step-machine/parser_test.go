package step_machine

import (
	"github.com/coinbase/step/machine"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Machine_Parser_FromJSON(t *testing.T) {
	json := []byte(`
  {
      "Comment": "Adds some coordinates to the input",
      "StartAt": "Coords",
      "States": {
        "Coords": {
          "Type": "Pass",
          "Result": {
            "x": 3.14,
            "y": 103.14159
          },
          "ResultPath": "$.coords",
          "End": true
        }
    }
  }`)

	_, err := machine.FromJSON(json)

	assert.Equal(t, err, nil)
}

func Test_Parser_Expands_TaskFn(t *testing.T) {
	json := []byte(`
  {
      "StartAt": "A",
      "States": {
        "A": {
          "Type": "TaskFn",
          "Next": "B"
        },
        "B": {
          "Type": "TaskFn",
          "End": true
        }
    }
  }`)

	sm, err := machine.FromJSON(json)
	assert.NoError(t, err)

	// Names and Types
	assert.Equal(t, len(sm.States), 2)
	assert.Equal(t, *sm.States["A"].GetType(), "Task")
	assert.Equal(t, *sm.States["B"].GetType(), "Task")

	ataskState := sm.States["A"].(*machine.TaskState)
	btaskState := sm.States["B"].(*machine.TaskState)

	// ORDER
	assert.Equal(t, ataskState.Parameters, map[string]interface{}{"Task": "A", "Input.$": "$"})
	assert.Equal(t, btaskState.Parameters, map[string]interface{}{"Task": "B", "Input.$": "$"})
}

func Test_Machine_Parser_FileNonexistantFile(t *testing.T) {
	_, err := machine.ParseFile("../asl/non_existent_file.json")
	assert.Error(t, err)
}

func Test_Machine_Parser_OfBadStateType(t *testing.T) {
	_, err := machine.ParseFile("../asl/bad_type.json")

	assert.Error(t, err)
	assert.Regexp(t, "Unknown State", err.Error())
}

func Test_Machine_Parser_OfBadPath(t *testing.T) {
	_, err := machine.ParseFile("../asl/bad_path.json")

	assert.Error(t, err)
	assert.Regexp(t, "Bad JSON path", err.Error())
}

// BASIC TYPE TESTS

func Test_Machine_Parser_AllTypes(t *testing.T) {
	sm, err := machine.ParseFile("../asl/all_types.json")
	assert.NoError(t, err)

	assert.NoError(t, sm.Validate())
}

func Test_Machine_Parser_BasicPass(t *testing.T) {
	sm, err := machine.ParseFile("../asl/basic_pass.json")
	assert.NoError(t, err)
	assert.NoError(t, sm.Validate())
}

func Test_Machine_Parser_BasicChoice(t *testing.T) {
	sm, err := machine.ParseFile("../asl/basic_choice.json")

	assert.Equal(t, err, nil)
	assert.NoError(t, sm.Validate())
}

func Test_Machine_Parser_TaskFn(t *testing.T) {
	sm, err := machine.ParseFile("../asl/taskfn.json")

	assert.Equal(t, err, nil)
	assert.NoError(t, sm.Validate())
}

func Test_Machine_Parser_Map(t *testing.T) {
	sm, err := machine.ParseFile("../asl/map.json")
	var mapState *machine.MapState
	mapState = sm.States["Start"].(*machine.MapState)
	assert.Equal(t, err, nil)
	assert.NoError(t, sm.Validate())
	assert.Equal(t, "$.detail", mapState.InputPath.String(), )
	assert.Equal(t, "$.shipped", mapState.ItemsPath.String(), )
	assert.Equal(t, "$.detail.shipped", mapState.ResultPath.String(), )
	assert.Equal(t, 1, len(mapState.Iterator.States))
	assert.Equal(t, "Task", *mapState.Iterator.States["Validate"].GetType(), )
}
