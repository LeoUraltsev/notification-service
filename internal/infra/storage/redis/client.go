package redis

import r "github.com/redis/go-redis/v9"

type Client struct {
	client *r.Client
}

func NewClient() *Client {
	client := r.NewClient(&r.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return &Client{
		client: client,
	}
}
