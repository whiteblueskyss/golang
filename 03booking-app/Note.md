```go func() {
defer wg.Done()
// ... work ...
}()
wg.Wait() // Waits until the goroutine calls Done()
```

# Goroutines and the `go` Keyword in Go

## What does the `go` keyword do?

- The `go` keyword starts a new goroutine, which is a lightweight thread managed by the Go runtime.
- The function after `go` runs concurrently (in the background) with the rest of your program.

## How does execution proceed?

- When you call `go someFunction()`, the current function does **not** wait for `someFunction` to finish.
- Execution continues immediately to the next line after the `go` statement.
- The goroutine runs independently.

## What if a goroutine fails (panics)?

- If a goroutine panics and the panic is not recovered inside that goroutine, only that goroutine stops.
- Other goroutines (including `main`) continue running unless the main goroutine exits, which ends the program and kills all running goroutines.

## Synchronizing Goroutines with WaitGroup

```go
var wg = sync.WaitGroup{}
wg.Add(1)
go func() {
    defer wg.Done()
    // ... work ...
}()
wg.Wait() // Waits until the goroutine calls Done()
```

## Key Points

- The `main` function does not wait for goroutines to finish unless you explicitly synchronize (e.g., with `WaitGroup` or channels).
- If `main()` returns, the program exits, and all running goroutines are killed immediately.
- Panics in a goroutine do not crash the whole program unless they occur in `main` or are not handled and `main` exits.

You should call wg.Wait() after you have started all your goroutines and want to wait for them to finish before moving on. Typically, wg.Wait() is placed at the end of your main function or just before any code that must run only after all goroutines are done.

What happens if you put code after wg.Wait()?
Any code after wg.Wait() will execute only after all goroutines have called wg.Done() and the WaitGroup counter is zero.
What if you put code between starting the goroutine and wg.Wait()?
That code will run immediately, in parallel with the goroutine.
wg.Wait() will block until all goroutines are done, then the code after wg.Wait() will run.
Summary:

Place wg.Wait() where you want to block until all goroutines finish.
Code after wg.Wait() runs only after all goroutines are complete.
Code before wg.Wait() runs concurrently with the goroutines.
