// Copyright 2022 Nickolas Kraus. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// The rss package provides utilities for handling RSS feeds. It leverages the
// RSS 2.0 Specification for determining struct fields used during marshaling
// and unmarshaling of RSS documents.
//
// See:
//   - https://validator.w3.org/feed/docs/rss2.html
package rss

import (
	"encoding/xml"
	"net/url"
	"reflect"
	"strconv"
	"time"
)

const RSSVERSION = "2.0"

type RSSElement interface {
	isValid() bool
}

// At the top level, a RSS document is a <rss> element, with a mandatory
// attribute called version, that specifies the version of RSS that the
// document conforms to.
//
// Subordinate to the <rss> element is a single <channel> element, which
// contains information about the channel (metadata) and its contents.
//
// The XMLName field name dictates the name of the XML element representing
// this struct.
//
// See: https://validator.w3.org/feed/docs/rss2.html#whatIsRss
type RSS struct {
	XMLName xml.Name `xml:"rss"`          // required
	Version Version  `xml:"version,attr"` // required
	Channel *Channel `xml:"channel"`      // required
}

type Version string

// Whether the RSS document version is valid.
//
// The <rss> element must have attribute 'version' with value '2.0'.
func (r Version) isValid() bool {
	return r == RSSVERSION
}

// TODO: Handle nil pointers.
//
// Whether the RSS document is valid.
//
// In order for the RSS document to be valid, it must contain all required
// elements and sub-elements.
//
// If the document contains optional sub-elements with required elements, these
// too must be valid.
//
// To accomplish this effectively, we recurse through all struct fields. If the
// struct field is of interface type RSSElement, the isValid method is called.
func (r RSS) isValid() bool {
	// ValueOf returns a new Value initialized to the concrete value
	// stored in the interface i. ValueOf(nil) returns the zero Value.
	v := reflect.ValueOf(r)
	// NumField returns the number of fields in the struct v.
	// It panics if v's Kind is not Struct.
	for i := 0; i < v.NumField(); i++ {
		// Field returns the i'th field of the struct v.
		// It panics if v's Kind is not Struct or i is out of range.
		//
		// Interface returns v's current value as an interface{}.
		// It is equivalent to:
		//
		//   var i interface{} = (v's underlying value)
		//
		// It panics if the Value was obtained by accessing
		// unexported struct fields.
		//
		// To test whether an interface value holds a specific type, a type
		// assertion can return two values: the underlying value and a boolean
		// value that reports whether the assertion succeeded.
		//
		//   t, ok := i.(T)
		//
		// If i holds a T, then t will be the underlying value and ok will be true.
		//
		// If not, ok will be false and t will be the zero value of type T, and no
		// panic occurs.
		if t, ok := v.Field(i).Interface().(RSSElement); ok {
			if t == nil || !t.isValid() {
				return false
			}
		}
	}
	return true
}

// The channel element has three required elements:
//   - title
//   - link
//   - description
//
// A list of optional channel elements can be found here:
//   - https://validator.w3.org/feed/docs/rss2.html#optionalChannelElements
type Channel struct {
	XMLName        xml.Name      `xml:"channel"`        // required
	Title          Title         `xml:"title"`          // required
	Link           Link          `xml:"link"`           // required
	Description    Description   `xml:"description"`    // required
	Language       Language      `xml:"language"`       // optional
	Copyright      string        `xml:"copyright"`      // optional
	ManagingEditor string        `xml:"managingEditor"` // optional
	WebMaster      string        `xml:"webMaster"`      // optional
	PubDate        PubDate       `xml:"pubDate"`        // optional
	LastBuildDate  LastBuildDate `xml:"lastBuildDate"`  // optional
	Category       Category      `xml:"category"`       // optional
	Generator      string        `xml:"generator"`      // optional
	Docs           string        `xml:"docs"`           // optional
	Cloud          string        `xml:"cloud"`          // optional
	TTL            TTL           `xml:"ttl"`            // optional
	Image          Image         `xml:"image"`          // optional
	Rating         string        `xml:"rating"`         // optional
	TextInput      TextInput     `xml:"textInput"`      // optional
	SkipHours      string        `xml:"skipHours"`      // optional
	SkipDays       string        `xml:"skipDays"`       // optional
	Item           []*Item       `xml:"item"`           // optional
}

type Title string

// Whether the RSS document channel element 'title' is valid.
//
// The <title> element should not be blank.
func (r Title) isValid() bool {
	return r != ""
}

type Link string

// Whether the RSS document channel element 'link' is valid.
//
// The <link> element must be a full and valid URL.
func (r Link) isValid() bool {
	if _, err := url.ParseRequestURI(string(r)); err != nil {
		return true
	}
	return false
}

