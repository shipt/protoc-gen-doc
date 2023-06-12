package gendoc

import "testing"

func TestGetContent(t *testing.T) {
	type args struct {
		files       []*File
		baseMessage string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "success",
			args: args{
				baseMessage: "Spec",
				files: []*File{
					{
						Messages: []*Message{
							{
								LongName: "Spec",
								Fields: []*MessageField{
									{
										Name:     "Spec.SpecTest",
										LongType: "SpecTest",
										FullType: "SpecTest",
									},
								},
							},
							{
								FullName: "SpecTest",
								Fields: []*MessageField{
									{
										Name:     "SpecTest.String",
										LongType: "string",
										FullType: "string",
									},
									{
										Name:     "SpecTest.Deploy",
										LongType: "Deploy",
										FullType: "Deploy",
									},
								},
							},
						},
					},
					{
						Messages: []*Message{
							{
								FullName: "Deploy",
								Fields: []*MessageField{
									{
										Name:     "Deploy.string",
										LongType: "string",
										FullType: "string",
									},
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetContent(tt.args.files, tt.args.baseMessage)
		})
	}
}
