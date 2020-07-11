# GO Company Seal Generate
golang  公司公章生成

```
go get -u github.com/sayhei/GOCompanySeal
```

###示例代码

```
package main

import (
	"github.com/sayhei/GOCompanySeal"
)

func main() {
    sealSvg:=seal.GenerateSeal("南京义途网络科技有限公司")
    fmt.Println(sealSvg)
}

```

![image](https://raw.githubusercontent.com/sayhei/GOCompanySeal/dd7ab472bf922204304d3bede1a9798bb27cd23b/seal.svg)
