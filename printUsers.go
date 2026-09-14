package main

import (
	"context"
	"fmt"
)

func (cfg *config) printUsers() error {
	users, err := cfg.queries.GetUsers(context.Background())
	if err != nil {
		return err
	}
	for _, u := range users {
		fmt.Println(u)
	}
	return nil
}
