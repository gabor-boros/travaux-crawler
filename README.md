# travaux-crawler

Travaux crawler is a CLI tool that crawls Travaux, OVH's incident report site to find matching cloud instance IDs in the incident description.

The crawler is built for [OpenCraft Instance Manager](https://ocim.opencraft.com/en/latest/) (Ocim) operators to make it easier finding out if an ongoing outage affects an Open edX installation.

[![Github license](https://img.shields.io/github/license/gabor-boros/travaux-crawler)](https://github.com/gabor-boros/travaux-crawler/)

## Demo

```plaintext
$ ./bin/travaux-crawler -a 19367 -p "Cloud" -s "In progress" --all-pages
Using config file: /Users/gabor/.travaux-crawler.toml
Getting app server info for 19367
App server #19367 (edxapp-***********-19991) (6409d284-512b-4003-9044-5a31cbe5e0dd) is potentically affected by:
        - FS#52478 — PCI - GRA1 GRA3 - rack G133A15
        - FS#52446 — PCI - GRA1 - host749715
```

## Installation

To install `travaux-crawler`, use one of the [release artifacts](https://github.com/gabor-boros/travaux-crawler/releases) or simply run `go install https://github.com/gabor-boros/travaux-crawler`

You must create a new configuration file `$HOME/.travaux-crawler.toml` with the following content:

```toml
ocim_url = "<Ocim installation URL>"
ocim_username = "<Ocim username>"
ocim_password = "<Ocim password>"
```

## Usage

```plaintext
Travaux crawler is a CLI tool that crawls Travaux, OVH's incident report site
to find matching cloud instance IDs in the incident description.

The crawler is built for OpenCraft Instance Manager (Ocim) operators to make it
easier finding out if an ongoing outage affects an Open edX installation.

Usage:
  travaux-crawler [flags]

Flags:
      --all-pages         crawl all pages using the paginator
  -a, --app-server ints   potentially affected app server ID (required)
      --config string     config file (default is $HOME/.travauxCrawler-crawler.yaml)
  -h, --help              help for travaux-crawler
  -p, --project string    set travaux project name (default "All")
  -s, --status string     set travaux project name (default "All")
      --verbose           print the visited page URLs
```

## Development

To install everything you need for development, run the following:

```shell
$ git clone git@github.com:gabor-boros/travaux-crawler.git
$ cd travaux-crawler
$ make prerequisites
$ make deps
```

## Contributors

- [gabor-boros](https://github.com/gabor-boros)