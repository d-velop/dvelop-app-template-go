// Package domain contains the heart of the domain model.
package domain

import (
	"context"
	"time"
)

type VacationType int

const (
	ANNUAL_VACATION VacationType = iota
	SPECIAL_VACATION
	COMPENSATORY_VACATION
)

var VacationTypes = [...]string{
	"annual",
	"special",
	"compensatory",
}

func (v VacationType) String() string {
	return VacationTypes[v]
}

type VacationRequestState int

const (
	REQUEST_NEW VacationRequestState = iota
	REQUEST_ACCEPTED
	REQUEST_REJECTED
	REQUEST_CANCELLED
)

var VacationRequestStates = [...]string{
	"new",
	"accepted",
	"rejected",
	"cancelled",
}

func (s VacationRequestState) String() string {
	return VacationRequestStates[s]
}

type VacationRequest struct {
	Id      string
	From    time.Time
	To      time.Time
	Type    VacationType
	State   VacationRequestState
	Comment string
}

// Repository provides access to vacation request repository
type VacationRequestRepository interface {
	FindAllVacationRequests(context.Context) ([]VacationRequest, error)
	FindById(context.Context, string) (VacationRequest, error)
	Add(context.Context, VacationRequest) error
	Update(context.Context, VacationRequest) error
}
