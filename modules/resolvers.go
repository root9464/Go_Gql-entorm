package graph

import (
	ent "root/database"

	graphql_types "root/graphql/out"
)

type Resolver struct {
	Client *ent.Client
}
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

func (r *Resolver) Mutation() graphql_types.MutationResolver { return &mutationResolver{r} }
func (r *Resolver) Query() graphql_types.QueryResolver       { return &queryResolver{r} }
