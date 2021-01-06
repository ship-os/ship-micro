# Log

DEPRECATED: use github.com/ship-os/ship-micro/v2/logger interface

This is the global logger for all micro based libraries.

## Set Logger

Set the logger for micro libraries

```go
// import go-micro/util/log
import "github.com/ship-os/ship-micro/util/log"

// SetLogger expects github.com/ship-os/ship-micro/debug/log.Log interface
log.SetLogger(mylogger)
```
