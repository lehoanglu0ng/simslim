package main

import (
	"testing"
	"time"

	"github.com/mobai-app/simslim"
)

func TestSortByRecentUse(t *testing.T) {
	now := time.Now()
	devices := []simslim.Device{
		{Name: "never-old", OSVersion: "17.5"},
		{Name: "yesterday", OSVersion: "17.5", LastBootedAt: now.Add(-24 * time.Hour)},
		{Name: "never-new", OSVersion: "26.0"},
		{Name: "just-now", OSVersion: "17.5", LastBootedAt: now},
	}
	sortByRecentUse(devices)
	want := []string{"just-now", "yesterday", "never-new", "never-old"}
	for i, name := range want {
		if devices[i].Name != name {
			t.Fatalf("position %d = %q, want %q", i, devices[i].Name, name)
		}
	}
}
