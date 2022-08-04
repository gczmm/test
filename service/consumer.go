package service

import (
	"context"
	"fmt"
)

func task(ctx context.Context, taskName string)  {
	fmt.Println("task")
}

func consumer(ctx context.Context, msg []byte)  {
	fmt.Println("consumer")
}

