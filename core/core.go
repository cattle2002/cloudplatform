package core

import (
	"cloudplatform/config"
	"cloudplatform/core/apiServer"
	"cloudplatform/log"
	"context"
	"sync"
)

type Server interface {
	Init() error
	Name() string
	Startup(ctx context.Context) error
	Close() error
}
type Core struct {
	servers map[string]Server
	wg      sync.WaitGroup
}

func New() *Core {
	c := Core{
		servers: make(map[string]Server),
	}
	apiServer := apiServer.New(config.ServerConfig())
	c.servers[apiServer.Name()] = apiServer
	return &c
}
func (c *Core) Init() (err error) {
	defer func() {
		if err == nil {
			return
		}
		for _, server := range c.servers {
			server.Close()
		}
	}()
	for _, server := range c.servers {
		err = server.Init()
		if err != nil {
			log.Error(err)
			return err
		}
	}
	return nil
}
func (c *Core) Run(ctx context.Context) {
	for _, serv := range c.servers {
		c.wg.Add(1)

		go func(serv Server, wg *sync.WaitGroup) {
			serv.Startup(ctx)
			serv.Close()
			wg.Done()
		}(serv, &c.wg)
	}
}
func (c *Core) Close() error {
	for _, serv := range c.servers {
		serv.Close()
	}
	return nil
}

func (c *Core) Wait() {
	c.wg.Wait()
}
