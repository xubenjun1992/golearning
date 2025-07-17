package test

import (
	"blog/controller"
	"blog/dao"
	"blog/reqdto"
	"blog/rout"
	"blog/service"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateArticle(t *testing.T) {
	// This is a placeholder for the test function.
	// You can implement your test logic here.
	t.Log("Create Article test executed")
	userDao := dao.NewUserDao()
	commentDao := dao.NewCommentDao()
	articleDao := dao.NewArticleDao()
	userService := service.NewUserService(userDao)
	commentService := service.NewCommentService(commentDao)
	articleService := service.NewArticleService(articleDao)

	r := rout.SetBlogRouts(
		controller.NewArticleController(articleService),
		controller.NewCommentController(commentService),
		controller.NewUserController(userService),
	)
	w := httptest.NewRecorder()
	articleDto := reqdto.CreateArticleDTO{
		Title:   "Test Article",
		Content: "This is a test article content.",
		UserId:  1, // Assuming user with ID 1 exists
	}
	body, _ := json.Marshal(articleDto)
	req := httptest.NewRequest("POST", "/api/blog/articles", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTI3NzU1NjAsInVzZXJJZCI6MSwidXNlck5hbWUiOiJ0ZXN0dXNlciJ9.d9MsMzn6GHXCNzygw5SJ8QyHOdsqgrlXQSUCKzkUgnI") // Assuming you have a valid token for testing
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d, body: %s", w.Code, w.Body)
	}
}

func TestGetAllArticles(t *testing.T) {
	// This is a placeholder for the test function.
	// You can implement your test logic here.
	t.Log("Get All Articles test executed")
	userDao := dao.NewUserDao()
	commentDao := dao.NewCommentDao()
	articleDao := dao.NewArticleDao()
	userService := service.NewUserService(userDao)
	commentService := service.NewCommentService(commentDao)
	articleService := service.NewArticleService(articleDao)

	r := rout.SetBlogRouts(
		controller.NewArticleController(articleService),
		controller.NewCommentController(commentService),
		controller.NewUserController(userService),
	)
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/blog/articles", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
	t.Logf("Login response: %s", w.Body.String())
}

func TestUpdateArticle(t *testing.T) {
	// This is a placeholder for the test function.
	// You can implement your test logic here.
	t.Log("Update Article test executed")
	userDao := dao.NewUserDao()
	commentDao := dao.NewCommentDao()
	articleDao := dao.NewArticleDao()
	userService := service.NewUserService(userDao)
	commentService := service.NewCommentService(commentDao)
	articleService := service.NewArticleService(articleDao)

	r := rout.SetBlogRouts(
		controller.NewArticleController(articleService),
		controller.NewCommentController(commentService),
		controller.NewUserController(userService),
	)
	w := httptest.NewRecorder()
	articleDto := reqdto.CreateArticleDTO{
		Id:      14, // Assuming article with ID 1 exists
		Title:   "Updated Test Article",
		Content: "This is updated content for the test article.",
	}
	body, _ := json.Marshal(articleDto)
	req := httptest.NewRequest("PUT", "/api/blog/articles/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTI3NzYxMjAsInVzZXJJZCI6MSwidXNlck5hbWUiOiJ0ZXN0dXNlciJ9.API1e6nBQwr3vTQBkNXivpFlbo7bkRRYQkyj0Ll2jUo")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d, body: %s", w.Code, w.Body.String())
	}
}

func TestDeleteArticle(t *testing.T) {
	// This is a placeholder for the test function.
	// You can implement your test logic here.
	t.Log("Delete Article test executed")
	userDao := dao.NewUserDao()
	commentDao := dao.NewCommentDao()
	articleDao := dao.NewArticleDao()
	userService := service.NewUserService(userDao)
	commentService := service.NewCommentService(commentDao)
	articleService := service.NewArticleService(articleDao)

	r := rout.SetBlogRouts(
		controller.NewArticleController(articleService),
		controller.NewCommentController(commentService),
		controller.NewUserController(userService),
	)
	w := httptest.NewRecorder()
	req := httptest.NewRequest("DELETE", "/api/blog/articles/1", nil) // Assuming article with ID 1 exists
	req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTI3NzYxMjAsInVzZXJJZCI6MSwidXNlck5hbWUiOiJ0ZXN0dXNlciJ9.API1e6nBQwr3vTQBkNXivpFlbo7bkRRYQkyj0Ll2jUo")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}

func TestGetAllArticlesByUserId(t *testing.T) {
	// This is a placeholder for the test function.
	// You can implement your test logic here.
	t.Log("Get Articles By User ID test executed")
	userDao := dao.NewUserDao()
	commentDao := dao.NewCommentDao()
	articleDao := dao.NewArticleDao()
	userService := service.NewUserService(userDao)
	commentService := service.NewCommentService(commentDao)
	articleService := service.NewArticleService(articleDao)

	r := rout.SetBlogRouts(
		controller.NewArticleController(articleService),
		controller.NewCommentController(commentService),
		controller.NewUserController(userService),
	)
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/blog/articles/user/1", nil) // Assuming user with ID 1 exists
	req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTI3NzYxMjAsInVzZXJJZCI6MSwidXNlck5hbWUiOiJ0ZXN0dXNlciJ9.API1e6nBQwr3vTQBkNXivpFlbo7bkRRYQkyj0Ll2jUo")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d, body: %s", w.Code, w.Body.String())
	}
}

func TestGetArticleById(t *testing.T) {
	// This is a placeholder for the test function.
	// You can implement your test logic here.
	t.Log("Get Article By ID test executed")
	userDao := dao.NewUserDao()
	commentDao := dao.NewCommentDao()
	articleDao := dao.NewArticleDao()
	userService := service.NewUserService(userDao)
	commentService := service.NewCommentService(commentDao)
	articleService := service.NewArticleService(articleDao)

	r := rout.SetBlogRouts(
		controller.NewArticleController(articleService),
		controller.NewCommentController(commentService),
		controller.NewUserController(userService),
	)
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/blog/articles/1", nil) // Assuming article with ID 1 exists
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d, body: %s", w.Code, w.Body.String())
	}
}
