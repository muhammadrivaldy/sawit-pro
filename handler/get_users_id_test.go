package handler_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/SawitProRecruitment/UserService/handler"
	"github.com/SawitProRecruitment/UserService/models"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestGetUsersIdSuccess(t *testing.T) {

	expected := `{"code":200,"message":"success","data":{"id":123,"full_name":"Rival","phone_number":"+6287789312891"}}`

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := repository.NewMockRepositoryInterface(ctrl)

	user := models.User{
		Id:           123,
		FullName:     "Rival",
		PhoneNumber:  "+6287789312891",
		PasswordHash: "password-hashed",
		CreatedBy:    123,
		CreatedAt:    time.Now(),
	}

	repo.EXPECT().
		SelectUsersById(gomock.Any(), 123).
		Times(1).Return(user, nil)

	serv := handler.NewServer(repo)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := serv.GetUsersId(c, 123)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, expected+"\n", rec.Body.String())

}

func TestGetUsersIdFailedBecauseUnexpectedError(t *testing.T) {

	expected := `{"code":500,"message":"unexpected error","data":null}`
	expectedErr := errors.New("unexpected error")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := repository.NewMockRepositoryInterface(ctrl)

	repo.EXPECT().
		SelectUsersById(gomock.Any(), 123).
		Times(1).Return(models.User{}, expectedErr)

	serv := handler.NewServer(repo)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := serv.GetUsersId(c, 123)

	assert.Equal(t, expectedErr, err)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.Equal(t, expected+"\n", rec.Body.String())

}

func TestGetUsersIdFailedBecauseUserNotFound(t *testing.T) {

	expected := `{"code":403,"message":"forbidden","data":null}`
	expectedErr := errors.New(strings.ToLower(http.StatusText(http.StatusForbidden)))

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := repository.NewMockRepositoryInterface(ctrl)

	repo.EXPECT().
		SelectUsersById(gomock.Any(), 123).
		Times(1).Return(models.User{}, gorm.ErrRecordNotFound)

	serv := handler.NewServer(repo)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := serv.GetUsersId(c, 123)

	assert.Equal(t, expectedErr, err)
	assert.Equal(t, http.StatusForbidden, rec.Code)
	assert.Equal(t, expected+"\n", rec.Body.String())

}
