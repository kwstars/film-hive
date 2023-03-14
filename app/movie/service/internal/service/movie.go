package service

import (
	"context"

	v1 "github.com/kwstars/film-hive/api/movie/service/v1"
	"github.com/kwstars/film-hive/app/movie/service/internal/biz"
)

type MovieService struct {
	v1.UnimplementedMovieServiceServer
	uc *biz.MovieUsecase
}

func NewMovieService(uc *biz.MovieUsecase) *MovieService {
	return &MovieService{
		uc: uc,
	}
}

func (m *MovieService) GetMovieDetail(ctx context.Context, req *v1.GetMovieDetailRequest) (resp *v1.GetMovieDetailResponse, err error) {
	var movieDetail *biz.MovieDetail
	movieDetail, err = m.uc.GetMovieDetail(ctx, req.GetId())
	resp = &v1.GetMovieDetailResponse{
		Rating: movieDetail.Rating,
		Metadata: &v1.GetMovieDetailResponse_Metadata{
			Id:          movieDetail.Metadata.ID,
			Title:       movieDetail.Metadata.Title,
			Description: movieDetail.Metadata.Description,
			Director:    movieDetail.Metadata.Director,
		},
	}
	return resp, err
}
