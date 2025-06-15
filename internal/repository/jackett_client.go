package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/dixon961/go-movie-downloader-backend/internal/domain"
)

type jackettResponse struct {
	Results []jackettResult `json:"Results"`
}

type jackettResult struct {
	Title     string `json:"Title"`
	Size      int64  `json:"Size"`
	MagnetURI string `json:"MagnetUri"`
	Seeders   int    `json:"Seeders"`
	Peers     int    `json:"Peers"`
	Tracker   string `json:"Tracker"`
}

type JackettClient struct {
	httpClient *http.Client
	baseURL    string
	apiKey     string
}

func NewJackettClient(httpClient *http.Client, baseURL, apiKey string) *JackettClient {
	return &JackettClient{
		httpClient: httpClient,
		baseURL:    baseURL,
		apiKey:     apiKey,
	}
}

func (c *JackettClient) Search(ctx context.Context, query string) ([]domain.Torrent, error) {
	fullURL, err := url.Parse(fmt.Sprintf("%s/all/results", c.baseURL))
	if err != nil {
		return nil, fmt.Errorf("invalid base url: %w", err)
	}

	params := url.Values{}
	params.Add("apikey", c.apiKey)
	params.Add("Query", query)
	fullURL.RawQuery = params.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fullURL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("jackett api returned non-200 status: %d", resp.StatusCode)
	}

	var jackettResp jackettResponse
	if err := json.NewDecoder(resp.Body).Decode(&jackettResp); err != nil {
		return nil, fmt.Errorf("failed to decode jackett response: %w", err)
	}

	torrents := make([]domain.Torrent, 0, len(jackettResp.Results))
	for _, r := range jackettResp.Results {
		torrents = append(torrents, domain.Torrent{
			Title:       r.Title,
			Size:        formatSize(r.Size),
			DownloadURL: r.MagnetURI,
			Seeders:     r.Seeders,
			Peers:       r.Peers,
			Indexer:     r.Tracker,
		})
	}

	return torrents, nil
}

func formatSize(bytes int64) string {
	if bytes == 0 {
		return "0 B"
	}
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %ciB", float64(bytes)/float64(div), "KMGTPE"[exp])
}
