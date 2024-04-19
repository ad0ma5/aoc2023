package main

import (
	"log"
	"os"
	"fmt"
	"strings"
	//"sort"
	//"slices"
	//"unicode"
	//"strconv"
)

type node struct{
  d string
  l string
  r string
}

func gonext(dest string, gt byte, nodes []node) string {
  done:=false
  i:=0
  for done !=true {
    if nodes[i].d == dest{
      if string(gt) == "L"{
        return nodes[i].l
      }
      return nodes[i].r
    }
    i++
  }
  return dest
}

///////
///////
func main() {
	data, err := os.ReadFile("input8.txt")
	//data, err := os.ReadFile("input8_small.txt")
	//data, err := os.ReadFile("input8_small1.txt")
	if err != nil {
		log.Fatal(err)
	}
	str:=string(data);
	l_arr:=strings.Split(str, "\n")
	l:= len(l_arr);
	///// Extract dat
	lri:=""
	nodes:=[]node{}
	for li:=0;li<l;li++{
	  if l_arr[li] == "" { continue }
	  if li == 0 {
	    lri=l_arr[li]
	    continue
	  }else{

          lp:=strings.Split(l_arr[li], " = (")
	  lpl:=strings.Split(lp[1], ", ")
	  lplp:=strings.Replace(lpl[1], ")", "", 1)
	  nid:=lp[0]
	  nl:=lpl[0]
	  nr:=lplp
	  dlro:=node{d:nid,l:nl,r:nr}
	  nodes=append(nodes,dlro)
	fmt.Println(" .d ",dlro)

          }

	}

	//Action
	cnod:=0
	dest:="AAA"
	lric:=0
	lril:=len(lri)
	for dest!="ZZZ"{
	  gt:=lri[lric]
	  lric++
	  if lric == lril {
	    lric=0
          }
	  dest=gonext(dest, gt, nodes)
	  cnod++

	}
	fmt.Println(lri)
	fmt.Println("cnod")
	fmt.Println("zzz")

	fmt.Println(cnod)

	/////
	//fmt.Println(l, l_arr)

        //os.Stdout.Write(l);
	//os.Stdout.Write(data)

}
