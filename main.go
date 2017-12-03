//go:generate goagen bootstrap -d github.com/sky0621/try-goa/fs

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/sky0621/try-goa/app"
)

func main() {
	// Create service
	service := goa.New("fs")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "movie" controller
	c := NewMovieController(service)
	app.MountMovieController(service, c)

	// Start service
	if err := service.ListenAndServe(":8585"); err != nil {
		service.LogError("startup", "err", err)
	}

}
