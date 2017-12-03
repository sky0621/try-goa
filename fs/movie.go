package fs

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("fs", func() {
	Title("Fan Movie")
	Description("A simple movie")
	Scheme("http")
	Host("localhost:8585")
})

var MovieMedia = MediaType("application/vnd.movie+json", func() {
	Description("Movie Response Type")
	Attributes(func() {
		Attribute("movieId", Integer, "Unique Movie ID")
		Attribute("movieName", String, "Movie Name")

		Required("movieId", "movieName")
	})
	View("default", func() {
		Attribute("movieId")
		Attribute("movieName")
	})
})

var _ = Resource("movie", func() {
	BasePath("/movies")
	DefaultMedia(MovieMedia)

	Action("show", func() {
		Description("Get 1 Movie Information")
		Routing(GET("/:movieId"))
		Params(func() {
			Param("movieId", Integer, "Unique Movie ID")
		})
		Response(OK)
		Response(NotFound)
	})
})
