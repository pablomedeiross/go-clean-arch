package db_test

import (
	"context"
	"testing"
	adapter "user-api/adapter/db"
	"user-api/external/db"
	"user-api/external/db/test/assertation"
	"user-api/helper"

	test_helper "user-api/external/db/test/helper"

	"github.com/stretchr/testify/assert"
)

var dbHelper *helper.InMemoryMongoDB
var userTest adapter.User

func init() {

	dbHelper = helper.NewInMemoryMongoDB()

	userTest = adapter.User{
		Name:  "test",
		Email: "email",
		Age:   13,
	}

}

func TestNewDBGateway(t *testing.T) {

	dbHelper.Start()
	defer dbHelper.Stop()

	dbGateway, err := db.NewNoSQLDB(dbHelper.URI(), dbHelper.Name())
	assert.NotNil(t, dbGateway)
	assert.Nil(t, err)
}

func TestNewDBGatewayError(t *testing.T) {

	dbGateway, err := db.NewNoSQLDB("", "")
	assert.Error(t, err)
	assert.Nil(t, dbGateway)
}

func TestSaveUser(t *testing.T) {

	dbHelper.Start()
	defer dbHelper.Stop()

	dbGateway, err := db.NewNoSQLDB(dbHelper.URI(), dbHelper.Name())
	id, err := dbGateway.SaveUser(context.Background(), userTest)

	assertation.AssertThatUserExistsInDB(t, id, dbHelper)
	assert.Nil(t, err)
}

func TestFindUserByName(t *testing.T) {

	err := dbHelper.Start()
	defer dbHelper.Stop()

	err = test_helper.InsertUser(dbHelper, userTest)

	dbGateway, err := db.NewNoSQLDB(dbHelper.URI(), dbHelper.Name())
	usr, err := dbGateway.FindUserByName(context.Background(), userTest.Name)

	assert.Nil(t, err)
	assertation.AssertThatUserExistsInDB(t, usr.Id, dbHelper)
	assertation.AssertThatUserEqualWithouId(t, userTest, usr)
}

func TestFindUserByNameWithoutNoneExistentName(t *testing.T) {

	err := dbHelper.Start()
	defer dbHelper.Stop()

	err = test_helper.InsertUser(dbHelper, userTest)

	dbGateway, err := db.NewNoSQLDB(dbHelper.URI(), dbHelper.Name())
	usr, err := dbGateway.FindUserByName(context.Background(), "nonexistent name")

	assert.Error(t, err)
	assert.Empty(t, usr)
}
