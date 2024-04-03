package openai

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"
)

func TestStop(t *testing.T) {
	var (
		stopSequence  Stop = StopSequence("test")
		stopSequences Stop = StopSequences{"test", "0401"}
	)
	if !reflect.DeepEqual(stopSequence.sequences(), []string{"test"}) {
		t.Errorf("stop: %v != %v", stopSequence.sequences(), []string{"test"})
		return
	}
	if !reflect.DeepEqual(stopSequences.sequences(), []string{"test", "0401"}) {
		t.Errorf("stop: %v != %v", stopSequences.sequences(), []string{"test", "0401"})
	}
}

func TestNullableType(t *testing.T) {
	t.Run("string", testNullableType[string]("test", false))
	t.Run("int", testNullableType[int](312, true))
	t.Run("float64", testNullableType[float64](3.12, true))
	t.Run("bool", testNullableType[bool](true, true))
}

func testNullableType[T interface {
	string | int | float64 | bool
}](instance T, isNull bool) func(*testing.T) {
	return func(t *testing.T) {
		serializedVal, err := json.Marshal(instance)
		if err != nil {
			t.Errorf("json.Marshal(%T): %s", instance, err)
			return
		}
		var val NullableType[T]
		if val.IsNull() != isNull {
			t.Errorf("NullableType[%T]: empty value must be null=%v, got null=%v", instance, isNull, !isNull)
			return
		}
		if err = json.Unmarshal([]byte("null"), &val); err != nil {
			t.Errorf("json.Unmarshal(nil): %s", err)
			return
		}
		if val != "" {
			t.Errorf("NullableType[%T]: unmarshal from nil should keep value empty, got %s", instance, val)
			return
		}
		val = "test"
		if err = json.Unmarshal([]byte("null"), &val); err != nil {
			t.Errorf("json.Unmarshal(nil): %s", err)
			return
		}
		if val != "test" {
			t.Errorf("NullableType[%T]: unmarshal from nil must be a noop, got %s", instance, val)
			return
		}
		if err = json.Unmarshal(serializedVal, &val); err != nil {
			t.Errorf("json.Unmarshal(%T): %s", instance, err)
			return
		}
		if val.IsNull() || val == "" {
			t.Errorf("NullableType[%[1]T]: expectes %[1]v, got null/empty value", instance)
			return
		}
		if val.Value() != instance {
			t.Errorf("NullableType[%[2]T]: %[1]s != %[2]v", val, instance)
			return
		}
		var reSerializedVal []byte
		reSerializedVal, err = json.Marshal(val)
		if err != nil {
			t.Errorf("json.Marshal(NullableType[%T]): %s", val, err)
			return
		}
		if !bytes.Equal(reSerializedVal, serializedVal) {
			t.Errorf("NullableType[%T]: serialized %s != %s", instance, string(reSerializedVal), string(serializedVal))
			return
		}
		if isNull {
			val = ""
			reSerializedVal, err = json.Marshal(val)
			if err != nil {
				t.Errorf("json.Marshal(NullableType[%T]): %s", val, err)
				return
			}
			if !bytes.Equal(reSerializedVal, []byte("null")) {
				t.Errorf("NullableType[%T]: serialized %s != null", instance, string(reSerializedVal))
				return
			}
		}
		if err = json.Unmarshal([]byte("[]"), &val); err == nil {
			t.Errorf("json.Unmarshal(\"[]\") expects errors, got nothing")
			return
		}
		if err = json.Unmarshal([]byte("{}"), &val); err == nil {
			t.Errorf("json.Unmarshal(\"{}\") expects errors, got nothing")
			return
		}
	}
}

func TestTokenType(t *testing.T) {
	t.Run("null", testTokenType(nil, "null"))
	t.Run("array", testTokenType(json.Delim('['), "array"))
	t.Run("object", testTokenType(json.Delim('{'), "object"))
	t.Run("bool", testTokenType(true, "bool"))
	t.Run("number", testTokenType(json.Number("0313"), "number"))
	t.Run("string", testTokenType("test", "string"))
	t.Run("unknown", testTokenType(struct{}{}, "unknown"))
}

func testTokenType(token json.Token, expect string) func(*testing.T) {
	return func(t *testing.T) {
		if tokType := tokenType(token); tokType != expect {
			t.Errorf("tokenType: %s != %s", tokType, expect)
			return
		}
	}
}

func TestDefProperty(t *testing.T) {
	const jsonProperty = `
	{
		"type": "array",
		"description": "test_array",
		"items": [
			{
				"type": "object",
				"description": "test_object",
				"properties": {
					"test_int": {
						"type": "integer",
						"description": "test_int",
						"maximum": 10
					}
				},
				"required": ["test_int"]
			}
		]
	}
	`
	defProperty := &DefProperty{
		Type:        DefPropertyTypeArray,
		Description: "test_array",
		Items: DefItems{
			&DefProperty{
				Type:        DefPropertyTypeObject,
				Description: "test_object",
				Properties: DefProperties{
					"test_int": &DefProperty{
						Type:        DefPropertyTypeInteger,
						Description: "test_int",
						ExtraDefs: Map[any]{
							"maximum": 10,
						},
					},
				},
				Required: DefRequired{"test_int"},
			},
		},
	}
	serializedDefProperty, err := json.MarshalIndent(defProperty, "", "    ")
	if err != nil {
		t.Errorf("DefProperty: %s", err)
		return
	}
	var (
		deserializedDefProperty  map[string]any
		deserializedJSONProperty map[string]any
	)
	if err = json.Unmarshal(serializedDefProperty, &deserializedDefProperty); err != nil {
		t.Errorf("DefProperty: %s", err)
		return
	}
	if err = json.Unmarshal([]byte(jsonProperty), &deserializedJSONProperty); err != nil {
		t.Errorf("DefProperty: %s", err)
		return
	}
	if !reflect.DeepEqual(deserializedDefProperty, deserializedJSONProperty) {
		t.Errorf(`DefProperty: 

============================================================
%s
============================================================
%s
============================================================
`, string(serializedDefProperty), jsonProperty)
		return
	}
}
