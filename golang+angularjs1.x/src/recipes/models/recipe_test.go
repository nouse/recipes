package models

import (
	"testing"
	"encoding/json"
	"bytes"
	"reflect"
)

func TestRecipeUnmarshall(t *testing.T) {
	body := []byte(`{"id":1,"title":"","description":"","ingredients":[{"amount":10},{"amount":20}],"instructions":""}`)
	expected := Recipe{
		ID: 1,
		Ingredients: jsonb([]byte(`[{"amount":10},{"amount":20}]`)),
	}
	actual := Recipe{}
	err := json.Unmarshal(body, &actual)
	if err != nil {
		t.Errorf("Unexpected error when marshalling, %s", err)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Unexpected unmarshalled result, actual: %+v", actual)
	}
}

func TestRecipeMarshall(t *testing.T) {
	recipe := Recipe{
		ID: 1,
		Ingredients: jsonb([]byte(`[{"amount": 10},{"amount": 20}]`)),
	}
	body, err := json.Marshal(recipe)
	if err != nil {
		t.Errorf("Unexpected error when marshalling, %s", err)
	}
	expected := []byte(`{"id":1,"title":"","description":"","ingredients":[{"amount":10},{"amount":20}],"instructions":""}`)
	if !bytes.Equal(body, expected) {
		t.Errorf("Unexpected body, actual: %s", string(body))
	}
}
