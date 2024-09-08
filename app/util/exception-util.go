package util

import (
	"errors"
	"fmt"
	"github.com/Savitar465/task-manager/app/constant"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func PanicException_(key string, message string) {
	err := errors.New(message)
	err = fmt.Errorf("%s: %w", key, err)
	if err != nil {
		panic(err)
	}
}

func PanicException(responseKey constant.ResponseStatus) {
	PanicException_(responseKey.GetResponseStatus(), responseKey.GetResponseMessage())
}

func PanicHandler(c *gin.Context) {
	if err := recover(); err != nil {
		str := fmt.Sprint(err)
		strArr := strings.SplitN(str, ":", 2) // Split into at most 2 parts

		key := strings.TrimSpace(strArr[0])
		msg := strings.TrimSpace(strArr[1])

		switch key {
		case constant.DataNotFound.GetResponseStatus():
			c.JSON(http.StatusBadRequest, BuildResponse_(key, msg, Null()))
		case constant.Unauthorized.GetResponseStatus():
			c.JSON(http.StatusUnauthorized, BuildResponse_(key, msg, Null()))
		default:
			c.JSON(http.StatusInternalServerError, BuildResponse_(key, msg, Null()))
		}
		c.Abort()
	}
}
