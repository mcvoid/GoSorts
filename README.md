GoSorts
=======

A collection of generic sorting routines implemented in Go.

Designed to work on any type that implements sort.Interface, which is such:

type Interface interface {
    Len() int
    Less(int, int) bool
    Swap(int, int)
}
