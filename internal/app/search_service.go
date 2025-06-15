package app

import (
	"context"
	"sort"

	"github.com/dixon961/go-movie-downloader-backend/internal/domain"
)

type SearchService struct {
	repo domain.TorrentRepository
}

func NewSearchService(repo domain.TorrentRepository) *SearchService {
	return &SearchService{repo: repo}
}

func (s *SearchService) SearchTorrents(ctx context.Context, query string) ([]domain.Torrent, error) {
	results, err := s.repo.Search(ctx, query)
	if err != nil {
		return nil, err
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].Seeders > results[j].Seeders
	})

	if len(results) > 25 {
		return results[:25], nil
	}

	return results, nil
}
