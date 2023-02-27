# tools/logger

- example

```
package main

import (
	"log"
	"github.com/kzeromiddle/tools/logger"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	l := logger.Logger{
		Prefix:    defineProgramName,
	}
	l.InitStandardLogger(logger.FILE)
}
```
