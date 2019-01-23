package memory

import (
	"fmt"
	"github.com/d-velop/dvelop-app-template-go/domain"
)

// MemoryStore keeps data in memory
type MemoryStore struct {
	vacationRequests []domain.VacationRequest
}

func (s *MemoryStore) Update(vr domain.VacationRequest) error {
	for i, r := range s.vacationRequests {
		if r.Id == vr.Id {
			s.vacationRequests[i] = vr
		}
	}
	return fmt.Errorf("vacation request with id '%v' does not exist", vr.Id)
}

func (s *MemoryStore) FindById(id string) (domain.VacationRequest, error) {
	for _, r := range s.vacationRequests {
		if r.Id == id {
			return r, nil
		}
	}
	return domain.VacationRequest{}, fmt.Errorf("vacation request with id '%v' does not exist", id)
}

func (s *MemoryStore) FindAllVacationRequests() ([]domain.VacationRequest, error) {
	var requests []domain.VacationRequest

	for _, r := range s.vacationRequests {
		requests = append(requests, r)
	}

	return requests, nil
}

func (v *MemoryStore) Add(vr domain.VacationRequest) {
	v.vacationRequests = append(v.vacationRequests, vr)
}

// NewStore creates a new MemoryStore
func NewStore() MemoryStore {
	return MemoryStore{}
}
