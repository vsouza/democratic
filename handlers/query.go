package handlers

import (
	"errors"

	"github.com/graphql-go/graphql"
	"github.com/labstack/echo"
	"github.com/vsouza/democratic/schemas"
)

func Query(c echo.Context) error {
	result := graphql.Do(graphql.Params{
		Schema:        schemas.Schema,
		RequestString: c.QueryParam("query"),
	})
	if len(result.Errors) > 0 {
		return errors.New("GraphQL ERROR")
	}
	return c.JSON(200, result)
}
