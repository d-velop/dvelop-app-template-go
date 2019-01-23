// Package applyForVacation provides the apply for vacation use-case
package applyForVacation

import (
	"github.com/d-velop/dvelop-app-template-go/domain"
	"github.com/satori/go.uuid"
)

// Service implements the apply for vacation use-case
type Service interface {
	Execute(vr domain.VacationRequest) string
}

type service struct {
	vacReqRepo domain.VacationRequestRepository
}

// NewService creates an applyForVacation service with the necessary dependencies
func NewService(r domain.VacationRequestRepository) Service {
	return &service{r}
}

// Executes the service
func (s *service) Execute(vr domain.VacationRequest) string {
	// business logic like validation goes here!
	vr.Id = uuid.Must(uuid.NewV4()).String()
	s.vacReqRepo.Add(vr)
	return vr.Id
}
