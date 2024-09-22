package request

import (
	"errors"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type (
	contextWrapperService interface {
		Bind(data any) error
	}

	contextWrapper struct {
		Context   *gin.Context
		validator *validator.Validate
	}
)

func ContextWrapper(ctx *gin.Context) contextWrapperService {
	v := validator.New()
	err := v.RegisterValidation("is_myanime_status", myAnimeListStatusValidator)
	if err != nil {
		log.Println("Error registering custom validation :", err.Error())
	}
	return &contextWrapper{
		Context:   ctx,
		validator: v,
	}
}

// Bind implements contextWrapperService.
func (c *contextWrapper) Bind(data any) error {
	fmt.Printf("c.Context.Request.Body: %v\n", c.Context.Request.Body)
	if err := c.Context.Bind(data); err != nil {
		log.Printf("Error: Bind data failed: %s", err.Error())
		return errors.New("errors: bad requset")

	}

	if err := c.validator.Struct(data); err != nil {
		log.Printf("Error: Validate data failed: %s", err.Error())
		return errors.New("errors: validate data failed")
	}
	return nil
}

func myAnimeListStatusValidator(fl validator.FieldLevel) bool {
	log.Println("myAnimeListStatusValidator()")
	status := fl.Field().String()
	log.Printf("status: %v\n", status)
	switch status {
	case "plan-to-watch", "watching", "completed", "on-hold", "dropped":
		return true
	default:
		return false
	}
}
