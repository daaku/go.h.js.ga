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
		`<script>(function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;`+
			`i[r]=i[r]||function(){(i[r].q=i[r].q||[]).push(arguments)},`+
			`i[r].l=1*new Date();a=s.createElement(o),`+
			`m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;`+
			`m.parentNode.insertBefore(a,m)})(window,document,'script',`+
			`'//www.google-analytics.com/analytics.js','ga');`+
			`ga('create','a42','auto');ga('send','pageview');</script>`)
}
