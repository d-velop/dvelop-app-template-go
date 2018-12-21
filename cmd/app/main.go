package main

import (
	"context"
	"fmt"
	"github.com/d-velop/dvelop-app-template-go/domain/plugins/conf"
	"github.com/d-velop/dvelop-app-template-go/domain/plugins/gui/assets"
	"github.com/d-velop/dvelop-app-template-go/domain/plugins/gui/templates"
	"github.com/d-velop/dvelop-app-template-go/domain/plugins/http"
	"github.com/d-velop/dvelop-sdk-go/log"
	"github.com/d-velop/dvelop-sdk-go/log/syslog"
	"github.com/d-velop/dvelop-sdk-go/requestid"
	"github.com/d-velop/dvelop-sdk-go/tenant"
	"net"
	"os"
	"strings"
)

func main() {
	setupLogging()

	// wire dependencies
	// todo var storage = memory.NewStore()
	//var listVacationRequestsSrv = listVacationRequests.NewService(&storage)

	resources := []http.Resource{
		{Pattern: conf.BasePath + "/", Handler: http.HandleRoot(conf.AssetBasePath(), templates.Render, conf.Version())},
		{Pattern: conf.BasePath + "/assets/", Handler: http.HandleAssets(conf.BasePath+"/assets/", assets.AssetFileSystem)},
		{Pattern: conf.BasePath + "/vacationrequest", Handler: http.HandleVacationRequest(conf.AssetBasePath(), templates.Render)},
	}

	socket, err := net.Listen("tcp", "localhost:")
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
