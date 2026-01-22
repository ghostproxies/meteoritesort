# MeteoriteSort

[![MeteoriteSort](https://repository-images.githubusercontent.com/1138710336/ef4aa968-17e2-454b-a06e-f680c5a42b73)](https://github.com/ghostproxies/meteoritesort)

## Table of Contents

- [Introduction](https://github.com/ghostproxies/meteoritesort?tab=readme-ov-file#introduction)
- [Author](https://github.com/ghostproxies/meteoritesort?tab=readme-ov-file#author)
- [License](https://github.com/ghostproxies/meteoritesort?tab=readme-ov-file#license)
- [Reference](https://github.com/ghostproxies/meteoritesort?tab=readme-ov-file#reference)
- [Speed](https://github.com/ghostproxies/meteoritesort?tab=readme-ov-file#speed)

## Introduction

MeteoriteSort is an efficient unstable sorting algorithm.

It's an in-place `insertionsort` and `shellsort` derivative that's intended to replace sorting algorithms with the [fastest speed](https://github.com/ghostproxies/meteoritesort?tab=readme-ov-file#speed) for a majority of array lengths when both the specific distribution of presorted elements isn't predictable and the presorted order of duplicate elements isn't relevant.

Furthermore, it has [simple implementation](https://github.com/ghostproxies/meteoritesort?tab=readme-ov-file#reference), no auxiliary subarrays and no vendor-specific SIMD.

## Author

MeteoriteSort was created by [William Stafford Parsons](https://github.com/williamstaffordparsons) as a product of [GhostProxies](https://ghostproxies.com).

## License

MeteoriteSort is licensed with the [BSD-3-Clause](https://github.com/ghostproxies/meteoritesort?tab=BSD-3-Clause-1-ov-file#readme) license.

The default phrase `the copyright holder` in the 3rd clause is replaced with `GhostProxies`.

## Reference

### C Implementation

[meteoritesort.c](https://github.com/ghostproxies/meteoritesort/blob/master/meteoritesort.c)

The `meteoritesort` function sorts an array of elements in ascending order without preserving the presorted order of duplicate elements.

`elements_length` must be the count of elements in the `elements` array.

The integral type of `element` must match the integral type of each element in the `elements` array.

### C# Implementation

[meteoritesort.cs](https://github.com/ghostproxies/meteoritesort/blob/master/meteoritesort.cs)

The `MeteoriteSort` function sorts an array of elements in ascending order without preserving the presorted order of duplicate elements.

The integral type of `element` must match the integral type of each element in the `elements` array.

### Go Implementation

[meteoritesort.go](https://github.com/ghostproxies/meteoritesort/blob/master/meteoritesort.go)

The `MeteoriteSort` function sorts an array of elements in ascending order without preserving the presorted order of duplicate elements.

The integral type of `element` must match the integral type of each element in the `elements` array.

## Speed

MeteoriteSort was benchmarked and refined extensively during development with varying 64-bit architectures, compilers, compiler optimization settings, devices, programming languages and sorting scenarios.

The result was 1st-place speed rankings that are likely consistent across a majority of implementations compared to competing algorithms such as `heapsort`, `mergesort`, `quicksort`, `radixsort` and `shellsort`.

It's still being benchmarked and refined for further speed improvements wherever possible.
