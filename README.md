# democratic
> Golang skeleton for GraphQL API

[![Go Version][go-version-image]][go-version-url]
[![Build Status][travis-image]][travis-url]
[![License][license-image]][license-url]
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](http://makeapullrequest.com)

One to two paragraph statement about your product and what it does.

![](header.png)

## Features

- [x] Configuration by enviroment
- [x] Resolvers
- [x] Schemas
- [x] Flag parser
- [ ] Custom Log middleware
- [ ] Auth middleware
- [ ] Metrics middleware

## Requirements

- Govendor
- Golang 1.7

## Installation

#### From source
You can use [Clone from Github](http://github.com/vsouza/democratic) and run `make run` 


To set the correct environment you should use: `-e` flag (dev is default) 

## Usage example

```
curl -XGET http://localhost:8888/graphql\?query\=%7Buser(id%3A%22test%22)%7Busername%7D%7D
```

## Contribute

We would love you for the contribution to **democratic**, check the ``LICENSE`` file for more info.

## Meta

Vinicius Souza – [@iamvsouza](https://twitter.com/iamvsouza) – hi@vsouza.com

Distributed under the MIT license. See ``LICENSE`` for more information.

[https://github.com/vsouza/democratic](https://github.com/vsouza/democratic)

[go-version-image]:https://img.shields.io/badge/Go-1.7-blue.svg
[go-version-url]: https://golang.org/
[license-image]: https://img.shields.io/badge/License-MIT-blue.svg
[license-url]: LICENSE
[travis-image]: https://img.shields.io/travis/dbader/node-datadog-metrics/master.svg?style=flat-square
[travis-url]: https://travis-ci.org/dbader/node-datadog-metrics
[codebeat-image]: https://codebeat.co/badges/c19b47ea-2f9d-45df-8458-b2d952fe9dad
[codebeat-url]: https://codebeat.co/projects/github-com-vsouza-democratic