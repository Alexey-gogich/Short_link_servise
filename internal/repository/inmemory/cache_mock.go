package inmemory

import (
	"reflect"
	"short_link_servise/internal/entity"

	"github.com/golang/mock/gomock"
)

type MockCache struct {
	ctrl     *gomock.Controller
	recorder *MockCacheMockRecorder
}

type MockCacheMockRecorder struct {
	mock *MockCache
}

func NewMockCache(ctrl *gomock.Controller) *MockCache {
	mock := &MockCache{ctrl: ctrl}
	mock.recorder = &MockCacheMockRecorder{mock}
	return mock
}

func (m *MockCache) EXPECT() *MockCacheMockRecorder {
	return m.recorder
}

func (m *MockCache) Get(link *entity.ShortLink) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", link)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCacheMockRecorder) Get(link *entity.ShortLink) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCache)(nil).Get), link)
}

func (m *MockCache) Insert(links *entity.LinkCreate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", links)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockCacheMockRecorder) Insert(links *entity.LinkCreate) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockCache)(nil).Insert), links)
}
