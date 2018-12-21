package listVacationRequests_test

import (
	"github.com/d-velop/dvelop-app-template-go/domain"
	"github.com/d-velop/dvelop-app-template-go/domain/listVacationRequests"
	"github.com/d-velop/dvelop-app-template-go/domain/mock"
	"testing"
	"time"
)

func TestNoVacationRequestExists_ReturnsNoVacationRequests(t *testing.T) {
	var vacRequestRepo mock.VacationRequestRepository
	vacRequestRepo.FindAllVacationRequestsFn = func() ([]domain.VacationRequest, error) {
		return nil, nil
	}

	s := listVacationRequests.NewService(&vacRequestRepo)
	vr := s.ListVacationRequests()

	if vr != nil {
		t.Errorf("got %v expected nil", vr)
	}
}

func TestOneVacationRequestExists_ReturnsExistingVacationRequest(t *testing.T) {
	var vacRequestRepo mock.VacationRequestRepository
	vacRequestRepo.FindAllVacationRequestsFn = func() ([]domain.VacationRequest, error) {
		return []domain.VacationRequest{{From: time.Now(), To: time.Now().Add(time.Hour * 24)}}, nil
	}

	s := listVacationRequests.NewService(&vacRequestRepo)
	vrs := s.ListVacationRequests()

	if vrs == nil || len(vrs) != 1 {
		t.Error("got no vacationrequest expected 1")
	}
}
