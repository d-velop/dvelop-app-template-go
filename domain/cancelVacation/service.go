// Package cancelVacation provides the cancel vacation use-case
package cancelVacation

import (
	"github.com/d-velop/dvelop-app-template-go/domain"
)

// Service implements the cancel vacation use-case
// todo extract common interface
type Service interface {
	Execute(vacationRequestId string)
}

type service struct {
	vacReqRepo domain.VacationRequestRepository
}

// NewService creates a cancelVacation service with the necessary dependencies
func NewService(r domain.VacationRequestRepository) Service {
	return &service{r}
}

// Executes the service
func (s *service) Execute(vacationRequestId string) {
	// business logic like validation goes here!
	// todo error handling
	vr, _ := s.vacReqRepo.FindById(vacationRequestId)
	vr.State = domain.REQUEST_CANCELLED
	s.vacReqRepo.Update(vr)
}