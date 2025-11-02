package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// handlerMock est un handler simple pour les tests
func handlerMock(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func TestCORSMiddleware_AllowedHost(t *testing.T) {
	config := &CORSConfig{
		AllowedOrigins: []string{"example.com", "api.example.com"},
	}

	middleware := CORSMiddleware(config)
	handler := middleware(http.HandlerFunc(handlerMock))

	req := httptest.NewRequest("GET", "/", nil)
	req.Host = "example.com"

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "OK"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestCORSMiddleware_UnauthorizedHost(t *testing.T) {
	config := &CORSConfig{
		AllowedOrigins: []string{"example.com", "api.example.com"},
	}

	middleware := CORSMiddleware(config)
	handler := middleware(http.HandlerFunc(handlerMock))

	req := httptest.NewRequest("GET", "/", nil)
	req.Host = "unauthorized.com"

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusForbidden {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusForbidden)
	}
}

func TestCORSMiddleware_HostWithPort(t *testing.T) {
	config := &CORSConfig{
		AllowedOrigins: []string{"example.com"},
	}

	middleware := CORSMiddleware(config)
	handler := middleware(http.HandlerFunc(handlerMock))

	req := httptest.NewRequest("GET", "/", nil)
	req.Host = "example.com:3000"

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code for host with port: got %v want %v", status, http.StatusOK)
	}
}

func TestCORSMiddleware_LocalhostAllowed(t *testing.T) {
	config := &CORSConfig{
		AllowedOrigins: []string{"example.com"},
	}

	middleware := CORSMiddleware(config)
	handler := middleware(http.HandlerFunc(handlerMock))

	testCases := []string{
		"localhost",
		"localhost:3000",
		"127.0.0.1",
		"127.0.0.1:8080",
		"[::1]",
		"[::1]:3000",
	}

	for _, host := range testCases {
		req := httptest.NewRequest("GET", "/", nil)
		req.Host = host

		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("localhost %s should be allowed: got %v want %v", host, status, http.StatusOK)
		}
	}
}

func TestCORSMiddleware_CORSHeaders(t *testing.T) {
	config := &CORSConfig{
		AllowedOrigins: []string{"example.com"},
	}

	middleware := CORSMiddleware(config)
	handler := middleware(http.HandlerFunc(handlerMock))

	req := httptest.NewRequest("GET", "/", nil)
	req.Host = "api.example.com"
	req.Header.Set("Origin", "https://example.com")

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Vérifier les headers CORS
	if origin := rr.Header().Get("Access-Control-Allow-Origin"); origin != "https://example.com" {
		t.Errorf("Access-Control-Allow-Origin = %v, want %v", origin, "https://example.com")
	}

	if methods := rr.Header().Get("Access-Control-Allow-Methods"); methods == "" {
		t.Error("Access-Control-Allow-Methods should be set")
	}

	if headers := rr.Header().Get("Access-Control-Allow-Headers"); headers == "" {
		t.Error("Access-Control-Allow-Headers should be set")
	}

	if credentials := rr.Header().Get("Access-Control-Allow-Credentials"); credentials != "true" {
		t.Errorf("Access-Control-Allow-Credentials = %v, want %v", credentials, "true")
	}
}

func TestCORSMiddleware_PreflightRequest(t *testing.T) {
	config := &CORSConfig{
		AllowedOrigins: []string{"example.com"},
	}

	middleware := CORSMiddleware(config)
	handler := middleware(http.HandlerFunc(handlerMock))

	req := httptest.NewRequest("OPTIONS", "/", nil)
	req.Host = "api.example.com"
	req.Header.Set("Origin", "https://example.com")
	req.Header.Set("Access-Control-Request-Method", "POST")

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("preflight request returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Le body ne devrait pas contenir "OK" car on répond directement à la preflight
	if rr.Body.String() == "OK" {
		t.Error("preflight request should not execute the handler")
	}
}

func TestCORSMiddleware_NoConfig(t *testing.T) {
	config := &CORSConfig{
		AllowedOrigins: []string{},
	}

	middleware := CORSMiddleware(config)
	handler := middleware(http.HandlerFunc(handlerMock))

	req := httptest.NewRequest("GET", "/", nil)
	req.Host = "any-domain.com"

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	// Sans configuration, tout devrait être autorisé
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("with no config, all should be allowed: got %v want %v", status, http.StatusOK)
	}
}

// Tests pour la validation du Domain configuré

func TestCORSMiddleware_DomainMatch(t *testing.T) {
	config := &CORSConfig{
		Domain:         "api.example.com",
		AllowedOrigins: []string{"example.com"},
	}

	middleware := CORSMiddleware(config)
	handler := middleware(http.HandlerFunc(handlerMock))

	req := httptest.NewRequest("GET", "/", nil)
	req.Host = "api.example.com"

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("request to configured domain should be allowed: got %v want %v", status, http.StatusOK)
	}
}

func TestCORSMiddleware_DomainMismatch(t *testing.T) {
	config := &CORSConfig{
		Domain:         "api.example.com",
		AllowedOrigins: []string{"example.com"},
	}

	middleware := CORSMiddleware(config)
	handler := middleware(http.HandlerFunc(handlerMock))

	req := httptest.NewRequest("GET", "/", nil)
	req.Host = "wrong.example.com"

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusForbidden {
		t.Errorf("request to wrong domain should be blocked: got %v want %v", status, http.StatusForbidden)
	}
}

func TestCORSMiddleware_DomainWithPort(t *testing.T) {
	config := &CORSConfig{
		Domain:         "api.example.com",
		AllowedOrigins: []string{"example.com"},
	}

	middleware := CORSMiddleware(config)
	handler := middleware(http.HandlerFunc(handlerMock))

	req := httptest.NewRequest("GET", "/", nil)
	req.Host = "api.example.com:3000"

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("request to configured domain with port should be allowed: got %v want %v", status, http.StatusOK)
	}
}

func TestCORSMiddleware_LocalhostWithDomainRestriction(t *testing.T) {
	config := &CORSConfig{
		Domain:         "api.example.com",
		AllowedOrigins: []string{"example.com"},
	}

	middleware := CORSMiddleware(config)
	handler := middleware(http.HandlerFunc(handlerMock))

	testCases := []string{
		"localhost",
		"localhost:3000",
		"127.0.0.1",
		"[::1]",
	}

	for _, host := range testCases {
		req := httptest.NewRequest("GET", "/", nil)
		req.Host = host

		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("localhost %s should be allowed even with domain restriction: got %v want %v", host, status, http.StatusOK)
		}
	}
}

func TestCORSMiddleware_DomainAndCORS(t *testing.T) {
	config := &CORSConfig{
		Domain:         "api.example.com",
		AllowedOrigins: []string{"example.com", "app.example.com"},
	}

	middleware := CORSMiddleware(config)
	handler := middleware(http.HandlerFunc(handlerMock))

	req := httptest.NewRequest("GET", "/", nil)
	req.Host = "api.example.com"
	req.Header.Set("Origin", "https://example.com")

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("request with valid domain and origin should be allowed: got %v want %v", status, http.StatusOK)
	}

	// Vérifier les headers CORS
	if origin := rr.Header().Get("Access-Control-Allow-Origin"); origin != "https://example.com" {
		t.Errorf("Access-Control-Allow-Origin = %v, want %v", origin, "https://example.com")
	}
}
