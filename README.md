# Skeleton

A data structures and algorithms library written in Go.

## Install

```sh
go get -u github.com/teleivo/skeleton
```

**Needs: export GOEXPERIMENT=rangefunc** as it uses the experimental [iterators](https://go.dev/wiki/RangefuncExperiment).

## Examples

* [frequency](./examples/frequency/main.go) - word frequency counter using an ordered map based on a
  left-leaning red-black binary search tree

## Disclaimer

I wrote this library for my personal projects. It is thus tailored to my needs. Feel free to use it!
That being said, my intention is not to adjust it to someone elses liking.

