package seal

import (
	"bytes"
	svg "github.com/sayhei/svgo"
)
//五角星x坐标
var PentagramXs=[]int{100,80,129,70,119}
//五角星y坐标
var PentagramYs=[]int{70,126,86,86,126}
//100,70 80,126 129,86 70,86 119,126
const (
	//svg宽度
	SVG_WIDTH=200
	//svg高度
	SVG_HEIGHT=200
	//公司名称轨迹
	TEXTPATH="M50,150 A70,70 0 1,1 150,150"
	//圆环中心点x坐标
	CIRCLE_X=100
	//圆环中心点y坐标
	CIRCLE_Y=100
)

//公章生成
func GenerateSeal(companyName string) string{
	buf5:=bytes.NewBuffer([]byte{})
	s := svg.New(buf5)
	s.Start(SVG_WIDTH, SVG_HEIGHT)
	s.Def()
	s.Path(TEXTPATH,`id="circle"`)
	s.DefEnd()
	s.Circle(CIRCLE_X, CIRCLE_Y, 96, "fill:none;stroke-width:6;stroke:red")
	s.Polygon(PentagramXs,PentagramYs,"fill:red;stroke:red;stroke-width:1;")
	s.Textpath(companyName, "50%","#circle",`style="fill:red;font-size:25px;font-weight:800;font-family:STSong"`,`text-anchor="middle"`)
	s.End()
	return buf5.String()
}