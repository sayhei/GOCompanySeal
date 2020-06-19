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

<svg width="200" height="200"
     xmlns="http://www.w3.org/2000/svg"
     xmlns:xlink="http://www.w3.org/1999/xlink">
<defs>
<path d="M45,150 A74,74 0 1,1 155,150" id="circle" />
</defs>
<circle cx="100" cy="100" r="96" style="fill:none;stroke-width:6;stroke:red" />
<polygon points="100,55 70,140 145,80 55,80 130,140" style="fill:red;stroke:red;stroke-width:1;" />
<text style="fill:red;" text-anchor="middle" ><textPath startOffset="50%" xlink:href="#circle">南京义途网络科技有限公司</textPath></text>
</svg>
