package middleware

import (
	"context"
	"strings"

	"github.com/gin-gonic/gin"
	constants "github.com/guncv/tech-exam-software-engineering/constant"
	"github.com/guncv/tech-exam-software-engineering/infras/log"
	"github.com/guncv/tech-exam-software-engineering/utils"
)

func AuthMiddleware(tokenMaker utils.IPasetoMaker, log *log.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.DebugWithID(ctx, "[Middleware: AuthMiddleware] Called")
		authorizationHeader := ctx.GetHeader(constants.AuthorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			log.ErrorWithID(ctx, "[Middleware: AuthMiddleware] Authorization header is not provided")
			utils.AbortWithErrorResponse(ctx, constants.ErrAuthHeaderMissing)
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			log.ErrorWithID(ctx, "[Middleware: AuthMiddleware] Invalid authorization header format")
			utils.AbortWithErrorResponse(ctx, constants.ErrAuthHeaderFormatInvalid)
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != constants.AuthorizationTypeBearer {
			log.ErrorWithID(ctx, "[Middleware: AuthMiddleware] Authorization header must start with "+constants.AuthorizationTypeBearer)
			utils.AbortWithErrorResponse(ctx, constants.ErrAuthHeaderMissingBearer)
			return
		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			log.ErrorWithID(ctx, "[Middleware: AuthMiddleware] Failed to verify token", err)
			utils.AbortWithErrorResponse(ctx, constants.ErrFailedToVerifyToken)
			return
		}

		newCtx := context.WithValue(ctx.Request.Context(), constants.AuthorizationPayloadKey, payload)
		ctx.Request = ctx.Request.WithContext(newCtx)
		log.DebugWithID(ctx, "[Middleware: AuthMiddleware] Token verified successfully", payload)
		log.DebugWithID(ctx, "[Middleware: AuthMiddleware] Next middleware", newCtx.Value(constants.AuthorizationPayloadKey))

		ctx.Next()
	}
}
