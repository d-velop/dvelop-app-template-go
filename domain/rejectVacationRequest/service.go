// Package rejectVacationRequest provides the accept vacation request use-case
package rejectVacationRequest

import (
	"github.com/d-velop/dvelop-app-template-go/domain"
)

// Service implements the reject vacation request use-case
// todo extract common interface
type Service interface {
	Execute(vacationRequestId string)
}

type service struct {
	vacReqRepo domain.VacationRequestRepository
}

// NewService creates an rejectVacationRequest service with the necessary dependencies
func NewService(r domain.VacationRequestRepository) Service {
	return &service{r}
}

// Executes the service
func (s *service) Execute(vacationRequestId string) {
	// business logic like validation goes here!
	// todo error handling
	vr, _ := s.vacReqRepo.FindById(vacationRequestId)
	vr.State = domain.REQUEST_REJECTED
	s.vacReqRepo.Update(vr)
}