type Description string

// Whether the RSS document channel element 'description' is valid.
//
// The <description> element must exist.
//
// NOTE: This method serves only implement the RSSElement interface.
func (r Description) isValid() bool { return true }

type Language string

// Whether the RSS document channel element 'language' is valid.
//
// The <language> element must be one of the identifiers specified in the
// current list of ISO 639 language codes:
//
// See:
//   - https://www.rssboard.org/rss-language-codes
//   - https://www.loc.gov/standards/iso639-2
func (r Language) isValid() bool { return true }

type PubDate string

// Whether the RSS document channel element 'pubDate' is valid.
//
// The <pubDate> element must conform to the Date and Time Specification of
// RFC 822, with the exception that the year may be expressed with two
// characters or four characters (four preferred).
func (r PubDate) isValid() bool { return true }

type LastBuildDate string

// Whether the RSS document channel element 'lastBuildDate' is valid.
//
// The <lastBuildDate> element must conform to the Date and Time Specification of
// RFC 822, with the exception that the year may be expressed with two
// characters or four characters (four preferred).
func (r LastBuildDate) isValid() bool { return true }

// Whether 'value' conforms to RFC 822.
func IsValidRFC822(value string) bool {
	if _, err := time.Parse(time.RFC822, value); err != nil {
		return true
	}
	return false
}

type Category string

// Whether the RSS document channel element is valid.
func (c Channel) isValid() bool {
	return c.Title != "" && c.Link != "" && c.Description != ""
}

// The <image> element is an optional sub-element of <channel>, which contains
// three required and three optional sub-elements.
//
// See: https://validator.w3.org/feed/docs/rss2.html#ltimagegtSubelementOfLtchannelgt
type Image struct {
	Url         string `xml:"url"`         // required
	Title       string `xml:"title"`       // required
	Link        string `xml:"link"`        // required
	Width       string `xml:"width"`       // optional
	Height      string `xml:"height"`      // optional
	Description string `xml:"description"` // optional
}

func (r Image) isValid() bool {
	// Required sub-elements: <url>, <title>, <link>
	//
	// NOTE: In practice the image <title> and <link> should have the same value
	// as the channel's <title> and <link>, but this is not a requirement.
	if r.Url == "" || r.Title == "" || r.Link == "" {
		return false
	}
	// Optional sub-elements: <width>, <height>, <description>
	//
	// The maximum value for width is 144, default value is 88.
	// The maximum value for height is 400, default value is 31.
	if i, err := strconv.ParseUint(r.Width, 10, 0); err != nil || i > 144 {
		return false
	}
	if i, err := strconv.ParseUint(r.Height, 10, 0); err != nil || i > 400 {
		return false
	}
	return true
}

// The <cloud> is an optional sub-element of <channel>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#ltcloudgtSubelementOfLtchannelgt
type Cloud struct {
	Domain            string `xml:"domain"`            // required
	Port              string `xml:"port"`              // required
	Path              string `xml:"path"`              // required
	RegisterProcedure string `xml:"registerProcedure"` // required
	Protocol          string `xml:"protocol"`          // required
}

func (r Cloud) isValid() bool { return true }

type TTL int

func (r TTL) isValid() bool {
	// TTL must be a positive integer.
	return r > 0
}

// The <textInput> is an optional sub-element of <channel>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#lttextinputgtSubelementOfLtchannelgt
type TextInput struct {
	Title       string `xml:"title"`       // required
	Description string `xml:"description"` // required
	Name        string `xml:"name"`        // required
	Link        string `xml:"link"`        // required
}

func (r TextInput) isValid() bool {
	return r.Title != "" && r.Description != "" && r.Name != "" && r.Link != ""
}

// All elements of an item are optional, however at least one of title or
// description must be present.
type Item struct {
	Title       string    `xml:"title"`       // conditionally required
	Link        string    `xml:"link"`        // conditionally required
	Description string    `xml:"description"` // optional
	Author      string    `xml:"author"`      // optional
	Category    string    `xml:"category"`    // optional
	Comments    string    `xml:"comments"`    // optional
	Enclosure   Enclosure `xml:"enclosure"`   // optional
	Guid        string    `xml:"guid"`        // optional
	PubDate     string    `xml:"pubDate"`     // optional
	Source      string    `xml:"source"`      // optional
}

type Enclosure struct {
	URL    string `xml:"url,attr"`    // required
	Length int64  `xml:"length,attr"` // required
	Type   string `xml:"type,attr"`   // required
}
