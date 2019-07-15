package memory

import (
	"context"
	"errors"
	"fmt"
	"git.d-velop.de/dvelopcloud/shop-middleware/domain"
	"github.com/d-velop/dvelop-sdk-go/tenant"
)

// store keeps data in memory
type store struct {
	// todo: maps and slices are not thread-safe cf. https://blog.golang.org/go-maps-in-action .
	//  so a mutex or something similar MUST be used as soon as the store is accessed concurrently
	//  which is NOT the case for local testing and AWS lambda deployments cf. 'Concurrency' in https://docs.aws.amazon.com/lambda/latest/dg/programming-model-v2.html
	vacationRequestsForTenant map[string][]domain.VacationRequest
}

func (s *store) Update(ctx context.Context, vr domain.VacationRequest) error {
	tenantId, err := tenant.IdFromCtx(ctx)
	if err != nil {
		return fmt.Errorf("can't update vacation request with id '%v' because context doesn't contain a tenant id", vr.Id)
	}
	vrs := s.vacationRequestsForTenant[tenantId]
	for i, r := range vrs {
		if r.Id == vr.Id {
			vrs[i] = vr
		}
	}
	return fmt.Errorf("vacation request with id '%v' does not exist", vr.Id)
}

func (s *store) FindById(ctx context.Context, id string) (domain.VacationRequest, error) {
	tenantId, err := tenant.IdFromCtx(ctx)
	if err != nil {
		return domain.VacationRequest{}, fmt.Errorf("can't search for vacation request with id '%v' because context doesn't contain a tenant id", id)
	}
	vrs := s.vacationRequestsForTenant[tenantId]
	for _, r := range vrs {
		if r.Id == id {
			return r, nil
		}
	}
	return domain.VacationRequest{}, fmt.Errorf("vacation request with id '%v' does not exist", id)
}

func (s *store) FindAllVacationRequests(ctx context.Context) ([]domain.VacationRequest, error) {
	tenantId, err := tenant.IdFromCtx(ctx)
	if err != nil {
		return nil, errors.New("can't get vacation requests because context doesn't contain a tenant id")
	}
	vrs := s.vacationRequestsForTenant[tenantId]
	var result []domain.VacationRequest
	for _, r := range vrs {
		result = append(result, r)
	}
	return result, nil
}

func (s *store) Add(ctx context.Context, vr domain.VacationRequest) error {
	tenantId, err := tenant.IdFromCtx(ctx)
	if err != nil {
		return fmt.Errorf("can't add vacation request '%v' because context doesn't contain a tenant id", vr)
	}
	if s.vacationRequestsForTenant == nil {
		s.vacationRequestsForTenant = make(map[string][]domain.VacationRequest)
	}
	s.vacationRequestsForTenant[tenantId] = append(s.vacationRequestsForTenant[tenantId], vr)
	return nil
}

// NewStore creates a new in memory store
func NewStore() domain.VacationRequestRepository {
	return &store{}
}
