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

func TestCreateComment(t *testing.T) {
	// This is a placeholder for the test function.
	// You can implement your test logic here.
	t.Log("Create Comment test executed")
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
	commentDto := reqdto.CreateCommentDTO{
		Content:   "This is a test comment.",
		ArticleId: 1, // Assuming article with ID 1 exists
	}
	body, _ := json.Marshal(commentDto)
	req := httptest.NewRequest("POST", "/api/blog/comments", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTI3NzYxMjAsInVzZXJJZCI6MSwidXNlck5hbWUiOiJ0ZXN0dXNlciJ9.API1e6nBQwr3vTQBkNXivpFlbo7bkRRYQkyj0Ll2jUo")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}

func TestGetCommentById(t *testing.T) {
	// This is a placeholder for the test function.
	// You can implement your test logic here.
	t.Log("Get Comment By ID test executed")
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
	req := httptest.NewRequest("GET", "/api/blog/comments/1", nil) // Assuming comment with ID 1 exists
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}

func TestUpdateComment(t *testing.T) {
	// This is a placeholder for the test function.
	// You can implement your test logic here.
	t.Log("Update Comment test executed")
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
	commentDto := reqdto.CreateCommentDTO{
		Id:        1, // Assuming comment with ID 1 exists
		Content:   "Updated comment content.",
		ArticleId: 1, // Assuming article with ID 1 exists
	}
	body, _ := json.Marshal(commentDto)
	req := httptest.NewRequest("PUT", "/api/blog/comments/1", bytes.NewBuffer(body)) // Assuming comment with ID 1 exists
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTI3NzYxMjAsInVzZXJJZCI6MSwidXNlck5hbWUiOiJ0ZXN0dXNlciJ9.API1e6nBQwr3vTQBkNXivpFlbo7bkRRYQkyj0Ll2jUo")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}

func TestDeleteComment(t *testing.T) {
	// This is a placeholder for the test function.
	// You can implement your test logic here.
	t.Log("Delete Comment test executed")
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
	req := httptest.NewRequest("DELETE", "/api/blog/comments/1", nil) // Assuming comment with ID 1 exists
	req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTI3NzYxMjAsInVzZXJJZCI6MSwidXNlck5hbWUiOiJ0ZXN0dXNlciJ9.API1e6nBQwr3vTQBkNXivpFlbo7bkRRYQkyj0Ll2jUo")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}

func TestGetCommentsByArticleId(t *testing.T) {
	// This is a placeholder for the test function.
	// You can implement your test logic here.
	t.Log("Get Comments By Article ID test executed")
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
	req := httptest.NewRequest("GET", "/api/blog/articles/1/comments", nil) // Assuming article with ID 1 exists
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}

func TestGetAllComments(t *testing.T) {
	// This is a placeholder for the test function.
	// You can implement your test logic here.
	t.Log("Get All Comments test executed")
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
	req := httptest.NewRequest("GET", "/api/blog/comments", nil) // Assuming this endpoint exists
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}

func TestGetCommentsByUserId(t *testing.T) {
	// This is a placeholder for the test function.
	// You can implement your test logic here.
	t.Log("Get Comments By User ID test executed")
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
	req := httptest.NewRequest("GET", "/api/blog/comments/user/1", nil) // Assuming user with ID 1 exists
	req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTI3NzYxMjAsInVzZXJJZCI6MSwidXNlck5hbWUiOiJ0ZXN0dXNlciJ9.API1e6nBQwr3vTQBkNXivpFlbo7bkRRYQkyj0Ll2jUo")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}
