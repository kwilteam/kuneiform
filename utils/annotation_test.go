package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseAnnotation(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name           string
		args           args
		wantAnnotation *Annotation
		wantErr        bool
	}{
		{
			name: "empty",
			args: args{
				input: "",
			},
			wantAnnotation: nil,
			wantErr:        true,
		},
		{
			name: "invalid not complete",
			args: args{
				input: "t",
			},
			wantAnnotation: nil,
			wantErr:        true,
		},
		{
			name: "invalid missing parenthesis",
			args: args{
				input: "t(a=1",
			},
			wantAnnotation: nil,
			wantErr:        true,
		},
		{
			name: "single int",
			args: args{
				input: "t(a=1)",
			},
			wantAnnotation: &Annotation{
				Target: "t",
				Attributes: []AnnotationAttribute{
					{
						Name:  "a",
						Value: "1",
					},
				},
			},
		},
		{
			name: "single string",
			args: args{
				input: "t(a='1')",
			},
			wantAnnotation: &Annotation{
				Target: "t",
				Attributes: []AnnotationAttribute{
					{
						Name:  "a",
						Value: "'1'",
					},
				},
			},
		},
		{
			name: "multiple int and string",
			args: args{
				input: "t(a=1,b='1',c=323,d='whatever')",
			},
			wantAnnotation: &Annotation{
				Target: "t",
				Attributes: []AnnotationAttribute{
					{
						Name:  "a",
						Value: "1",
					},
					{
						Name:  "b",
						Value: "'1'",
					},
					{
						Name:  "c",
						Value: "323",
					},
					{
						Name:  "d",
						Value: "'whatever'",
					},
				},
			},
		},
		{
			name: "multiple int and string with random space",
			args: args{
				input: "t(a= 1, b='1',c =323,d = 'whatever' )",
			},
			wantAnnotation: &Annotation{
				Target: "t",
				Attributes: []AnnotationAttribute{
					{
						Name:  "a",
						Value: "1",
					},
					{
						Name:  "b",
						Value: "'1'",
					},
					{
						Name:  "c",
						Value: "323",
					},
					{
						Name:  "d",
						Value: "'whatever'",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAnnotation, err := ParseAnnotation(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseAnnotation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.EqualValues(t, tt.wantAnnotation, gotAnnotation)
		})
	}
}
