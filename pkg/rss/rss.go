// Copyright 2023 Nickolas Kraus. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// The rss package provides utilities for handling RSS documents. It leverages
// the RSS 2.0 Specification for determining struct fields used during
// marshaling, unmarshaling, and validation.
//
// See:
//   - https://validator.w3.org/feed/docs/rss2.html
package rss

import (
	"encoding/xml"
	"reflect"
	"strconv"
)

const RSSVERSION = "2.0"

// The RSSElement interface specifies a single method, IsValid. IsValid checks
// whether the element conforms to the RSS 2.0 Specification.
type RSSElement interface {
	IsValid() bool
}

// At the top level, a RSS document is a <rss> element, with a mandatory
// attribute called version, that specifies the version of RSS that the
// document conforms to. If it conforms to this specification, the version
// attribute must be 2.0.
//
// Subordinate to the <rss> element is a single <channel> element, which
// contains information about the channel (metadata) and its contents.
//
// NOTE: The XMLName field name dictates the name of the XML element
// representing this struct.
//
// See:
//   - https://validator.w3.org/feed/docs/rss2.html#whatIsRss
type RSS struct {
	XMLName xml.Name `xml:"rss"`          // required
	Version Version  `xml:"version,attr"` // required
	Channel *Channel `xml:"channel"`      // required
}

// version is a required attribute of <rss>.
//
// See:
//   - https://validator.w3.org/feed/docs/rss2.html#whatIsRss
type Version string

// Whether version is valid.
//
// <rss> element must have attribute version with value "2.0".
func (r Version) IsValid() bool {
	return r == RSSVERSION
}

// Whether the RSS document is valid.
//
// In order for the RSS document to be valid, it must comprise all required
// elements and sub-elements.
//
// If the RSS document contains optional sub-elements with required elements,
// these too must be valid.
//
// To accomplish this, we recurse through all struct fields. If the struct
// field is of interface type RSSElement, the IsValid method is called. Each
// RSSElement is responsible for implementing its IsValid method in accordance
// with the RSS 2.0 Specification.
func (r RSS) IsValid() bool {
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
		//	var i interface{} = (v's underlying value)
		//
		// It panics if the Value was obtained by accessing
		// unexported struct fields.
		//
		// To test whether an interface value holds a specific type, a type
		// assertion can return two values: the underlying value and a boolean
		// value that reports whether the assertion succeeded.
		//
		//  t, ok := i.(T)
		//
		// If i holds a T, then t will be the underlying value and ok will be true.
		//
		// If not, ok will be false and t will be the zero value of type T, and no
		// panic occurs.
		if t, ok := v.Field(i).Interface().(RSSElement); ok {
			// Indirect returns the value that v points to.
			// If v is a nil pointer, Indirect returns a zero Value.
			// If v is not a pointer, Indirect returns v.
			v := reflect.Indirect(reflect.ValueOf(t))
			if v.IsNil() || !v.IsValid() {
				return false
			}
		}
	}
	return true
}

// <channel> is a required sub-element of <rss>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#requiredChannelElements
type Channel struct {
	XMLName        xml.Name       `xml:"channel"`        // required
	Title          Title          `xml:"title"`          // required
	Link           Link           `xml:"link"`           // required
	Description    Description    `xml:"description"`    // required
	Language       Language       `xml:"language"`       // optional
	Copyright      Copyright      `xml:"copyright"`      // optional
	ManagingEditor ManagingEditor `xml:"managingEditor"` // optional
	WebMaster      WebMaster      `xml:"webMaster"`      // optional
	PubDate        PubDate        `xml:"pubDate"`        // optional
	LastBuildDate  LastBuildDate  `xml:"lastBuildDate"`  // optional
	Category       Category       `xml:"category"`       // optional
	Generator      Generator      `xml:"generator"`      // optional
	Docs           Docs           `xml:"docs"`           // optional
	Cloud          Cloud          `xml:"cloud"`          // optional
	TTL            TTL            `xml:"ttl"`            // optional
	Image          Image          `xml:"image"`          // optional
	Rating         Rating         `xml:"rating"`         // optional
	TextInput      TextInput      `xml:"textInput"`      // optional
	SkipHours      SkipHours      `xml:"skipHours"`      // optional
	SkipDays       SkipDays       `xml:"skipDays"`       // optional
	Item           []*Item        `xml:"item"`           // optional
}

