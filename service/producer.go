package service

import (
	"context"
	"fmt"
)

func producer(ctx context.Context)  {
	fmt.Println("producer")
}

func TestSyncSend(ctx context.Context, msg []byte)  {
	fmt.Println("TestSyncSend")
}