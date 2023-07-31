package src

import (
	"fmt"
	"log"
	"net/http"
	"restapi-altera/config"
	"restapi-altera/src/delivery"
	"restapi-altera/src/helper"
	"restapi-altera/src/helper/validator"
	"restapi-altera/src/model"
	"restapi-altera/src/repository"
	"restapi-altera/src/usecase"
	"time"

	validatorEngine "github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server interface {
	Run()
}

type server struct {
	httpServer *echo.Echo
	config     config.Config
}

func InitServer(config config.Config) Server {

	e := echo.New()
	e.Validator = &validator.GoPlaygroundValidator{
		Validator: validatorEngine.New(),
	}
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	return &server{
		config:     config,
		httpServer: e,
	}

}

func (s *server) Run() {

	s.httpServer.GET("/", func(c echo.Context) error {
		return helper.WriteResponse(c, http.StatusOK, "hallo dari echo", nil)
	})

	s.httpServer.POST("/rahasia", func(c echo.Context) error {

		temp := &model.Shoes{
			ShoesId: "9fb0af00-18f2-4025-9132-7f19ee25b162",
			Name:    "Nike SB Dunk Low Pro Prod",
			Size:    40,
			AddAt:   time.Now().Format(time.RFC822),
		}

		err := s.config.Database().Create(temp).Error
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, "succes")

	})

	repo := repository.NewRepository(s.config)
	usecase := usecase.NewUsecase(repo)
	delivery := delivery.NewDelivery(usecase)
	delivery.Auth(s.httpServer.Group("auth"))
	delivery.User(s.httpServer.Group("user"))

	if err := s.httpServer.Start(fmt.Sprintf(":%d", s.config.Port())); err != nil {
		log.Panic(err)
	}
}
