package client

import (
	"testing"
)

func TestGenderFetcher_Fetch(t *testing.T) {
	fetcher := NewGenderFetcher()
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "with name",
			args:    args{name: "Ivan"},
			wantErr: false,
		},
		{
			name:    "without name",
			args:    args{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := fetcher.Fetch(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenderFetcher.Fetch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
