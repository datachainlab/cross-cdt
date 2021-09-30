# CDT

![Test](https://github.com/datachainlab/cross-cdt/workflows/Test/badge.svg)
[![GoDoc](https://godoc.org/github.com/datachainlab/cross-cdt?status.svg)](https://pkg.go.dev/github.com/datachainlab/cross-cdt?tab=doc)

CDT (Conflict-free data type) provides a data structure that can be read and written in concurrency in distributed transactions. It also allows complex concurrent states to be represented by small data structure. It is used in distributed transactions between DLTs that execute transactions in serial.

When there are multiple distributed transactions accessing the same state in concurrency, CDT provides the following properties and operations:

1. A transaction can update the state without acquiring an exclusive lock
2. Subsequent transactions can update the state and perform the comparision operations under certain conditions
3. if operation 2 is failed, an "indefinite" error is raised and the transaction is aborted

CDT ensures that data is always merged in a consistent state, regardless of the existence of concurrent transactions and their success or failure. These merge operations are automatically performed.

## Related work

- [CRDT](https://crdt.tech/)
  - CRDT is a data structure used in multi-master databases or applications. CDT differs from CRDT in that it is a data structure that provides conflict-free operations for updating incomplete transactions in concurrency.

## Supported types

- Integer
  - operations
    - Add, Sub
    - Compares: LT(E), GT(E)
- Set
  - Grow-only Set
    - operations
      - Add
      - Lookup

## Authors

- Jun Kimura (https://github.com/bluele)
