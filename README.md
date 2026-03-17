# LeetCode Algorithms (Go)

A clean base repository for solving LeetCode problems in Go.

## Structure

- `problems/` - one folder per problem, named as `NNNN_slug`
- `scripts/new_problem.sh` - scaffolds a new problem folder with starter files

## Quick Start

1. Create a new problem scaffold:

   ```bash
   make new ID=1 SLUG=two-sum
   ```

2. Run tests:

   ```bash
   make test
   ```

## Naming Convention

- Problem folder: `problems/0001_two_sum`
- Package name: `p0001_two_sum`

This keeps packages unique and valid in Go.
