package routes

import "github.com/mberlanda/impostor/api/controller"

type UserRoute struct {
	controller *controller.UserController
	router     *GinRouter
}

func NewUserRoute(c *controller.UserController, r *GinRouter) UserRoute {
	return UserRoute{controller: c, router: r}
}

func (u *UserRoute) Setup() {
	user := u.router.Gin.Group("/users")
	{
		user.GET("/", u.controller.Index)
		user.POST("/", u.controller.Create)
		user.GET("/:id", u.controller.Show)
	}
}
