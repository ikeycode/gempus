# gempus

> golang implementation of crontab parser for a recruitment question

[![Report](https://goreportcard.com/badge/github.com/ikeycode/gempus)](https://goreportcard.com/report/github.com/ikeycode/gempus) [![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

This 'project' is initially an implementation of a programming question set
by a recruiter. Given an input file in a simplified crontab format, output
the time-ascending order of events (i.e. execution paths)

I personally thought it had some potential, figured, FCUK it, create
a more 'complete' crond in Golang around the notion. More than likely
I'll never complete the project, but hey, I like code.

Oh, and I know, it has literally no real world purpose, and vixie-style
cronds the world unite, systemd ate the crond, etc, etc.

But seriously, don't rely on this for anything other than reference. It
doesn't do anything useful and it's deliberately tied to the original
design per the question.

## Name

Perverse merging of `tempus` (Latin: time) with `golang`.

## Running it, requirements

OS capable of running `make`, and a recent version of `golang` (1.12 perhaps..)

### Building it

    make

### Running it

    make check
    ./bin/demo
    ./bin/gempusd

## Authors

What, too cool for `git blame`? OK, Ikey Doherty is to blame for the code.

## License

`gempus` is available under the terms of the `Apache-2.0` license
