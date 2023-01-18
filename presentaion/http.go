package presentaion

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"vending-mechine/domain"
)

type HttpServer struct {
	host     string
	port     string
	server   *echo.Echo
	handlers *HttpHandlers
}

func NewHttpServer(host string, port string, domain *domain.VendingMachineDomain) *HttpServer {
	server := echo.New()
	handlers := NewHttpHandlers(domain)
	instance := &HttpServer{
		host:     host,
		port:     port,
		server:   server,
		handlers: handlers,
	}
	instance.registerRoutes()
	return instance
}

func (h *HttpServer) Start() {
	url := fmt.Sprintf("%s:%s", h.host, h.port)
	h.server.Logger.Fatal(h.server.Start(url))
}

func (h *HttpServer) registerRoutes() {
	h.server.GET("/", h.handlers.machineList)
	h.server.GET("/coffee", h.handlers.buyCaffe)
	h.server.GET("/coca", h.handlers.buyCoca)
}
