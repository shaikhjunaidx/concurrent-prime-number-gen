# Prime Number Generator

This Go program generates prime numbers up to a specified upper limit. It can run both with and without goroutines to compare performance.

## Features

- Generates prime numbers up to a given upper limit.
- Runs with concurrency using goroutines for faster computation on larger limits.
- Compares execution time between concurrent and non-concurrent implementations.

## Installation

Ensure you have Go installed on your system. You can check by running `go version`.

Clone the repository or copy the code into a file named `main.go`.

```bash
git clone <repository-url>
cd <repository-directory>
```

## Usage

### Running with Goroutines

To generate prime numbers up to a specified limit using goroutines, run:

```bash
go run main.go -limit <upper-limit>
```

Example:

```bash
go run main.go -limit 1000
```

## Measuring Performance

The program will output the time taken to generate prime numbers with and without goroutines. Compare the times to understand the performance benefits of concurrency.