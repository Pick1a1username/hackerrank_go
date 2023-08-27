# StacksAndQueuesBenchmark

Run by same machine.

* StacksAndQueues

```
% go test -v -bench . -count 5
goos: darwin
goarch: arm64
pkg: StacksAndQueues
Benchmark_processQueries_testcase15
Benchmark_processQueries_testcase15-12          1000000000               0.5982 ns/op
Benchmark_processQueries_testcase15-12          1000000000               0.5834 ns/op
Benchmark_processQueries_testcase15-12          1000000000               0.5779 ns/op
Benchmark_processQueries_testcase15-12          1000000000               0.5972 ns/op
Benchmark_processQueries_testcase15-12          1000000000               0.5903 ns/op
PASS
ok      StacksAndQueues 88.699s
% 
```

* StacksAndQueues2

```
% go test -v -bench . -count 5
goos: darwin
goarch: arm64
pkg: StacksAndQueues2
Benchmark_processQueries_testcase15
Benchmark_processQueries_testcase15-12          1000000000               0.009768 ns/op
Benchmark_processQueries_testcase15-12          1000000000               0.008874 ns/op
Benchmark_processQueries_testcase15-12          1000000000               0.008817 ns/op
Benchmark_processQueries_testcase15-12          1000000000               0.008969 ns/op
Benchmark_processQueries_testcase15-12          1000000000               0.008879 ns/op
PASS
ok      StacksAndQueues2        0.645s
% 
```

In StacksAndQueues, memory reallocation must be bottleneck. StacksAndQueues reallocates memory everytime when adding new value to the stack(`Enqueue`). On the other hand, StacksAndQueues2 reallocates memory much less than StacksAndQueues.