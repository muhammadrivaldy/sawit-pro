package repository_test

import (
	"context"
	"database/sql/driver"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/SawitProRecruitment/UserService/models"
	"github.com/SawitProRecruitment/UserService/repository"
	_ "github.com/lib/pq"
	goutil "github.com/muhammadrivaldy/go-util"
	"github.com/stretchr/testify/assert"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

var (
	mock sqlmock.Sqlmock
	repo repository.Repository
)

func init() {

	sqlDb, sqlMock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}

	gormDb, err := goutil.NewGorm(sqlDb, "postgres", goutil.LoggerSilent)
	if err != nil {
		panic(err)
	}

	repo = repository.Repository{gormDb}
	mock = sqlMock

}

func TestInsertUsersSuccess(t *testing.T) {

	user := models.User{
		Id:           123,
		FullName:     "Rival",
		PhoneNumber:  "+6287789312891",
		PasswordHash: "password-hashed",
		CreatedBy:    123,
		CreatedAt:    time.Now(),
	}

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "mst_users" ("full_name","phone_number","password_hash","created_by","created_at","updated_by","updated_at","deleted_at","id") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING "id"`)).
		WithArgs(user.FullName, user.PhoneNumber, user.PasswordHash, user.CreatedBy, user.CreatedAt, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), user.Id).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	_, err := repo.InsertUsers(context.TODO(), user)
	if err != nil {
		t.Fatal(err)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Fatal(err)
	}

}

func TestInsertUsersFailed(t *testing.T) {

	user := models.User{
		Id:           123,
		FullName:     "Rival",
		PhoneNumber:  "+6287789312891",
		PasswordHash: "password-hashed",
		CreatedBy:    123,
		CreatedAt:    time.Now(),
	}

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "mst_users" ("full_name","phone_number","password_hash","created_by","created_at","updated_by","updated_at","deleted_at","id") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING "id"`)).
		WithArgs(user.FullName, user.PhoneNumber, user.PasswordHash, user.CreatedBy, user.CreatedAt, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), user.Id).
		WillReturnError(gorm.ErrRecordNotFound)
	mock.ExpectRollback()

	_, err := repo.InsertUsers(context.TODO(), user)
	assert.Equal(t, gorm.ErrRecordNotFound, err)

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Fatal(err)
	}

}

func TestUpdateUsersSuccess(t *testing.T) {

	user := models.User{
		Id:        123,
		FullName:  "Raynaldy",
		UpdatedBy: null.IntFrom(123),
		UpdatedAt: null.TimeFrom(time.Now()),
	}

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "mst_users" SET "full_name"=$1,"updated_by"=$2,"updated_at"=$3 WHERE "id" = $4`)).
		WithArgs(user.FullName, user.UpdatedBy, sqlmock.AnyArg(), user.Id).
		WillReturnResult(driver.RowsAffected(1))
	mock.ExpectCommit()

	_, err := repo.UpdateUsers(context.TODO(), user)
	if err != nil {
		t.Fatal(err)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Fatal(err)
	}

}

func TestUpdateUsersFailed(t *testing.T) {

	user := models.User{
		Id:        123,
		FullName:  "Raynaldy",
		UpdatedBy: null.IntFrom(123),
		UpdatedAt: null.TimeFrom(time.Now()),
	}

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "mst_users" SET "full_name"=$1,"updated_by"=$2,"updated_at"=$3 WHERE "id" = $4`)).
		WithArgs(user.FullName, user.UpdatedBy, sqlmock.AnyArg(), user.Id).
		WillReturnError(gorm.ErrRecordNotFound)
	mock.ExpectRollback()

	_, err := repo.UpdateUsers(context.TODO(), user)
	assert.Equal(t, gorm.ErrRecordNotFound, err)

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Fatal(err)
	}

}

func TestSelectUsersByIdSuccess(t *testing.T) {

	timeNow := time.Now()
	expectedReturn := sqlmock.NewRows([]string{"id", "full_name", "phone_number", "password_hash", "created_by", "created_at", "updated_by", "updated_at", "deleted_at"}).
		AddRow(123, "Rival", "+6287789312891", "password-hashed", 123, &timeNow, 123, &timeNow, &timeNow)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "mst_users" WHERE id = $1 ORDER BY "mst_users"."id" LIMIT 1`)).
		WithArgs(123).WillReturnRows(expectedReturn)

	_, err := repo.SelectUsersById(context.TODO(), 123)
	if err != nil {
		t.Fatal(err)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Fatal(err)
	}

}

func TestSelectUsersByIdFailed(t *testing.T) {

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "mst_users" WHERE id = $1 ORDER BY "mst_users"."id" LIMIT 1`)).
		WithArgs(123).WillReturnError(gorm.ErrRecordNotFound)

	_, err := repo.SelectUsersById(context.TODO(), 123)
	assert.Equal(t, gorm.ErrRecordNotFound, err)

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Fatal(err)
	}

}

func TestSelectUsersByPhoneNumberSuccess(t *testing.T) {

	timeNow := time.Now()
	expectedReturn := sqlmock.NewRows([]string{"id", "full_name", "phone_number", "password_hash", "created_by", "created_at", "updated_by", "updated_at", "deleted_at"}).
		AddRow(123, "Rival", "+6287789312891", "password-hashed", 123, &timeNow, 123, &timeNow, &timeNow)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "mst_users" WHERE phone_number = $1 ORDER BY "mst_users"."id" LIMIT 1`)).
		WithArgs("+6287789312891").WillReturnRows(expectedReturn)

	_, err := repo.SelectUsersByPhoneNumber(context.TODO(), "+6287789312891")
	if err != nil {
		t.Fatal(err)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Fatal(err)
	}

}

func TestSelectUsersByPhoneNumberFailed(t *testing.T) {

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "mst_users" WHERE phone_number = $1 ORDER BY "mst_users"."id" LIMIT 1`)).
		WithArgs("+6287789312891").WillReturnError(gorm.ErrRecordNotFound)

	_, err := repo.SelectUsersByPhoneNumber(context.TODO(), "+6287789312891")
	assert.Equal(t, gorm.ErrRecordNotFound, err)

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Fatal(err)
	}

}

func TestInsertSessionsSuccess(t *testing.T) {

	session := models.Session{
		Id:      123,
		UserId:  123,
		LoginAt: time.Now(),
	}

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "trx_sessions" ("user_id","login_at","id") VALUES ($1,$2,$3) RETURNING "id"`)).
		WithArgs(session.UserId, sqlmock.AnyArg(), session.Id).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	_, err := repo.InsertSessions(context.TODO(), session)
	if err != nil {
		t.Fatal(err)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Fatal(err)
	}

}

func TestInsertSessionsFailed(t *testing.T) {

	session := models.Session{
		Id:      123,
		UserId:  123,
		LoginAt: time.Now(),
	}

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "trx_sessions" ("user_id","login_at","id") VALUES ($1,$2,$3) RETURNING "id"`)).
		WithArgs(session.UserId, sqlmock.AnyArg(), session.Id).
		WillReturnError(gorm.ErrRecordNotFound)
	mock.ExpectRollback()

	_, err := repo.InsertSessions(context.TODO(), session)
	assert.Equal(t, gorm.ErrRecordNotFound, err)

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Fatal(err)
	}

}
