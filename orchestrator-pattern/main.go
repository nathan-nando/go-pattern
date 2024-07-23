package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"sync/atomic"
	"time"
)

type Service func(ctx context.Context) (int64, error)

func main() {
	svcs := []Service{
		serviceA,
		serviceB,
		serviceC,
	}

	var sum int64

	start := time.Now()

	g, ctx := errgroup.WithContext(context.Background())

	for _, svc := range svcs {
		sv := svc

		g.Go(func() error {
			v, err := sv(ctx)

			if err != nil {
				return err
			}

			atomic.AddInt64(&sum, v)

			return nil
		})
	}

	if err := g.Wait(); err != nil {
		log.Fatalln("Group", err)
	}

	fmt.Println("Total", sum, "Duration", time.Now().Sub(start))
}

func serviceA(ctx context.Context) (int64, error) {
	time.Sleep(200 * time.Millisecond)
	return 1, nil
}

func serviceB(ctx context.Context) (int64, error) {
	time.Sleep(300 * time.Millisecond)
	return 2, nil
}

func serviceC(ctx context.Context) (int64, error) {
	time.Sleep(400 * time.Millisecond)
	return 3, nil
}
