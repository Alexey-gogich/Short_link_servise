package inmemory

import (
	"context"
	"fmt"
	"short_link_servise/internal/entity"
	"short_link_servise/internal/repository/inmemory"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	type fields struct {
		db *inmemory.MockCache
	}

	type args struct {
		ctx       context.Context
		shortLink *entity.ShortLink
	}
	ctx := context.Background()

	tests := []struct {
		name    string
		args    args
		setup   func(a args, f fields)
		want    *entity.Link
		wantErr bool
	}{
		{
			name: "Successful call of repositoryLinks.Get()",
			args: args{
				ctx: ctx,
				shortLink: &entity.ShortLink{
					Link: "vjka91njL_",
				},
			},
			setup: func(a args, f fields) {
				f.db.EXPECT().Get(a.shortLink).Return("http://localhost", nil)
			},
			want: &entity.Link{
				Link: "http://localhost",
			},
			wantErr: false,
		},
		{
			name: "Error in query repositoryLinks.Get()",
			args: args{
				ctx: ctx,
				shortLink: &entity.ShortLink{
					Link: "vjka91njL_",
				},
			},
			setup: func(a args, f fields) {
				f.db.EXPECT().Get(a.shortLink).Return("", fmt.Errorf("Error in cache database"))
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
				db: inmemory.NewMockCache(ctrl),
			}

			linkRepo := inmemory.NewLinksRepo(f.db)

			tt.setup(tt.args, f)

			got, err := linkRepo.Get(tt.args.ctx, tt.args.shortLink)
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
		db *inmemory.MockCache
	}

	type args struct {
		ctx   context.Context
		links *entity.LinkCreate
	}
	ctx := context.Background()

	tests := []struct {
		name    string
		args    args
		setup   func(a args, f fields)
		wantErr bool
	}{
		{
			name: "Successful call of repositoryLinks.Create()",
			args: args{
				ctx: ctx,
				links: &entity.LinkCreate{
					Link:      "http://localhost",
					ShortLink: "vjka91njL_",
				},
			},
			setup: func(a args, f fields) {
				f.db.EXPECT().Insert(a.links).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "Error in query repositoryLinks.Get()",
			args: args{
				ctx: ctx,
				links: &entity.LinkCreate{
					Link:      "http://localhost",
					ShortLink: "vjka91njL_",
				},
			},
			setup: func(a args, f fields) {
				f.db.EXPECT().Insert(a.links).Return(fmt.Errorf("Error in cache database"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				db: inmemory.NewMockCache(ctrl),
			}

			linkRepo := inmemory.NewLinksRepo(f.db)

			tt.setup(tt.args, f)

			err := linkRepo.Create(tt.args.ctx, tt.args.links)
			if tt.wantErr == true {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
