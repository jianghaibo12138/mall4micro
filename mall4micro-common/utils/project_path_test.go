package utils

import (
	"fmt"
	"testing"
)

func TestProjectBasePath(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "c1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ProjectBasePath()
			fmt.Println(got)
		})
	}
}

func Test_getCurrentAbPathByExecutable(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "c1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getCurrentAbPathByExecutable()
			t.Log(got)
		})
	}
}
