package test_plugins //nolint: stylecheck,revive

import (
	"net/http"
)

// PluginMiddleware test
type PluginMiddleware struct {
}

// Init test
func (p *PluginMiddleware) Init() error {
	return nil
}

// Middleware test
func (p *PluginMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/halt" {
			w.WriteHeader(500)
			_, err := w.Write([]byte("halted"))
			if err != nil {
				panic("error writing the data to the http reply")
			}
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

// Name test
func (p *PluginMiddleware) Name() string {
	return "pluginMiddleware"
}

// PluginMiddleware2 test
type PluginMiddleware2 struct {
}

// Init test
func (p *PluginMiddleware2) Init() error {
	return nil
}

// Middleware test
func (p *PluginMiddleware2) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/boom" {
			w.WriteHeader(555)
			_, err := w.Write([]byte("boom"))
			if err != nil {
				panic("error writing the data to the http reply")
			}
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

// Name test
func (p *PluginMiddleware2) Name() string {
	return "pluginMiddleware2"
}
