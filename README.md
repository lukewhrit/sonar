# Sonar

Sonar is a modern and privacy-focused search engine with a somewhat decentralized design. Sonar is in no way ready to be used by anyone in production.

This repository includes the source code for the Sonar server, which in turn provides the essential components of Sonar: the crawler and the indexer. The crawler was created using [Colly](https://github.com/gocolly/colly) &mdash; a crawler and scraper framework. And the indexer was made with [Badger](https://github.com/dgraph-io/badger) &mdash; a fast and embedded key-value database in Go.

A Sonar instance is usable entirely by it's API. Simply `POST` to the search endpoint with a JSON object of keywords and criteria and you will get a paginated array of results.

The creation of Sonar was induced primarily by Drew DeVault's article ["We can do better than DuckDuckGo"](https://drewdevault.com/2020/11/17/Better-than-DuckDuckGo.html).

## Features

* RESTful API.
* Easy to install, deploy, and maintain.
* Super fast search experience.
* Actually a search engine &mdash; doesn't pull results from any other search engine.
* Free of advertisements or paid results.
* Typo-tolerant.
* Supports non-HTTP data sources like Gopher or Man pages and third-party name systems like OpenNIC.

Most of these features are not yet implemented, but they will be eventually.

## Table Of Contents

- [Sonar](#sonar)
  - [Features](#features)
  - [Table Of Contents](#table-of-contents)
  - [Documentation](#documentation)
    - [Self-hosting](#self-hosting)
    - [Running searches on the instance](#running-searches-on-the-instance)
    - [More Documentation and API Reference](#more-documentation-and-api-reference)
  - [Contributing new features](#contributing-new-features)
  - [License](#license)

## Documentation

### Self-hosting
### Running searches on the instance

### More Documentation and API Reference

The rest of the documentation for Sonar is available on this repository's [Github Wiki](https://github.com/lukewhrit/sonar/wiki). Additionally this page includes an in-depth reference for the server API and instructions for self-hosting and configuring an instance.

## Contributing new features

We're always looking for new volunteers for the project, we've got a lot to work to do and can definitely use help.

If you want to work on Sonar feel free to simply create a pull request, just make sure you have read our contributing guide and follow its rules. More information and the guide are available in the [`CONTRIBUTING.md`](CONTRIBUTING.md) file.

## License

The source code of Sonar, available in this repository, is licensed under the Apache 2.0 open-source license. A copy of Apache 2.0 is available in the repository in the [`LICENSE`](LICENSE) text file or on the web at both: [choosealicense.com](https://choosealicense.com/licenses/apache-2.0), and [apache.org](https://www.apache.org/licenses/LICENSE-2.0).
