# Files fully read into memory

Loading files fully into memory is normally not done by `grep` implementations for several practical, understandable reasons:

* Memory limitations
* Unknown size of files
* Inability to filter files

The `bongrep` command doesn't have these limitations because stateful configuration persists across calls and can be used to throttle all of these variables, with sane defaults from the beginning (such as only operating on text files of a limited size).
