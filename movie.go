package main

import (
	"github.com/goadesign/goa"
	"github.com/sky0621/try-goa/app"
)

// MovieController implements the movie resource.
type MovieController struct {
	*goa.Controller
}

// NewMovieController creates a movie controller.
func NewMovieController(service *goa.Service) *MovieController {
	return &MovieController{Controller: service.NewController("MovieController")}
}

// Show runs the show action.
func (c *MovieController) Show(ctx *app.ShowMovieContext) error {
	// MovieController_Show: start_implement

	// Put your logic here

	// MovieController_Show: end_implement
	res := &app.Movie{}
	return ctx.OK(res)
}
