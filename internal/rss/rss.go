package rss

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
)

func (c *Client) FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {

	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating HTTP GET request to %s: %w", feedURL, err)
	}

	req.Header.Set("User-Agent", "gator")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error makiing HTTP GET request to %s: %w", feedURL, err)
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading HTTP response body: %w", err)
	}

	feedResponse := RSSFeed{}
	err = xml.Unmarshal(data, &feedResponse)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling HTTP response body to XML RSSFeed struct: %w", err)
	}

	feedResponse.Channel.Title = html.UnescapeString(feedResponse.Channel.Title)
	feedResponse.Channel.Description = html.UnescapeString(feedResponse.Channel.Description)

	for i := 0; len(feedResponse.Channel.Item) > i; i++ {
		feedResponse.Channel.Item[i].Title = html.UnescapeString(feedResponse.Channel.Item[i].Title)
		feedResponse.Channel.Item[i].Description = html.UnescapeString(feedResponse.Channel.Item[i].Description)
	}

	return &feedResponse, nil
}
