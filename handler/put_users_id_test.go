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
	"gorm.io/gorm"
)

func TestPutUsersIdSuccess(t *testing.T) {

	expected := `{"code":200,"message":"success","data":null}`

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
		SelectUsersById(gomock.Any(), 123).
		Times(1).Return(user, nil)

	repo.EXPECT().
		SelectUsersByPhoneNumber(gomock.Any(), "+62812283910041").
		Times(1).Return(models.User{}, nil)

	repo.EXPECT().
		UpdateUsers(gomock.Any(), gomock.Any()).
		Times(1).Return(models.User{}, nil)

	serv := handler.NewServer(repo)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	urlValues := url.Values{}
	urlValues.Add("phone_number", "+62812283910041")
	urlValues.Add("full_name", "Rivaldy")
	req.PostForm = urlValues
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("user-id", 123)

	err := serv.PutUsersId(c, 123)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, expected+"\n", rec.Body.String())

}

func TestPutUsersIdFailedBecauseUnexpectedErrorWhenUpdate(t *testing.T) {

	expected := `{"code":500,"message":"unexpected error","data":null}`
	expectedErr := errors.New("unexpected error")

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
		SelectUsersById(gomock.Any(), 123).
		Times(1).Return(user, nil)

	repo.EXPECT().
		SelectUsersByPhoneNumber(gomock.Any(), "+62812283910041").
		Times(1).Return(models.User{}, nil)

	repo.EXPECT().
		UpdateUsers(gomock.Any(), gomock.Any()).
		Times(1).Return(models.User{}, expectedErr)

	serv := handler.NewServer(repo)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	urlValues := url.Values{}
	urlValues.Add("phone_number", "+62812283910041")
	urlValues.Add("full_name", "Rivaldy")
	req.PostForm = urlValues
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("user-id", 123)

	err := serv.PutUsersId(c, 123)

	assert.Equal(t, expectedErr, err)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.Equal(t, expected+"\n", rec.Body.String())

}

func TestPutUsersIdFailedBecauseConflict(t *testing.T) {

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
		SelectUsersById(gomock.Any(), 123).
		Times(1).Return(user, nil)

	user.Id = 321 // trigger conflict
	repo.EXPECT().
		SelectUsersByPhoneNumber(gomock.Any(), "+62812283910041").
		Times(1).Return(user, nil)

	serv := handler.NewServer(repo)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	urlValues := url.Values{}
	urlValues.Add("phone_number", "+62812283910041")
	urlValues.Add("full_name", "Rivaldy")
	req.PostForm = urlValues
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("user-id", 123)

	err := serv.PutUsersId(c, 123)

	assert.Equal(t, expectedErr, err)
	assert.Equal(t, http.StatusConflict, rec.Code)
	assert.Equal(t, expected+"\n", rec.Body.String())

}

func TestPutUsersIdFailedBecauseUnexpectedErrorWhenSelectByPhone(t *testing.T) {

	expected := `{"code":500,"message":"unexpected error","data":null}`
	expectedErr := errors.New("unexpected error")

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
		SelectUsersById(gomock.Any(), 123).
		Times(1).Return(user, nil)

	repo.EXPECT().
		SelectUsersByPhoneNumber(gomock.Any(), "+62812283910041").
		Times(1).Return(models.User{}, expectedErr)

	serv := handler.NewServer(repo)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	urlValues := url.Values{}
	urlValues.Add("phone_number", "+62812283910041")
	urlValues.Add("full_name", "Rivaldy")
	req.PostForm = urlValues
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("user-id", 123)

	err := serv.PutUsersId(c, 123)

	assert.Equal(t, expectedErr, err)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.Equal(t, expected+"\n", rec.Body.String())

}

func TestPutUsersIdFailedBecauseUnexpectedErrorWhenSelectById(t *testing.T) {

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
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	urlValues := url.Values{}
	urlValues.Add("phone_number", "+62812283910041")
	urlValues.Add("full_name", "Rivaldy")
	req.PostForm = urlValues
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("user-id", 123)

	err := serv.PutUsersId(c, 123)

	assert.Equal(t, expectedErr, err)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.Equal(t, expected+"\n", rec.Body.String())

}

func TestPutUsersIdFailedBecauseUserNotFound(t *testing.T) {

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
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	urlValues := url.Values{}
	urlValues.Add("phone_number", "+62812283910041")
	urlValues.Add("full_name", "Rivaldy")
	req.PostForm = urlValues
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("user-id", 123)

	err := serv.PutUsersId(c, 123)

	assert.Equal(t, expectedErr, err)
	assert.Equal(t, http.StatusForbidden, rec.Code)
	assert.Equal(t, expected+"\n", rec.Body.String())

}

func TestPutUsersIdFailedBecauseUnexpectedParams(t *testing.T) {

	expected := `{"code":400,"message":"phone_number format must be valid, full_name must be at least 3 characters in length","data":null}`
	expectedErr := errors.New("phone_number format must be valid, full_name must be at least 3 characters in length")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := repository.NewMockRepositoryInterface(ctrl)
	serv := handler.NewServer(repo)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	urlValues := url.Values{}
	urlValues.Add("phone_number", "+6281228391004123121")
	urlValues.Add("full_name", "dy")
	req.PostForm = urlValues
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("user-id", 123)

	err := serv.PutUsersId(c, 123)

	assert.Equal(t, expectedErr, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, expected+"\n", rec.Body.String())

}

func TestPutUsersIdFailedBecauseUserTryToUpdateAnotherUser(t *testing.T) {

	expected := `{"code":403,"message":"forbidden","data":null}`
	expectedErr := errors.New(strings.ToLower(http.StatusText(http.StatusForbidden)))

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := repository.NewMockRepositoryInterface(ctrl)
	serv := handler.NewServer(repo)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	urlValues := url.Values{}
	urlValues.Add("phone_number", "+6281228391004123121")
	urlValues.Add("full_name", "dy")
	req.PostForm = urlValues
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("user-id", 321)

	err := serv.PutUsersId(c, 123)

	assert.Equal(t, expectedErr, err)
	assert.Equal(t, http.StatusForbidden, rec.Code)
	assert.Equal(t, expected+"\n", rec.Body.String())

}
