# profiling-demo

Based on Dave Cheney's [profiling exercise](https://dave.cheney.net/high-performance-go-workshop/gophercon-2019.html#cpu_profiling_exercise)
from GopherCon 2019

## Two little demos

- Simple word count app: `words_test.go`.
  - Hint 1 - `syscall`s can be slow, try CPU profiling.
  - Hint 2 - Once you've fixed the CPU usage, try using the memory profiler to see
    if there's any stray memory allocations.
- Mutex profiling: `mutex_test.go`.
