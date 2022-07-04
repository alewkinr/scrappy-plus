package main

import (
	"context"
)

// Handler – обработчик для запросов Yandex.Cloud Functions
//nolint:unparam,deadcode
func Handler(ctx context.Context) (struct{}, error) {
	if runErr := run(ctx); runErr != nil {
		return struct{}{}, runErr
	}

	return struct{}{}, nil
}
