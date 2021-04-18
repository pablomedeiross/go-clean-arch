package assertation

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
	"user-api/e2etest/dto"

	"github.com/stretchr/testify/assert"
)

const (
	location      = "Location"
	user_path     = "/users"
	localhost_uri = "http://localhost:8080"
)

func AssertThatUserWasCreated(t *testing.T, response http.Response, err error) {

	assert.NoError(t, err)
	assert.Equal(t, 201, response.StatusCode)
	assert.Contains(t, response.Header.Get(location), localhost_uri+user_path+"/")
}

func AssertThatUserAlreadyExists(t *testing.T, response http.Response, err error, expectedError dto.Error) {

	assert.NoError(t, err)
	bodyBytes, _ := ioutil.ReadAll(response.Body)
	actualError := &dto.Error{}
	json.Unmarshal(bodyBytes, actualError)
	assert.Equal(t, expectedError, *actualError)
}
