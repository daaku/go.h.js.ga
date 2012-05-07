// Package ga provides go.h compatible async loading for
// Google Analytics.
package ga

import (
	"fmt"
	"log"
)

// Loadable for a Page Track event using Google Analytics.
type Track struct {
	ID string
}

func (g *Track) URLs() []string {
	return []string{"https://ssl.google-analytics.com/ga.js"}
}

func (g *Track) Script() string {
	if g.ID == "" {
		log.Fatal("GoogleAnalyics requires an ID.")
	}
	return fmt.Sprintf(
		`try{`+
			`var pageTracker=_gat._getTracker("%s");`+
			`pageTracker._trackPageview();`+
			`}catch(e){}`, g.ID)
}
