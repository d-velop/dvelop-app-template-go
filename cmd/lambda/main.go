package main

import (
	"context"
	"fmt"
	"github.com/d-velop/dvelop-app-template-go/domain/acceptVacationRequest"
	"github.com/d-velop/dvelop-app-template-go/domain/applyForVacation"
	"github.com/d-velop/dvelop-app-template-go/domain/cancelVacation"
	"github.com/d-velop/dvelop-app-template-go/domain/plugins/conf"
	"github.com/d-velop/dvelop-app-template-go/domain/plugins/gui/templates"
	"github.com/d-velop/dvelop-app-template-go/domain/plugins/http"
	"github.com/d-velop/dvelop-app-template-go/domain/plugins/storage/memory"
	"github.com/d-velop/dvelop-app-template-go/domain/rejectVacationRequest"
	"github.com/d-velop/dvelop-sdk-go/idp"
	"github.com/d-velop/dvelop-sdk-go/lambda"
	"github.com/d-velop/dvelop-sdk-go/log"
	"github.com/d-velop/dvelop-sdk-go/requestid"
	"github.com/d-velop/dvelop-sdk-go/tenant"
	"strings"
)

func main() {
	setupLogging()

	// wire dependencies
	storage := memory.NewStore() // todo Use dynamodb
	applyForVacationService := applyForVacation.NewService(storage)
	cancelVacationService := cancelVacation.NewService(storage)
	rejectVacationRequestService := rejectVacationRequest.NewService(storage)
	acceptVacationRequestService := acceptVacationRequest.NewService(storage)
	vacationRequestHandler := http.NewVacationRequestHandler(
		conf.AssetBasePath(),
		templates.Render,
		storage,
		applyForVacationService,
		cancelVacationService,
		rejectVacationRequestService,
		acceptVacationRequestService)

	logError := func(ctx context.Context, logmessage string) { log.Error(ctx, logmessage) }
	logInfo := func(ctx context.Context, logmessage string) { log.Info(ctx, logmessage) }

	resources := []http.Resource{
		{Pattern: conf.BasePath + "/", Handler: http.HandleRoot(conf.AssetBasePath(), templates.Render, conf.Version())},
		{Pattern: conf.BasePath + "/vacationrequest", Handler: vacationRequestHandler.HandleNewForm()},
		{Pattern: conf.BasePath + "/vacationrequest/", Handler: vacationRequestHandler.Handle(conf.BasePath + "/vacationrequest/")},
		{Pattern: conf.BasePath + "/features", Handler: http.HandleFeatures()},
		{Pattern: conf.BasePath + "/idpdemo", Handler: idp.HandleAuth(tenant.SystemBaseUriFromCtx, tenant.IdFromCtx, false, logError, logInfo)(http.HandleIdpDemo(conf.AssetBasePath(), templates.Render))},
	}

	lambda.Serve(http.Handle(resources),
		func(ctx context.Context, logmessage string) { log.Error(ctx, logmessage) },
		func(ctx context.Context, logmessage string) { log.Info(ctx, logmessage) })
}

func setupLogging() {
	log.StdError.SetWriteMessage(newWriteHeaderFunc("ERROR"), writeMessage)
	log.StdInfo.SetWriteMessage(newWriteHeaderFunc("INFO"), writeMessage)
	log.StdDebug.SetWriteMessage(newWriteHeaderFunc("DEBUG"), writeMessage)
}

func newWriteHeaderFunc(severity string) func(ctx context.Context, buf []byte, message string) []byte {
	fn := func(ctx context.Context, buf []byte, message string) []byte {
		// no timestamp because cloudwatch adds timestamps
		buf = append(buf, severity...)
		buf = append(buf, ' ')
		return buf
	}
	return fn
}

func writeMessage(ctx context.Context, buf []byte, message string) []byte {
	// STRUCTURED-DATA
	ten, _ := tenant.IdFromCtx(ctx)
	rid, _ := requestid.FromCtx(ctx)
	lambdaReqId, _ := lambda.ReqIdFromCtx(ctx)
	buf = append(buf, fmt.Sprintf("[ctx@49610 rid=\"%v\" lrid=\"%v\" tn=\"%v\"]",
		rid, lambdaReqId, ten)...)

	// MSG
	if message != "" {
		msgBegin := strings.LastIndex(message, "]") + 1
		// add part of the message as additional structured data
		buf = append(buf, message[:msgBegin]...)
		if msgBegin+1 <= len(message)-1 && message[msgBegin] != ' ' {
			// add space between structured data and message
			buf = append(buf, ' ')
		}
		// add message
		buf = append(buf, message[msgBegin:]...)
	}

	return buf
}
