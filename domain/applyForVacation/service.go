// Package applyForVacation provides the apply for vacation use-case
package applyForVacation

import (
	"github.com/d-velop/dvelop-app-template-go/domain"
)

// Service implements the apply for vacation use-case
type Service interface {
	ApplyForVacation()
}

type service struct {
	vacReqRepo domain.VacationRequestRepository
}

// NewService creates an applyForVacation service with the necessary dependencies
func NewService(r domain.VacationRequestRepository) Service {
	return &service{r}
}

// applies for vacation
func (s *service) ApplyForVacation() {
}
