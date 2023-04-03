package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/favtuts/golang-mux-api/cache"
	"github.com/favtuts/golang-mux-api/entity"
	"github.com/favtuts/golang-mux-api/errors"
	"github.com/favtuts/golang-mux-api/service"
)

type controller struct{}

var (
	postService service.PostService
	postCache   cache.PostCache
)

type PostController interface {
	GetPostByID(response http.ResponseWriter, request *http.Request)
	GetPosts(response http.ResponseWriter, request *http.Request)
	AddPost(response http.ResponseWriter, request *http.Request)
}

// More effiction way to use Dependency Injection
func NewPostController(service service.PostService, cache cache.PostCache) PostController {
	postService = service
	postCache = cache
	return &controller{}
}

func (*controller) GetPosts(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")
	posts, err := postService.FindAll()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error getting the posts"})
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(posts)
}

func (*controller) GetPostByID(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	postID := strings.Split(request.URL.Path, "/")[2]
	var post *entity.Post = postCache.Get(postID)
	if post == nil {
		post, err := postService.FindByID(postID)
		if err != nil {
			response.WriteHeader(http.StatusNotFound)
			json.NewEncoder(response).Encode(errors.ServiceError{Message: "No posts found!"})
			return
		}
		postCache.Set(postID, post)
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(post)
	} else {
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(post)
	}

}

func (*controller) AddPost(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")
	var post entity.Post
	err := json.NewDecoder(request.Body).Decode(&post)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error unmarshalling data"})
		return
	}
	err1 := postService.Validate(&post)
	if err1 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: err1.Error()})
		return
	}

	result, err2 := postService.Create(&post)
	if err2 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error saving the post"})
		return
	}
	postCache.Set(strconv.FormatInt(post.ID, 10), &post)
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(result)
}
