package rpc

import (
	"context"
	"fmt"
	"github.com/bcessa/sample-twirp/proto"
	"github.com/chzyer/readline"
	"github.com/golang/protobuf/ptypes/empty"
	"log"
)

type ClientConsole struct {
	client sample.BusinessCase
	rl     *readline.Instance
}

func NewConsole(c sample.BusinessCase, prompt string) *ClientConsole {
	rl, _ := readline.New(prompt)
	return &ClientConsole{
		client: c,
		rl:     rl,
	}
}

func (c *ClientConsole) Start() error {
	c.usage()
	for {
		line, err := c.rl.Readline()
		if err != nil {
			return err
		}
		switch line {
		case "p":
			pong, err := c.client.Ping(context.TODO(), &empty.Empty{})
			if err != nil {
				log.Println("error", err)
			}
			log.Printf("%+v", pong)
		case "q":
			fmt.Println("closing console")
			return nil
		case "h":
			c.usage()
		default:
			fmt.Println("invalid command")
		}
	}
	return nil
}

func (c *ClientConsole) Close() {
	c.rl.Close()
}

func (c *ClientConsole) usage() {
	fmt.Println("p = ping")
	fmt.Println("h = help")
	fmt.Println("q = quit")
}
