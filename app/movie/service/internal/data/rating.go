package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/kwstars/film-hive/api/movie/service/v1"
	"github.com/kwstars/film-hive/app/movie/service/internal/biz"
)

type movieRepo struct {
	data *Data
	log  *log.Helper
}

func (r *movieRepo) ListMovies(ctx context.Context, recordType, recordID string) (movies []biz.Movie, err error) {
	var (
		ok bool
		rs []Movie
	)
	if _, ok = r.data.movie[recordType]; !ok {
		return nil, v1.ErrorMovieNotFound("no record of record_type(%s)", recordType)
	}
	if rs, ok = r.data.movie[recordType][recordID]; !ok || len(rs) == 0 {
		return nil, v1.ErrorMovieNotFound("no record of record_type(%s) and record_id(%s)", recordType, recordID)
	}

	movies = make([]biz.Movie, 0, len(rs))
	for _, movie := range rs {
		t := biz.Movie{
			RecordId:   movie.RecordID,
			RecordType: movie.RecordType,
			UserId:     movie.UserID,
			Value:      movie.Value,
		}
		movies = append(movies, t)
	}
	return
}

func (r *movieRepo) CreateMovie(ctx context.Context, recordType, recordID string, movie *biz.Movie) (err error) {
	t := Movie{
		RecordID:   movie.RecordId,
		RecordType: movie.RecordType,
		UserID:     movie.UserId,
		Value:      movie.Value,
	}
	if _, ok := r.data.movie[recordType]; !ok {
		r.data.movie[recordType] = map[string][]Movie{}
	}
	r.data.movie[recordType][recordID] = append(r.data.movie[recordType][recordID], t)
	return
}

// NewMovieRepo .
func NewMovieRepo(data *Data, logger log.Logger) biz.MovieRepo {
	return &movieRepo{
		data: data,
		log:  log.NewHelper(logger, log.WithMessageKey("movie data")),
	}
}
