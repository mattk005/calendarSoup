package config

import (
	"context"
	"fmt"
	"strings"
)

func (cfg *Config) PrintUsers() (string, error) {
	users, err := cfg.Queries.GetUsers(context.Background())
	if err != nil {
		return "", err
	}
	mySlice := make([]string, len(users))
	for _, u := range users {
		mySlice = append(mySlice, fmt.Sprint(u))
	}
	return strings.Join(mySlice, " "), nil
}
