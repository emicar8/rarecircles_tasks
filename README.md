# rarecircles_tasks

This repo contains the tasks for the rarecircles technical assesment. Cobra-cli was used to create the cli interface available in the project.

## Task 1 implementation of the “adapter” structural design pattern
This task simulates a use case where a payment is usually created based on a authorization. Even though that is the normal logic there are border cases where the payments must be created from a capture. To cover that border case, an adapter is implemented to allow a capture to be passed to the process payment function. The command to run task 1 is shown below.

```
go run main.go task1
```

## Task 2 implementation of concurrency and parallelism
The simulated situation shown below is a simplified version where when a transaction is received the balance and card validations are done. These validations use external APIs and are independent from eachother, making them prime candidates for parallelization. To futher simulate the situation, both repos can be configured with what response should be given (if valid or invalid) and how much time in seconds the process should take. Another consideration is that both validations must be valid to allow the transactions to continue. The following code creates an errgroup with context to be able to parrallelize, but also propagate errors and mainting synchronization. Then to goroutines are called to execute each repo function. These functions are built to be able to return either when the repo finishes or the context is cancelled. This is important because it allows to not waste time waiting for one validation if the other has already returned with error or with an invalid result. If an error is encountered in any of the repos then all goroutines finish. The command to run task 2 is shown below
```
go run main.go task2 --cardSleep=2 --balanceSleep=3 --validBalance=true --validCard=true
```
