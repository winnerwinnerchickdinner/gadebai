package auth

import (
	"auth/pkg/handlers"
)

func Start() {
	handlers.RegistAll()
}
