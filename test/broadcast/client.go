package main

type Client struct {
	Name     string
	Messages chan string
}

func NewClient(name string) *Client {
	return &Client{
		Name:     name,
		Messages: make(chan string),
	}
}
