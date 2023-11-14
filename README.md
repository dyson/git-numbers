# git-numbers: Select files in git commands by number

![version](https://img.shields.io/github/v/tag/dyson/git-numbers?label=version)
[![release](https://github.com/dyson/git-numbers/actions/workflows/release.yml/badge.svg)](https://github.com/dyson/git-numbers/actions/workflows/release.yml)
[![build](https://github.com/dyson/git-numbers/actions/workflows/build.yml/badge.svg)](https://github.com/dyson/git-numbers/actions/workflows/build.yml)
[![test coverage](https://coveralls.io/repos/github/dyson/git-numbers/badge.svg?branch=main)](https://coverallsio/github/dyson/git-numbers?branch=main)
[![maintainability](https://api.codeclimate.com/v1/badges/a9de05463178f58c181f/maintainability)](https://codeclimate.com/github/dyson/git-numbers/maintainability)
[![go report](https://goreportcard.com/badge/github.com/dyson/git-numbers)](https://goreportcard.com/report/github.com/dyson/git-numbers)
[![license](https://img.shields.io/github/license/dyson/git-numbers.svg)](https://github.com/dyson/git-numbers/blob/master/LICENSE)

*Pipe* because it's similar to unix like pipes and *sore* because the initial
hackathon version of this project was an eyesore.

Born from a proof of concept in using
[bitfield/script](https://github.com/bitfield/script) directly in the CLI
pipesore provides a number of text filters that you can pipe together to
process text. It takes input from stdin and writes the pipeline output to
stdout allowing it to be used alongside unix pipes.

## Motivation

## Installation

Download a stable [release](https://github.com/dyson/git-numbers/releases) or install the latest development version with:

```bash
$ go install github.com/dyson/git-numbers/cmd/git-numbers@main
```
TODO: ensure in path
TODO: git alias

Optionally alias git-numbers to something quicker to type, eg:

```bash
echo 'alias gn="git numbers"' >> ~/.bash_profile
```

```

## Basic Usage

A contrived example:

```bash
$ echo "cat cat cat dog bird bird bird bird" | pipesore 'Replace(" ", "\n") | Frequency() | First(1)'
4 bird
```

## License
See [LICENSE](https://github.com/dyson/git-numbers/blob/master/LICENSE) file.
