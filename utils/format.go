package utils

import (
	"bytes"
	"encoding/base64"
	"io"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	constants "github.com/guncv/tech-exam-software-engineering/constant"
	"github.com/guncv/tech-exam-software-engineering/entities"
)

func ErrorResponse(ctx *gin.Context, err error, details ...interface{}) {
	statusCode, ok := constants.ErrorMapWithStatusCode[err]
	if !ok {
		statusCode = http.StatusInternalServerError
	}

	code, ok := constants.ErrorMapWithCode[err]
	if !ok {
		code = constants.CodeInternalServerError
	}

	var detail interface{}
	if len(details) > 0 {
		detail = details[0]
	}

	ctx.JSON(statusCode, gin.H{
		"error": entities.ErrorResponse{
			Code:    int(code),
			Message: err.Error(),
			Details: detail,
		},
	})
}

func AbortWithErrorResponse(ctx *gin.Context, err error) {
	ctx.AbortWithStatusJSON(constants.ErrorMapWithStatusCode[err], gin.H{
		"error": entities.ErrorResponse{
			Code:    int(constants.ErrorMapWithCode[err]),
			Message: err.Error(),
		},
	})
}

func ConvertFileHeaderToBase64(fileHeader *multipart.FileHeader) (string, error) {
	file, err := fileHeader.Open() // Get multipart.File
	if err != nil {
		return "", constants.ErrOpenFileContext
	}
	defer file.Close()

	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, file); err != nil {
		return "", constants.ErrConvertFileHeaderToBase64
	}

	base64Str := base64.StdEncoding.EncodeToString(buf.Bytes())
	return base64Str, nil
}

func FormatBangkokRFC3339(t time.Time) string {
	loc, _ := time.LoadLocation(constants.CurrentTimeLocation)
	return t.In(loc).Format("2006-01-02T15:04:05.000Z07:00")
}
