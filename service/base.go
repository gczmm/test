package service

import (
	"context"
	"fmt"
)

func GetMemberID(ctx context.Context)  {
	fmt.Println("GetMemberID")
}

func GetMember(ctx context.Context, memberID int64)  {
	fmt.Println("GetMember")
}

func GetInviteList(ctx context.Context, memberID int64)  {
	fmt.Println("GetInviteList")
}

func Join(ctx context.Context, memberID int64)  {
	fmt.Println("Join")
}

func Submit(ctx context.Context, memberID int64)  {
	fmt.Println("Submit")
}