package service

import (
	"context"
	"fmt"
)

func producer(ctx context.Context)  {
	fmt.Println("producer")
}

func SyncSend(ctx context.Context, msg []byte)  {
	fmt.Println("SyncSend")
}
