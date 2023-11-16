package contexts

import (
	"context"
	"fmt"
)

func ContextWithValues() {
	// Another way to use context is to pass values between goroutines
	// This is useful when you want to pass some metadata between goroutines
	// For example, you want to pass the request ID between goroutines
	authenticate(context.Background())
}

func authenticate(ctx context.Context) {
	// Get the request ID from the context
	ctx = context.WithValue(ctx, "ID", "123456")
	// Do some authentication with the request ID
	func1(ctx)
}

func func1(ctx context.Context) {
	// Get the request ID from the context
	fmt.Println("Request ID From FUNCTION 1:", ctx.Value("ID"))
	// Do some authentication with the request ID
	func2(ctx)
}

func func2(ctx context.Context) {
	fmt.Println("Request ID From FUNCTION 2:", ctx.Value("ID"))
}
