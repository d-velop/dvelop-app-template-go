package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/d-velop/dvelop-app-template-go/domain/acceptVacationRequest"
	"github.com/d-velop/dvelop-app-template-go/domain/applyForVacation"
	"github.com/d-velop/dvelop-app-template-go/domain/cancelVacation"
	"github.com/d-velop/dvelop-app-template-go/domain/plugins/conf"
	"github.com/d-velop/dvelop-app-template-go/domain/plugins/gui/assets"
	"github.com/d-velop/dvelop-app-template-go/domain/plugins/gui/templates"
	"github.com/d-velop/dvelop-app-template-go/domain/plugins/http"
	"github.com/d-velop/dvelop-app-template-go/domain/plugins/storage/memory"
	"github.com/d-velop/dvelop-app-template-go/domain/rejectVacationRequest"
	"github.com/d-velop/dvelop-sdk-go/idp"
	"github.com/d-velop/dvelop-sdk-go/idp/idpclient"
	"github.com/d-velop/dvelop-sdk-go/log"
	"github.com/d-velop/dvelop-sdk-go/log/syslog"
	"github.com/d-velop/dvelop-sdk-go/requestid"
	"github.com/d-velop/dvelop-sdk-go/tenant"
)

func main() {
	setupLogging()

	// wire dependencies
	storage := memory.NewStore()
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

	idpClient, err := idpclient.New()
	if err != nil {
		log.Error(context.Background(), err)
		os.Exit(1)
	}

	authenticate := idp.Authenticate(idpClient, tenant.SystemBaseUriFromCtx, tenant.IdFromCtx, false, logError, logInfo)

	resources := []http.Resource{
		{Pattern: conf.BasePath + "/", Handler: http.HandleRoot(conf.AssetBasePath(), templates.Render, conf.Version())},
		{Pattern: conf.BasePath + "/assets/", Handler: http.HandleAssets(conf.BasePath+"/assets/", assets.AssetFileSystem)},
		{Pattern: conf.BasePath + "/vacationrequest", Handler: vacationRequestHandler.HandleNewForm()},
		{Pattern: conf.BasePath + "/vacationrequest/", Handler: vacationRequestHandler.Handle(conf.BasePath + "/vacationrequest/")},
		{Pattern: conf.BasePath + "/features", Handler: http.HandleFeatures()},
		{Pattern: conf.BasePath + "/idpdemo", Handler: authenticate(http.HandleIdpDemo(conf.AssetBasePath(), templates.Render))},
	}

	socket, err := net.Listen("tcp", "localhost:5000")
	if err != nil {
		log.Error(context.Background(), err)
		os.Exit(1)
	}
	log.Info(context.Background(), fmt.Sprintf("Listening on http://%s/%s", socket.Addr().String(), conf.AppName))

	err = http.Serve(socket, http.Handle(resources))
	if err != nil {
		log.Error(context.Background(), err)
	}

}

func setupLogging() {
	log.StdError.SetWriteMessage(syslog.NewWriteHeaderFunc(conf.AppName, syslog.ERROR), writeMessage)
	log.StdInfo.SetWriteMessage(syslog.NewWriteHeaderFunc(conf.AppName, syslog.INFO), writeMessage)
	log.StdDebug.SetWriteMessage(syslog.NewWriteHeaderFunc(conf.AppName, syslog.DEBUG), writeMessage)
	// Uncomment to enable writing to a local syslog server
	//syslogwriter, err := syslog.NewWriter(conf.SyslogEndpoint())
	//if err != nil {
	//	log.Infof(context.Background(), "Could not connect to syslogserver '%v'. Writing log to STDERR.", conf.SyslogEndpoint())
	//} else {
	//	log.Infof(context.Background(), "Writing logs to syslogserver '%v'", conf.SyslogEndpoint())
	//	log.StdError.SetOutput(syslogwriter)
	//	log.StdInfo.SetOutput(syslogwriter)
	//	log.StdDebug.SetOutput(syslogwriter)
	//}
}

func writeMessage(ctx context.Context, buf []byte, message string) []byte {
	// STRUCTURED-DATA
	ten, _ := tenant.IdFromCtx(ctx)
	rid, _ := requestid.FromCtx(ctx)
	buf = append(buf, fmt.Sprintf("[ctx@49610 rid=\"%v\" tn=\"%v\"]", rid, ten)...)

	// MSG
	if message != "" {
		msgBegin := strings.LastIndex(message, "]") + 1
		// add part of the message as additional structured data
		buf = append(buf, message[:msgBegin]...)
		if msgBegin+1 <= len(message)-1 && message[msgBegin] != ' ' {
			// add space and BOM between structured data and message
			buf = append(buf, ' ')
			const BOM = "\xef\xbb\xbf" // cf. https://de.wikipedia.org/wiki/Byte_Order_Mark
			buf = append(buf, BOM...)
		}
		// add message
		buf = append(buf, message[msgBegin:]...)
	}

	return buf
}
