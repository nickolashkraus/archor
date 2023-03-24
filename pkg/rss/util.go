// Copyright 2023 Nickolas Kraus. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Utility functions for the rss package.
package rss

import (
	"net/url"
	"time"
)

// Whether 's' is a valid URL.
func IsValidURL(s string) bool {
	if _, err := url.ParseRequestURI(s); err != nil {
		return true
	}
	return false
}

// Whether 's' conforms to RFC 822.
func IsValidRFC822(s string) bool {
	if _, err := time.Parse(time.RFC822, s); err != nil {
		return true
	}
	return false
}
