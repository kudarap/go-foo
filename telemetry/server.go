package telemetry

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	semconv "go.opentelemetry.io/otel/semconv/v1.24.0"
	"go.opentelemetry.io/otel/trace"
)

type Server struct {
	service string
}

func (s *Server) Middleware() func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return otelhttp.NewHandler(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				h.ServeHTTP(w, r)

				ctx := r.Context()
				routePattern := chi.RouteContext(ctx).RoutePattern()

				span := trace.SpanFromContext(ctx)
				span.SetName(routePattern)
				span.SetAttributes(semconv.HTTPTarget(r.URL.String()), semconv.HTTPRoute(routePattern))

				labeler, ok := otelhttp.LabelerFromContext(ctx)
				if ok {
					labeler.Add(semconv.HTTPRoute(routePattern))
				}
			}),
			"",
		)
	}
}

// NewServerInstrumentation automatic instrumentation for server.
func NewServerInstrumentation(service string) *Server {
	return &Server{service: fmt.Sprintf("%s-server", service)}
}
