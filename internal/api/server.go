package api

import (
	"fmt"
	"log"
	"os"

	"github.com/darchlabs/backoffice/internal/api/context"
	v1 "github.com/darchlabs/backoffice/internal/api/v1"
	"github.com/darchlabs/backoffice/internal/api/v1/user"
	"github.com/darchlabs/backoffice/internal/application"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type ServerConfig struct {
	Port string
	App  *application.App
}

type Server struct {
	server *fiber.App
	app    *application.App
	port   string
}

func NewServer(config *ServerConfig) *Server {
	server := fiber.New()
	server.Use(logger.New())
	server.Use(logger.New(logger.Config{
		Format:     "[${ip}]:${port} ${status} - ${method} ${path}\n",
		TimeFormat: "2006-01-02 15:04:05",
		Output:     os.Stdout,
	}))

	return &Server{
		server: server,
		app:    config.App,
		port:   config.Port,
	}
}

func (s *Server) Start(app *application.App) error {
	go func() {
		ctx := context.New(&context.Config{
			Server:   s.server,
			App:      s.app,
			SqlStore: app.SqlStore,
		})
		log.Println("[API] context", ctx)
		// route endpoints
		v1.HealthRoute(ctx)
		user.Route("/api/v1/users", ctx)

		// sever listen
		log.Println("[API] server running")
		err := s.server.Listen(fmt.Sprintf(":%s", s.port))
		if err != nil {
			panic(err)
		}
	}()

	return nil
}
