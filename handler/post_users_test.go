package handler_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/SawitProRecruitment/UserService/handler"
	"github.com/SawitProRecruitment/UserService/models"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestPostUsersSuccess(t *testing.T) {

	expected := `{"code":201,"message":"success","data":{"id":123,"full_name":"Rival","phone_number":"+6287789312891"}}`

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := repository.NewMockRepositoryInterface(ctrl)

	user := models.User{
		Id:           123,
		FullName:     "Rival",
		PhoneNumber:  "+6287789312891",
		PasswordHash: "$2a$10$MXWbKFotaINb8seF5ybpPu1V43r4MxPXtjseSzsJbUbT8q7qAT2j2",
		CreatedBy:    123,
		CreatedAt:    time.Now(),
	}

	repo.EXPECT().
		SelectUsersByPhoneNumber(gomock.Any(), "+62812283910041").
		Times(1).Return(models.User{}, nil)

	repo.EXPECT().
		InsertUsers(gomock.Any(), gomock.Any()).
		Times(1).Return(user, nil)

	serv := handler.NewServer(repo)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	urlValues := url.Values{}
	urlValues.Add("phone_number", "+62812283910041")
	urlValues.Add("full_name", "Rivaldy")
	urlValues.Add("password", "hahi37#A")
	req.PostForm = urlValues
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := serv.PostUsers(c)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusCreated, rec.Code)
	assert.Equal(t, expected+"\n", rec.Body.String())

}

func TestPostUsersFailedBecauseUnexpectedErrorWhenInsert(t *testing.T) {

	expected := `{"code":500,"message":"unexpected error","data":null}`
	expectedErr := errors.New("unexpected error")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := repository.NewMockRepositoryInterface(ctrl)

	repo.EXPECT().
		SelectUsersByPhoneNumber(gomock.Any(), "+62812283910041").
		Times(1).Return(models.User{}, nil)

	repo.EXPECT().
		InsertUsers(gomock.Any(), gomock.Any()).
		Times(1).Return(models.User{}, expectedErr)

	serv := handler.NewServer(repo)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	urlValues := url.Values{}
	urlValues.Add("phone_number", "+62812283910041")
	urlValues.Add("full_name", "Rivaldy")
	urlValues.Add("password", "hahi37#A")
	req.PostForm = urlValues
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := serv.PostUsers(c)

	assert.Equal(t, expectedErr, err)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.Equal(t, expected+"\n", rec.Body.String())

}

func TestPostUsersFailedBecauseConflictPhoneNumber(t *testing.T) {

	expected := `{"code":409,"message":"conflict","data":null}`
	expectedErr := errors.New(strings.ToLower(http.StatusText(http.StatusConflict)))

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := repository.NewMockRepositoryInterface(ctrl)

	user := models.User{
		Id:           123,
		FullName:     "Rival",
		PhoneNumber:  "+62812283910041",
		PasswordHash: "$2a$10$MXWbKFotaINb8seF5ybpPu1V43r4MxPXtjseSzsJbUbT8q7qAT2j2",
		CreatedBy:    123,
		CreatedAt:    time.Now(),
	}

	repo.EXPECT().
		SelectUsersByPhoneNumber(gomock.Any(), "+62812283910041").
		Times(1).Return(user, nil)

	serv := handler.NewServer(repo)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	urlValues := url.Values{}
	urlValues.Add("phone_number", "+62812283910041")
	urlValues.Add("full_name", "Rivaldy")
	urlValues.Add("password", "hahi37#A")
	req.PostForm = urlValues
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := serv.PostUsers(c)

	assert.Equal(t, expectedErr, err)
	assert.Equal(t, http.StatusConflict, rec.Code)
	assert.Equal(t, expected+"\n", rec.Body.String())

}

func TestPostUsersFailedBecauseUnexpectedErrorWhenSelect(t *testing.T) {

	expected := `{"code":500,"message":"unexpected error","data":null}`
	expectedErr := errors.New("unexpected error")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := repository.NewMockRepositoryInterface(ctrl)

	repo.EXPECT().
		SelectUsersByPhoneNumber(gomock.Any(), "+62812283910041").
		Times(1).Return(models.User{}, expectedErr)

	serv := handler.NewServer(repo)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	urlValues := url.Values{}
	urlValues.Add("phone_number", "+62812283910041")
	urlValues.Add("full_name", "Rivaldy")
	urlValues.Add("password", "hahi37#A")
	req.PostForm = urlValues
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := serv.PostUsers(c)

	assert.Equal(t, expectedErr, err)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.Equal(t, expected+"\n", rec.Body.String())

}

func TestPostUsersFailedBecauseUnexpectedParams(t *testing.T) {

	expected := `{"code":400,"message":"phone_number format must be valid, full_name must be at least 3 characters in length, password format must be valid","data":null}`
	expectedErr := errors.New("phone_number format must be valid, full_name must be at least 3 characters in length, password format must be valid")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := repository.NewMockRepositoryInterface(ctrl)
	serv := handler.NewServer(repo)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	urlValues := url.Values{}
	urlValues.Add("phone_number", "+6281228391004121312")
	urlValues.Add("full_name", "Ri")
	urlValues.Add("password", "hahi37A")
	req.PostForm = urlValues
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := serv.PostUsers(c)

	assert.Equal(t, expectedErr, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, expected+"\n", rec.Body.String())

}
