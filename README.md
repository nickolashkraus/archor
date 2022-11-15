# Archor

**TODO**
* Add coverage
* Add releases
* Add GoDoc

[![MIT License](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/nickolashkraus/archer/blob/master/LICENSE)
[![Go Test](https://github.com/github/docs/actions/workflows/main.yml/badge.svg)](https://github.com/nickolashkraus/archor/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/nickolashkraus/archor)](https://goreportcard.com/report/github.com/nickolashkraus/archor)

Archor (**Arch**ive + Mirr**or**) is a program for creating, archiving, and mirroring podcasts.

## Objective

Have you ever found a podcast you loved only to find that past episodes have been removed? Using Archor you can ensure that you always have a backup.

## Philosophy

>"All information should be free"

As stated in Steven Levy's *Hackers: Heroes of the Computer Revolution*, the second principle of hacker ethic is free and open access to information.

## Caution

Archor can be used for evil. It is possible to create a parallel feed, thereby diverting traffic from the original creators.

## Usage

To start using Archor right away, just point it at a public RSS feed.

```bash
$ archor init https://podcast.com/rss.xml --local
```

## Creating, updating and merging RSS feeds

### What is RSS?

At its core, RSS is a standard for distributing content on the internet.

The [RSS specification](https://validator.w3.org/feed/docs/rss2.html) is extremely terse requiring only 5 elements:
* `<rss>`: Top level element with a mandatory attribute called `version` (ex. `version="2.0"`). Specifies a RSS document and the version of RSS that the document conforms to.
* `<channel>`: Subordinate to the `<rss>` element. Contains information about the channel (metadata) and its contents. There is only one `<channel>` element per RSS document.

The following is a list of the required channel elements:
* `<title>`: The name of the channel.
* `<link>`:	The URL to the HTML website corresponding to the channel.
* `<description>`: A Phrase or sentence describing the channel.

The list of optional channel elements comprising the meat of the RSS document and the enumerated [here](https://validator.w3.org/feed/docs/rss2.html#optionalChannelElements).

**Example**
```xml
<?xml version="1.0" encoding="UTF-8" ?>
<rss version="2.0">
<channel>
 <title>RSS Title</title>
 <description>This is an example of an RSS feed</description>
 <link>http://www.example.com/main.html</link>
 <copyright>2020 Example.com All rights reserved</copyright>
 <lastBuildDate>Mon, 6 September 2010 00:01:00 +0000</lastBuildDate>
 <pubDate>Sun, 6 September 2009 16:20:00 +0000</pubDate>
 <ttl>1800</ttl>

 <item>
  <title>Example entry</title>
  <description>Here is some text containing an interesting description.</description>
  <link>http://www.example.com/blog/post/1</link>
  <guid isPermaLink="false">7bd204c6-1655-4c27-aeee-53f933c5395f</guid>
  <pubDate>Sun, 6 September 2009 16:20:00 +0000</pubDate>
 </item>

</channel>
</rss>
```

## Archiving vs. Mirroring

An archive system backs up your hard drive, keeps files even after they are deleted, and retains historical versions of your files.  Often, the process to restore and access your data is quite time intensive.

A mirroring backup system is one that simply mirrors your existing hard drive.  When you delete a file from your hard drive, your online backup does the same.
