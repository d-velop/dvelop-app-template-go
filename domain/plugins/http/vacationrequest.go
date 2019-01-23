package http

// todo File needs code cleanup/refactoring to reduce duplication

import (
	"encoding/json"
	"fmt"
	"github.com/d-velop/dvelop-app-template-go/domain"
	"github.com/d-velop/dvelop-app-template-go/domain/acceptVacationRequest"
	"github.com/d-velop/dvelop-app-template-go/domain/applyForVacation"
	"github.com/d-velop/dvelop-app-template-go/domain/cancelVacation"
	"github.com/d-velop/dvelop-app-template-go/domain/plugins/conf"
	"github.com/d-velop/dvelop-app-template-go/domain/rejectVacationRequest"
	"github.com/d-velop/dvelop-sdk-go/contentnegotiation/mediatype"
	"github.com/d-velop/dvelop-sdk-go/log"
	"io"
	"net/http"
	"strings"
	"time"
)

type vacationRequestHandler struct {
	assetBasePath         string
	renderhtml            func(w io.Writer, data interface{}, templatename string) error
	storage               domain.VacationRequestRepository
	applyForVacation      applyForVacation.Service
	cancelVacation        cancelVacation.Service
	rejectVacationRequest rejectVacationRequest.Service
	acceptVacationRequest acceptVacationRequest.Service
}

func NewVacationRequestHandler(assetBasePath string,
	renderhtml func(w io.Writer, data interface{}, templatename string) error,
	storage domain.VacationRequestRepository,
	applyForVacation applyForVacation.Service,
	cancelVacation cancelVacation.Service,
	rejectVacationRequest rejectVacationRequest.Service,
	acceptVacationRequest acceptVacationRequest.Service) *vacationRequestHandler {
	return &vacationRequestHandler{
		assetBasePath:         assetBasePath,
		renderhtml:            renderhtml,
		storage:               storage,
		applyForVacation:      applyForVacation,
		cancelVacation:        cancelVacation,
		rejectVacationRequest: rejectVacationRequest,
		acceptVacationRequest: acceptVacationRequest,
	}
}

func (h *vacationRequestHandler) HandleNewForm() http.Handler {
	fn := func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodGet:
			negotiatedType, err := mediatype.Negotiate(req.Header.Get("Accept"), []string{"text/html"})
			if err != nil {
				log.Error(req.Context(), err)
				http.Error(w, http.StatusText(http.StatusNotAcceptable), http.StatusNotAcceptable)
				return
			}
			w.Header().Set("content-type", negotiatedType.String()+";charset=utf-8")
			d := &VacationRequestHtmlDto{Title: "Apply for vacation", Mode: "new"}
			d.AssetBasePath = h.assetBasePath
			err = h.renderhtml(w, d, "vacationrequest.html")
			if err != nil {
				log.Error(req.Context(), err)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return

			}
		default:
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		}
	}
	return http.HandlerFunc(fn)
}

func (h *vacationRequestHandler) Handle(pattern string) http.Handler {
	fn := func(w http.ResponseWriter, req *http.Request) {
		subresource := req.URL.Path[len(pattern):]
		if subresource == "" {
			h.handleList().ServeHTTP(w, req)
		} else if strings.Contains(subresource, "/") {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		} else {
			h.handleSingle(subresource).ServeHTTP(w, req)
		}
	}
	return http.HandlerFunc(fn)
}

func (h *vacationRequestHandler) handleList() http.Handler {
	fn := func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodGet:
			negotiatedType, err := mediatype.Negotiate(req.Header.Get("Accept"), []string{"text/html"})
			if err != nil {
				log.Error(req.Context(), err)
				http.Error(w, http.StatusText(http.StatusNotAcceptable), http.StatusNotAcceptable)
				return
			}
			w.Header().Set("content-type", negotiatedType.String()+";charset=utf-8")

			vr, err := h.storage.FindAllVacationRequests()
			dto := toListHtmlDto(vr)
			dto.Title = "Vacationrequests"
			dto.AssetBasePath = h.assetBasePath
			fmt.Println(len(dto.Requests))
			err = h.renderhtml(w, dto, "vacationrequests.html")
			if err != nil {
				log.Error(req.Context(), err)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return

			}
		case http.MethodPost:
			if req.Header.Get("content-type") != "application/json" {
				http.Error(w, http.StatusText(http.StatusUnsupportedMediaType), http.StatusUnsupportedMediaType)
				return
			}
			var dto VacationRequestDto
			err := json.NewDecoder(req.Body).Decode(&dto)
			if err != nil {
				log.Error(req.Context(), err)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
			vr, err := dto.ToDomain()
			if err != nil {
				log.Error(req.Context(), err)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
			id := h.applyForVacation.Execute(vr)
			w.Header().Set("Location", conf.BasePath+"/vacationrequest/"+id)
			w.WriteHeader(http.StatusCreated)
		default:
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		}
	}
	return http.HandlerFunc(fn)
}

