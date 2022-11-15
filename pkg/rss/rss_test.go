// Copyright 2022 Nickolas Kraus. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package rss

import (
	"encoding/xml"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

var TESTDIR = filepath.Join("..", "..", "test")

func TestRSS(t *testing.T) {
	t.Run("test RSS - rss-0.xml", func(t *testing.T) {
		data, _ := os.ReadFile(filepath.Join(TESTDIR, "data", "rss-0.xml"))
		ret := &RSS{}
		err := xml.Unmarshal(data, ret)
		exp := RSS{
			XMLName: xml.Name{Space: "", Local: "rss"},
			Version: "2.0",
			Channel: &Channel{
				XMLName:     xml.Name{Space: "", Local: "channel"},
				Title:       "GoUpstate.com News Headlines",
				Link:        "http://www.goupstate.com",
				Description: "The latest news from GoUpstate.com, a Spartanburg Herald-Journal Web site.",
			},
		}
		assert.Nil(t, err)
		assert.Equal(t, exp, *ret)
	})
}

func TestVersionIsValid(t *testing.T) {
	t.Run("check version is valid - pass", func(t *testing.T) {
		data := []byte(`<rss version="2.0"></rss>`)
		ret := &RSS{}
		err := xml.Unmarshal(data, ret)
		assert.Nil(t, err)
		assert.True(t, ret.Version.isValid())
	})
	t.Run("check version is valid - fail", func(t *testing.T) {
		data := []byte(`<rss version="1.0"></rss>`)
		ret := &RSS{}
		err := xml.Unmarshal(data, ret)
		assert.Nil(t, err)
		assert.True(t, ret.Version.isValid())
	})
}

func TestChannelIsValid(t *testing.T) {
	t.Run("check channel is valid - pass", func(t *testing.T) {
		data := []byte(`<rss version="2.0"></rss>`)
		ret := &RSS{}
		err := xml.Unmarshal(data, ret)
		assert.Nil(t, err)
		assert.True(t, ret.Version.isValid())
	})
	t.Run("check channel is valid - fail", func(t *testing.T) {
		data := []byte(`<rss version="1.0"></rss>`)
		ret := &RSS{}
		err := xml.Unmarshal(data, ret)
		assert.Nil(t, err)
		assert.True(t, ret.Version.isValid())
	})
}

func TestRSSIsValid(t *testing.T) {
	t.Run("check rss is valid", func(t *testing.T) {
		data := []byte(`<?xml version="1.0" encoding="UTF-8"?><rss version="2.0"></rss>`)
		ret := &RSS{}
		err := xml.Unmarshal(data, ret)
		assert.Nil(t, err)
		assert.True(t, ret.isValid())
		// // Fail: RSS.Version == nil
		// data = []byte(`<?xml version="1.0" encoding="UTF-8"?><rss></rss>`)
		// ret = &RSS{}
		// err = xml.Unmarshal(data, ret)
		// assert.Nil(t, err)
		// assert.False(t, ret.isValid())
		// // Fail: RSS.Version != "2.0"
		// data = []byte(`<?xml version="1.0" encoding="UTF-8"?><rss version="1.0"></rss>`)
		// ret = &RSS{}
		// err = xml.Unmarshal(data, ret)
		// assert.Nil(t, err)
		// assert.False(t, ret.isValid())
	})
}
