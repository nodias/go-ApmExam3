package service

import "context"

func Example_PanicGenerator() {
	ctx := context.Background()
	PanicGenerator(ctx, "ㅁ")
	//Output:
	//panic
}
