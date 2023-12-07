package repository

import (
	"context"
	"reflect"
	"short_link_servise/internal/entity"

	"github.com/golang/mock/gomock"
)

type MockLinkRepository struct {
	ctrl     *gomock.Controller
	recorder *MockLinkRepositoryMockRecorder
}

type MockLinkRepositoryMockRecorder struct {
	mock *MockLinkRepository
}

func NewMockLinkRepository(ctrl *gomock.Controller) *MockLinkRepository {
	mock := &MockLinkRepository{ctrl: ctrl}
	mock.recorder = &MockLinkRepositoryMockRecorder{mock}
	return mock
}

func (m *MockLinkRepository) EXPECT() *MockLinkRepositoryMockRecorder {
	return m.recorder
}

func (m *MockLinkRepository) Get(ctx context.Context, link *entity.ShortLink) (*entity.Link, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, link)
	ret0, _ := ret[0].(*entity.Link)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLinkRepositoryMockRecorder) Get(ctx context.Context, link *entity.ShortLink) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockLinkRepository)(nil).Get), ctx, link)
}

func (m *MockLinkRepository) Create(ctx context.Context, link *entity.LinkCreate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, link)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockLinkRepositoryMockRecorder) Create(ctx context.Context, link *entity.LinkCreate) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockLinkRepository)(nil).Create), ctx, link)
}
