# User Stories

## The Hacker

As a hacker, I revel in the ability to archive all the Internet's information. As such, I want to be able to archive the the content provided by a public RSS feed simply by providing the program with the URL to the feed:

```bash
$ archor mirror https://podcast.com/rss.xml
```

Furthermore, I want the option to store content locally (on disk) or remotely (on a Cloud provider) simply by passing a flag to the command:

```bash
$ archor mirror https://podcast.com/rss.xml --local
```

```bash
$ archor mirror https://podcast.com/rss.xml --s3-bucket
```

**NOTE**: Mirroring provides a one-to-one duplication of the upstream RSS feed.

## The Collector

As a podcast enjoyer, I want to rest assured that my precious podcast episodes will not someday disappear. As such, I want to keep my owe archive of a podcast generated from a public (or private) RSS feed. In addition, I want to be merge RSS feeds to my heart's content either from other RSS feeds or from a hand-made configuration:

```bash
$ archor archive https://podcast.com/rss.xml --local
```

```bash
$ archor archive https://podcast.com/rss.xml --s3
```

**NOTE**: Archiving provides a fork of the of the upstream RSS feed.

I want to be able to merge and update an RSS feed using both the upstream RSS feed and configuration files:

```bash
$ archor merge https://podcast.com/rss.xml --local
```

## The Creator

As a podcast creator, I want full control over my podcast creation workflow. I want to be able to create the content, episode notes, etc., and generate the RSS feed automatically.

```bash
$ archor init
```

```bash
$ archor update
```
