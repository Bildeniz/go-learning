package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "Main Context", "Main context is here")

	F1(ctx)

}

func F1(ctx context.Context) {
	fmt.Println("F1 is runned", ctx.Value("Main Context"))
	F2(ctx)
}
func F2(ctx context.Context) {
	fmt.Println("F2 is runned", ctx.Value("Main Context"))
	F3(ctx)
}
func F3(ctx context.Context) {
	fmt.Println("F3 is runned", ctx.Value("Main Context"))
}
