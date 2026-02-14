package controller

import (
	"Auth-Service/internal/domain"
	"Auth-Service/internal/dtos"
	"Auth-Service/internal/http/handlers"
	"Auth-Service/internal/parser"
	"Auth-Service/internal/service"
	"Auth-Service/pkg/logger"
	"Auth-Service/pkg/logger/console"
	"Auth-Service/pkg/obfuscate"
	"encoding/json"
	"net/http"
)

const (
	registerControllerTitle = "Register endpoint: "
)

type registerController struct {
	logger  logger.ILogger
	service service.IAuthService
	parsers parser.IFactory
}

func NewRegisterController(logger logger.ILogger, service service.IAuthService, parsers parser.IFactory) *registerController {
	return &registerController{
		logger:  logger,
		service: service,
		parsers: parsers,
	}
}

func (c *registerController) controller(w http.ResponseWriter, r *http.Request) {
	var req *dtos.RegisterRequest

	ctx := r.Context()
	c.logger.Info(ctx, registerControllerTitle+console.StartKey)

	if r.Method != http.MethodPost {
		c.logger.Error(ctx, registerControllerTitle+console.ErrorKey, "Invalid method")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		c.logger.Error(ctx, registerControllerTitle, console.ErrorKey, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx = console.SetContextWithRegister(ctx, req)
	c.logger.Info(ctx, registerControllerTitle, console.RequestKey, obfuscate.RegisterController(*req))

	Parser, _ := c.parsers.Get(parser.UserDtoToUserDomainParser)
	user, err := Parser.Parser(req)
	if err != nil {
		c.logger.Error(ctx, registerControllerTitle, console.ErrorKey, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	serviceResp, err := c.service.Register(ctx, user.(*domain.User))
	if err != nil {
		handlers.HandleHttpError(w, err)
		return
	}

	resp := c.buildResponse(serviceResp)
	c.logger.Info(ctx, registerControllerTitle+console.ResponseKey, console.ResponseKey, resp)

	w.Header().Set("Content-Type", "application/json; charset=uft-8")
	_ = json.NewEncoder(w).Encode(resp)
	c.logger.Info(ctx, registerControllerTitle+console.EndKey)
	return
}

func (c *registerController) buildResponse(serviceResp *domain.RegisterResult) *dtos.RegisterResponse {
	return &dtos.RegisterResponse{
		Code:    serviceResp.Code,
		Message: serviceResp.Message,
	}
}
