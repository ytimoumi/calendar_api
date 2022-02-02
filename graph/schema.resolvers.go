package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"giskard_api/controllers"
	"giskard_api/graph/generated"
	"giskard_api/middlewares"
	"giskard_api/models"
	"log"
)

func (r *mutationResolver) CreateReservation(ctx context.Context, idAvailability int, input models.ToReserve) (*models.Response, error) {
	ginContext, err := middlewares.GinContextFromContext(ctx)
	var response *models.Response
	if err != nil {
		return nil, err
	}
	res, er := controllers.CreateReservation(ginContext, idAvailability, input)
	if er != nil {
		log.Println("CreateReservation error")
		return response, er
	}
	return res, nil
}

func (r *mutationResolver) DeleteReservation(ctx context.Context, id int, email string) (*models.Response, error) {
	ginContext, err := middlewares.GinContextFromContext(ctx)
	var response *models.Response
	if err != nil {
		return nil, err
	}
	res, er := controllers.DeleteReservation(ginContext, id, email)
	if er != nil {
		log.Println("DeleteReservation error")
		return response, er
	}
	return res, nil
}

func (r *mutationResolver) CreateAvailability(ctx context.Context, input models.Availability) (*models.Response, error) {
	ginContext, err := middlewares.GinContextFromContext(ctx)
	var response *models.Response
	if err != nil {
		return nil, err
	}
	res, er := controllers.CreateAvailability(ginContext, input)
	if er != nil {
		log.Println("CreateAvailability error")
		return response, er
	}
	return res, nil
}

func (r *mutationResolver) DeleteAvailability(ctx context.Context, id int) (*models.Response, error) {
	ginContext, err := middlewares.GinContextFromContext(ctx)
	var response *models.Response
	if err != nil {
		return nil, err
	}
	res, er := controllers.DeleteAvailability(ginContext, id)
	if er != nil {
		log.Println("DeleteAvailability error")
		return response, er
	}
	return res, nil
}

func (r *queryResolver) GetCalendar(ctx context.Context) ([]*models.Calendar, error) {
	ginContext, err := middlewares.GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}
	res, er := controllers.GetCalendar(ginContext)
	if er != nil {
		log.Println("GetCalendar error")
		return nil, er
	}
	return res, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
