// Package acceptVacationRequest provides the accept vacation request use-case
package acceptVacationRequest

import (
	"context"
	"git.d-velop.de/dvelopcloud/shop-middleware/domain"
)

// Service implements the accept vacation request use-case
// todo extract common interface
type Service interface {
	Execute(ctx context.Context, vacationRequestId string)
}

type service struct {
	vacReqRepo domain.VacationRequestRepository
}

// NewService creates an acceptVacationRequest service with the necessary dependencies
func NewService(r domain.VacationRequestRepository) Service {
	return &service{r}
}

// Executes the service
func (s *service) Execute(ctx context.Context, vacationRequestId string) {
	// business logic like validation goes here!
	// todo error handling
	vr, _ := s.vacReqRepo.FindById(ctx, vacationRequestId)
	vr.State = domain.REQUEST_ACCEPTED
	s.vacReqRepo.Update(ctx, vr)
}
