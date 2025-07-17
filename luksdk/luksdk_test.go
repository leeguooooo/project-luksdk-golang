package luksdk_test

import (
	"fmt"
	"github.com/CFGameTech/project-luksdk-golang/luksdk"
	"os"
	"strconv"
	"testing"
	"time"
)

var AppId int64
var GameId int64 = 999999
var AppSecret = os.Getenv("TEST_APP_SECRET")
var NowUnix = time.Now().Unix()
var LukSDK = func() *luksdk.LukSDK {
	var appId = os.Getenv("TEST_APP_ID")
	var intAppId int
	var err error
	intAppId, err = strconv.Atoi(appId)
	if err != nil {
		panic(err)
	}
	AppId = int64(intAppId)

	return luksdk.NewLukSDKWithConfigurators(luksdk.ConfiguratorFN(func(config *luksdk.Config) {
		config.
			WithDebug(true).
			WithDomain(os.Getenv("TEST_DOMAIN")).
			WithAppId(AppId).
			WithAppSecret(AppSecret)
	}))
}()

func checkFatal(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

func assertionInt(t *testing.T, code, assert int) {
	if code != assert {
		t.Fatal(fmt.Errorf("code except: %d, got: %d", assert, code))
	}
}
