package mxgraph

import (
	"fmt"
	"testing"
)

func TestSM01(t *testing.T) {
	data := []byte(`
<mxGraphModel dx="1838" dy="1210" grid="1" gridSize="10" guides="1" tooltips="1" connect="1" arrows="1" fold="1" page="0" pageScale="1" pageWidth="827" pageHeight="1169" math="0" shadow="0">
  <root>
    <mxCell id="0" />
    <mxCell id="1" parent="0" />
    <mxCell id="EBAIrEwSQ_sO8G7dM4pI-14" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;exitX=1;exitY=0.5;exitDx=0;exitDy=0;entryX=0;entryY=0.5;entryDx=0;entryDy=0;" edge="1" parent="1" source="EBAIrEwSQ_sO8G7dM4pI-9" target="EBAIrEwSQ_sO8G7dM4pI-10">
      <mxGeometry relative="1" as="geometry" />
    </mxCell>
    <object label="" type="initialstate" id="EBAIrEwSQ_sO8G7dM4pI-9">
      <mxCell style="ellipse;whiteSpace=wrap;html=1;aspect=fixed;glass=0;sketch=0;fillColor=#000000;" vertex="1" parent="1">
        <mxGeometry x="-630" y="100" width="30" height="30" as="geometry" />
      </mxCell>
    </object>
    <mxCell id="EBAIrEwSQ_sO8G7dM4pI-15" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;exitX=1;exitY=0.5;exitDx=0;exitDy=0;entryX=0;entryY=0.5;entryDx=0;entryDy=0;" edge="1" parent="1" source="EBAIrEwSQ_sO8G7dM4pI-10" target="EBAIrEwSQ_sO8G7dM4pI-12">
      <mxGeometry relative="1" as="geometry" />
    </mxCell>
    <mxCell id="EBAIrEwSQ_sO8G7dM4pI-18" value="Push button" style="edgeLabel;html=1;align=center;verticalAlign=middle;resizable=0;points=[];" vertex="1" connectable="0" parent="EBAIrEwSQ_sO8G7dM4pI-15">
      <mxGeometry x="0.3111" y="1" relative="1" as="geometry">
        <mxPoint x="-19" y="-14" as="offset" />
      </mxGeometry>
    </mxCell>
    <object label="Wait" type="state" id="EBAIrEwSQ_sO8G7dM4pI-10">
      <mxCell style="swimlane;rounded=1;whiteSpace=wrap;html=1;glass=0;sketch=0;" vertex="1" parent="1">
        <mxGeometry x="-510" y="75" width="130" height="80" as="geometry" />
      </mxCell>
    </object>
    <mxCell id="EBAIrEwSQ_sO8G7dM4pI-16" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;exitX=0.5;exitY=1;exitDx=0;exitDy=0;entryX=1;entryY=0.5;entryDx=0;entryDy=0;" edge="1" parent="1" source="EBAIrEwSQ_sO8G7dM4pI-12" target="EBAIrEwSQ_sO8G7dM4pI-13">
      <mxGeometry relative="1" as="geometry" />
    </mxCell>
    <mxCell id="EBAIrEwSQ_sO8G7dM4pI-21" value="Push button" style="edgeLabel;html=1;align=center;verticalAlign=middle;resizable=0;points=[];" vertex="1" connectable="0" parent="EBAIrEwSQ_sO8G7dM4pI-16">
      <mxGeometry x="-0.0875" y="1" relative="1" as="geometry">
        <mxPoint as="offset" />
      </mxGeometry>
    </mxCell>
    <object label="Run" type="state" id="EBAIrEwSQ_sO8G7dM4pI-12">
      <mxCell style="swimlane;rounded=1;whiteSpace=wrap;html=1;glass=0;sketch=0;" vertex="1" parent="1">
        <mxGeometry x="-290" y="75" width="130" height="80" as="geometry" />
      </mxCell>
    </object>
    <mxCell id="EBAIrEwSQ_sO8G7dM4pI-17" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;exitX=0;exitY=0.5;exitDx=0;exitDy=0;entryX=0.5;entryY=1;entryDx=0;entryDy=0;" edge="1" parent="1" source="EBAIrEwSQ_sO8G7dM4pI-13" target="EBAIrEwSQ_sO8G7dM4pI-10">
      <mxGeometry relative="1" as="geometry" />
    </mxCell>
    <mxCell id="EBAIrEwSQ_sO8G7dM4pI-22" value="Push button" style="edgeLabel;html=1;align=center;verticalAlign=middle;resizable=0;points=[];" vertex="1" connectable="0" parent="EBAIrEwSQ_sO8G7dM4pI-17">
      <mxGeometry x="0.1286" y="-1" relative="1" as="geometry">
        <mxPoint as="offset" />
      </mxGeometry>
    </mxCell>
    <object label="Stop" type="state" id="EBAIrEwSQ_sO8G7dM4pI-13">
      <mxCell style="swimlane;rounded=1;whiteSpace=wrap;html=1;glass=0;sketch=0;" vertex="1" parent="1">
        <mxGeometry x="-410" y="220" width="130" height="80" as="geometry" />
      </mxCell>
    </object>
    <mxCell id="EBAIrEwSQ_sO8G7dM4pI-20" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;exitX=0;exitY=0.5;exitDx=0;exitDy=0;exitPerimeter=0;entryX=0.5;entryY=0;entryDx=0;entryDy=0;endArrow=none;endFill=0;dashed=1;" edge="1" parent="1" source="EBAIrEwSQ_sO8G7dM4pI-19" target="EBAIrEwSQ_sO8G7dM4pI-10">
      <mxGeometry relative="1" as="geometry" />
    </mxCell>
    <object label="" type="note" id="EBAIrEwSQ_sO8G7dM4pI-19">
      <mxCell style="shape=note;whiteSpace=wrap;html=1;backgroundOutline=1;darkOpacity=0.05;glass=0;sketch=0;fillColor=#ffffff;" vertex="1" parent="1">
        <mxGeometry x="-410" y="-120" width="80" height="100" as="geometry" />
      </mxCell>
    </object>
  </root>
</mxGraphModel>
`)
	result := Parse(data)

	fmt.Printf("result type is %T\n", result)

	for _, v := range result {
		fmt.Println("--------------------------------------")
		for _, x := range v.States {
			fmt.Println(x)
		}
		for _, x := range v.Events {
			fmt.Println(x)
		}
		for _, x := range v.Transitions {
			fmt.Println(x.Src, x.Dest, x.Event)
		}
		fmt.Println(v.Initial)
	}
}

