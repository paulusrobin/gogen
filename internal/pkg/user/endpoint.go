package user

import (
	"github.com/go-kit/kit/endpoint"
)

// Endpoint interface of user package.
type Endpoint interface {
	Create() endpoint.Endpoint
	GetByID() endpoint.Endpoint
}
