package task2

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

var (
	ValidCard, ValidBalance bool
	CardSleep, BalanceSleep int
)

// Task2 implementation of concurrency and parallelism
// The simulated situation shown below is a simplified version where when a transaction is received the balance and card validations are done.
// These validations use external APIs and are independent from eachother, making them prime candidates for parallelization. To futher simulate
// the situation, both repos can be configured with what response should be given (if valid or invalid) and how much time in seconds the process should take.
// Another consideration is that both validations must be valid to allow the transactions to continue.
// The following code creates an errgroup with context to be able to parrallelize, but also propagate errors and mainting synchronization.
// Then to goroutines are called to execute each repo function. These functions are built to be able to return either when the repo finishes or
// the context is cancelled. This is important because it allows to not waste time waiting for one validation if the other has already returned with error
// or with an invalid result. If an error is encountered in any of the repos then all goroutines finish.
func Task2() (bool, bool, error) {
	var isValidBalance, isValidCard bool

	balanceRepo := BalanceRepo{validBalance: ValidBalance, seconds: BalanceSleep}
	cardRepo := CardRepo{validCard: ValidCard, seconds: CardSleep}

	balanceChan := make(chan bool)
	cardChan := make(chan bool)

	group, ctx := errgroup.WithContext(context.Background())

	group.Go(func() error {
		var err error
		if isValidBalance, err = balanceRepo.cancellableValidBalance(ctx, balanceChan); err != nil {
			return err
		}
		return nil
	})

	group.Go(func() error {
		var err error
		if isValidCard, err = cardRepo.cancellableValidCard(ctx, cardChan); err != nil {
			return err
		}
		return nil
	})

	if err := group.Wait(); err != nil {
		fmt.Println("group ended with error")
		return isValidBalance, isValidCard, err
	}

	fmt.Printf("validations: %t, %t", isValidBalance, isValidCard)
	return isValidBalance, isValidCard, nil
}

type BalanceRepo struct {
	validBalance bool
	seconds      int
}

func (repo BalanceRepo) cancellableValidBalance(ctx context.Context, responseChan chan bool) (bool, error) {
	var err error
	go func() {
		err = repo.isValidBalance(responseChan)
	}()
	select {
	case <-ctx.Done():
		return false, nil
	case result := <-responseChan:
		return result, err
	}
}

func (repo BalanceRepo) isValidBalance(responseChan chan bool) error {
	fmt.Println("entered valid balance")
	time.Sleep(time.Duration(repo.seconds) * time.Second)
	if !repo.validBalance {
		responseChan <- repo.validBalance
		return fmt.Errorf("valid balance error")
	}
	fmt.Println("exiting valid balance")
	responseChan <- repo.validBalance
	return nil
}

type CardRepo struct {
	validCard bool
	seconds   int
}

func (repo CardRepo) cancellableValidCard(ctx context.Context, responseChan chan bool) (bool, error) {
	var err error
	go func() {
		err = repo.isValidCard(responseChan)
	}()
	select {
	case <-ctx.Done():
		return false, nil
	case result := <-responseChan:
		return result, err
	}
}

func (repo CardRepo) isValidCard(responseChan chan bool) error {
	fmt.Println("entered valid card")
	time.Sleep(time.Duration(repo.seconds) * time.Second)
	if !repo.validCard {
		responseChan <- repo.validCard
		return fmt.Errorf("valid card error")
	}
	fmt.Println("exiting valid card")
	responseChan <- repo.validCard
	return nil
}
