package memory

import (
	"github.com/d-velop/dvelop-app-template-go/domain"
	"time"
)

// MemoryStore keeps data in memory
type MemoryStore struct {
	vacationRequests []domain.VacationRequest
}

func (s *MemoryStore) FindAllVacationRequests() ([]domain.VacationRequest, error) {
	var requests []domain.VacationRequest

	for _, r := range s.vacationRequests {
		requests = append(requests,
			domain.VacationRequest{
				From: r.From,
				To:   r.To,
			})
	}

	return requests, nil
}

func (s *MemoryStore) FindVacationRequests(from time.Time, to time.Time) ([]domain.VacationRequest, error) {
	panic("implement me")
}

// NewStore creates a new MemoryStore
func NewStore() MemoryStore {
	return MemoryStore{}
}
