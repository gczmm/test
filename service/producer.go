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

func TestSyncSend2(ctx context.Context, msg []byte)  {
	fmt.Println("TestSyncSend2")
}

func TestASyncSend2(ctx context.Context, msg []byte) {
	fmt.Println("TestASyncSend2")
}

func TestSyncSend3(ctx context.Context, msg []byte)  {
	fmt.Println("TestSyncSend3")
}

func TestASyncSend3(ctx context.Context, msg []byte) {
	fmt.Println("TestASyncSend3")
}

func TestSyncSend5(ctx context.Context, msg []byte)  {
	fmt.Println("TestSyncSend5")
}

func TestASyncSend5(ctx context.Context, msg []byte) {
	fmt.Println("TestASyncSend5")
}
func TestSyncSend4(ctx context.Context, msg []byte)  {
	fmt.Println("TestSyncSend4")
}

func TestASyncSend4(ctx context.Context, msg []byte) {
	fmt.Println("TestASyncSend3")
}