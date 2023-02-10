package ginmiddleware

import (
	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
)

// SentryTracingMiddleware makes middleware, that record request as transaction.
func SentryTracingMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		transaction := sentry.StartTransaction(
			ctx, ctx.FullPath(),
			sentry.ContinueFromRequest(ctx.Request),
		)

		defer func(){
			// TODO: recover panic and finish transaction with error

			transaction.Status = matchTransactionStatus(ctx.Request.Response.StatusCode)
			transaction.Finish()
		}()

		ctx.Next()
	}
}

// SentryDefaultMiddleware is a wrapper around official Sentry library with little improvements.
func SentryDefaultMiddleware() gin.HandlerFunc {
	return sentrygin.New(sentrygin.Options{
		Repanic:         true,
	})
}

// matchTransactionStatus mathes HTTP status code to Sentry transaction status.
//
// More infofmation here: https://develop.sentry.dev/sdk/event-payloads/span/
func matchTransactionStatus(statusCode int) string {
	if statusCode >= 200 && statusCode < 300 {
		return "ok"
	}

	switch statusCode {
	case 400:
		return "failed_precondition"
	case 401:
		return "unauthenticated"
	case 403:
		return "permission_denied"
	case 404:
		return "not_found"
	case 409:
		return "already_exists"
	case 429:
		return "resource_exhausted"
	case 499:
		return "cancelled"
	case 500:
		return "internal_error"
	case 501:
		return "unimplemented"
	case 503:
		return "unavailable"
	case 504:
		return "deadline_exceeded"
	default:
		return "unknown"
	}

	switch {
	case statusCode >= 200 && statusCode < 300:
		return "ok"
	case statusCode ==
	case statusCode == 499:
		return "cancelled"
	case
	}
}
