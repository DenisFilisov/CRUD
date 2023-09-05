// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	model "CRUD/pkg/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	logrus "github.com/sirupsen/logrus"
)

// MockAuthorisation is a mock of Authorisation interface.
type MockAuthorisation struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorisationMockRecorder
}

// MockAuthorisationMockRecorder is the mock recorder for MockAuthorisation.
type MockAuthorisationMockRecorder struct {
	mock *MockAuthorisation
}

// NewMockAuthorisation creates a new mock instance.
func NewMockAuthorisation(ctrl *gomock.Controller) *MockAuthorisation {
	mock := &MockAuthorisation{ctrl: ctrl}
	mock.recorder = &MockAuthorisationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthorisation) EXPECT() *MockAuthorisationMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockAuthorisation) CreateUser(user model.User) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", user)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockAuthorisationMockRecorder) CreateUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockAuthorisation)(nil).CreateUser), user)
}

// FindUserByUsernameAndPswd mocks base method.
func (m *MockAuthorisation) FindUserByUsernameAndPswd(username, password string) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByUsernameAndPswd", username, password)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByUsernameAndPswd indicates an expected call of FindUserByUsernameAndPswd.
func (mr *MockAuthorisationMockRecorder) FindUserByUsernameAndPswd(username, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByUsernameAndPswd", reflect.TypeOf((*MockAuthorisation)(nil).FindUserByUsernameAndPswd), username, password)
}

