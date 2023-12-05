package servise

import (
	"fmt"
	"short_link_servise/internal/entity"
	"short_link_servise/internal/repository"
	"short_link_servise/internal/servise"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	type fields struct {
		repo *repository.MockLinkRepository
	}

	type args struct {
		link *entity.ShortLink
	}

	tests := []struct {
		name    string
		args    args
		setup   func(a args, f fields)
		want    *entity.Link
		wantErr bool
	}{
		{
			name: "Successful call of serviseLinks.Get()",
			args: args{
				link: &entity.ShortLink{
					Link: "vjka91njL_",
				},
			},
			setup: func(a args, f fields) {
				f.repo.EXPECT().Get(a.link).
					Return(&entity.Link{
						Link: "http://localhost",
					}, nil)
			},
			want: &entity.Link{
				Link: "http://localhost",
			},
			wantErr: false,
		},
		{
			name: "Error in repositoryLinks.Get()",
			args: args{
				link: &entity.ShortLink{
					Link: "vjka91njL_",
				},
			},
			setup: func(a args, f fields) {
				f.repo.EXPECT().Get(a.link).Return(nil, fmt.Errorf("Error in repositoryLinks.Get()"))
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
				repo: repository.NewMockLinkRepository(ctrl),
			}

			linkRepo := servise.NewLinksServise(f.repo)

			tt.setup(tt.args, f)

			got, err := linkRepo.Get(tt.args.link)
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
		repo *repository.MockLinkRepository
	}

	type args struct {
		links *entity.LinkCreate
	}

	tests := []struct {
		name    string
		args    args
		setup   func(a args, f fields)
		wantErr bool
	}{
		{
			name: "Successful call of repositoryLinks.Create()",
			args: args{
				links: &entity.LinkCreate{
					Link:      "http://localhost",
					ShortLink: "vjka91njL_",
				},
			},
			setup: func(a args, f fields) {
				f.repo.EXPECT().Create(a.links).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "Error in repositoryLinks.Get()",
			args: args{
				links: &entity.LinkCreate{
					Link:      "http://localhost",
					ShortLink: "vjka91njL_",
				},
			},
			setup: func(a args, f fields) {
				f.repo.EXPECT().Create(a.links).Return(fmt.Errorf("Error in repositoryLinks.Get()"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			f := fields{
				repo: repository.NewMockLinkRepository(ctrl),
			}

			linkRepo := servise.NewLinksServise(f.repo)

			tt.setup(tt.args, f)

			err := linkRepo.Create(tt.args.links)
			if tt.wantErr == true {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
