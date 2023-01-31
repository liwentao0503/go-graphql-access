package common

import (
	"context"
	"errors"
	"testing"
)

func TestGoAndWait(t *testing.T) {
	type args struct {
		ctx     context.Context
		handles []Handle
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				context.Background(),
				[]Handle{
					func() error {
						return nil
					},
					func() error {
						return errors.New("test1")
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GoAndWait(tt.args.ctx, tt.args.handles); (err != nil) != tt.wantErr {
				t.Errorf("GoAndWait() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
