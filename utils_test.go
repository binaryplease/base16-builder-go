package main

import (
	"reflect"
	"testing"

	yaml "gopkg.in/yaml.v2"
)

func Test_color_UnmarshalYAML(t *testing.T) {
	type fields struct {
		R int
		G int
		B int
	}
	type args struct {
		f func(interface{}) error
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &color{
				R: tt.fields.R,
				G: tt.fields.G,
				B: tt.fields.B,
			}
			if err := c.UnmarshalYAML(tt.args.f); (err != nil) != tt.wantErr {
				t.Errorf("color.UnmarshalYAML() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_readSourcesList(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name    string
		args    args
		want    yaml.MapSlice
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readSourcesList(tt.args.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("readSourcesList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readSourcesList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateMapSlice(t *testing.T) {
	type args struct {
		sources yaml.MapSlice
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateMapSlice(tt.args.sources); (err != nil) != tt.wantErr {
				t.Errorf("validateMapSlice() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_errorOrFatal(t *testing.T) {
	type args struct {
		preferError bool
		msg         string
		data        []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errorOrFatal(tt.args.preferError, tt.args.msg, tt.args.data...)
		})
	}
}
