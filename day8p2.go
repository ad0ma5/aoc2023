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
func gonextl(dest string, gt byte, nodes []node) []string {
  done:=false
  i:=0
  out:=[]string{}
  for done !=true {
    nd:=nodes[i].d
    if nd[len(nd)-1] == dest[len(dest)-1]{
      fmt.Println(" a ",nd)
      if string(gt) == "L"{
	out=append(out, nodes[i].l)
      }else{
        out=append(out, nodes[i].r)
      }
    }
    i++
    if len(nodes) == i{ done=true }
  }
  return out
}


///////
///////
func main() {
	data, err := os.ReadFile("input8.txt")
	//data, err := os.ReadFile("input8_small.txt")
	//data, err := os.ReadFile("input8_small1.txt")
	//data, err := os.ReadFile("input8_smallp2.txt")
	if err != nil {
		log.Fatal(err)
	}
	str:=string(data);
	l_arr:=strings.Split(str, "\n")
	l:= len(l_arr);
	///// Extract dat
	lri:=""
	nodes:=[]node{}
	smap:=map[string]string{}
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
          smap[nid+"L"] = nl
          smap[nid+"R"] = nr


          }

	}
	//fmt.Println(smap)
	//os.Exit(3)

	//Action
	cnod:=0
	dest:="A"
	dest_arr:=[]string{}
	lric:=0
	lril:=len(lri)
	done:=false
	fmt.Println(lri)
	for !done {
	  gt:=lri[lric]
	//fmt.Println(" gt ",string(gt) )
	  lric++
	  if lric == lril {
	    lric=0
          }
	  if len(dest_arr) == 0 {
	    dest_arr=gonextl(dest, gt, nodes)
	    cnod++
	    fmt.Println(dest_arr)
	    //done=true
	    //break
          }else{
	    endz:=true
	    for di:=0;di<len(dest_arr);di++{

	      //ndi:=gonext(dest_arr[di],gt,nodes)
	      ndi:=smap[dest_arr[di]+string(gt)]
              dest_arr[di]=ndi
	      if string(ndi[len(ndi)-1]) != "Z"{
	        endz=false
	      }else{

	//	break
	      }
	    }
		//if cnod%1000000 ==0{
	        //fmt.Println(" z "," c ",cnod," ", dest_arr )
	        //}
	    //fmt.Println(dest_arr)
	    //if cnod > 100 {
	      //break
            //}
	    cnod++
	    if endz { done=true }
          }

	}
	fmt.Println("cnod")

	fmt.Println(cnod)

	/////
	//fmt.Println(l, l_arr)

        //os.Stdout.Write(l);
	//os.Stdout.Write(data)

}
