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

func TestRegister(t *testing.T) {
	// This is a placeholder for the test function.
	// You can implement your test logic here.
	t.Log("Register test executed")
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
	userDto := reqdto.CreateUserDTO{
		Username: "testuser",
		Password: "testpassword",
		Email:    "513469155@qq.com",
	}
	body, _ := json.Marshal(userDto)
	req := httptest.NewRequest("POST", "/api/blog/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}

}

func TestLogin(t *testing.T) {
	// This is a placeholder for the test function.
	// You can implement your test logic here.
	t.Log("Login test executed")
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
	userDto := reqdto.CreateUserDTO{
		Username: "testuser",
		Password: "testpassword",
	}
	body, _ := json.Marshal(userDto)
	req := httptest.NewRequest("POST", "/api/blog/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d, body: %s", w.Code, w.Body.String())
	}
	t.Logf("Login response: %s", w.Body.String())
}
