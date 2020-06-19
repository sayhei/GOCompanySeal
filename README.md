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

![image](https://github.com/sayhei/GOCompanySeal/seal.svg)
