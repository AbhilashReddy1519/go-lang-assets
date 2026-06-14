package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

// errorsDemo runs short demos of Go error-handling patterns.
func errorsDemo() {
	fmt.Println("== basic error return ==")
	if err := basicError(); err != nil {
		fmt.Println("basicError returned:", err)
	}

	fmt.Println("== sentinel error and os.IsNotExist ==")
	sentinelDemo()

	fmt.Println("== wrapping and errors.Is / errors.As ==")
	wrapAndCheckDemo()

	fmt.Println("== errors.Join (multiple errors) ==")
	joinErrorDemo()

	fmt.Println("== io.EOF handling when reading ==")
	eofDemo()

	fmt.Println("== panic and recover ==")
	panicRecoverDemo()

	fmt.Println("== context cancellation ==")
	contextCancelDemo()

	fmt.Println("== timeout-like interface assertion ==")
	timeoutInterfaceDemo()
}

// 1) Basic error return pattern
func basicError() error {
	// Validate or attempt an operation and return an error value when something goes wrong.
	return errors.New("something went wrong")
}

// 2) Sentinel errors and using stdlib helpers (os.IsNotExist)
var ErrNotFound = errors.New("not found")

func sentinelDemo() {
	// Simulate a function that returns a sentinel error.
	if err := findItem(false); err != nil {
		// callers can compare with errors.Is
		if errors.Is(err, ErrNotFound) {
			fmt.Println("item was not found (sentinel)")
		} else if os.IsNotExist(err) {
			fmt.Println("os: not exist")
		} else {
			fmt.Println("other error:", err)
		}
	}
}

func findItem(exists bool) error {
	if !exists {
		// return sentinel error
		return ErrNotFound
	}
	return nil
}

// 3) Wrapping errors to add context; checking with errors.Is and errors.As
type MyError struct {
	Code int
	Msg  string
}

func (e *MyError) Error() string { return fmt.Sprintf("%d:%s", e.Code, e.Msg) }

func wrapAndCheckDemo() {
	// produce a typed error
	base := &MyError{Code: 404, Msg: "user missing"}
	// wrap with context
	wrapped := fmt.Errorf("loading user: %w", base)

	// errors.Is checks the chain for equality by Unwrap and Is methods
	if errors.Is(wrapped, base) {
		fmt.Println("errors.Is: wrapped contains base")
	}

	// errors.As can extract a typed error from the chain
	var me *MyError
	if errors.As(wrapped, &me) {
		fmt.Println("errors.As: extracted MyError code=", me.Code)
	}

	// show unwrapping
	fmt.Println("unwrapped:", errors.Unwrap(wrapped))
}

// 4) Joining multiple errors (Go 1.20+)
func joinErrorDemo() {
	e1 := errors.New("first")
	e2 := errors.New("second")
	ej := errors.Join(e1, e2)
	fmt.Println("joined error:", ej)
}

// 5) io.EOF handling example
func eofDemo() {
	r := strings.NewReader("abc")
	buf := make([]byte, 2)
	for {
		n, err := r.Read(buf)
		if n > 0 {
			fmt.Print(string(buf[:n]))
		}
		if err != nil {
			if errors.Is(err, io.EOF) {
				fmt.Println(" <- reached EOF")
				break
			}
			fmt.Println("read error:", err)
			break
		}
	}
}

// 6) Panic and recover — use sparingly (to isolate a failure boundary)
func panicRecoverDemo() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from panic:", r)
		}
	}()

	// deliberate panic
	causePanic()
	fmt.Println("this line will not run")
}

func causePanic() {
	panic("unexpected state")
}

// 7) Context cancellation and returning errors
func contextCancelDemo() {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	ch := make(chan error, 1)
	go func() {
		// simulate work that respects ctx
		select {
		case <-time.After(100 * time.Millisecond):
			ch <- nil // finished
		case <-ctx.Done():
			ch <- ctx.Err()
		}
	}()

	if err := <-ch; err != nil {
		fmt.Println("work ended with:", err)
	} else {
		fmt.Println("work completed successfully")
	}
}

// 8) Timeout/Temporary style interface assertion
type timeoutErr struct{}

func (timeoutErr) Error() string   { return "temporary timeout" }
func (timeoutErr) Timeout() bool   { return true }
func (timeoutErr) Temporary() bool { return true }

func timeoutInterfaceDemo() {
	var err error = timeoutErr{}
	// assert to an interface that provides Timeout()
	if te, ok := err.(interface{ Timeout() bool }); ok && te.Timeout() {
		fmt.Println("error reports timeout")
	} else {
		fmt.Println("no timeout information")
	}
}

// Extra: demonstrate mapping os errors (e.g., file not exist)
func osErrorDemo() {
	_, err := os.ReadFile("file-that-does-not-exist.txt")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("file is missing")
		} else {
			fmt.Println("other file error:", err)
		}
	}
}