// <title> is a required sub-element of <channel>, <textInput>, and <item>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#requiredChannelElements
type Title string

// Whether <title> is valid.
func (r Title) IsValid() bool {
	return r != ""
}

// <link> is a required sub-element of <channel>, <image>, <textInput>, and
// <item>.
//
// See:
//   - https://validator.w3.org/feed/docs/rss2.html#requiredChannelElements
//   - https://validator.w3.org/feed/docs/rss2.html#ltimagegtSubelementOfLtchannelgt
//   - https://validator.w3.org/feed/docs/rss2.html#lttextinputgtSubelementOfLtchannelgt
//   - https://validator.w3.org/feed/docs/rss2.html#hrelementsOfLtitemgt
type Link string

// Whether <link> is valid.
func (r Link) IsValid() bool {
	return IsValidURL(string(r))
}

// <description> is a required sub-element of <channel> and <textInput>.
//
// <description> is an optional sub-element of <image> and <item>
//
// See:
//   - https://validator.w3.org/feed/docs/rss2.html#requiredChannelElements
//   - https://validator.w3.org/feed/docs/rss2.html#ltimagegtSubelementOfLtchannelgt
//   - https://validator.w3.org/feed/docs/rss2.html#lttextinputgtSubelementOfLtchannelgt
//   - https://validator.w3.org/feed/docs/rss2.html#hrelementsOfLtitemgt
type Description string

// Whether <descripton> is valid.
func (r Description) IsValid() bool {
	return r != ""
}

// <language> is an optional sub-element of <channel>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#optionalChannelElements
type Language string

// TODO
// Whether <language> is valid.
//
// The <language> element must be one of the identifiers specified in the
// current list of ISO 639 language codes:
//
// See:
//   - https://www.rssboard.org/rss-language-codes
//   - https://www.loc.gov/standards/iso639-2
func (r Language) IsValid() bool { return true }

// <copyright> is an optional sub-element of <channel>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#optionalChannelElements
type Copyright string

// Whether <copyright> is valid.
func (r Copyright) IsValid() bool { return true }

// <managingEditor> is an optional sub-element of <channel>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#optionalChannelElements
type ManagingEditor string

// Whether <managingEditor> is valid.
func (r ManagingEditor) IsValid() bool { return true }

// <webMaster> is an optional sub-element of <channel>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#optionalChannelElements
type WebMaster string

// Whether <webMaster> is valid.
func (r WebMaster) IsValid() bool { return true }

// <pubDate> is an optional sub-element of <channel> and <item>.
//
// See:
//   - https://validator.w3.org/feed/docs/rss2.html#optionalChannelElements
//   - https://validator.w3.org/feed/docs/rss2.html#ltpubdategtSubelementOfLtitemgt
type PubDate string

// Whether <pubDate> is valid.
//
// <pubDate> must conform to the Date and Time Specification of RFC 822, with
// the exception that the year may be expressed with two characters or four
// characters (four preferred).
//
// See: http://asg.web.cmu.edu/rfc/rfc822.html
func (r PubDate) IsValid() bool {
	return IsValidRFC822(string(r))
}

// <lastBuildDate> is an optional sub-element of <channel> and <item>.
type LastBuildDate string

// Whether <lastBuildDate> is valid.
//
// <lastBuildDate> must conform to the Date and Time Specification of RFC 822,
// with the exception that the year may be expressed with two characters or
// four characters (four preferred).
//
// See: http://asg.web.cmu.edu/rfc/rfc822.html
func (r LastBuildDate) IsValid() bool {
	return IsValidRFC822(string(r))
}

// <category> is an optional sub-element of <channel> and <item>.
//
// See:
//   - https://validator.w3.org/feed/docs/rss2.html#optionalChannelElements
//   - https://validator.w3.org/feed/docs/rss2.html#ltcategorygtSubelementOfLtitemgt
type Category struct {
	XMLName xml.Name `xml:"category"`    // required
	Domain  Domain   `xml:"domain,attr"` // optional
}

// Whether <category> is valid.
func (r Category) IsValid() bool {
	if r.Domain != "" {
		return r.Domain.IsValid()
	}
	return true
}

// <domain> is an optional attribute of <category> of <item.
//
// See: https://validator.w3.org/feed/docs/rss2.html#ltcategorygtSubelementOfLtitemgt
type Domain string

// Whether <generator> is valid.
func (r Domain) IsValid() bool { return true }

// <generator> is an optional sub-element of <channel>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#optionalChannelElements
type Generator string

