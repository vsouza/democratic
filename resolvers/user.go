package resolvers

import (
	"github.com/graphql-go/graphql"
	"github.com/vsouza/democratic/models"
)

func GetUserByID(p graphql.ResolveParams) (interface{}, error) {
	return models.GetUserByID(""), nil
}
