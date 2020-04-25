package actions

import "mangagram/models"

// AvailableFeeds defines information for
// all current available manga feeds.
var AvailableFeeds = []models.MangaFeed{
	{
		Code: 1,
		Name: "MangaReader",
		URL:  "https://mangareader.pw",
	},
	{
		Code: 2,
		Name: "Manganelo",
		URL:  "https://manganelo.com",
	},
	{
		Code: 3,
		Name: "Manga Eden",
		URL:  "https://mangaeden.com",
	},
	{
		Code: 4,
		Name: "Kissmanga",
		URL:  "https://kissmanga.org",
	},
	{
		Code: 5,
		Name: "Mangadex",
		URL:  "https://mangadex.org",
	},
}
