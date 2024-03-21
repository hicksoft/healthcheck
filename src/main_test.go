package main

import (
	"testing"
)

func TestReadConfig(t *testing.T) {
	t.Setenv("CONFIG_FILE", "../config.yml")

	monitors := readConfig()

	if monitors["Monitor1"].Target != "http://myservice1.com" {
		t.Error("Check is not populated for Monitor 1")
	}
	if monitors["Monitor1"].Ping != "http://healthceck.xyz" {
		t.Error("Ping is not populated for Monitor 1")
	}

	if monitors["Monitor2"].Target != "http://myservice2.com" {
		t.Error("Check is not populated for Monitor 2")
	}
	if monitors["Monitor2"].Ping != "http://healthceck.abc" {
		t.Error("Ping is not populated for Monitor 2")
	}
	if monitors["Monitor2"].Period != "30s" {
		t.Error("Period is not populated for Monitor 2")
	}
	if monitors["Monitor2"].Status != 201 {
		t.Error("Status is not populated for Monitor 2")
	}
}

func TestCheckStatusCode(t *testing.T) {
	result, code := checkStatusCode("http://google.com", 200)
	if !result {
		t.Error("checkStatusCode failed")
	}
	if code != 200 {
		t.Error("checkStatusCode failed")
	}
}

func TestCheckStatusCodeWrongCode(t *testing.T) {
	result, _ := checkStatusCode("http://google.com", 201)
	if result {
		t.Error("checkStatusCode succeeded for incorrect code")
	}
}

func TestCheckStatusCodeInvalidTarget(t *testing.T) {
	result, _ := checkStatusCode("google.com", 200)
	if result {
		t.Error("checkStatusCode succeeded for invalid address")
	}
}

func TestCreateJobPopulateDefaults(t *testing.T) {
	monitor := Monitor{
		Target: "abc",
		Ping:   "abc",
	}

	createJob("name", &monitor)

	if monitor.Status != 200 {
		t.Error("Status not populated with default")
	}
	if monitor.Period != "10m" {
		t.Error("Period not populated with default")
	}
}

func TestCreateJobNoTarget(t *testing.T) {
	monitor := Monitor{
		Ping: "abc",
	}

	_, err := createJob("name", &monitor)

	if err == nil {
		t.Error("Monitor without a 'Target' successfully created job")
	}
}

func TestCreateJobNoPing(t *testing.T) {
	monitor := Monitor{
		Target: "abc",
	}

	_, err := createJob("name", &monitor)

	if err == nil {
		t.Error("Monitor without a 'Ping' successfully created job")
	}
}
