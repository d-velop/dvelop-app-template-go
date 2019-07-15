package mock

import (
	"git.d-velop.de/dvelopcloud/shop-middleware/domain"
	"time"
)

// VacationRequestRepository is a mock vacation request repository.
type VacationRequestRepository struct {
	FindVacationRequestsFn      func(from time.Time, to time.Time) ([]domain.VacationRequest, error)
	FindVacationRequestsInvoked bool

	FindAllVacationRequestsFn      func() ([]domain.VacationRequest, error)
	FindAllVacationRequestsInvoked bool
}

func (r *VacationRequestRepository) FindAllVacationRequests() ([]domain.VacationRequest, error) {
	r.FindAllVacationRequestsInvoked = true
	return r.FindAllVacationRequestsFn()
}

func (r *VacationRequestRepository) FindVacationRequests(from time.Time, to time.Time) ([]domain.VacationRequest, error) {
	r.FindVacationRequestsInvoked = true
	return r.FindVacationRequestsFn(from, to)
}
