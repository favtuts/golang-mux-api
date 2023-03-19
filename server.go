package main

import (
	"os"

	"github.com/favtuts/golang-mux-api/controller"
	router "github.com/favtuts/golang-mux-api/http"
	"github.com/favtuts/golang-mux-api/repository"
	"github.com/favtuts/golang-mux-api/service"
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

	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPost)

	httpRouter.SERVE(":" + os.Getenv("PORT"))
}
