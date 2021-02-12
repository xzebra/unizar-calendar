# unizar-calendar

[![Tests](https://github.com/xzebra/unizar-calendar/workflows/tests/badge.svg)](https://github.com/xzebra/unizar-calendar/actions?query=workflow%3Atests)
[![GoReportCard](https://goreportcard.com/badge/github.com/xzebra/unizar-calendar?.svg)](https://goreportcard.com/report/github.com/xzebra/unizar-calendar)
[![Docs](https://godoc.org/github.com/xzebra/unizar-calendar?status.svg)](https://godoc.org/github.com/xzebra/unizar-calendar)

## Usage

You can interact with the CLI by running the executable from a
terminal. If you want to check the available options, you can use the
flag `-h`.

    $ ./unizar-calendar -h

In case you generate for Google Calendar, make sure to create a new
calendar with the correct timezone before importing. To import a csv
file as a calendar, check [this guide](https://support.google.com/calendar/answer/37118?co=GENIE.Platform%3DDesktop&hl=en).

## Compile

Compilation is automatic thanks to Go modules. That means you have to
enable modules support by setting the environment variable
`GO111MODULE=on`.

    go build

## Requirements

-   Go 1.14 (or higher).
-   Go modules enabled.

If you are going to use the `webdata` module or `pkg/gcal`, you need
the following:
-   Project with [Google Calendar API site](https://developers.google.com/calendar/quickstart/go) enabled.
-   Create a Service Account.
-   Download JSON credentials of Service Account.
-   Set the path to the JSON file in the
    `GOOGLE_APPLICATION_CREDENTIALS` environment variable.
