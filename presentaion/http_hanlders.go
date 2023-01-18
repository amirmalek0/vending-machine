package presentaion

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"vending-mechine/domain"
)

type HttpHandlers struct {
	domain *domain.VendingMachineDomain
}

func NewHttpHandlers(domain *domain.VendingMachineDomain) *HttpHandlers {
	return &HttpHandlers{domain}
}

type HttpResponse struct {
	Error interface{} `json:"error"`
	Data  interface{} `json:"data"`
}

func (h *HttpHandlers) machineList(c echo.Context) error {
	list := h.domain.MachineList()
	return c.JSON(http.StatusOK, HttpResponse{Data: list})
}

func (h *HttpHandlers) buyCaffe(c echo.Context) error {
	type Request struct {
		MachineName string `json:"machine_name"`
		Coin        int32  `json:"coin"`
	}
	req := &Request{}
	err := c.Bind(req)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, HttpResponse{Error: "invalid input"})
	}
	message, err := h.domain.BuyCoffee(req.MachineName, req.Coin)
	if err != nil {
		return c.JSON(http.StatusBadRequest, HttpResponse{Error: err.Error()})
	}
	return c.JSON(http.StatusOK, HttpResponse{Data: message})
}

func (h *HttpHandlers) buyCoca(c echo.Context) error {
	type Request struct {
		MachineName string `json:"machine_name"`
		Coin        int32  `json:"coin"`
	}
	req := &Request{}
	err := c.Bind(req)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, HttpResponse{Error: "invalid input"})
	}
	message, err := h.domain.BuyCoca(req.MachineName, req.Coin)
	if err != nil {
		return c.JSON(http.StatusBadRequest, HttpResponse{Error: err.Error()})
	}
	return c.JSON(http.StatusOK, HttpResponse{Data: message})
}
