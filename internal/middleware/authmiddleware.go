package middleware

import (
	"context"
	"net/http"

	"github.com/samuelevilla/hasnet-api/internal/httputil"
)

// HTTP middleware setting a value on the request context
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// create a new context with a value
		userId := r.Header.Get("userId")
		if userId == "" {
			httputil.WriteError(w, http.StatusUnauthorized, "missing Authorization header")
			return
		}

		ctx := context.WithValue(r.Context(), "userId", userId)

		// call the next handler in the chain, passing the response writer and
		// the updated request object with the new context value.
		//
		// note: context.Context values are nested, so any previously set
		// values will be accessible as well, and the new `"user"` key
		// will be accessible from this point forward.
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
