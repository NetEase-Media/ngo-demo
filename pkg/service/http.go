package service

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/NetEase-Media/ngo/adapter/log"
	"github.com/NetEase-Media/ngo/client/httplib"
	"github.com/valyala/fasthttp"
)

// HttpRequest
func HttpRequest(ctx context.Context) (string, error) {
	var rs string
	code, err := httplib.Get("http://localhost:8080/hello/" + strconv.Itoa(time.Now().Second())).BindString(&rs).Do(ctx)
	if err != nil {
		log.Error("http request failed ", err)
		return "", err
	}
	if code != fasthttp.StatusOK {
		log.WithField("status", code).Error("code error")
		return "", errors.New("response code error")
	}
	return rs, nil
}
