package request_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gawbsouza/boot-help/request"
	"github.com/gawbsouza/boot-help/response"
	"github.com/stretchr/testify/require"
)

type Product struct {
	Name  string `json:"name" validate:"required"`
	Price float64 `json:"price" validate:"required"`
}

func TestRequest_ParseJSON(t *testing.T) {
	t.Run("ParseJSON with valid body and content-type", func(t *testing.T) {
		var product Product

		req := httptest.NewRequest("POST", "/", strings.NewReader(`{"name":"Gopher", "price":122.22}`))
		req.Header.Set("Content-Type", "application/json")
	
		res := httptest.NewRecorder()
	
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if err := request.From(r).ParseJSON(&product); err != nil {
				response.To(w).BadErr(err.Error()).Send()
				return
			}
	
			response.To(w).Status(http.StatusNoContent).SendJSON()
		})
		handler.ServeHTTP(res, req)
	
		expectedCode := 204
	
		require.Equal(t, expectedCode, res.Code)
	})
	
	t.Run("ParseJSON with invalid body syntax", func(t *testing.T) {
		var product Product

		req := httptest.NewRequest("POST", "/", strings.NewReader(`{"name":"Gopher", "price":122.22`))
		req.Header.Set("Content-Type", "application/json")
	
		res := httptest.NewRecorder()
	
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if err := request.From(r).ParseJSON(&product); err != nil {
				response.To(w).BadErr(err.Error()).Send()
				return
			}
	
			response.To(w).Status(http.StatusNoContent).SendJSON()
		})
		handler.ServeHTTP(res, req)
	
		expectedCode := 400
		expectedMessage := "status_code: 400, message: unexpected EOF"
		
		require.Equal(t, expectedCode, res.Code)
		require.Equal(t, expectedMessage, res.Body.String())
	})

	t.Run("ParseJSON with invalid content-type", func(t *testing.T) {
		var product Product

		req := httptest.NewRequest("POST", "/", strings.NewReader(`{"name":"Gopher", "price":122.22"`))
		req.Header.Set("Content-Type", "ghg/jso")
	
		res := httptest.NewRecorder()
	
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if err := request.From(r).ParseJSON(&product); err != nil {
				response.To(w).BadErr(err.Error()).Send()
				return
			}
	
			response.To(w).Status(http.StatusNoContent).SendJSON()
		})
		handler.ServeHTTP(res, req)
	
		expectedCode := 400
		expectedMessage := "status_code: 400, message: request content-type is not application/json"
		
		require.Equal(t, expectedCode, res.Code)
		require.Equal(t, expectedMessage, res.Body.String())
	})
}

func TestRequest_ParseValidJSON(t *testing.T) {
	t.Run("ParseValidJSON with valid body and content-type", func(t *testing.T) {
		var product Product

		req := httptest.NewRequest("POST", "/", strings.NewReader(`{"name":"Gopher", "price":122.22}`))
		req.Header.Set("Content-Type", "application/json")
	
		res := httptest.NewRecorder()
	
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			val, err := request.From(r).ParseValidJSON(&product)
			if err != nil {
				response.To(w).BadErr(err.Error()).Send()
				return
			}

			if val != nil {
				response.To(w).UnprocessableErr(strings.Join(val, "-")).Send()
				return
			}
	
			response.To(w).Status(http.StatusNoContent).SendJSON()
		})
		handler.ServeHTTP(res, req)
	
		expectedCode := 204
	
		require.Equal(t, expectedCode, res.Code)
	})
	
	t.Run("ParseValidJSON with invalid body syntax", func(t *testing.T) {
		var product Product

		req := httptest.NewRequest("POST", "/", strings.NewReader(`{"name":"Gopher", "price":122.22`))
		req.Header.Set("Content-Type", "application/json")
	
		res := httptest.NewRecorder()
	
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			val, err := request.From(r).ParseValidJSON(&product)
			if err != nil {
				response.To(w).BadErr(err.Error()).Send()
				return
			}

			if val != nil {
				response.To(w).UnprocessableErr(strings.Join(val, "-")).Send()
				return
			}
	
			response.To(w).Status(http.StatusNoContent).SendJSON()
		})
		handler.ServeHTTP(res, req)
	
		expectedCode := 400
		expectedMessage := "status_code: 400, message: unexpected EOF"
		
		require.Equal(t, expectedCode, res.Code)
		require.Equal(t, expectedMessage, res.Body.String())
	})

	t.Run("ParseValidJSON with invalid content-type", func(t *testing.T) {
		var product Product

		req := httptest.NewRequest("POST", "/", strings.NewReader(`{"name":"Gopher", "price":122.22}`))
		req.Header.Set("Content-Type", "application/")
	
		res := httptest.NewRecorder()
	
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			val, err := request.From(r).ParseValidJSON(&product)
			if err != nil {
				response.To(w).BadErr(err.Error()).Send()
				return
			}

			if val != nil {
				response.To(w).UnprocessableErr(strings.Join(val, "-")).Send()
				return
			}
	
			response.To(w).Status(http.StatusNoContent).SendJSON()
		})
		handler.ServeHTTP(res, req)
	
		expectedCode := 400
		expectedMessage := "status_code: 400, message: request content-type is not application/json"
		
		require.Equal(t, expectedCode, res.Code)
		require.Equal(t, expectedMessage, res.Body.String())
	})

	t.Run("ParseValidJSON with invalid value in required param", func(t *testing.T) {
		var product Product

		req := httptest.NewRequest("POST", "/", strings.NewReader(`{"name":"", "price":122.22}`))
		req.Header.Set("Content-Type", "application/json")
	
		res := httptest.NewRecorder()
	
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			val, err := request.From(r).ParseValidJSON(&product)
			if err != nil {
				response.To(w).BadErr(err.Error()).Send()
				return
			}

			if val != nil {
				response.To(w).UnprocessableErr(strings.Join(val, "-")).Send()
				return
			}
	
			response.To(w).Status(http.StatusNoContent).SendJSON()
		})
		handler.ServeHTTP(res, req)
	
		expectedCode := 422
		expectedMessage := "status_code: 422, message: Key: 'Product.Name' Error:Field validation for 'Name' failed on the 'required' tag"
		
		require.Equal(t, expectedCode, res.Code)
		require.Equal(t, expectedMessage, res.Body.String())
	})
}