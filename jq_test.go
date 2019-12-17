package jq

import (
	"encoding/json"
	"testing"
)

const testdata = `
{
	"cats": 2,
	"data": [
		{
			"type": "cat",
			"name": "Willow",
			"attributes": {
				"whiskers": 12
			},
			"friends": [
				"Hitch",
				"Olive"
			],
			"meals": [
				{
					"day": "monday",
					"time": "8:30",
					"type": "dry food"
				},
				{
					"day": "monday",
					"time": "17:30",
					"type": "dry food"
				}
			]
		},
		{
			"type": "cat",
			"name": "Hitch",
			"attributes": {
				"whiskers": 15
			},
			"friends": [
				"Willow"
			],
			"meals": [
				{
					"day": "monday",
					"time": "8:00",
					"type": "dry food"
				},
				{
					"day": "monday",
					"time": "18:15",
					"type": "dry food"
				}
			]
		}
	],
	"meta": {
		"errors": [],
		"response": 200
	}
}
`

func TestBasic(t *testing.T) {
	unsafeData := make(map[string]interface{})
	err := json.Unmarshal([]byte(testdata), &unsafeData)
	if err != nil {
		t.Error(err)
		return
	}

	results, err := Query("cats", unsafeData, nil)
	if err != nil {
		t.Error(err)
		return
	}

	if len(results) != 1 {
		t.Error("Wrong number of results returned")
		return
	}

	if results[0] != float64(2) {
		t.Error("Wrong result value returned")
		return
	}
}

func TestBasicNest(t *testing.T) {
	unsafeData := make(map[string]interface{})
	err := json.Unmarshal([]byte(testdata), &unsafeData)
	if err != nil {
		t.Error(err)
		return
	}

	results, err := Query("meta.response", unsafeData, nil)
	if err != nil {
		t.Error(err)
		return
	}

	if len(results) != 1 {
		t.Error("Wrong number of results returned")
		return
	}

	if results[0] != float64(200) {
		t.Error("Wrong result value returned")
		return
	}
}

func TestLoopBasicNest(t *testing.T) {
	unsafeData := make(map[string]interface{})
	err := json.Unmarshal([]byte(testdata), &unsafeData)
	if err != nil {
		t.Error(err)
		return
	}

	results, err := Query("data.[].attributes.whiskers", unsafeData, nil)
	if err != nil {
		t.Error(err)
		return
	}

	if len(results) != 2 {
		t.Error("Wrong number of results returned")
		return
	}

	if results[0] != float64(12) || results[1] != float64(15) {
		t.Error("Wrong result value returned")
		return
	}
}

func TestLoopBasicNestOne(t *testing.T) {
	unsafeData := make(map[string]interface{})
	err := json.Unmarshal([]byte(testdata), &unsafeData)
	if err != nil {
		t.Error(err)
		return
	}

	results, err := Query("data.[i].attributes.whiskers", unsafeData, Options{OptionVarIndexAt: 1})
	if err != nil {
		t.Error(err)
		return
	}

	if len(results) != 1 {
		t.Error("Wrong number of results returned")
		return
	}

	if results[0] != float64(15) {
		t.Error("Wrong result value returned")
		return
	}
}

func TestLoopSlice(t *testing.T) {
	unsafeData := make(map[string]interface{})
	err := json.Unmarshal([]byte(testdata), &unsafeData)
	if err != nil {
		t.Error(err)
		return
	}

	results, err := Query("data.[].friends", unsafeData, nil)
	if err != nil {
		t.Error(err)
		return
	}

	if len(results) != 2 {
		t.Error("Wrong number of results returned")
		return
	}

	if results[0].([]interface{})[0] != "Hitch" || results[0].([]interface{})[1] != "Olive" {
		t.Error("Wrong result value returned")
		return
	}

	if results[1].([]interface{})[0] != "Willow" {
		t.Error("Wrong result value returned")
		return
	}
}

func TestLoopLoop(t *testing.T) {
	unsafeData := make(map[string]interface{})
	err := json.Unmarshal([]byte(testdata), &unsafeData)
	if err != nil {
		t.Error(err)
		return
	}

	results, err := Query("data.[].meals.[].time", unsafeData, nil)
	if err != nil {
		t.Error(err)
		return
	}

	if len(results) != 4 {
		t.Error("Wrong number of results returned")
		return
	}

	if results[0] != "8:30" || results[1] != "17:30" || results[2] != "8:00" || results[3] != "18:15" {
		t.Error("Wrong result value returned")
		return
	}

}
