package main

import (
	"sync"

	flagsmith "github.com/Flagsmith/flagsmith-go-client"
)

var (
	client *flagsmith.Client
	once   sync.Once
)

func NewFlagClient() *flagsmith.Client {
	once.Do(func() {
		client = flagsmith.DefaultClient("ser.RKeb3BPYV2YERESrHxK92L")
	})
	return client
}
