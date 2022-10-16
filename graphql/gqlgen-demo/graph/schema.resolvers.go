package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"gqlgen-demo/graph/generated"
	"gqlgen-demo/graph/model"
	"strconv"
)

func (r *mutationResolver) UpsertCharacter(ctx context.Context, input model.CharacterInput) (*model.Character, error) {
	if input.ID == nil {
		newID := strconv.Itoa(len(r.Resolver.CharacterStore) + 1)
		input.ID = &newID
	}

	newChar := &model.Character{
		ID:   *input.ID,
		Name: input.Name,
	}
	r.CharacterStore[*input.ID] = newChar
	return newChar, nil
}

func (r *queryResolver) Character(ctx context.Context, id string) (*model.Character, error) {
	return r.CharacterStore[id], nil
}

func (r *queryResolver) Poques(ctx context.Context) ([]*model.Character, error) {
	return []*model.Character{}, nil
}

func (r *queryResolver) Kooks(ctx context.Context) ([]*model.Character, error) {
	return []*model.Character{}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
