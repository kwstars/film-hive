package service

import (
	"context"
	"flag"
	"os"
	"testing"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/kwstars/film-hive/api/rating/service/v1"
	"github.com/kwstars/film-hive/app/rating/service/internal/conf"
	"github.com/stretchr/testify/assert"
)

var (
	ratingService *RatingService
	flagconf      = flag.String("conf", "../../configs", "config path, eg: -conf config.yaml")
)

func TestMain(m *testing.M) {
	flag.Parse()
	var (
		bc      conf.Bootstrap
		closeFn func()
		err     error
	)
	c := config.New(
		config.WithSource(
			file.NewSource(*flagconf),
		),
	)
	if err = c.Load(); err != nil {
		panic(err)
	}
	if err = c.Scan(&bc); err != nil {
		panic(err)
	}
	if ratingService, closeFn, err = createRatingService(&bc, log.DefaultLogger); err != nil {
		log.Fatal(err)
	}
	defer closeFn()

	log.Log(log.LevelInfo, "start testing")
	exitVal := m.Run()
	log.Log(log.LevelInfo, "stop testing")

	os.Exit(exitVal)
}

func TestRatingService(t *testing.T) {
	t.Run("CreateRating", func(t *testing.T) {
		type testCase struct {
			name          string
			input         *v1.CreateRatingRequest
			expectedError error
		}
		tests := []testCase{
			{"Success", &v1.CreateRatingRequest{Rating: &v1.Rating{RecordType: 1, RecordId: 1, UserId: 1234, RatingValue: 9}}, nil},
			{"Repeat", &v1.CreateRatingRequest{Rating: &v1.Rating{RecordType: 1, RecordId: 1, UserId: 1234, RatingValue: 9}}, nil},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				_, err := ratingService.CreateRating(context.Background(), tt.input)
				assert.Equal(t, tt.expectedError, err)
			})
		}
	})

	t.Run("GetAggregatedRating", func(t *testing.T) {
		req := &v1.GetAggregatedRatingRequest{RecordType: 1, RecordId: 1}
		if resp, err := ratingService.GetAggregatedRating(context.Background(), req); err != nil {
			t.Errorf("unexpected result: %v (expected nil)", err)
		} else {
			assert.Equal(t, float64(9), resp.GetAvgRating())
		}

		req = &v1.GetAggregatedRatingRequest{RecordType: 1, RecordId: 100}
		if _, err := ratingService.GetAggregatedRating(context.Background(), req); err != nil && !v1.IsRatingNotFound(err) {
			t.Errorf("unexpected result: %v (expected RatingNotFound)", err)
		}
	})
}
