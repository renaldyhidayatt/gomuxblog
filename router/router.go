package router

import (
	"context"
	"fmt"
	db "muxblog/db/sqlc"
	"muxblog/handler"
	"muxblog/middleware"
	"muxblog/repository"
	"muxblog/services"
	"net/http"

	"github.com/gorilla/mux"
)

var AppRoutes []RoutePrefix

type RoutePrefix struct {
	Prefix    string
	SubRoutes []Route
}

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	Protected   bool
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome")
}

func NewCategoryRouter(db *db.Queries) *mux.Router {
	context := context.Background()

	repositoryCategory := repository.NewCategoryRepository(db, context)
	serviceCategory := services.NewCategoryService(repositoryCategory)
	handlerCategory := handler.NewCategoryHandler(serviceCategory)

	repositoryUser := repository.NewUserRepository(db, context)
	serviceUser := services.NewUserService(repositoryUser)
	handlerUser := handler.NewUserHandler(serviceUser)

	repositoryPosts := repository.NewPostsRepository(db, context)
	servicePosts := services.NewPostService(repositoryPosts)
	handlerPosts := handler.NewPostsHandler(servicePosts)

	repositoryComment := repository.NewCommentRepository(db, context)
	serviceComment := services.NewCommentService(repositoryComment)
	handlerComment := handler.NewCommentHandler(serviceComment)

	router := mux.NewRouter()
	var RoutesCategory = RoutePrefix{
		"/",
		[]Route{
			{
				"Index",
				"GET",
				"/",
				indexHandler,
				false,
			},
			{
				"GetAll",
				"GET",
				"/category",
				handlerCategory.GetAll,
				false,
			},
			{
				"GetID",
				"GET",
				"/category/{id}",
				handlerCategory.GetID,
				false,
			},
			{
				"Create",
				"POST",
				"/category/create",
				handlerCategory.Create,
				false,
			},
			{
				"Update",
				"PUT",
				"/category/update/{id}",
				handlerCategory.Update,
				false,
			},
			{
				"Delete",
				"DELETE",
				"/category/{id}",
				handlerCategory.Delete,
				false,
			},
			{
				"UserAll",
				"GET",
				"/users",
				handlerUser.GetAll,
				true,
			},
			{
				"GetID",
				"GET",
				"/users/{id}",
				handlerUser.GetBYID,
				false,
			},
			{
				"Update",
				"PUT",
				"/users/update/{id}",
				handlerUser.Update,
				false,
			},
			{
				"Delete",
				"DELETE",
				"/users/{id}",
				handlerUser.Delete,
				false,
			},
			{
				"GETALLPOSTS",
				"GET",
				"/posts",
				handlerPosts.GetAll,
				false,
			},
			{
				"GETBYID",
				"GET",
				"/posts/{id}",
				handlerPosts.GetID,
				false,
			},
			{
				"GetPostRelation",
				"GET",
				"/posts/relation/{id}",
				handlerPosts.GetIDRelationJoin,
				false,
			},
			{
				"CreatePosts",
				"POST",
				"/posts/create",
				handlerPosts.Create,
				false,
			},
			{
				"UpdatePosts",
				"PUT",
				"/posts/update/{id}",
				handlerPosts.Update,
				false,
			},
			{
				"DeletePosts",
				"DELETE",
				"/posts/{id}",
				handlerPosts.Delete,
				false,
			},
			// Comments
			{
				"GETALLCOMMENTS",
				"GET",
				"/comments",
				handlerComment.GetAll,
				false,
			},
			{
				"GETBYIDCOMMENT",
				"GET",
				"/comment/{id}",
				handlerComment.GetID,
				false,
			},
			{
				"CREATECOMMENT",
				"POST",
				"/comment/create",
				handlerComment.Create,
				false,
			},
			{
				"UpdateComments",
				"PUT",
				"/comment/{id}",
				handlerComment.Update,
				false,
			},
			{
				"DeleteComment",
				"DELETE",
				"/comment/{id}",
				handlerComment.Delete,
				false,
			},
			{
				"register",
				"POST",
				"/auth/register",
				handlerUser.Create,
				false,
			},
			{
				"login",
				"POST",
				"/auth/login",
				handlerUser.Login,
				false,
			},
		},
	}

	AppRoutes = append(AppRoutes, RoutesCategory)

	for _, route := range AppRoutes {
		routerPrefix := router.PathPrefix(route.Prefix).Subrouter()

		for _, r := range route.SubRoutes {
			var handler http.Handler
			handler = r.HandlerFunc

			if r.Protected {
				handler = middleware.AuthMiddleware(r.HandlerFunc)
			}

			routerPrefix.
				Path(r.Pattern).
				Handler(handler).
				Methods(r.Method).
				Name(r.Name)
		}
	}

	return router

}
