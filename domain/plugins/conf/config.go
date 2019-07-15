// Package conf provides the config for the app
package conf

import (
	"context"
	"encoding/base64"
	"github.com/d-velop/dvelop-sdk-go/log"
	"os"
)

const AppName = "shop-middleware"
const BasePath = "/" + AppName

// Will be used if request contains no information about the SystemBaseUri
func DefaultSystemBaseURI() string {
	return os.Getenv("systemBaseUri")
}

func AssetBasePath() string {
	e := os.Getenv("ASSET_BASE_PATH")
	if e == "" {
		return "/" + AppName + "/assets"
	} else {
		return e
	}
}

// Uncomment to enable writing to a local syslog server
//func SyslogEndpoint() string {
//	e := os.Getenv("SYSLOG_ENDPOINT")
//	if e == "" {
//		return "localhost:514"
//	} else {
//		return e
//	}
//}

func SecretKey() []byte {
	sig, err := base64.StdEncoding.DecodeString(os.Getenv("SIGNATURE_SECRET"))
	if err != nil {
		log.Info(context.Background(), "error while decoding secretkey env", err)
	}
	return sig
}

func Version() string {
	const baseVer = "1.0.0"
	v := os.Getenv("BUILD_VERSION")
	if v != "" {
		return baseVer + "+" + v
	}
	return baseVer
}
