package middleware

import (
	"context"
	"net/http"
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
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse(constants.ErrAuthorizationHeaderNotProvided))
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			log.ErrorWithID(ctx, "[Middleware: AuthMiddleware] Invalid authorization header format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse(constants.ErrInvalidAuthorizationHeaderFormat))
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != constants.AuthorizationTypeBearer {
			log.ErrorWithID(ctx, "[Middleware: AuthMiddleware] Authorization header must start with "+constants.AuthorizationTypeBearer)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse(constants.ErrAuthorizationHeaderMustStartWithBearer))
			return
		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			log.ErrorWithID(ctx, "[Middleware: AuthMiddleware] Failed to verify token", err)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse(constants.ErrFailedToVerifyToken))
			return
		}

		newCtx := context.WithValue(ctx.Request.Context(), constants.AuthorizationPayloadKey, payload)
		ctx.Request = ctx.Request.WithContext(newCtx)
		log.DebugWithID(ctx, "[Middleware: AuthMiddleware] Token verified successfully", payload)
		log.DebugWithID(ctx, "[Middleware: AuthMiddleware] Next middleware", newCtx.Value(constants.AuthorizationPayloadKey))

		ctx.Next()
	}
}
