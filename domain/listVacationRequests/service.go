// Package listVacationRequests provides the list existing vacation requests use-case
package listVacationRequests

import (
	"github.com/d-velop/dvelop-app-template-go/domain"
)

// Service implements the list vacation requests use-case
type Service interface {
	ListVacationRequests() []domain.VacationRequest
}

type service struct {
	vacReqRepo domain.VacationRequestRepository
}

// NewService creates a list vacation request service with the necessary dependencies
func NewService(r domain.VacationRequestRepository) Service {
	return &service{r}
}

// list existing vacation requests
func (s *service) ListVacationRequests() []domain.VacationRequest {
	vrs, _ := s.vacReqRepo.FindAllVacationRequests()
	return vrs
}
