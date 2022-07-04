package add

import (
	"testing"

	"github.com/joerdav/zapray"
)

func TestHandler_Handle(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name   string
		args   args
		result int
	}{
		{
			name: "Happy path 1+1=2",
			args: args{
				a: 1,
				b: 1,
			},
			result: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Handler{
				Log: zapray.NewNop(),
			}
			if gotX := h.Handle(tt.args.a, tt.args.b); gotX != tt.result {
				t.Errorf("Handler.Handle() = %v, want %v", gotX, tt.result)
			}
		})
	}
}
