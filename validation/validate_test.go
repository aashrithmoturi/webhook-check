package validation

import "testing"

func Test_validateOwner(t *testing.T) {
	type args struct {
		file string
		s    string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test-1",
			args: args{
				file: "scripts/SimpleConcordPlaybook.yaml",
				s:    "a@walmart.com",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateOwner(tt.args.file, tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("validateOwner() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
