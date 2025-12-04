package main

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestDatabaseConnection(t *testing.T) {
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, "postgres://postgres:password@localhost:5432/app_db?sslmode=disable")
	assert.NoError(t, err)
	defer pool.Close()

	var version string
	err = pool.QueryRow(ctx, "SELECT version()").Scan(&version)
	assert.NoError(t, err)
	assert.Contains(t, version, "PostgreSQL")
	t.Logf("Succesfully connected to: %s", version)
}

func TestCountUsers(t *testing.T) {
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, "postgres://postgres:password@localhost:5432/app_db?sslmode=disable")
	assert.NoError(t, err)
	defer pool.Close()

	var count int
	err = pool.QueryRow(ctx, "SELECT COUNT(*) FROM users").Scan(&count)
	assert.NoError(t, err)
	t.Logf("Succesfully count users: %d", count)
	assert.GreaterOrEqual(t, count, 0)
}

func TestInsertAndSelectUser(t *testing.T) {
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, "postgres://postgres:password@localhost:5432/app_db?sslmode=disable")
	assert.NoError(t, err)
	defer pool.Close()

	_, err = pool.Exec(ctx, "INSERT INTO users (username, email) VALUES ($1, $2)",
		"test_user", "test@example.com")
	assert.NoError(t, err)

	var username string
	err = pool.QueryRow(ctx, "SELECT username FROM users WHERE email = $1", "test@example.com").Scan(&username)
	assert.NoError(t, err)
	assert.Equal(t, "test_user", username)

	_, err = pool.Exec(ctx, "DELETE FROM users WHERE email = $1", "test@example.com")
	assert.NoError(t, err)

	t.Log("Insert → Select → Delete works!")
}

func TestPoolConcurrency(t *testing.T) {
	ctx := context.Background()

	config, err := pgxpool.ParseConfig("postgres://postgres:password@localhost:5432/app_db?sslmode=disable")
	assert.NoError(t, err)
	config.MaxConns = 20
	config.MinConns = 5

	pool, err := pgxpool.NewWithConfig(ctx, config)
	assert.NoError(t, err)
	defer pool.Close()

	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			var count int
			pool.QueryRow(ctx, "SELECT COUNT(*) FROM users").Scan(&count)
		}()
	}
	wg.Wait()

	stats := pool.Stat()
	assert.Equal(t, int32(20), stats.MaxConns())
	t.Logf("Pool stats: %d/%d/%d", stats.AcquiredConns(), stats.IdleConns(), stats.MaxConns())
}