// Whether <generator> is valid.
func (r Generator) IsValid() bool { return true }

// <docs> is an optional sub-element of <channel>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#optionalChannelElements
type Docs string

// Whether <docs> is valid.
func (r Docs) IsValid() bool { return true }

// <cloud> is an optional sub-element of <channel>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#ltcloudgtSubelementOfLtchannelgt
type Cloud struct {
	XMLName           xml.Name `xml:"cloud"`             // required
	Domain            string   `xml:"domain"`            // required
	Port              string   `xml:"port"`              // required
	Path              string   `xml:"path"`              // required
	RegisterProcedure string   `xml:"registerProcedure"` // required
	Protocol          string   `xml:"protocol"`          // required
}

// Whether <cloud> is valid.
//
// TODO: https://www.rssboard.org/rsscloud-interface
func (r Cloud) IsValid() bool { return true }

// <ttl> is an optional sub-element of <channel>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#ltttlgtSubelementOfLtchannelgt
type TTL string

// Whether <ttl> is valid.
//
// TTL must be a positive integer.
func (r TTL) IsValid() bool {
	if i, err := strconv.ParseUint(string(r), 10, 0); err != nil || i < 0 {
		return false
	}
	return true
}

// <image> is an optional sub-element of <channel>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#ltimagegtSubelementOfLtchannelgt
//
// TODO: Set default values for width and height.
type Image struct {
	XMLName     xml.Name    `xml:"image"`       // required
	URL         URL         `xml:"url"`         // required
	Title       Title       `xml:"title"`       // required
	Link        Link        `xml:"link"`        // required
	Width       Width       `xml:"width"`       // optional
	Height      Height      `xml:"height"`      // optional
	Description Description `xml:"description"` // optional
}

// <url> is a required sub-element of <image>.
//
// <url> is a required attribute of <enclosure>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#ltimagegtSubelementOfLtchannelgt
type URL string

// Whether <url> is valid.
func (r URL) IsValid() bool {
	return IsValidURL(string(r))
}

// <width> is an optional sub-element of <image>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#ltimagegtSubelementOfLtchannelgt
type Width string

// Whether <width> is valid.
//
// The maximum value for width is 144, default value is 88.
func (r Width) IsValid() bool {
	if i, err := strconv.ParseUint(string(r), 10, 0); err != nil || i > 144 {
		return false
	}
	return true
}

// <height> is an optional sub-element of <image>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#ltimagegtSubelementOfLtchannelgt
type Height string

// Whether <height> is valid.
//
// The maximum value for height is 400, default value is 31.
func (r Height) IsValid() bool {
	if i, err := strconv.ParseUint(string(r), 10, 0); err != nil || i > 400 {
		return false
	}
	return true
}

// Whether <image> is valid.
func (r Image) IsValid() bool {
	// Required sub-elements: <url>, <title>, <link>
	//
	// NOTE: In practice the image <title> and <link> should have the same value
	// as the channel's <title> and <link>.
	if r.URL.IsValid() || r.Title.IsValid() || r.Link.IsValid() {
		return false
	}
	// Optional sub-elements: <width>, <height>, <description>
	return r.Width.IsValid() && r.Height.IsValid()
}

// <rating> is an optional sub-element of <channel>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#optionalChannelElements
type Rating string

// Whether <rating> is valid.
func (r Rating) IsValid() bool { return true }

// <textInput> is an optional sub-element of <channel>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#lttextinputgtSubelementOfLtchannelgt
type TextInput struct {
	XMLName     xml.Name    `xml:"textInput"`   // required
	Title       Title       `xml:"title"`       // required
	Description Description `xml:"description"` // required
	Name        Name        `xml:"name"`        // required
	Link        Link        `xml:"link"`        // required
}

// Whether <textInput> is valid.
func (r TextInput) IsValid() bool {
	return r.Title != "" && r.Description != "" && r.Name != "" && r.Link != ""
}

// <name> is a required sub-element of <textInput>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#lttextinputgtSubelementOfLtchannelgt
type Name string

// Whether <name> is valid.
func (r Name) IsValid() bool { return true }

// <skipHours> is an optional sub-element of <channel>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#optionalChannelElements
type SkipHours struct {
	XMLName xml.Name `xml:"skipHours"` // required
	Hour    []*Hour  `xml:"hour"`      // required
}

