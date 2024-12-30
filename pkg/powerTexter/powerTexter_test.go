package powertexter

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"
)

func TestResponseUnmarshall(t *testing.T) {
	input, err := os.ReadFile("response.json")
	if err != nil {
		t.Errorf("Could not reaf response.json: %v", err)
	}
	response := &ResponseSchema{}
	json.Unmarshal(input, response)
	expected := ResponseSchema{
		Results: map[string]ResultSchema{
			"A": {Status: 200, Frames: []FrameSchema{
				{Schema: SchemaSchema{Name: "measurement", Fields: []FieldSchema{
					{Name: "Time", Type: "time", Labels: map[string]string(nil)},
					{Name: "Value", Type: "number", Labels: map[string]string{"location": "PDU 28 (Chaos West)"}},
				}}, Data: DataSchema{Values: [][]float64{{1.7355216e+12, 1.7355219e+12}, {4216.593627255858, 4315.89920435353}}}},
			}},
		},
	}
	if reflect.DeepEqual(response, expected) {
		t.Error("Result does not match expected")
	}
	return
}
