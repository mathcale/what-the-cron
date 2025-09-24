# What The Cron?

A simple Go CLI tool to parse Cron expressions into human-readable descriptions and show the next execution time.

## Usage

```sh
what-the-cron "<cron-expression>"
```

### Example

```sh
./bin/what-the-cron "*/5 * * * *"
```

Output:

```plaintext
At every 5th minute
Next execution: 2025-09-24 18:25:00
```

## Building

Just run:

```sh
make
```

To build the binary at `bin/what-the-cron`.
