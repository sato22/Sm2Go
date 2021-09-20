package mxgraph

import (
	"fmt"
	"testing"
)

func TestDecode01(t *testing.T) {
	s := []byte(`5VhRb5swEP41eewUoJDmsUnTblInRcqkrY9uuIJbwyHjJNBfv3MwEOK1zbYmbZSn5D7ujP1992FMzxsnxY1kWfwdQxA9tx8WPe+q57pOfzigH42UFRI4QQVEkocmqQVm/BnqSoMueAh5J1EhCsWzLjjHNIW56mBMSlx10x5QdO+asQgsYDZnwkZ/8lDFFXrhDlr8K/Aoru/sBMPqSsLqZLOSPGYhrjYgb9LzxhJRVf+SYgxCk1fzUtVdv3C1mZiEVO1ScFte3/ojPHviNwWbPH5Llrx/5pm5qbJeMIS0fhOiVDFGmDIxadGRxEUagh61T1Gbc4uYEegQ+AhKlUZMtlBIUKwSYa5CwdUvXf7FN9GdGUz/vyo2g7IOUiXLjSId3tXj6aAtW0d1XbU+vagXaTNQjgs5h1e4qtuPyQjUK3luIy65AjABmg/VSRBM8WV3Hsy0Z9TktQrSHyPiXwhqxl0ysTB3Sm2JhSD7aClXMVcwy9h63StycFcolmeVpx54oQUfkcsU4ynIDrtLkAqK1/m1+agLzo0/zAPCGZh41drNM1C84bQae3cGg2O3RP9wlnB3tIT/n5ZYl15KycqNhAx5qvKNkacaaDvLc/1uZ3n+VnNUI7at0kzt37vHtfxntVO3Wd6woERFFGFK4fCdHNfsrQ0vtuP8PzjO2Zfj6oE/g+Wcz245f0fLDT9yF/LfdsGn2oS8bUt89CbkODZhx2WJA76YDXe0xAs9cBhLDI9vY2iOLh+2Mdi76Q861HGKLfJo2WqLIcj5M7tfJ2iGdL/nVeuvnyuCR5q9ObGjnyUjzR2nk9+luZDwMFz7yrxo0G38Uc+/IuSBCzFGgfoRlGKqk3Il8Qm2wK6ie5DoYjeF3L0p5NkK5SCXnLr2RBXqd9863cCWKDioROeWRNP5Ilc0vDxRjYKtg8HAtyQaHFQi+31pql0EklZ9mhJ51meBvWlEYfsVrjoHtt8yvclv`)
	expected := `<mxGraphModel dx="1097" dy="616" grid="1" gridSize="10" guides="1" tooltips="1" connect="1" arrows="1" fold="1" page="1" pageScale="1" pageWidth="827" pageHeight="1169" math="0" shadow="0"><root><mxCell id="0"/><mxCell id="1" parent="0"/><mxCell id="LyFL5Bo-kiGxaEjImvi0-3" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;exitX=0.5;exitY=0;exitDx=0;exitDy=0;entryX=0.5;entryY=1;entryDx=0;entryDy=0;" edge="1" parent="1" source="LyFL5Bo-kiGxaEjImvi0-1" target="LyFL5Bo-kiGxaEjImvi0-2"><mxGeometry relative="1" as="geometry"/></mxCell><mxCell id="LyFL5Bo-kiGxaEjImvi0-1" value="n" style="ellipse;whiteSpace=wrap;html=1;aspect=fixed;container=0;" vertex="1" parent="1"><mxGeometry x="140" y="170" width="30" height="30" as="geometry"/></mxCell><mxCell id="LyFL5Bo-kiGxaEjImvi0-6" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;exitX=0.5;exitY=0;exitDx=0;exitDy=0;entryX=0.5;entryY=0;entryDx=0;entryDy=0;" edge="1" parent="1" source="LyFL5Bo-kiGxaEjImvi0-2" target="LyFL5Bo-kiGxaEjImvi0-5"><mxGeometry relative="1" as="geometry"><Array as="points"><mxPoint x="325" y="135"/></Array></mxGeometry></mxCell><mxCell id="LyFL5Bo-kiGxaEjImvi0-2" value="" style="rounded=0;whiteSpace=wrap;html=1;rotation=90;" vertex="1" parent="1"><mxGeometry x="210" y="130" width="50" height="10" as="geometry"/></mxCell><mxCell id="LyFL5Bo-kiGxaEjImvi0-10" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;exitX=0.5;exitY=1;exitDx=0;exitDy=0;entryX=0.5;entryY=0;entryDx=0;entryDy=0;" edge="1" parent="1" source="LyFL5Bo-kiGxaEjImvi0-5" target="LyFL5Bo-kiGxaEjImvi0-9"><mxGeometry relative="1" as="geometry"/></mxCell><mxCell id="LyFL5Bo-kiGxaEjImvi0-5" value="" style="ellipse;whiteSpace=wrap;html=1;aspect=fixed;container=0;" vertex="1" parent="1"><mxGeometry x="310" y="170" width="30" height="30" as="geometry"/></mxCell><mxCell id="LyFL5Bo-kiGxaEjImvi0-11" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;exitX=0.5;exitY=1;exitDx=0;exitDy=0;entryX=0.5;entryY=1;entryDx=0;entryDy=0;" edge="1" parent="1" source="LyFL5Bo-kiGxaEjImvi0-9" target="LyFL5Bo-kiGxaEjImvi0-1"><mxGeometry relative="1" as="geometry"/></mxCell><mxCell id="LyFL5Bo-kiGxaEjImvi0-9" value="" style="rounded=0;whiteSpace=wrap;html=1;rotation=90;" vertex="1" parent="1"><mxGeometry x="210" y="220" width="50" height="10" as="geometry"/></mxCell><mxCell id="LyFL5Bo-kiGxaEjImvi0-12" value="Tarrival" style="text;html=1;resizable=0;autosize=1;align=center;verticalAlign=middle;points=[];fillColor=none;strokeColor=none;rounded=0;" vertex="1" parent="1"><mxGeometry x="210" y="80" width="50" height="20" as="geometry"/></mxCell><mxCell id="LyFL5Bo-kiGxaEjImvi0-13" value="Tservice" style="text;html=1;resizable=0;autosize=1;align=center;verticalAlign=middle;points=[];fillColor=none;strokeColor=none;rounded=0;" vertex="1" parent="1"><mxGeometry x="205" y="260" width="60" height="20" as="geometry"/></mxCell><mxCell id="LyFL5Bo-kiGxaEjImvi0-14" value="Pcustomer" style="text;html=1;resizable=0;autosize=1;align=center;verticalAlign=middle;points=[];fillColor=none;strokeColor=none;rounded=0;" vertex="1" parent="1"><mxGeometry x="65" y="175" width="70" height="20" as="geometry"/></mxCell><mxCell id="LyFL5Bo-kiGxaEjImvi0-15" value="Pservered" style="text;html=1;resizable=0;autosize=1;align=center;verticalAlign=middle;points=[];fillColor=none;strokeColor=none;rounded=0;" vertex="1" parent="1"><mxGeometry x="340" y="175" width="70" height="20" as="geometry"/></mxCell></root></mxGraphModel>`
	w, err := decode(s)
	if err != nil {
		fmt.Println("error")
	}
	if expected != string(w) {
		t.Error("Error in decode test.")
	}
}

func TestMxgraph01(t *testing.T) {
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
	result, err := GetGraphModel(data)
	if err != nil {
		t.Error(err)
	}
	for _, x := range result {
		fmt.Println(x)
	}
}
