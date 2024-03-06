package matchinghandler

import (
	"github.com/abbasfisal/game-app/service/authservice"
	"github.com/abbasfisal/game-app/service/matchingservice"
	"github.com/abbasfisal/game-app/validtor/matchingvalidator"
)

type Handler struct {
	authConfig        authservice.Config
	authSvc           authservice.Service
	matchingSvc       matchingservice.Service
	matchingValidator matchingvalidator.Validator
}

func New(authConfig authservice.Config, authSvc authservice.Service, matchingSvc matchingservice.Service, matchingValidator matchingvalidator.Validator) Handler {

	return Handler{
		authConfig:        authConfig,
		authSvc:           authSvc,
		matchingSvc:       matchingSvc,
		matchingValidator: matchingValidator,
	}
}
