package handlers

import (
	"context"
	"fmt"
	"short_link_servise/internal/entity"
	"short_link_servise/internal/handlers"
	"short_link_servise/internal/servise"
	"testing"

	pb "short_link_servise/internal/proto"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	type fields struct {
		serv *servise.MockLinkServise
	}

	type args struct {
		ctx     context.Context
		request *pb.GetRequest
	}
	ctx := context.Background()

	tests := []struct {
		name    string
		args    args
		setup   func(a args, f fields)
		want    *pb.GetResponse
		wantErr bool
	}{
		{
			name: "Successful call of HandlersLinks.Get()",
			args: args{
				ctx: ctx,
				request: &pb.GetRequest{
					ShortUrl: "vjka91njL_",
				},
			},
			setup: func(a args, f fields) {
				link := &entity.ShortLink{Link: a.request.ShortUrl}
				f.serv.EXPECT().Get(a.ctx, link).
					Return(&entity.Link{
						Link: "http://localhost",
					}, nil)
			},
			want: &pb.GetResponse{
				Url: "http://localhost",
			},
			wantErr: false,
		},
		{
			name: "Error in repositoryLinks.Get()",
			args: args{
				ctx: ctx,
				request: &pb.GetRequest{
					ShortUrl: "vjka91njL_",
				},
			},
			setup: func(a args, f fields) {
				link := &entity.ShortLink{Link: a.request.ShortUrl}
				f.serv.EXPECT().Get(a.ctx, link).Return(nil, fmt.Errorf("Error in repositoryLinks.Get()"))
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			f := fields{
				serv: servise.NewMockLinkServise(ctrl),
			}

			linkservise := handlers.NewLinksHandler(f.serv)

			tt.setup(tt.args, f)

			got, err := linkservise.Get(tt.args.ctx, tt.args.request)
			if tt.wantErr == true {
				assert.Error(t, err)
			} else {
				if assert.NoError(t, err) {
					assert.Equal(t, tt.want, got)
				}
			}
		})
	}
}

func TestCreate(t *testing.T) {
	type fields struct {
		serv *servise.MockLinkServise
	}

	type args struct {
		ctx     context.Context
		request *pb.CreateRequest
	}
	ctx := context.Background()

	tests := []struct {
		name    string
		args    args
		setup   func(a args, f fields)
		want    *pb.CreateResponse
		wantErr bool
	}{
		{
			name: "Successful call of repositoryLinks.Create()",
			args: args{
				ctx: ctx,
				request: &pb.CreateRequest{
					Url: "http://localhost",
				},
			},
			setup: func(a args, f fields) {
				links := &entity.LinkCreate{Link: a.request.Url, ShortLink: "vjka91njL_"}
				f.serv.EXPECT().Create(a.ctx, links).Return(nil)
			},
			want:    &pb.CreateResponse{},
			wantErr: false,
		},
		{
			name: "Error in repositoryLinks.Get()",
			args: args{
				request: &pb.CreateRequest{
					Url: "http://localhost",
				},
			},
			setup: func(a args, f fields) {
				links := &entity.LinkCreate{Link: a.request.Url, ShortLink: "vjka91njL_"}
				f.serv.EXPECT().Create(a.ctx, links).Return(fmt.Errorf("Error in repositoryLinks.Get()"))
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			f := fields{
				serv: servise.NewMockLinkServise(ctrl),
			}

			linkHandler := handlers.NewLinksHandler(f.serv)

			tt.setup(tt.args, f)

			got, err := linkHandler.Create(tt.args.ctx, tt.args.request)
			if tt.wantErr == true {
				assert.Error(t, err)
			} else {
				if assert.NoError(t, err) {
					assert.Equal(t, tt.want, got)
				}
			}
		})
	}
}
