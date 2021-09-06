# Dragon

Mini framework for Go

---

* [Install](#install)
* [Examples](#examples)

---

## Install

Run this command to install.
```bash
go get github.com/michaelwp/dragon
```

## Examples

```go
package main

import (
	"github.com/michaelwp/dragon"
	"log"
)

type response struct {
	Message string `json:"message"`
}

func main() {
	r := dragon.NewRouter()
	
	api := r.Group("/api/v1")
	api.GET("/home", func(d *dragon.Dragon) error {
		var resp response
		resp.Message = "Hello World"
		return d.ResponseJSON(200, resp)
	})

	log.Fatal(r.Run(":8090"))
}
```

## License

BSD licensed. See the LICENSE file for details.

