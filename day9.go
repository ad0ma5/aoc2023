package main

import (
	"log"
	"os"
	"fmt"
	"strings"
	//"sort"
	//"slices"
	//"unicode"
	"strconv"
)


func getdeeper(in []int) ([]int) {
  out:=[]int{}
  for i:=1;i<len(in);i++{
    s:=in[i]-in[i-1]
    out = append(out, s)
  }
  return out;
}
///////
///////
func main() {
	//data, err := os.ReadFile("input9.txt")
	data, err := os.ReadFile("input9_small.txt")
	//data, err := os.ReadFile("input8_small1.txt")
	if err != nil {
	  log.Fatal(err)
	}
	str:=string(data);
	l_arr:=strings.Split(str, "\n")
	l:= len(l_arr);
	///// Extract dat
	gransum:=0
	for li:=0;li<l;li++{
	  if l_arr[li] == "" { continue }


	  seq:=[][]int{}
	  no:=strings.Split(l_arr[li], " ")

	  seqin:=[]int{}

	  for al:=0;al<len(no);al++{
	    n,_:= strconv.Atoi(no[al])
	    seqin=append(seqin,n)
	  }
	  seq=append(seq,seqin)
	  done:=false
	  for !done{
	    seqins:=[]int{}
            seqins = getdeeper(seqin)
	    seq=append(seq,seqins)
	    done = true
	    for s:=0;s<len(seqins);s++{
	      if seqins[s]!=0{
	        done = false
	      }
            }
	    seqin = seqins
          }
	  //go backwards
	  next:=0
	  for z:=len(seq)-2;z>-1;z--{
	    //prev last + cur last
	    //prev:=seq[z+1]
	    cur:=seq[z]
	    fmt.Println(cur)
	    curl:=cur[len(cur)-1]
	    fmt.Println("curl")
	    fmt.Println(curl)
	    fmt.Println("next")
	    fmt.Println(next)
	    fmt.Println("next+curl")
	    next=next +curl
	    fmt.Println(next)

          }
	  gransum+=next
	    /*
//
	    if al > 0 {
	      sub:=n-seqin[al-1]
	      seqins=append(seqins,sub)
	    }
	      if sub > 0 && len(seqins) > 0{
		      seqin, seqins = getdeeper(seqin,seqins,n,al)
	      }
	  seq=append(seq,seqins)
*/
	  fmt.Println(seq)

	}

	//Action
	fmt.Println("gransum")
	fmt.Println(gransum)


	/////
	//fmt.Println(l, l_arr)

        //os.Stdout.Write(l);
	//os.Stdout.Write(data)

}
