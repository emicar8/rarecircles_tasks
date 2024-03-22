package task2_test

import (
	"fmt"
	"testing"

	"github.com/ecaron/rarecircles_tasks/task2"
	"github.com/go-playground/assert/v2"
)

func Test_task2(t *testing.T) {
	type args struct {
		validCard    bool
		validBalance bool
		cardSleep    int
		balanceSleep int
	}

	type wants struct {
		isValidCard    bool
		isValidBalance bool
		err            error
	}

	tests := []struct {
		name  string
		args  func() *args
		wants func() *wants
	}{
		{
			name: "Valid card and balance",
			args: func() *args {
				return &args{
					validCard:    true,
					validBalance: true,
					cardSleep:    1,
					balanceSleep: 2,
				}
			},
			wants: func() *wants {
				return &wants{
					isValidCard:    true,
					isValidBalance: true,
					err:            nil,
				}
			},
		},
		{
			name: "Valid card and invalid balance",
			args: func() *args {
				return &args{
					validCard:    true,
					validBalance: false,
					cardSleep:    1,
					balanceSleep: 2,
				}
			},
			wants: func() *wants {
				return &wants{
					err: fmt.Errorf("valid balance error"),
				}
			},
		},
		{
			name: "Invalid card and valid balance",
			args: func() *args {
				return &args{
					validCard:    false,
					validBalance: true,
					cardSleep:    1,
					balanceSleep: 2,
				}
			},
			wants: func() *wants {
				return &wants{
					err: fmt.Errorf("valid card error"),
				}
			},
		},
		{
			name: "Invalid card and invalid balance",
			args: func() *args {
				return &args{
					validCard:    false,
					validBalance: false,
					cardSleep:    1,
					balanceSleep: 20,
				}
			},
			wants: func() *wants {
				return &wants{
					err: fmt.Errorf("valid card error"),
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args, wants := tt.args(), tt.wants()
			task2.ValidCard = args.validCard
			task2.ValidBalance = args.validBalance
			task2.CardSleep = args.cardSleep
			task2.BalanceSleep = args.balanceSleep

			resultValidBalance, resultValidCard, resultErr := task2.Task2()

			if wants.err != nil {
				assert.Equal(t, wants.err, resultErr)
			} else {
				assert.Equal(t, wants.isValidBalance, resultValidBalance)
				assert.Equal(t, wants.isValidCard, resultValidCard)
			}
		})
	}
}
