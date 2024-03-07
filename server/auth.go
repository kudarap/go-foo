package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

// JWTAuth represents json web token bearer authentication.
type JWTAuth struct {
	NoVerify bool
}

type AuthClaims struct {
	ID string `json:"id"`
	jwt.RegisteredClaims
}

func (v *JWTAuth) VerifyToken(ctx context.Context, tokenStr string) (claims map[string]interface{}, err error) {
	token, err := jwt.ParseWithClaims(tokenStr, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(``), nil
	})
	if err != nil {
		return nil, err
	}

	ac, ok := token.Claims.(*AuthClaims)
	if !ok {
		return nil, errors.New("could not parse JWT claims")
	}
	return map[string]interface{}{
		"id": ac.ID,
	}, nil
}

type authenticator interface {
	VerifyToken(ctx context.Context, token string) (claims map[string]interface{}, err error)
}

// authMiddleware is middleware that looks for authorization bearer token
// from header request and process verification. Verified token provides authorized
// user id and adds to request context that can be used for validating authorized requests.
//
// When authorization header is not present it skips the verification.
func authMiddleware(auth authenticator) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			tokenStr, err := tokenFromHeader(r.Header)
			if err != nil {
				encodeJSONError(w, err, http.StatusForbidden)
				return
			}

			ctx := r.Context()
			claims, err := auth.VerifyToken(ctx, tokenStr)
			if err != nil {
				encodeJSONError(w, err, http.StatusForbidden)
				return
			}

			id, _ := claims["id"].(string)
			if id == "" {
				encodeJSONError(w, fmt.Errorf("unauthorized request"), http.StatusForbidden)
				return
			}
			r = r.WithContext(userToContext(ctx, id))
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}

// Key to use when setting the user id.
type ctxKeyUserID int

// requestIDKey is the key that holds the unique user id in a request context.
const userIDKey ctxKeyUserID = iota

// userToContext sets user id to context.
func userToContext(parent context.Context, userID string) context.Context {
	return context.WithValue(parent, userIDKey, userID)
}

// userFromContext returns user id from the context if one is present.
func userFromContext(ctx context.Context) (userID string) {
	if ctx == nil {
		return ""
	}
	id, _ := ctx.Value(userIDKey).(string)
	return id
}

func tokenFromHeader(h http.Header) (tokenStr string, err error) {
	ah := h.Get("Authorization")
	if ah == "" {
		return "", errors.New("missing bearer token")
	}

	t := strings.TrimPrefix(ah, "Bearer ")
	if t == "" {
		return "", errors.New("malformed bearer token")
	}

	return t, nil
}
