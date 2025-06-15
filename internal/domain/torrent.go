package domain

type Torrent struct {
	Title       string `json:"title"`
	Size        string `json:"size"`
	DownloadURL string `json:"downloadUrl"`
	Seeders     int    `json:"seeders"`
	Peers       int    `json:"peers"`
	Indexer     string `json:"indexer"`
}
