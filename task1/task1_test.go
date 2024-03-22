package task1_test

import (
	"testing"

	"github.com/ecaron/rarecircles_tasks/task1"
	"github.com/go-playground/assert/v2"
)

func Test_processPayment(t *testing.T) {
	type args struct {
		payment task1.Payment
	}

	type want struct {
		result string
	}

	tests := []struct {
		name string
		args func() *args
		want func() *want
	}{
		{
			name: "Authorization",
			args: func() *args {
				auth := task1.Authorization{AuthorizedAmount: 123.0}
				return &args{payment: &auth}
			},
			want: func() *want {
				return &want{result: "approved amount: 123.000000"}
			},
		},
		{
			name: "Capture",
			args: func() *args {
				cap := task1.Capture{CaptureAmount: 23.0}
				capAdapter := task1.CaptureAdapter{Capture: &cap}
				return &args{payment: &capAdapter}
			},
			want: func() *want {
				return &want{result: "approved amount: 23.000000"}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args, want := tt.args(), tt.want()
			client := task1.Client{}
			result := client.ProcessPayment(args.payment)
			assert.Equal(t, want.result, result)
		})
	}
}