func (h *vacationRequestHandler) handleSingle(subresource string) http.Handler {
	fn := func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodGet:
			negotiatedType, err := mediatype.Negotiate(req.Header.Get("Accept"), []string{"text/html"})
			if err != nil {
				log.Error(req.Context(), err)
				http.Error(w, http.StatusText(http.StatusNotAcceptable), http.StatusNotAcceptable)
				return
			}
			w.Header().Set("content-type", negotiatedType.String()+";charset=utf-8")
			vr, err := h.storage.FindById(subresource)
			dto := toHtmlDto(vr)
			dto.Title = fmt.Sprintf("Vacation Request %v", vr.Id)
			dto.State = "show"
			dto.AssetBasePath = h.assetBasePath
			err = h.renderhtml(w, dto, "vacationrequest.html")
			if err != nil {
				log.Error(req.Context(), err)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return

			}
		case http.MethodPatch:
			// cf. https://tools.ietf.org/html/rfc7386
			// cf. https://tools.ietf.org/html/rfc5789
			// todo: Needs improvement concerning error handling an compliance to rfc5789
			//  as soon as this is part of the public API. Right now it's only used by the html frontend
			// 	so we know the structure of the http request
			if req.Header.Get("content-type") != "application/merge-patch+json" {
				http.Error(w, http.StatusText(http.StatusUnsupportedMediaType), http.StatusUnsupportedMediaType)
				return
			}
			var dto VacationRequestPatchDto
			err := json.NewDecoder(req.Body).Decode(&dto)
			if err != nil {
				log.Error(req.Context(), err)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
			switch dto.State {
			case domain.VacationRequestStates[domain.REQUEST_ACCEPTED]:
				h.acceptVacationRequest.Execute(subresource)
			case domain.VacationRequestStates[domain.REQUEST_REJECTED]:
				h.rejectVacationRequest.Execute(subresource)
			case domain.VacationRequestStates[domain.REQUEST_CANCELLED]:
				h.cancelVacation.Execute(subresource)
			}
			w.Header().Set("Content-Location", conf.BasePath+"/vacationrequest/"+subresource)
			// todo E-Tag cf. https://tools.ietf.org/html/rfc5789#section-2.1
			w.WriteHeader(http.StatusNoContent)
		default:
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		}
	}
	return http.HandlerFunc(fn)
}

type VacationRequestListHtmlDto struct {
	BaseHtmlDto
	Title    string
	Requests []VacationRequestHtmlDto
}

func toListHtmlDto(vrs []domain.VacationRequest) VacationRequestListHtmlDto {
	var dto VacationRequestListHtmlDto
	for _, vr := range vrs {
		dto.Requests = append(dto.Requests, toHtmlDto(vr))
	}
	return dto
}

type VacationRequestDto struct {
	From    time.Time `json:"from"`
	To      time.Time `json:"to"`
	Type    string    `json:"type"`
	State   string    `json:"state"`
	Comment string    `json:"comment"`
	Id      string    `json:"id"`
}

type VacationRequestPatchDto struct {
	State string `json:"state"`
}

type VacationRequestHtmlDto struct {
	BaseHtmlDto
	Title   string
	Mode    string
	From    string
	To      string
	Type    string
	State   string
	Comment string
	Id      string
}

func toHtmlDto(vr domain.VacationRequest) VacationRequestHtmlDto {
	const dateFormat = "2006-01-02"
	return VacationRequestHtmlDto{
		From:    vr.From.Format(dateFormat),
		To:      vr.To.Format(dateFormat),
		Type:    domain.VacationTypes[vr.Type],
		State:   domain.VacationRequestStates[vr.State],
		Comment: vr.Comment,
		Id:      vr.Id,
	}
}

func (dto *VacationRequestDto) ToDomain() (domain.VacationRequest, error) {
	vr := domain.VacationRequest{
		From:    dto.From,
		To:      dto.To,
		Comment: dto.Comment,
	}
	value, ok := map[string]domain.VacationType{
		"Annual":       domain.ANNUAL_VACATION,
		"Special":      domain.SPECIAL_VACATION,
		"Compensatory": domain.COMPENSATORY_VACATION,
	}[dto.Type]
	if !ok {
		return domain.VacationRequest{}, fmt.Errorf("%v is no valid enum type", dto.Type)
	}
	vr.Type = value
	return vr, nil
}
