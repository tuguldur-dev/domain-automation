package controller

import (
	"context"
	"fmt"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/gofiber/fiber/v2"
	"github.com/johandui/domain-automation/domain"
	"github.com/johandui/domain-automation/utils"
)

func Create(c *fiber.Ctx) error {
	defer func() {
		if r := recover(); r != nil {
			c.SendString(fmt.Sprintf("Алдаа гарлаа, %v", r.(error)))
		}
	}()
	redis := utils.InitRedis()
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	type RequestBody struct {
		Image  string `json:"image" `
		Port   string `json:"port"`
		Expose string `json:"expose"`
		Name   string `json:"name"`
	}

	p := new(RequestBody)

	if err := c.BodyParser(p); err != nil {
		return err
	}
	subdomain := fmt.Sprintf("%s.%s", p.Name, os.Getenv("aws.domain"))
	config := &container.Config{
		Image: p.Image,
	}
	hostConfig := &container.HostConfig{
		PortBindings: nat.PortMap{
			nat.Port(fmt.Sprintf("%s/tcp", p.Expose)): []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: p.Port,
				},
			},
		},
	}

	resp, err := cli.ContainerCreate(ctx, config, hostConfig, nil, nil, "")
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}
	if err := domain.CreateDomain(subdomain, p.Port); err != nil {
		panic(err)
	}
	if _, err := utils.RSet(redis, subdomain, fmt.Sprintf("%s:%s", os.Getenv("aws.ip"), p.Port)); err != nil {
		panic(err)
	}

	return c.SendString("Амжилттай үүслээ")
}
