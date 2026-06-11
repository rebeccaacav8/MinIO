package cmd

import (
	"net/url"
	"strings"
)

// getCanonicalURI returns the canonical URI for the request.
// It ensures that consecutive spaces and other special characters are preserved
// by using the raw escaped path from the URL.
func getCanonicalURI(u *url.URL) string {
	// Use EscapedPath() to get the original encoded path from the request,
	// which preserves consecutive spaces (e.g., %20%20) as expected by AWS SigV4.
	path := u.EscapedPath()
	if path == "" {
		path = "/"
	}
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	return path
}