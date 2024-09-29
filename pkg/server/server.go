package server

import (
	"fmt"
	"log"
	"log/slog"

	_ "github.com/joho/godotenv/autoload"

	"go-groc-store/pkg/customer"
	"go-groc-store/pkg/database"

	"github.com/gofiber/fiber/v2"
)

type Handlers struct {
	customerHandler *customer.CustomerHandler
}

type Server struct {
	port     string
	handlers *Handlers
	logger   *slog.Logger

	dbService database.DbService
	webserver *fiber.App
}

func NewServer(app *fiber.App, logger *slog.Logger, port string, dbService database.DbService) *Server {
	// portInt, _ := strconv.Atoi(port)
	fmt.Println(port)
	customerNewHandler := customer.CustomerNewHandler(logger, dbService)
	customerNewHandler.CreateCustomerTable()
	handlers := &Handlers{
		customerHandler: customerNewHandler,
	}

	// RouteInit(app, handlers)
	return &Server{
		handlers:  handlers,
		port:      port,
		logger:    logger,
		dbService: dbService,
		webserver: app,
	}
	//
}

func (s *Server) Start() {
	routes(s.webserver, s.handlers)
	log.Fatal(s.webserver.Listen(":" + s.port))
}

func routes(webserver *fiber.App, handlers *Handlers) {
	customerApi := webserver.Group("/customers")
	customerApi.Get("/", handlers.customerHandler.GetAll)
	customerApi.Post("/", handlers.customerHandler.CreateCustomer)
	customerApi.Get("/:id", handlers.customerHandler.GetCustomer)
	// customerApi.Delete("/create", handlers.customerHandler.DeleteCustomer)
	// customerApi.Update("/create", handlers.customerHandler.UpdateCustomer)
}
