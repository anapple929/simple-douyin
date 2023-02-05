package handlers

import (
	"api-gateway/pkg/logging"
	"errors"
)

// user包装错误
func PanicIfUserError(err error) {
	if err != nil {
		err = errors.New("userService--" + err.Error())
		logging.Info(err)
		panic(err)
	}
}

// feed包装错误
func PanicIfFeedError(err error) {
	if err != nil {
		err = errors.New("feedService--" + err.Error())
		logging.Info(err)
		panic(err)
	}
}

// publish包装错误
func PanicIfPublishError(err error) {
	if err != nil {
		err = errors.New("publishService--" + err.Error())
		logging.Info(err)
		panic(err)
	}
}