func TestSM02(t *testing.T) {
	data := []byte(`
<mxGraphModel dx="1838" dy="1210" grid="1" gridSize="10" guides="1" tooltips="1" connect="1" arrows="1" fold="1" page="0" pageScale="1" pageWidth="827" pageHeight="1169" math="0" shadow="0">
  <root>
    <mxCell id="0" />
    <mxCell id="1" parent="0" />
    <mxCell id="EBAIrEwSQ_sO8G7dM4pI-14" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;exitX=1;exitY=0.5;exitDx=0;exitDy=0;entryX=0;entryY=0.5;entryDx=0;entryDy=0;" parent="1" source="EBAIrEwSQ_sO8G7dM4pI-9" target="EBAIrEwSQ_sO8G7dM4pI-10" edge="1">
      <mxGeometry relative="1" as="geometry" />
    </mxCell>
    <object label="" type="initialstate" id="EBAIrEwSQ_sO8G7dM4pI-9">
      <mxCell style="ellipse;whiteSpace=wrap;html=1;aspect=fixed;glass=0;sketch=0;fillColor=#000000;" parent="1" vertex="1">
        <mxGeometry x="-630" y="100" width="30" height="30" as="geometry" />
      </mxCell>
    </object>
    <mxCell id="EBAIrEwSQ_sO8G7dM4pI-15" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;exitX=1;exitY=0.5;exitDx=0;exitDy=0;entryX=0;entryY=0.5;entryDx=0;entryDy=0;" parent="1" source="EBAIrEwSQ_sO8G7dM4pI-10" target="EBAIrEwSQ_sO8G7dM4pI-12" edge="1">
      <mxGeometry relative="1" as="geometry" />
    </mxCell>
    <mxCell id="EBAIrEwSQ_sO8G7dM4pI-18" value="Push button" style="edgeLabel;html=1;align=center;verticalAlign=middle;resizable=0;points=[];" parent="EBAIrEwSQ_sO8G7dM4pI-15" vertex="1" connectable="0">
      <mxGeometry x="0.3111" y="1" relative="1" as="geometry">
        <mxPoint x="-19" y="-14" as="offset" />
      </mxGeometry>
    </mxCell>
    <object label="Wait" type="state" id="EBAIrEwSQ_sO8G7dM4pI-10">
      <mxCell style="swimlane;rounded=1;whiteSpace=wrap;html=1;glass=0;sketch=0;" parent="1" vertex="1">
        <mxGeometry x="-510" y="75" width="130" height="80" as="geometry" />
      </mxCell>
    </object>
    <mxCell id="EBAIrEwSQ_sO8G7dM4pI-16" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;exitX=0.5;exitY=1;exitDx=0;exitDy=0;entryX=1;entryY=0.5;entryDx=0;entryDy=0;" parent="1" source="EBAIrEwSQ_sO8G7dM4pI-12" target="EBAIrEwSQ_sO8G7dM4pI-13" edge="1">
      <mxGeometry relative="1" as="geometry" />
    </mxCell>
    <mxCell id="EBAIrEwSQ_sO8G7dM4pI-21" value="Push button" style="edgeLabel;html=1;align=center;verticalAlign=middle;resizable=0;points=[];" parent="EBAIrEwSQ_sO8G7dM4pI-16" vertex="1" connectable="0">
      <mxGeometry x="-0.0875" y="1" relative="1" as="geometry">
        <mxPoint as="offset" />
      </mxGeometry>
    </mxCell>
    <object label="Run" type="state" id="EBAIrEwSQ_sO8G7dM4pI-12">
      <mxCell style="swimlane;rounded=1;whiteSpace=wrap;html=1;glass=0;sketch=0;" parent="1" vertex="1">
        <mxGeometry x="-290" y="75" width="130" height="80" as="geometry" />
      </mxCell>
    </object>
    <mxCell id="EBAIrEwSQ_sO8G7dM4pI-17" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;exitX=0;exitY=0.5;exitDx=0;exitDy=0;entryX=0.5;entryY=1;entryDx=0;entryDy=0;" parent="1" source="EBAIrEwSQ_sO8G7dM4pI-13" target="EBAIrEwSQ_sO8G7dM4pI-10" edge="1">
      <mxGeometry relative="1" as="geometry" />
    </mxCell>
    <mxCell id="EBAIrEwSQ_sO8G7dM4pI-22" value="Push button" style="edgeLabel;html=1;align=center;verticalAlign=middle;resizable=0;points=[];" parent="EBAIrEwSQ_sO8G7dM4pI-17" vertex="1" connectable="0">
      <mxGeometry x="0.1286" y="-1" relative="1" as="geometry">
        <mxPoint as="offset" />
      </mxGeometry>
    </mxCell>
    <object label="Stop" type="state" id="EBAIrEwSQ_sO8G7dM4pI-13">
      <mxCell style="swimlane;rounded=1;whiteSpace=wrap;html=1;glass=0;sketch=0;" parent="1" vertex="1">
        <mxGeometry x="-410" y="220" width="280" height="250" as="geometry">
          <mxRectangle x="-410" y="220" width="60" height="23" as="alternateBounds" />
        </mxGeometry>
      </mxCell>
    </object>
    <object label="Run1" type="state" id="EBAIrEwSQ_sO8G7dM4pI-2">
      <mxCell style="swimlane;rounded=1;whiteSpace=wrap;html=1;glass=0;sketch=0;" parent="EBAIrEwSQ_sO8G7dM4pI-13" vertex="1">
        <mxGeometry x="150" y="60" width="70" height="40" as="geometry" />
      </mxCell>
    </object>
    <object label="Run2" type="state" id="EBAIrEwSQ_sO8G7dM4pI-3">
      <mxCell style="swimlane;rounded=1;whiteSpace=wrap;html=1;glass=0;sketch=0;" parent="EBAIrEwSQ_sO8G7dM4pI-13" vertex="1">
        <mxGeometry x="30" y="60" width="70" height="40" as="geometry" />
      </mxCell>
    </object>
    <mxCell id="EBAIrEwSQ_sO8G7dM4pI-4" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;exitX=1;exitY=0.25;exitDx=0;exitDy=0;entryX=0;entryY=0.25;entryDx=0;entryDy=0;endArrow=classic;endFill=1;" parent="EBAIrEwSQ_sO8G7dM4pI-13" source="EBAIrEwSQ_sO8G7dM4pI-3" target="EBAIrEwSQ_sO8G7dM4pI-2" edge="1">
      <mxGeometry relative="1" as="geometry" />
    </mxCell>
    <mxCell id="EBAIrEwSQ_sO8G7dM4pI-6" value="a" style="edgeLabel;html=1;align=center;verticalAlign=middle;resizable=0;points=[];" parent="EBAIrEwSQ_sO8G7dM4pI-4" vertex="1" connectable="0">
      <mxGeometry relative="1" as="geometry">
        <mxPoint as="offset" />
      </mxGeometry>
    </mxCell>
    <mxCell id="EBAIrEwSQ_sO8G7dM4pI-5" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;exitX=0;exitY=0.75;exitDx=0;exitDy=0;entryX=1;entryY=0.75;entryDx=0;entryDy=0;endArrow=classic;endFill=1;" parent="EBAIrEwSQ_sO8G7dM4pI-13" source="EBAIrEwSQ_sO8G7dM4pI-2" target="EBAIrEwSQ_sO8G7dM4pI-3" edge="1">
      <mxGeometry relative="1" as="geometry" />
    </mxCell>
    <mxCell id="EBAIrEwSQ_sO8G7dM4pI-7" value="b" style="edgeLabel;html=1;align=center;verticalAlign=middle;resizable=0;points=[];" parent="EBAIrEwSQ_sO8G7dM4pI-5" vertex="1" connectable="0">
      <mxGeometry x="-0.25" y="2" relative="1" as="geometry">
        <mxPoint as="offset" />
      </mxGeometry>
    </mxCell>
    <mxCell id="EBAIrEwSQ_sO8G7dM4pI-8" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;exitX=0.5;exitY=0;exitDx=0;exitDy=0;entryX=0.5;entryY=1;entryDx=0;entryDy=0;endArrow=classic;endFill=1;" edge="1" parent="EBAIrEwSQ_sO8G7dM4pI-13" source="EBAIrEwSQ_sO8G7dM4pI-1" target="EBAIrEwSQ_sO8G7dM4pI-3">
      <mxGeometry relative="1" as="geometry" />
    </mxCell>
    <object label="" type="initialstate" id="EBAIrEwSQ_sO8G7dM4pI-1">
      <mxCell style="ellipse;whiteSpace=wrap;html=1;aspect=fixed;glass=0;sketch=0;fillColor=#000000;" vertex="1" parent="EBAIrEwSQ_sO8G7dM4pI-13">
        <mxGeometry x="50" y="140" width="30" height="30" as="geometry" />
      </mxCell>
    </object>
    <mxCell id="EBAIrEwSQ_sO8G7dM4pI-20" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;exitX=0;exitY=0.5;exitDx=0;exitDy=0;exitPerimeter=0;entryX=0.5;entryY=0;entryDx=0;entryDy=0;endArrow=none;endFill=0;dashed=1;" parent="1" source="EBAIrEwSQ_sO8G7dM4pI-19" target="EBAIrEwSQ_sO8G7dM4pI-10" edge="1">
      <mxGeometry relative="1" as="geometry" />
    </mxCell>
    <object label="" type="note" id="EBAIrEwSQ_sO8G7dM4pI-19">
      <mxCell style="shape=note;whiteSpace=wrap;html=1;backgroundOutline=1;darkOpacity=0.05;glass=0;sketch=0;fillColor=#ffffff;" parent="1" vertex="1">
        <mxGeometry x="-410" y="-120" width="80" height="100" as="geometry" />
      </mxCell>
    </object>
  </root>
</mxGraphModel>
`)
	result := Parse(data)
	for _, v := range result {
		fmt.Println("--------------------------------------")
		for _, x := range v.States {
			fmt.Println(x)
			// 	fmt.Println(reflect.TypeOf(x))
		}
		for _, x := range v.Events {
			fmt.Println(x)
			// 	fmt.Println(reflect.TypeOf(x))
		}
		for _, x := range v.Transitions {
			fmt.Println(x.Src, x.Dest, x.Event)
			// 	fmt.Println(reflect.TypeOf(x.Src))
			// 	fmt.Println(reflect.TypeOf(x.Dest))
			// 	fmt.Println(reflect.TypeOf(x.Event))
		}
		fmt.Println(v.Initial)
		// fmt.Println(reflect.TypeOf(v.Initial))
	}
}
