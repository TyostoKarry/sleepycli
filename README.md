<h1 align="center"> sleepycli </h1>

<p align="center">
    <img src="https://img.shields.io/badge/go-1.26.1-00ADD8?logo=go">
    <img src="https://github.com/TyostoKarry/sleepycli/actions/workflows/verify-pr.yml/badge.svg?branch=main" />
</p>

A small Go CLI for calculating sleep and wake times around 90-minute sleep cycles.

`sleepycli` helps you answer questions like:
- “If I sleep now, when should I wake up?”
- “If I want to wake up at 07:00, when should I go to bed?”
- “How many full sleep cycles fit between two times?”

The tool supports four main modes: calculate from **now**, from a target **wake** time, from a target **sleep** time, or across a **window** of time.

## Features

- Calculate wake times from the current time
- Calculate bedtimes for a target wake time
- Calculate wake times for a target sleep time
- Calculate full sleep cycles within a sleep window
- Configurable fall-asleep buffer
- Configurable minimum and maximum cycle count
- Input validation for modes, time format, and cycle ranges
- Custom help output
- Bash completion script
- Pull request CI for linting and tests

## How it works

The calculator uses a **90-minute sleep cycle** and adds an optional **buffer** to represent the time it takes to fall asleep.

Defaults:
- Sleep cycle: `90 minutes`
- Buffer: `15 minutes`
- Displayed cycles: `4` to `6`

## Installation

### Build from source

```bash
git clone https://github.com/TyostoKarry/sleepycli.git
cd sleepycli
go build -o sleepycli .
```

### Run without installing

```bash
go run . --help
```

## Usage

```bash
sleepycli [mode] [options]
```

Choose exactly **one** mode.

### Modes

#### `--now`
Calculate wake times starting from the current time.

```bash
sleepycli --now
```

#### `--wake HH:MM`
Calculate suggested bedtimes for a target wake time.

```bash
sleepycli --wake 07:00
```

#### `--sleep HH:MM`
Calculate suggested wake times for a target sleep time.

```bash
sleepycli --sleep 22:30
```

#### `--from HH:MM --to HH:MM`
Calculate how many complete sleep cycles fit inside a time window.

```bash
sleepycli --from 22:00 --to 07:00
```

### Options examples

#### Use a custom fall-asleep buffer

```bash
sleepycli --wake 07:00 --buffer 30
```

#### Show a different cycle range

```bash
sleepycli --sleep 23:00 --cycles-min 3 --cycles-max 7
```

### Sample output

#### `--wake 07:00`

```text
To wake up at 07:00:
───────────────────
Assuming 15 min to fall asleep

  - 6 cycles, go to sleep at 21:45 (9h 00m)
  - 5 cycles, go to sleep at 23:15 (7h 30m)
  - 4 cycles, go to sleep at 00:45 (6h 00m)
```

#### `--from 22:00 --to 07:00`

```text
Between 22:00 and 07:00:
───────────────────────
Assuming 15 min to fall asleep

  - 5 complete cycles (7h 30m)
  - 75 minutes remaining
```

## Flags

| Flag | Description | Default |
|---|---|---:|
| `-n, --now` | Calculate wake times from the current time | - |
| `-w, --wake HH:MM` | Calculate bedtimes for a target wake time | - |
| `-s, --sleep HH:MM` | Calculate wake times for a target sleep time | - |
| `-f, --from HH:MM` | Start of sleep window, used with `--to` | - |
| `-t, --to HH:MM` | End of sleep window, used with `--from` | - |
| `-b, --buffer int` | Minutes needed to fall asleep | `15` |
| `-m, --cycles-min int` | Minimum cycles to show | `4` |
| `-x, --cycles-max int` | Maximum cycles to show | `6` |
| `-g, --good-night` | Print random good night ASCII art | - |
| `-v, --version` | Print version | - |
| `-h, --help` | Show help | - |

## Input rules

- Time format is 24-hour `HH:MM`
- Short hour values like `7:00` are accepted
- Modes cannot be combined
- `--from` and `--to` must be used together
- `--now`, `--wake` and `--sleep` are mutually exclusive
- `--cycles-min` cannot be greater than `--cycles-max`
- Buffer and cycle values cannot be negative

## Bash completion

A completion script is included at:

```bash
scripts/sleepycli-completion.bash
```

To load it in your current shell:

```bash
source scripts/sleepycli-completion.bash
```

To make it persistent, add that line to your shell profile.

## CI

The repository includes a pull request workflow that:
- runs `golangci-lint`
- runs `go test ./...`
- triggers on pull requests targeting `main`
