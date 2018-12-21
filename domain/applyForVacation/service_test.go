package applyForVacation_test

import (
	"errors"
	"github.com/d-velop/dvelop-app-template-go/domain"
	"github.com/d-velop/dvelop-app-template-go/domain/applyForVacation"
	"github.com/d-velop/dvelop-app-template-go/domain/mock"
	"testing"
	"time"
)

func TestOneVacationDay_ReducesNumberOfVacationDaysByOne(t *testing.T) {
	var vacRequestRepo mock.VacationRequestRepository

	vacRequestRepo.FindVacationRequestsFn = func(from time.Time, to time.Time) ([]domain.VacationRequest, error) {
		return nil, errors.New("NotImplemented")
	}

	s := applyForVacation.NewService(&vacRequestRepo)
	s.ApplyForVacation()
}
