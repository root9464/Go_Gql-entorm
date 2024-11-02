package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.55

import (
	"context"
	"fmt"
	"root/graphql/model"
	graph_types "root/graphql/out"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented: CreateUser - createUser"))
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, id int, input model.CreateUserInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented: UpdateUser - updateUser"))
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id int) (*model.User, error) {
	panic(fmt.Errorf("not implemented: DeleteUser - deleteUser"))
}

// GetUser is the resolver for the getUser field.
func (r *queryResolver) GetUser(ctx context.Context, id int) (*model.User, error) {
	panic(fmt.Errorf("not implemented: GetUser - getUser"))
}

// GetAllUsers is the resolver for the getAllUsers field.
func (r *queryResolver) GetAllUsers(ctx context.Context) ([]*model.User, error) {
	users, err := r.Client.User.Query().All(ctx)
	if err != nil {
		log.WithError(err).Error("Failed to fetch users")
		return nil, fmt.Errorf("failed to get users: %w", err)
	}

	modelUsers := []*model.User{}

	for _, u := range users {
		modelUser := &model.User{
			ID:       u.ID,
			Name:     u.Name,
			Email:    u.Email,
			Password: u.Password,
		}
		modelUsers = append(modelUsers, modelUser)
	}

	return modelUsers, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() graph_types.MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() graph_types.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }