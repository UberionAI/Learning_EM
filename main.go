package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"log"
	"sync"
	"time"
)

func main() {
	ctx := context.Background()

	//custom config pgxpool
	config, err := pgxpool.ParseConfig("postgres://postgres:password@localhost:5432/app_db?sslmode=disable")
	if err != nil {
		panic(err)
	}
	config.MaxConns = 20
	config.MinConns = 2
	config.MaxConnLifetime = 50 * time.Second
	config.MaxConnIdleTime = 30 * time.Second
	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		panic(err)
	}
	defer pool.Close()

	//creating poll workers for testing
	var wg sync.WaitGroup
	wg.Add(20)

	for i := 0; i < 20; i++ {
		go func() {
			defer wg.Done()
			var count int
			err := pool.QueryRow(ctx, "SELECT COUNT(*) FROM users").Scan(&count)
			if err != nil {
				log.Printf("QueryRow failed: %v", err)
			}
			fmt.Printf("Pool in work! Total users: %d\n", count)
		}()
	}
	wg.Wait()

	stats := pool.Stat()
	fmt.Printf("Pool stats: %d/%d/%d active/idle/max\n",
		stats.AcquiredConns(), stats.IdleConns(), stats.MaxConns())
	fmt.Printf("Lifetime: %v, Idle time: %v\n",
		config.MaxConnLifetime, config.MaxConnIdleTime)
}
