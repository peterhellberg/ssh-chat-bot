package hi

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewScraper(t *testing.T) {
	s := NewScraper("foobar")

	if s.URL != "https://twitter.com/hashtag/foobar?f=images" {
		t.Errorf(`unexpected url %q`, s.URL)
	}
}

func TestFindImages(t *testing.T) {
	for _, tt := range []struct {
		body     []byte
		count    int
		imageURL string
	}{
		{twoImagesBody, 2, "bar"},
		{sixImagesBody, 6, "corge"},
		{noImagesBody, 0, ""},
	} {
		server, scraper := testServerAndScraper(tt.body)
		defer server.Close()

		images, err := scraper.FindImages()

		if err != nil {
			t.Fatalf(`unexpected error: %v`, err)
		}

		if got := len(images); got != tt.count {
			t.Errorf(`unexpected number of images: %d, want %d`, got, tt.count)
		}

		if tt.count > 0 {
			if img := images[len(images)-1]; img.URL != tt.imageURL {
				t.Errorf(`unexpected image URL: %q, want %q`, img.URL, tt.imageURL)
			}
		}
	}
}

func TestFindImage(t *testing.T) {
	for _, tt := range []struct {
		body     []byte
		err      error
		imageURL string
	}{
		{twoImagesBody, nil, "foo"},
		{sixImagesBody, nil, "foo"},
		{noImagesBody, ErrNoImagesFound, ""},
	} {
		server, scraper := testServerAndScraper(tt.body)
		defer server.Close()

		img, err := scraper.FindImage()

		if err != tt.err {
			t.Fatalf(`unexpected error: %v`, err)
		}

		if img.URL != tt.imageURL {
			t.Errorf(`unexpected image URL: %q, want %q`, img.URL, tt.imageURL)
		}
	}
}

func TestImage(t *testing.T) {
	server, scraper := testServerAndScraper(twoImagesBody)
	defer server.Close()

	img, err := scraper.FindImage()

	if err != nil {
		t.Fatalf(`unexpected error: %v`, err)
	}

	var (
		imageURL           = "foo"
		imageItemID        = "123"
		imageTweetID       = "456"
		imageHeight        = "100"
		imageWidth         = "200"
		imageName          = "Full Name"
		imageScreenName    = "screen-name"
		imagePermalinkPath = "/permalink/path/"
	)

	if img.URL != imageURL {
		t.Errorf(`unexpected image URL: %q, want %q`, img.URL, imageURL)
	}

	if img.ItemID != imageItemID {
		t.Errorf(`unexpected image ItemID: %q, want %q`, img.ItemID, imageItemID)
	}

	if img.TweetID != imageTweetID {
		t.Errorf(`unexpected image TweetID: %q, want %q`, img.TweetID, imageTweetID)
	}

	if img.Height != imageHeight {
		t.Errorf(`unexpected image Height: %q, want %q`, img.Height, imageHeight)
	}

	if img.Width != imageWidth {
		t.Errorf(`unexpected image Width: %q, want %q`, img.Width, imageWidth)
	}

	if img.Name != imageName {
		t.Errorf(`unexpected image Name: %q, want %q`, img.Name, imageName)
	}

	if img.ScreenName != imageScreenName {
		t.Errorf(`unexpected image ScreenName: %q, want %q`, img.ScreenName, imageScreenName)
	}

	if img.PermalinkPath != imagePermalinkPath {
		t.Errorf(`unexpected image PermalinkPath: %q, want %q`, img.PermalinkPath, imagePermalinkPath)
	}
}

func TestFindShuffledImage(t *testing.T) {
	for _, tt := range []struct {
		body []byte
		err  error
	}{
		{twoImagesBody, nil},
		{sixImagesBody, nil},
		{noImagesBody, ErrNoImagesFound},
	} {
		server, scraper := testServerAndScraper(tt.body)
		defer server.Close()

		img, err := scraper.FindShuffledImage()

		if err != tt.err {
			t.Fatalf(`unexpected error: %v`, err)
		}

		if err == nil && img.URL == "" {
			t.Errorf(`no image url`)
		}
	}
}

func TestSingleImage_withError(t *testing.T) {
	testErr := errors.New("test error for singleImage")

	if _, err := singleImage(nil, testErr); err != testErr {
		t.Fatalf(`unexpected error: %v`, err)
	}
}

func TestShuffledImages_withError(t *testing.T) {
	testErr := errors.New("test error for shuffledImages")

	if _, err := shuffledImages(nil, testErr); err != testErr {
		t.Fatalf(`unexpected error: %v`, err)
	}
}

func testServerAndScraper(body []byte) (*httptest.Server, *Scraper) {
	server := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Write(body)
		},
	))

	return server, &Scraper{URL: server.URL}
}

var noImagesBody = []byte(`
<html>
	<body>
	</body>
</html>`)

var twoImagesBody = []byte(`
<html>
	<body>
		<span class="AdaptiveStreamGridImage" data-url="foo" data-item-id="123" data-tweet-id="456" data-height="100" data-width="200" data-name="Full Name" data-screen-name="screen-name" data-permalink-path="/permalink/path/"></span>
		<span class="AdaptiveStreamGridImage" data-url="bar"></span>
	</body>
</html>`)

var sixImagesBody = []byte(`
<html>
	<body>
		<span class="AdaptiveStreamGridImage" data-url="foo"></span>
		<span class="AdaptiveStreamGridImage" data-url="bar"></span>
		<span class="AdaptiveStreamGridImage" data-url="baz"></span>
		<span class="AdaptiveStreamGridImage" data-url="qux"></span>
		<span class="AdaptiveStreamGridImage" data-url="quux"></span>
		<span class="AdaptiveStreamGridImage" data-url="corge"></span>
	</body>
</html>`)
