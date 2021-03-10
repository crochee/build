package main

import "testing"

func TestInitConfig(t *testing.T) {
	cf := &Config{
		Product: &Product{
			Type: "docker",
			Name: "obs:latest",
		},
		Script: "build.sh",
	}
	y := Yml{path: "obs.yml"}
	if err := y.Encode(cf); err != nil {
		t.Error(err)
	}
}
