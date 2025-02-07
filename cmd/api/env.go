package main

import (
	"os"
)

func getString(key, fallback string) string {
	value := os.Getenv(key)

	if value == "" {
		return fallback
	}
	return value
}

func getInt(key string, fallback int) int {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return fallback
}
