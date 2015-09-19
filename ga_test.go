package ga

import (
	"testing"

	"github.com/daaku/go.h"
	"github.com/facebookgo/ensure"

	"golang.org/x/net/context"
)

func TestTrackMissingAccount(t *testing.T) {
	h, err := (&Track{}).HTML(context.Background())
	ensure.True(t, err == errMissingAccount)
	ensure.Nil(t, h)
}

func TestTrack(t *testing.T) {
	s, err := h.Render(context.Background(), &Track{Account: "a42"})
	ensure.Nil(t, err)
	ensure.DeepEqual(t, s,
		`<script>var _gaq = _gaq || [];_gaq.push(['_setAccount', 'a42']);`+
			`_gaq.push(['_trackPageview']);</script><script src="`+
			`https://ssl.google-analytics.com/ga.js" async></script>`)
}
