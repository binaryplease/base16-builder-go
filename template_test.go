package main

import (
	"reflect"
	"testing"
)

func Test_templatesFromFile(t *testing.T) {
	type args struct {
		templatesDir string
	}
	tests := []struct {
		name    string
		args    args
		want    []*template
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := templatesFromFile(tt.args.templatesDir)
			if (err != nil) != tt.wantErr {
				t.Errorf("templatesFromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("templatesFromFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_template_Render(t *testing.T) {
	type fields struct {
		Name      string
		Dir       string
		Extension string
		OutputDir string
	}
	type args struct {
		schemes []*scheme
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
			t := &template{
				Name:      tt.fields.Name,
				Dir:       tt.fields.Dir,
				Extension: tt.fields.Extension,
				OutputDir: tt.fields.OutputDir,
			}
			if err := t.Render(tt.args.schemes); (err != nil) != tt.wantErr {
				t.Errorf("template.Render() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_loadTemplates(t *testing.T) {
	type args struct {
		templateFile string
		targets      []string
	}
	tests := []struct {
		name  string
		args  args
		want  []*template
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := loadTemplates(tt.args.templateFile, tt.args.targets)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadTemplates() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("loadTemplates() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
