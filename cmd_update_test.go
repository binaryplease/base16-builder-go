package main

import (
	"testing"

	"github.com/Sirupsen/logrus"
)

func Test_downloadSourceList(t *testing.T) {
	type args struct {
		sourceFile string
		targetDir  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := downloadSourceList(tt.args.sourceFile, tt.args.targetDir); got != tt.want {
				t.Errorf("downloadSourceList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cloneRepo(t *testing.T) {
	type args struct {
		src  string
		dest string
		key  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cloneRepo(tt.args.src, tt.args.dest, tt.args.key); got != tt.want {
				t.Errorf("cloneRepo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_handleVcsError(t *testing.T) {
	type args struct {
		logger *logrus.Entry
		err    error
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handleVcsError(tt.args.logger, tt.args.err)
		})
	}
}
