package seal

import (
	"bytes"
	svg "github.com/sayhei/svgo"
)
//五角星x坐标
var PentagramXs=[]int{100,70,145,55,130}
//五角星y坐标
var PentagramYs=[]int{55,140,80,80,140}

const (
	//svg宽度
	SVG_WIDTH=200
	//svg高度
	SVG_HEIGHT=200
	//公司名称轨迹
	TEXTPATH="M45,150 A74,74 0 1,1 155,150"
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
	s.Textpath(companyName, "#circle","50%",`style="fill:red;"`,`text-anchor="middle"`)
	s.End()
	return buf5.String()
}