// Whether <skipHours> is valid.
//
// This element contains up to 24 <hour> sub-elements whose value is a number
// between 0 and 23.
func (r SkipHours) IsValid() bool {
	if len(r.Hour) > 24 {
		return false
	} else {
		for _, h := range r.Hour {
			if !h.IsValid() {
				return false
			}
		}
	}
	return true
}

// <hour> is an optional sub-element of <skipHours>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#optionalChannelElements
type Hour int

// Whether <hour> is valid.
func (r Hour) IsValid() bool {
	if r < 0 || r > 23 {
		return false
	}
	return true
}

// <skipDays> is an optional sub-element of <channel>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#optionalChannelElements
type SkipDays struct {
	XMLName xml.Name `xml:"skipDays"` // required
	Day     []*Day   `xml:"hour"`     // required
}

// Whether <skipDays> is valid.
//
// This element contains up to seven <day> sub-elements whose value is
// Monday, Tuesday, Wednesday, Thursday, Friday, Saturday or Sunday.
func (r SkipDays) IsValid() bool {
	if len(r.Day) > 7 {
		return false
	} else {
		for _, d := range r.Day {
			if !d.IsValid() {
				return false
			}
		}
	}
	return true
}

// <day> is an optional sub-element of <skipDays>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#optionalChannelElements
type Day string

// Whether <day> is valid.
//
// TODO: Check if Monday - Sunday.
func (r Day) IsValid() bool { return true }

// Whether <channel> is valid.
func (c Channel) IsValid() bool {
	return c.Title != "" && c.Link != "" && c.Description != ""
}

// <item> is an optional sub-element of <channel>.
//
// A channel may contain any number of <item>s.
//
// All elements of an item are optional, however at least one of title or
// description must be present.
//
// See: https://validator.w3.org/feed/docs/rss2.html#hrelementsOfLtitemgt
type Item struct {
	XMLName     xml.Name    `xml:"item"`        // required
	Title       Title       `xml:"title"`       // conditionally required
	Link        Link        `xml:"link"`        // optional
	Description Description `xml:"description"` // conditionally required
	Source      Source      `xml:"source"`      // optional
	Enclosure   Enclosure   `xml:"enclosure"`   // optional
	Category    Category    `xml:"category"`    // optional
	PubDate     PubDate     `xml:"pubDate"`     // optional
	GUID        GUID        `xml:"guid"`        // optional
	Comments    Comments    `xml:"comments"`    // optional
	Author      Author      `xml:"author"`      // optional
}

// Whether <item> is valid.
func (r Item) IsValid() bool {
	return r.Title.IsValid() || r.Description.IsValid()
}

// <source> is an optional sub-element of <item>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#ltsourcegtSubelementOfLtitemgt
type Source struct {
	URL URL `xml:"url,attr"` // required
}

// Whether <source> is valid.
func (r Source) IsValid() bool {
	return r.URL.IsValid()
}

// <enclosure> is an optional sub-element of <item>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#ltenclosuregtSubelementOfLtitemgt
type Enclosure struct {
	URL    URL    `xml:"url,attr"`    // required
	Length Length `xml:"length,attr"` // required
	Type   Type   `xml:"type,attr"`   // required
}

// <length> is a required attribute of <enclosure>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#ltenclosuregtSubelementOfLtitemgt
type Length string

// Whether <length> is valid.
func (r Length) IsValid() bool { return true }

// <type> is a required attribute of <enclosure>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#ltenclosuregtSubelementOfLtitemgt
type Type string

// Whether <type> is valid.
func (r Type) IsValid() bool { return true }

// Whether <enclosure> is valid.
func (r Enclosure) IsValid() bool {
	return r.URL.IsValid() && r.Length.IsValid() && r.Type.IsValid()
}

// <guid> is an optional sub-element of <item>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#ltguidgtSubelementOfLtitemgt
type GUID struct {
	IsPermaLink string `xml:"isPermaLink,attr"` // optional
}

// Whether <guid> is valid.
func (r GUID) IsValid() bool { return true }

// <comments> is an optional sub-element of <item>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#ltcommentsgtSubelementOfLtitemgt
type Comments string

// Whether <comments> is valid.
func (r Comments) IsValid() bool { return true }

// <author> is an optional sub-element of <item>.
//
// See: https://validator.w3.org/feed/docs/rss2.html#ltauthorgtSubelementOfLtitemgt
type Author string

// Whether <author> is valid.
func (r Author) IsValid() bool { return true }
