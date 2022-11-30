# Mime-type detection overhead

The benchmarks for Gabriel's `mimetype` package[^1] are incontrovertible. For most types they are more than half the time to establish. There are no benchmarks comparing it to the `file` command that uses a compiled database of MIME types for it's examination, but Gabriel's seems to be the fastest.

At first, creating a single `IsText` heuristic based on the one used by the perl `-T` operator was considered, but as it became clear that any overhead from a full mime-type examination of files that are *already* filtered out and loaded into memory was negligible --- especially since every file is examined in a separate goroutine allowing all to be examined concurrently (something that traditional `grep` has never had by comparison).

[^1]: gabriel-vasile/mimetype: A fast Golang library for media type and file extension detection, based on magic numbers <https://github.com/gabriel-vasile/mimetype>
