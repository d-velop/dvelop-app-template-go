package applyForVacation_test

import (
	"github.com/d-velop/dvelop-app-template-go/domain"
	"github.com/d-velop/dvelop-app-template-go/domain/applyForVacation"
	"github.com/d-velop/dvelop-app-template-go/domain/plugins/storage/memory"
	"testing"
	"time"
)

func TestNewAndValidRequest_NewRequestIsStored(t *testing.T) {
	vacRequestRepo := memory.NewStore()
	s := applyForVacation.NewService(&vacRequestRepo)
	newVR := domain.VacationRequest{
		From:    time.Date(2018, 10, 10, 0, 0, 0, 0, time.UTC),
		To:      time.Date(2018, 10, 11, 0, 0, 0, 0, time.UTC),
		Type:    domain.ANNUAL_VACATION,
		Comment: "I realy need a day of",
	}

	s.Execute(newVR)

	vrs, e := vacRequestRepo.FindAllVacationRequests()
	if e != nil {
		t.Error(e)
	}
	var foundVR domain.VacationRequest
	found := false
	for _, vr := range vrs {
		if vr.From == newVR.From && vr.To == newVR.To && vr.Type == newVR.Type && vr.Comment == newVR.Comment {
			foundVR = vr
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Stored vacation requests %v didn't contain new request %v", vrs, newVR)
	}
	if foundVR.State != domain.REQUEST_NEW {
		t.Errorf("Stored vacation requests %v has wrong state: got %v wanted %v", foundVR, foundVR.State, domain.REQUEST_NEW)
	}
}
