package businesslogic

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"

	"github.com/NicholasMantovani/rssaggregator/internal/models"
)

func UrlToFeed(url string) (models.RSSFeed, error) {
	httpCLient := http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := httpCLient.Get(url)
	if err != nil {
		return models.RSSFeed{}, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.RSSFeed{}, err
	}

	rssFeed := models.RSSFeed{}

	err = xml.Unmarshal(data, &rssFeed)
	if err != nil {
		return models.RSSFeed{}, err
	}

	return rssFeed, err
}
