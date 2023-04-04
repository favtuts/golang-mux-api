package main

import (
	"os"

	"github.com/favtuts/golang-mux-api/cache"
	"github.com/favtuts/golang-mux-api/controller"
	router "github.com/favtuts/golang-mux-api/http"
	"github.com/favtuts/golang-mux-api/repository"
	"github.com/favtuts/golang-mux-api/service"
)

var (
	//postRepository repository.PostRepository = repository.NewFirestoreRepository()
	//postRepository repository.PostRepository = repository.NewSQLiteRepository()
	postRepository repository.PostRepository = repository.NewDynamoDBRepository()
	postService    service.PostService       = service.NewPostService(postRepository)
	postCache      cache.PostCache           = cache.NewRedisCache("localhost:6379", 1, 100)
	postController controller.PostController = controller.NewPostController(postService, postCache)
	httpRouter     router.Router             = router.NewMuxRouter()
	//httpRouter router.Router = router.NewChiRouter()
)

func main() {

	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPost)
	httpRouter.GET("/posts/{id}", postController.GetPostByID)

	httpRouter.SERVE(":" + os.Getenv("PORT"))
}
