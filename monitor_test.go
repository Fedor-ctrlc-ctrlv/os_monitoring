package main

import (
	"testing"
)

func TestGetCPUload(t *testing.T) {
	load := getCPUload()
	if load < 0 || load > 100 {
		t.Errorf("CPU load should be between 0 and 100, got %.2f", load)
	}
}

func TestGetRamUsage(t *testing.T) {
	usage := getRamusage()
	if usage < 0 || usage > 100 {
		t.Errorf("RAM usage should be between 0 and 100, got %.2f", usage)
	}
}