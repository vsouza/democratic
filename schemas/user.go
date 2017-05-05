package schemas

import "github.com/graphql-go/graphql"

var userType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id":        &graphql.Field{Type: graphql.String},
		"email":     &graphql.Field{Type: graphql.String},
		"username":  &graphql.Field{Type: graphql.String},
		"fullName":  &graphql.Field{Type: graphql.String},
		"gender":    &graphql.Field{Type: graphql.String},
		"status":    &graphql.Field{Type: graphql.String},
		"userType":  &graphql.Field{Type: graphql.String},
		"birthDate": &graphql.Field{Type: graphql.String},
		"createdAt": &graphql.Field{Type: graphql.String},
		"updatedAt": &graphql.Field{Type: graphql.String},
		"address": &graphql.Field{
			Type: graphql.NewObject(graphql.ObjectConfig{
				Name: "UserAddress",
				Fields: graphql.Fields{
					"neighborhood":   &graphql.Field{Type: graphql.String},
					"zipCode":        &graphql.Field{Type: graphql.String},
					"addressLineOne": &graphql.Field{Type: graphql.String},
					"addressLineTwo": &graphql.Field{Type: graphql.String},
					"addressType":    &graphql.Field{Type: graphql.String},
					"number":         &graphql.Field{Type: graphql.String},
					"city":           &graphql.Field{Type: graphql.String},
					"state":          &graphql.Field{Type: graphql.String},
					"country":        &graphql.Field{Type: graphql.String},
				},
			}),
		},
	},
})
