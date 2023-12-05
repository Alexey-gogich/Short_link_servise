package repository

import (
	"fmt"
	"short_link_servise/internal/entity"
	"short_link_servise/internal/repository"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	type fields struct {
		db sqlmock.Sqlmock
	}

	type args struct {
		shortLink *entity.ShortLink
	}

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
				shortLink: &entity.ShortLink{
					Link: "vjka91njL_",
				},
			},
			setup: func(a args, f fields) {
				rows := sqlmock.
					NewRows([]string{
						"url",
					}).
					AddRow(
						"http://localhost",
					)
				f.db.ExpectQuery("Select url from links where short_url = $1").WithArgs(a.shortLink.Link).WillReturnRows(rows)
			},
			want: &entity.Link{
				Link: "http://localhost",
			},
			wantErr: false,
		},
		{
			name: "Error in query repositoryLinks.Get()",
			args: args{
				shortLink: &entity.ShortLink{
					Link: "vjka91njL_",
				},
			},
			setup: func(a args, f fields) {
				f.db.ExpectQuery("Select url from links where short_url = $1").WithArgs(a.shortLink.Link).WillReturnError(fmt.Errorf("Not found"))
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			database, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			assert.NoError(t, err)
			defer database.Close()

			f := fields{
				db: mock,
			}

			linkRepo := repository.NewLinksRepo(sqlx.NewDb(database, "sqlmock"))

			tt.setup(tt.args, f)

			got, err := linkRepo.Get(tt.args.shortLink)
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
		db sqlmock.Sqlmock
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
				f.db.ExpectExec("Insert into links (url, short_url) values ($1, $2)").WithArgs(a.links.Link, a.links.ShortLink).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			wantErr: false,
		},
		{
			name: "Error in query repositoryLinks.Get()",
			args: args{
				links: &entity.LinkCreate{
					Link:      "http://localhost",
					ShortLink: "vjka91njL_",
				},
			},
			setup: func(a args, f fields) {
				f.db.ExpectExec("Insert into links (url, short_url) values ($1, $2)").WithArgs(a.links.Link, a.links.ShortLink).
					WillReturnError(fmt.Errorf("Exec error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			database, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			assert.NoError(t, err)
			defer database.Close()

			f := fields{
				db: mock,
			}

			linkRepo := repository.NewLinksRepo(sqlx.NewDb(database, "sqlmock"))

			tt.setup(tt.args, f)

			err = linkRepo.Create(tt.args.links)
			if tt.wantErr == true {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
