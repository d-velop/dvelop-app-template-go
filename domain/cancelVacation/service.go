// Package cancelVacation provides the cancel vacation use-case
package cancelVacation

import (
	"context"
	"git.d-velop.de/dvelopcloud/approuter-echoapp/domain"
)

// Service implements the cancel vacation use-case
// todo extract common interface
type Service interface {
	Execute(ctx context.Context, vacationRequestId string)
}

type service struct {
	vacReqRepo domain.VacationRequestRepository
}

// NewService creates a cancelVacation service with the necessary dependencies
func NewService(r domain.VacationRequestRepository) Service {
	return &service{r}
}

// Executes the service
func (s *service) Execute(ctx context.Context, vacationRequestId string) {
	// business logic like validation goes here!
	// todo error handling
	vr, _ := s.vacReqRepo.FindById(ctx, vacationRequestId)
	vr.State = domain.REQUEST_CANCELLED
	s.vacReqRepo.Update(ctx, vr)
}
