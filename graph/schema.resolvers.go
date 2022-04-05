package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/RickardA/stop-watch/graph/generated"
)

func (r *mutationResolver) SendStopSignal(ctx context.Context, lane int) (int, error) {
	*r.StopChan <- lane
	return lane, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
