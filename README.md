# StrategySort
StrategySort is an unstable adaptive hybrid sorting algorithm that treats sorting as a strategy selection problem, dynamically choosing the most appropriate approach at runtime.

# Rationale
The (unstable) sorting algorithm in Go's standard library is primarily a pattern-defeating quicksort (pdqsort). It is fast and works well for a wide variety of inputs without requiring any prior knowledge about the data. This general, data-independent approach is also typical of most high-performance sorting algorithms.

StrategySort takes a different approach.

Instead of relying on a single algorithm that works quite well in many scenarios, StrategySort performs an initial analysis of the input to extract information about the data. This information is then used to select the most appropriate sorting strategy between customized variants of mergesort and quicksort.

At first glance, this may seem counterproductive. Most high-performance sorting algorithms deliberately avoid additional passes over the data unless absolutely necessary. However, StrategySort shows that this assumption is not always true: an explicit presort analysis phase can still lead to competitive, and in some cases, even superior performance.

# Performance
Benchmarks were run on an Intel i5-9600K.

| Input Type                            | 1M integers        | 1K integers        |
| ------------------------------------- | ------------------ | ------------------ |
| Random data                           | ~20% faster        | ~35% faster        |
| Fully ascending                       | ~35% faster        | ~50% faster        |
| Fully descending                      | ~30% faster        | ~50% faster        |
| Nearly sorted (≥80%, local disorder)  | up to ~300% faster | up to ~200% faster |
| Nearly sorted (≥80%, global disorder) | up to ~10% slower  | roughly equal      |
| Duplicate-heavy data                  | roughly equal      | roughly equal      |

These results reflect the intended trade-off: when a suitable structure is present, the analysis phase often pays off; when it is not, performance remains competitive with Go's standard library.

# Novelty

Both adaptive and hybrid sorting are long-established concepts. What is novel about StrategySort, however, is its treatment of explicit presort analysis as a primary design decision.

Instead of indirectly inferring the structure during the sorting process, StrategySort deliberately performs up to O(n) steps before sorting in order to make data-driven strategy decisions. This design challenges the common assumption that such presort analysis is inherently wasteful and shows that it can be beneficial in practice for general sorting operations.

# Timsort

StrategySort is not the first algorithm to perform a prepass analysis on the input. Timsort, for example, performs such an analysis in order to:

- identify monotonic runs,
- reverse descending runs,
- determine the order of merges.

In this sense, the idea of paying O(n) in advance is not unique. However, this is where the designs differ:

Timsort
- Always performs a merge-based stable sort
- Uses the analysis only to optimize the execution of merges

StrategySort
- Chooses between fundamentally different sorting algorithms
- Uses analysis to decide which algorithm to execute

Instead of optimizing a single algorithm, StrategySort treats sorting as a problem of strategy selection.

# Implementation

Non-recursive implementation (manual stack instead of recursion) using stack-allocated buffers + one stack-allocated buffer to increase merge performance.

The implementations are inspired by Peters’ Pattern-Defeating Quicksort, Peters’ TimSort, Yaroslavskiy’s Dual-Pivot Quicksort, and Astrelin’s GrailSort, as well as other hybrid adaptive sorting techniques. Credit is also due to Igor van den Hoven for his Conjoined Triple Reversal rotation, which is used in the mergesort implementation.

# Memory

If in-place means that memory consumption does not scale with input size n, then StrategySort satisfies this definition. This is because it uses a fixed [512]int buffer for the quicksort stack, which corresponds to either 2048 B (32-bit) or 4096 B (64-bit), depending on the system architecture. The Mergesort implementation uses a fixed [192]int stack buffer (768 B or 1536 B) together with a [320]T buffer to speed up merging. The size of this generic buffer depends on T and ranges from 320 B (for int8, uint8) to 5120 B (for strings). While it would be possible to further reduce memory usage without sacrificing performance or functionality, this was not a primary design goal.