// GenerateTokens mocks base method.
func (m *MockAuthorisation) GenerateTokens(oldToken string, user model.User) (string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateTokens", oldToken, user)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GenerateTokens indicates an expected call of GenerateTokens.
func (mr *MockAuthorisationMockRecorder) GenerateTokens(oldToken, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateTokens", reflect.TypeOf((*MockAuthorisation)(nil).GenerateTokens), oldToken, user)
}

// ParseToken mocks base method.
func (m *MockAuthorisation) ParseToken(token string) (int, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseToken", token)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ParseToken indicates an expected call of ParseToken.
func (mr *MockAuthorisationMockRecorder) ParseToken(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseToken", reflect.TypeOf((*MockAuthorisation)(nil).ParseToken), token)
}

// RefreshToken mocks base method.
func (m *MockAuthorisation) RefreshToken(refreshToken string) (string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshToken", refreshToken)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// RefreshToken indicates an expected call of RefreshToken.
func (mr *MockAuthorisationMockRecorder) RefreshToken(refreshToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshToken", reflect.TypeOf((*MockAuthorisation)(nil).RefreshToken), refreshToken)
}

// MockNews is a mock of News interface.
type MockNews struct {
	ctrl     *gomock.Controller
	recorder *MockNewsMockRecorder
}

// MockNewsMockRecorder is the mock recorder for MockNews.
type MockNewsMockRecorder struct {
	mock *MockNews
}

// NewMockNews creates a new mock instance.
func NewMockNews(ctrl *gomock.Controller) *MockNews {
	mock := &MockNews{ctrl: ctrl}
	mock.recorder = &MockNewsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNews) EXPECT() *MockNewsMockRecorder {
	return m.recorder
}

// FindNewsById mocks base method.
func (m *MockNews) FindNewsById(id int) (model.News, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindNewsById", id)
	ret0, _ := ret[0].(model.News)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindNewsById indicates an expected call of FindNewsById.
func (mr *MockNewsMockRecorder) FindNewsById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindNewsById", reflect.TypeOf((*MockNews)(nil).FindNewsById), id)
}

// GetAllNews mocks base method.
func (m *MockNews) GetAllNews() ([]model.News, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllNews")
	ret0, _ := ret[0].([]model.News)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllNews indicates an expected call of GetAllNews.
func (mr *MockNewsMockRecorder) GetAllNews() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllNews", reflect.TypeOf((*MockNews)(nil).GetAllNews))
}

// PostNews mocks base method.
func (m *MockNews) PostNews(news model.News) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostNews", news)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostNews indicates an expected call of PostNews.
func (mr *MockNewsMockRecorder) PostNews(news interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostNews", reflect.TypeOf((*MockNews)(nil).PostNews), news)
}

// RemoveNews mocks base method.
func (m *MockNews) RemoveNews(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveNews", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveNews indicates an expected call of RemoveNews.
func (mr *MockNewsMockRecorder) RemoveNews(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveNews", reflect.TypeOf((*MockNews)(nil).RemoveNews), id)
}

// UpdateNews mocks base method.
func (m *MockNews) UpdateNews(id int, news string) (model.News, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateNews", id, news)
	ret0, _ := ret[0].(model.News)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateNews indicates an expected call of UpdateNews.
func (mr *MockNewsMockRecorder) UpdateNews(id, news interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateNews", reflect.TypeOf((*MockNews)(nil).UpdateNews), id, news)
}

// MockSubscribers is a mock of Subscribers interface.
type MockSubscribers struct {
	ctrl     *gomock.Controller
	recorder *MockSubscribersMockRecorder
}

// MockSubscribersMockRecorder is the mock recorder for MockSubscribers.
type MockSubscribersMockRecorder struct {
	mock *MockSubscribers
}

// NewMockSubscribers creates a new mock instance.
func NewMockSubscribers(ctrl *gomock.Controller) *MockSubscribers {
	mock := &MockSubscribers{ctrl: ctrl}
	mock.recorder = &MockSubscribersMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSubscribers) EXPECT() *MockSubscribersMockRecorder {
	return m.recorder
}

// GetAllSubscribersByNewsID mocks base method.
func (m *MockSubscribers) GetAllSubscribersByNewsID(id int) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllSubscribersByNewsID", id)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllSubscribersByNewsID indicates an expected call of GetAllSubscribersByNewsID.
func (mr *MockSubscribersMockRecorder) GetAllSubscribersByNewsID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllSubscribersByNewsID", reflect.TypeOf((*MockSubscribers)(nil).GetAllSubscribersByNewsID), id)
}

// SubscribeToNews mocks base method.
func (m *MockSubscribers) SubscribeToNews(userId, newsId int) (model.News, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeToNews", userId, newsId)
	ret0, _ := ret[0].(model.News)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubscribeToNews indicates an expected call of SubscribeToNews.
func (mr *MockSubscribersMockRecorder) SubscribeToNews(userId, newsId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeToNews", reflect.TypeOf((*MockSubscribers)(nil).SubscribeToNews), userId, newsId)
}

// UnsubscribeFromNews mocks base method.
func (m *MockSubscribers) UnsubscribeFromNews(userId any, newsId int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnsubscribeFromNews", userId, newsId)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnsubscribeFromNews indicates an expected call of UnsubscribeFromNews.
func (mr *MockSubscribersMockRecorder) UnsubscribeFromNews(userId, newsId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnsubscribeFromNews", reflect.TypeOf((*MockSubscribers)(nil).UnsubscribeFromNews), userId, newsId)
}

// MockKafka is a mock of Kafka interface.
type MockKafka struct {
	ctrl     *gomock.Controller
	recorder *MockKafkaMockRecorder
}

// MockKafkaMockRecorder is the mock recorder for MockKafka.
type MockKafkaMockRecorder struct {
	mock *MockKafka
}

// NewMockKafka creates a new mock instance.
func NewMockKafka(ctrl *gomock.Controller) *MockKafka {
	mock := &MockKafka{ctrl: ctrl}
	mock.recorder = &MockKafkaMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKafka) EXPECT() *MockKafkaMockRecorder {
	return m.recorder
}

// SentMessage mocks base method.
func (m *MockKafka) SentMessage(message *logrus.Entry) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SentMessage", message)
}

// SentMessage indicates an expected call of SentMessage.
func (mr *MockKafkaMockRecorder) SentMessage(message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SentMessage", reflect.TypeOf((*MockKafka)(nil).SentMessage), message)
}
