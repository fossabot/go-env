package goenv

import (
	"testing"
)

func TestEnvLoadJson(t *testing.T) {
	err := Init("test", "json", "env")
	if err != nil {
		t.Fail()
	}
}

func TestEnvReadJson(t *testing.T) {
	err := Init("test", "json", "env")
	if err != nil {
		t.Fail()
	}
	// test string
	var result = Env.GetString("test")
	if result != "success-json" {
		t.Fail()
	}
}
