# go-monit-log
Logging for Go Monitor


### Getting started
Current version of the library requires a latest stable Go release.
If you don't have the Go compiler installed, read the official Go install guide.

Use go tool to install the package in your packages tree:
```bash
go get github.com/phuongdo/go-monit-log
```

Then you can use it in import section of your Go programs:
```golang
import "github.com/phuongdo/go-monit-log"
```

### Basic Example

```golang
package main

import (
    "fmt"
    "github.com/phuongdo/go-monit-log"
)

func main(){
    var API_LOG = ...// your API log
    var serviceId = "8eeb3c67-fd93-11e6-947e-0cc47a7e0172"
    logErro := monitlog.LogMesg{serviceId, monitlog.ERROR, "your error message"}
	monitlog.PostToMonitor(logErro, API_LOG)

}
```
