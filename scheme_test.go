package main

import (
	"reflect"
	"testing"
)

func Test_schemeFromFile(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name  string
		args  args
		want  *scheme
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := schemeFromFile(tt.args.fileName)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("schemeFromFile() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("schemeFromFile() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_scheme_mustacheContext(t *testing.T) {
	type fields struct {
		Name   string
		Slug   string
		Scheme string
		Author string
		Colors map[string]color
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &scheme{
				Name:   tt.fields.Name,
				Slug:   tt.fields.Slug,
				Scheme: tt.fields.Scheme,
				Author: tt.fields.Author,
				Colors: tt.fields.Colors,
			}
			if got := s.mustacheContext(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("scheme.mustacheContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loadSchemes(t *testing.T) {
	type args struct {
		schemeFile string
	}
	tests := []struct {
		name  string
		args  args
		want  []*scheme
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := loadSchemes(tt.args.schemeFile)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadSchemes() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("loadSchemes() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
