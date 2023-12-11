package handler_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/SawitProRecruitment/UserService/handler"
	"github.com/SawitProRecruitment/UserService/models"
	"github.com/SawitProRecruitment/UserService/payloads"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/SawitProRecruitment/UserService/utils"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestPostUsersLoginSuccess(t *testing.T) {

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
		Times(1).Return(user, nil)

	repo.EXPECT().
		InsertSessions(gomock.Any(), gomock.Any()).
		Times(1).Return(models.Session{}, nil)

	serv := handler.NewServer(repo)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	req.SetBasicAuth("+62812283910041", "hahi37#A")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := serv.PostUsersLogin(c)
	if err != nil {
		t.Fatal(err)
	}

	response := payloads.Response{}
	json.Unmarshal(rec.Body.Bytes(), &response)
	responseData := response.Data.(map[string]interface{})
	userId, err := utils.ParseJWT(responseData["token"].(string))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, user.Id, userId)

}

func TestPostUsersLoginSuccessButInsertSessionGotUnexpectedError(t *testing.T) {

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
		Times(1).Return(user, nil)

	repo.EXPECT().
		InsertSessions(gomock.Any(), gomock.Any()).
		Times(1).Return(models.Session{}, errors.New("unexpected error"))

	serv := handler.NewServer(repo)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	req.SetBasicAuth("+62812283910041", "hahi37#A")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := serv.PostUsersLogin(c)
	if err != nil {
		t.Fatal(err)
	}

	response := payloads.Response{}
	json.Unmarshal(rec.Body.Bytes(), &response)
	responseData := response.Data.(map[string]interface{})
	userId, err := utils.ParseJWT(responseData["token"].(string))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, user.Id, userId)

}

func TestPostUsersLoginFailedBecauseUnexpectedParams(t *testing.T) {

	expected := `{"code":400,"message":"phone_number format must be valid, password format must be valid","data":null}`
	expectedErr := errors.New("phone_number format must be valid, password format must be valid")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := repository.NewMockRepositoryInterface(ctrl)
	serv := handler.NewServer(repo)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	req.SetBasicAuth("+6281228391", "hahi37A")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := serv.PostUsersLogin(c)

	assert.Equal(t, expectedErr, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, expected+"\n", rec.Body.String())

}

func TestPostUsersLoginFailedBecauseUserNotFound(t *testing.T) {

	expected := `{"code":400,"message":"your email \u0026 password are wrong","data":null}`

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := repository.NewMockRepositoryInterface(ctrl)

	repo.EXPECT().
		SelectUsersByPhoneNumber(gomock.Any(), "+62812283910041").
		Times(1).Return(models.User{}, gorm.ErrRecordNotFound)

	serv := handler.NewServer(repo)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	req.SetBasicAuth("+62812283910041", "hahi37#A")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := serv.PostUsersLogin(c)

	assert.Equal(t, gorm.ErrRecordNotFound, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, expected+"\n", rec.Body.String())

}

func TestPostUsersLoginFailedBecausePasswordNotMatched(t *testing.T) {

	expected := `{"code":400,"message":"your email \u0026 password are wrong","data":null}`
	expectedErr := errors.New("crypto/bcrypt: hashedPassword is not the hash of the given password")

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
		Times(1).Return(user, nil)

	serv := handler.NewServer(repo)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	req.SetBasicAuth("+62812283910041", "hahi37#A_123")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := serv.PostUsersLogin(c)

	assert.Equal(t, expectedErr, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, expected+"\n", rec.Body.String())

}
