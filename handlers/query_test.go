package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/vsouza/democratic/handlers"
)

var (
	queryGraphQL  = `{user(id:"test"){username}}`
	resultGraphQL = `{"data":{"user":{"username":"example"}}}`
)

func TestQueryUser(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.GET, fmt.Sprintf("/graphql?query=%s", queryGraphQL), nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, handlers.Query(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, resultGraphQL, rec.Body.String())
	}

}
