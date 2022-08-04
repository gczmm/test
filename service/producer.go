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

func AsyncSend(ctx context.Context, msg []byte)  {
	fmt.Println("asyncSend")
}

func TestSyncSend(ctx context.Context, msg []byte)  {
	fmt.Println("TestSyncSend")
}

func TestASyncSend(ctx context.Context, msg []byte) {
	fmt.Println("TestASyncSend")
}

func TestSyncSend3(ctx context.Context, msg []byte)  {
	fmt.Println("TestSyncSend3")
}

func TestASyncSend3(ctx context.Context, msg []byte) {
	fmt.Println("TestASyncSend3")
}