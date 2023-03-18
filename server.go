package main

import (
	"fmt"
	"golang-mux-api/controller"
	router "golang-mux-api/http"
	"golang-mux-api/repository"
	"golang-mux-api/service"
	"net/http"
)

var (
	//postRepository repository.PostRepository = repository.NewFirestoreRepository()
	postRepository repository.PostRepository = repository.NewSQLiteRepository()
	postService    service.PostService       = service.NewPostService(postRepository)
	postController controller.PostController = controller.NewPostController(postService)
	httpRouter     router.Router             = router.NewMuxRouter()
	//httpRouter router.Router = router.NewChiRouter()
)

func main() {
	const port string = ":8000"

	httpRouter.GET("/", func(resp http.ResponseWriter, request *http.Request) {
		fmt.Fprint(resp, "Up and running...")
	})

	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPost)

	httpRouter.SERVE(port)
}
