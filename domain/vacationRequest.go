// Package domain contains the heart of the domain model.
package domain

import "time"

type VacationRequest struct {
	From time.Time `json:"from"`
	To   time.Time `json:"to"`
}

// Repository provides access to vacation request repository
type VacationRequestRepository interface {
	FindAllVacationRequests() ([]VacationRequest, error)
	FindVacationRequests(from time.Time, to time.Time) ([]VacationRequest, error)
}
