package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go func() {
		time.Sleep(time.Second * 10)
		cancel()
	}()

	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Time exceeded")
	case <-time.After(time.Second * 5):
		fmt.Print("Room reserved with success")
	}
}
