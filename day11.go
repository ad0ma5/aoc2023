package main

import (
	"log"
	"os"
	"fmt"
	//"unicode"
	"strconv"
	"strings"
	//"math"
)

type dot struct {
  x int
  y int
  n int
}

func iscross(d1 dot, d2 dot, cross dot) bool {
  if d1.x < cross.x && cross.x  < d2.x {
    return true
  }
  if d1.y < cross.y  && cross.y < d2.y {
    return true
  }
  if d1.x > cross.x && cross.x  > d2.x {
    return true
  }
  if d1.y > cross.y  && cross.y > d2.y {
    return true
  }
  return false
}

func dotdiff(d1 dot, d2 dot) dot {
  d:=dot{}
  d.x=d1.x-d2.x
  d.y=d1.y-d2.y
  return d
}

func abs ( i int ) int {
  if i < 0 { i=-i }
  return i
}

func dotlen(d dot ) int{
  return int ( abs(d.x)+abs(d.y) )
}

func main() {
  fmt.Println("Hello, day11.go!")
  data, err := os.ReadFile("input11.txt")
  //data, err := os.ReadFile("input11s.txt")
  if err != nil {
	log.Fatal(err)
  }
  str:=string(data);
  ds:=[]dot{}
  larr:=strings.Split(str,"\n")
  l:= len(larr[0]);
  fmt.Println(l)
  yhassharp:=make([]bool, l)//[l]bool{false}
  for ind := range yhassharp {
    yhassharp[ind] = false
  }
  xhassharp:=make([]bool, len(larr)-1 )//[l]bool{false}
  expandedu:=[]string{}
  //FIND NOSHARP COLLS AND ROWS TO EXPAND THEM DOUBLE
  for i, v := range larr {
    if v == "" { continue }
    fmt.Println(i," ",v)
    xhassharp[i]=false
    for ii, vv := range v {
      if string(vv) == "#" {
	xhassharp[i]=true
	yhassharp[ii]=true
        d:=dot{x: ii, y:i, n: len(ds)+1}
        ds = append(ds, d)
      }
    }
    if !xhassharp[i] {

      fmt.Println(" NOHAS SHARP", i)
      theline:=""//[l]bool{false}
      for indi:=0; indi<len(v);indi++ {
        theline+= "*"
      }
      expandedu = append(expandedu, theline)
    }else{
      expandedu = append(expandedu, v)
    }
  }

  expandeduy:=[]string{}
  for _, v := range expandedu {
    nout:=v
    for ii:=len(yhassharp)-1;ii > 0;ii-- {
      if !yhassharp[ii]{
	nout=nout[:ii] + "*" + nout[ii+1:]
	//nout[ii] = "*"
        //add extra "."
      }
    }
    expandeduy = append(expandeduy, nout)
  }
  fmt.Println("xhassharp")
  fmt.Println(xhassharp)
  fmt.Println("yhassharp")
  fmt.Println(yhassharp)
  fmt.Println(ds)
  ds=[]dot{}

  for i, v := range expandeduy {
    for ii, vv := range v {
        //fmt.Print(v[ii])
      if string(vv) == "#" {
        d:=dot{x: ii, y:i, n: len(ds)+1}
        ds = append(ds, d)
	//expandeduy[i][ii] = rune(d.n)
	//expandeduy[i][ii] = d.n
        fmt.Print(strconv.Itoa(d.n))
      }else{
        fmt.Print(string(vv))
      }
    }
    fmt.Println(" ]",i," ")
  }

  fmt.Println(yhassharp)
  for _, line := range expandeduy {
    fmt.Println(line)
  }
  fmt.Println(ds)
  total:=0
  for id, dsi:= range ds {
    for dsiind:=id+1;dsiind<len(ds);dsiind++ {
      res:=dotlen(dotdiff(dsi, ds[dsiind]))
      crossdot:=dot{x:7,y:3}
      crosc:=0
      for xhi, xh := range yhassharp {
	if !xh {
	  crossdot=dot{x:xhi, y:0}
	  fmt.Println("x",dsi,ds[dsiind],crossdot)
          if iscross(dsi, ds[dsiind], crossdot){
	    fmt.Println("is cross x",dsi,ds[dsiind],crossdot)
	    crosc++
          }
	}
      }
      for yhi, yh := range xhassharp {
	if !yh {
	  crossdot=dot{x:0,y:yhi}
	  fmt.Println("y",dsi,ds[dsiind],crossdot)
          if iscross(dsi, ds[dsiind], crossdot){
	    fmt.Println("is cross y",dsi,ds[dsiind],crossdot)
	    crosc++
          }
	}
      }

      //check d1x < *x < d2x or d1y < *y < d2y and then multiply tiles needed number of expanding plains crossed
      fmt.Println("ttt",crosc)
      expandno:=1000000
      res=res-crosc+(crosc*expandno)
      total+=res
      fmt.Println(" r=",res, " id1=", id+1, " id2=",dsiind+1)

    }

  }
  fmt.Println(total)
}
