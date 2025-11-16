# Benchmarking performance in Sequential vs Concurrent with Single & MultiCore in Golang

Prerequisite knowledge in Go:
- `go` channels
    - To run concurrent code
- `sync.waitGroup`
    - To run concurrent code
- `testing` module
    - To benchmark performance results
- `runtime.GOMAXPROCS`
    - Restrict to use a single thread


**Note:** 
1. While the theory on concurrency mentioned Single Core vs Multi Core, the test was done using `GOMAXPROCS(1)` vs `GOMAXPROCS(8)` (8 is the default number of threads present in the system used to test). i.e. it is Single Thread vs Multi Thread comparison. Threads and Cores are two different topics but it is out of the scope of this discussion so we are moving ahead with it.
2. The results are from a Go program which compared single core vs multi core. It is possible chaning the `GOMAXPROCS` for each case might add som OS-level noise to the results. In order to get an even finer result, consider the Go program which compares Sequential vs Concurrency in 1 core and Sequential vs Concurrency using multiple cores. But,  


## Results:

### Go Concurrency Benchmarks Summary

These results (from an 8-core system) compare sequential vs. concurrent workloads, each iteration executing 4x `cpuWork()` (CPU-bound loop) or 4x `ioWork()` (500ms sleep). Benchmarks ran for a fixed time budget, so higher iterations (b.N) and lower ns/op indicate faster performance. Key takeaways:
- **CPU-Bound:**
    - Sequential is consistent (~33M ns/op)
    - Concurrent on single-core serializes (no gain)
    - Multi-core enables ~2.4x speedup for concurrent via true parallelism
- **I/O-Bound:**
    - Sequential blocks linearly (~2s/op)
    - Concurrent overlaps sleeps (~500ms/op, 4x faster) even on single-core
        - Go scheduler yields during blocking operations
    - Multi-core shows no extra gain for I/O workloads
- **Overall:**
    - Concurrency benefits I/O-bound tasks universally and doesn't benefit with additional cores
    - CPU-bound tasks need multi-core for speedup
    - Sequential workloads have negligible scaling with additional cores (single thread)

#### Raw Results Table
| Benchmark Type                  | Iterations (b.N) | ns/op (per iteration) | Approx. Time per Iteration |
|---------------------------------|------------------|-----------------------|----------------------------|
| **Single-Core Sequential CPU**  | 208              | 33,691,748           | ~34ms                     |
| **Single-Core Concurrent CPU**  | 170              | 33,760,320           | ~34ms                     |
| **Single-Core Sequential I/O**  | 3                | 2,002,026,290        | ~2.00s                    |
| **Single-Core Concurrent I/O**  | 10               | 500,572,596          | ~501ms                    |
| **Multi-Core Sequential CPU**   | 190              | 32,345,160           | ~32ms                     |
| **Multi-Core Concurrent CPU**   | 807              | 13,950,460           | ~14ms                     |
| **Multi-Core Sequential I/O**   | 3                | 2,002,324,078        | ~2.00s                    |
| **Multi-Core Concurrent I/O**   | 10               | 500,597,294          | ~501ms                    |

#### Comparison 1: Sequential vs. Concurrent
Concurrent outperforms sequential for I/O (4x faster) due to overlapping blocks; CPU concurrent only wins on multi-core.

| Workload Type | Sequential (ns/op) | Concurrent (ns/op) | Speedup (Concurrent / Sequential) |
|---------------|--------------------|--------------------|-----------------------------------|
| **CPU-Bound** | ~33M (avg. both cores) | 34M (single) / 14M (multi) | 1.0x (single, no gain) / 2.4x (multi) |
| **I/O-Bound** | ~2B (avg. both cores) | ~501M (both cores) | 4.0x (both configs) |

#### Comparison 2: Concurrent on 1 Core vs. Multi-Core
Multi-core boosts CPU concurrent (2.4x faster); I/O concurrent is identical (sleeps don't scale with cores).

| Workload Type | Single-Core Concurrent (ns/op) | Multi-Core Concurrent (ns/op) | Speedup (Multi / Single) |
|---------------|--------------------------------|-------------------------------|---------------------------|
| **CPU-Bound** | 33,760,320                    | 13,950,460                   | 2.4x                     |
| **I/O-Bound** | 500,572,596                   | 500,597,294                  | 1.0x (negligible)        |

#### Comparison 3: Sequential on 1 Core vs. Multi-Core
Minimal difference—sequential uses 1 thread regardless of cores (slight variance from noise).

| Workload Type | Single-Core Sequential (ns/op) | Multi-Core Sequential (ns/op) | Speedup (Multi / Single) |
|---------------|--------------------------------|-------------------------------|---------------------------|
| **CPU-Bound** | 33,691,748                    | 32,345,160                   | 1.0x (negligible)        |
| **I/O-Bound** | 2,002,026,290                 | 2,002,324,078                | 1.0x (negligible)        |


