meituan-union-go
---

美团联盟golang-sdk

### Quickstart

```go
package main

import (
	"fmt"
	"github.com/yzchan/meituan-union-go"
	"github.com/yzchan/meituan-union-go/request"
)

func main() {
	client := meituan.NewClient("xxx", "yyy")
	client.Debug = false

	req := request.NewGenerateLinkRequest()
	req.SetActId(33)
	req.SetSid("xxx")
	req.SetLinkType(1)
	req.SetShortLink(true)
	response, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	fmt.Print(string(response))
}
```