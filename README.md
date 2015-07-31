## Walkers

Several concurrent or linear implementation of walker functions.

* unbounded: creates an unbounded number of goroutines to scan a tree
* bounded: scans a tree using a fixed number of workers and a buffer
* zero: standard library implementation with sorted results

Unlike other implementation, this only aims at correctness, not speed
at all costs.

See also:
* https://github.com/MichaelTJones/walk - Fast concurrent walker
* https://github.com/kr/fs - Non-recursive implementation
