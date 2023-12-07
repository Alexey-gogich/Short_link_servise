package servise

import (
	"context"
	"reflect"

	"short_link_servise/internal/entity"

	"github.com/golang/mock/gomock"
)

type MockLinkServise struct {
	ctrl     *gomock.Controller
	recorder *MockLinkServiseMockRecorder
}

type MockLinkServiseMockRecorder struct {
	mock *MockLinkServise
}

func NewMockLinkServise(ctrl *gomock.Controller) *MockLinkServise {
	mock := &MockLinkServise{ctrl: ctrl}
	mock.recorder = &MockLinkServiseMockRecorder{mock}
	return mock
}

func (m *MockLinkServise) EXPECT() *MockLinkServiseMockRecorder {
	return m.recorder
}

func (m *MockLinkServise) Get(ctx context.Context, link *entity.ShortLink) (*entity.Link, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, link)
	ret0, _ := ret[0].(*entity.Link)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLinkServiseMockRecorder) Get(ctx context.Context, link *entity.ShortLink) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockLinkServise)(nil).Get), ctx, link)
}

func (m *MockLinkServise) Create(ctx context.Context, links *entity.LinkCreate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, links)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockLinkServiseMockRecorder) Create(ctx context.Context, links *entity.LinkCreate) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockLinkServise)(nil).Create), ctx, links)
}
