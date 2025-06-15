package domain

import "context"

type TorrentRepository interface {
	Search(ctx context.Context, query string) ([]Torrent, error)
}
