# unizar-calendar

[![Tests](https://github.com/xzebra/unizar-calendar/workflows/tests/badge.svg)](https://github.com/xzebra/unizar-calendar/actions?query=workflow%3Atests)
[![GoReportCard](https://goreportcard.com/badge/github.com/xzebra/unizar-calendar?.svg)](https://goreportcard.com/report/github.com/xzebra/unizar-calendar)
[![Docs](https://godoc.org/github.com/xzebra/unizar-calendar?status.svg)](https://godoc.org/github.com/xzebra/unizar-calendar)

## Compile

Compilation is automatic thanks to Go modules. That means you have to
enable modules support by setting the environment variable
`GO111MODULE=on`.

    go build

## Setup

First of all, go to [Google Calendar API site](https://developers.google.com/calendar/quickstart/go), create a new project
with the Calendar API enabled, and save the `credentials.json` file in
this project root folder.

After that, compile and run the application. It will ask you to login
and paste the authorization code returned by Google.

Once `token.json` and `credentials.json` are present in project root
folder, you are good to go.

## Usage

You can interact with the CLI by running the executable from a
terminal. If you want to check the available options, you can use the
flag `-h`.

    $ ./unizar-calendar -h


## Requirements

-   Go 1.14 (or higher).
-   Go modules enabled.
-   Project with Google Calendar API enabled.
-   `credentials.json` and `token.json` generated.

