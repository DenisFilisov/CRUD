package handler

import (
	"CRUD/pkg/model"
	"CRUD/pkg/service"
	mock_service "CRUD/pkg/service/mocks"
	"bytes"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"net/http/httptest"
	"testing"
)

func TestHandler_singUp(t *testing.T) {
	type mockBehavior func(s *mock_service.MockAuthorisation, user model.User)

	testTable := []struct {
		name                string
		inputBody           string
		inputUser           model.User
		mockBehavior        mockBehavior
		expectedStatus      int
		expectedRequestBody string
	}{
		{name: "OK",
			inputBody: `{"name": "Test", "username":"testUsername", "password":"testPassword"}`,
			inputUser: model.User{
				Name:     "Test",
				Username: "testUsername",
				Password: "testPassword",
			},
			mockBehavior: func(s *mock_service.MockAuthorisation, user model.User) {
				s.EXPECT().CreateUser(user).Return(1, nil)
			},
			expectedStatus:      201,
			expectedRequestBody: `{"id":1}`,
		}, {name: "Empty fields",
			inputBody: `{"username":"testUsername", "password":"testPassword"}`,
			inputUser: model.User{
				Name:     "Test",
				Username: "testUsername",
				Password: "testPassword",
			},
			mockBehavior: func(s *mock_service.MockAuthorisation, user model.User) {
			},
			expectedStatus:      400,
			expectedRequestBody: `{"message":"Key: 'User.Name' Error:Field validation for 'Name' failed on the 'required' tag"}`,
		},
		{name: "Service Failure",
			inputBody: `{"name": "Test", "username":"testUsername", "password":"testPassword"}`,
			inputUser: model.User{
				Name:     "Test",
				Username: "testUsername",
				Password: "testPassword",
			},
			mockBehavior: func(s *mock_service.MockAuthorisation, user model.User) {
				s.EXPECT().CreateUser(user).Return(1, errors.New("service failure"))
			},
			expectedStatus:      500,
			expectedRequestBody: `{"message":"service failure"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			auth := mock_service.NewMockAuthorisation(c)
			testCase.mockBehavior(auth, testCase.inputUser)
			services := &service.Service{Authorisation: auth}
			handler := NewHandler(services)

			//Test service
			r := gin.New()
			r.POST("/sing-up", handler.singUp)

			//Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sing-up", bytes.NewBufferString(testCase.inputBody))

			// Perform Request
			r.ServeHTTP(w, req)

			//	Assert
			assert.Equal(t, w.Code, testCase.expectedStatus)
			assert.Equal(t, w.Body.String(), testCase.expectedRequestBody)
		})
	}
}

func TestHandler_singIn(t *testing.T) {
	type mockBehavior func(s *mock_service.MockAuthorisation, userName, password string)

	testTable := []struct {
		name                string
		inputBody           string
		username            string
		password            string
		mockBehavior        mockBehavior
		expectedStatus      int
		expectedRequestBody string
	}{
		{name: "OK",
			inputBody: `{"username":"testUsername", "password":"testPassword"}`,
			username:  "testUsername",
			password:  "testPassword",
			mockBehavior: func(s *mock_service.MockAuthorisation, userName, password string) {
				s.EXPECT().FindUserByUsernameAndPswd(userName, password).Return(model.User{
					Name:     "MockUser",
					Username: "testUsername",
					Password: "testPassword",
				}, nil)
				s.EXPECT().GenerateTokens("", model.User{
					Name:     "MockUser",
					Username: "testUsername",
					Password: "testPassword",
				}).Return("newToken", "oldToken", nil)
				s.EXPECT().ParseToken("newToken").Return(1, int64(1), nil)
			},
			expectedStatus:      200,
			expectedRequestBody: `{"token":"newToken"}`,
		},
		{name: "missing fields",
			inputBody: `{"password":"testPassword"}`,
			password:  "testPassword",
			mockBehavior: func(s *mock_service.MockAuthorisation, userName, password string) {
			},
			expectedStatus:      400,
			expectedRequestBody: `{"message":"Key: 'singIn.Username' Error:Field validation for 'Username' failed on the 'required' tag"}`,
		},
		{name: "Problem with generate Token",
			inputBody: `{"username":"testUsername", "password":"testPassword"}`,
			username:  "testUsername",
			password:  "testPassword",
			mockBehavior: func(s *mock_service.MockAuthorisation, userName, password string) {
				s.EXPECT().FindUserByUsernameAndPswd(userName, password).Return(model.User{
					Name:     "MockUser",
					Username: "testUsername",
					Password: "testPassword",
				}, nil)
				s.EXPECT().GenerateTokens("", model.User{
					Name:     "MockUser",
					Username: "testUsername",
					Password: "testPassword",
				}).Return("", "", errors.New("Problem with generate token"))
			},
			expectedStatus:      500,
			expectedRequestBody: `{"message":"Problem with generate token"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			auth := mock_service.NewMockAuthorisation(c)
			testCase.mockBehavior(auth, testCase.username, testCase.password)
			services := &service.Service{Authorisation: auth}
			handler := NewHandler(services)

			//Test service
			r := gin.New()
			r.POST("/sing-in", handler.singIn)

			//Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sing-in", bytes.NewBufferString(testCase.inputBody))

			// Perform Request
			r.ServeHTTP(w, req)

			//	Assert
			assert.Equal(t, w.Code, testCase.expectedStatus)
			assert.Equal(t, w.Body.String(), testCase.expectedRequestBody)
		})
	}
}

func TestHandler_refreshToken(t *testing.T) {
	type mockBehavior func(s *mock_service.MockAuthorisation)

	testTable := []struct {
		name                string
		refreshToken        string
		mockBehavior        mockBehavior
		expectedStatus      int
		expectedRequestBody string
	}{
		{name: "OK",
			refreshToken: "refresh-token=af74d412b9ffe987d3e86527e8db5639a600450c97c5796cce71e38a171b07a5",
			mockBehavior: func(s *mock_service.MockAuthorisation) {
				s.EXPECT().RefreshToken("af74d412b9ffe987d3e86527e8db5639a600450c97c5796cce71e38a171b07a5").Return("newToken", "oldToken", nil)
			},
			expectedStatus:      200,
			expectedRequestBody: `{"token":"newToken"}`,
		},
		{name: "Wrong refreshToken in Cookies",
			refreshToken: "",
			mockBehavior: func(s *mock_service.MockAuthorisation) {
			},
			expectedStatus:      400,
			expectedRequestBody: `{"message":"Wrong refreshToken in Cookies"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			auth := mock_service.NewMockAuthorisation(c)
			testCase.mockBehavior(auth)
			services := &service.Service{Authorisation: auth}
			handler := NewHandler(services)

			//Test service
			r := gin.New()
			r.POST("/refresh-token", handler.refreshToken)

			//Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/refresh-token", nil)
			req.Header.Set("Cookie", testCase.refreshToken)

			// Perform Request
			r.ServeHTTP(w, req)

			//	Assert
			assert.Equal(t, w.Code, testCase.expectedStatus)
			assert.Equal(t, w.Body.String(), testCase.expectedRequestBody)
		})
	}
}
