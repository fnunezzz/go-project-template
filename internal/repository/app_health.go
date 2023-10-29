package repository

import (
	"context"
	"fmt"
	"os"
	"time"
)

type AppHealthRepository interface {
	GetCurrentTimestamp() (time.Time, error)
}

type appHealthRepository struct{}

func (u *appHealthRepository) GetCurrentTimestamp() (time.Time, error) {
	var now time.Time
	err := DB.QueryRow(context.Background(), "select current_timestamp").Scan(&now)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return time.Now(), err
	}
	return now, nil
}