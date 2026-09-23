# Goroutine HTTP Probe ⚡🔍

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://golang.org)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![CI](https://github.com/williamthompson8446/latency-pulse/actions/workflows/ci.yml/badge.svg)](https://github.com/williamthompson8446/latency-pulse/actions)

A high-performance endpoint latency prober demonstrating clean Go concurrency with goroutines and channels.

## Features
- 🚀 **Concurrent Probing**: Uses worker goroutines with `sync.WaitGroup`.
- ⏱️ **Precision Timing**: Microsecond-accurate latency measurements.
- 🛡️ **Graceful Timeouts**: Client timeout controls prevent lingering sockets.

## Build & Run
```bash
go run main.go https://api.github.com https://cloudflare.com
```

## License
MIT.
