package main

import (
	"log"
	"os"
	"fmt"
	"strings"
	"sort"
	"slices"
	//"unicode"
	"strconv"
)

type hc struct{
  label byte
  count int
}

type hrht struct{
  hand string
  rank int
  handtype int
  bid int
}
var htypes=[7]string{
	  "Five of a kind",
	  "Four of a kind",
	  "Full house",
	  "Three of a kind",
	  "Two pair",
	  "One pair",
	  "High card"}
var htv=[7][5]int{
	  {5,0,0,0,0},
	  {4,1,0,0,0},
	  {3,2,0,0,0},
	  {3,1,1,0,0},
	  {2,2,1,0,0},
	  {2,1,1,1,0},
	  {1,1,1,1,1} }


var hlabel=[]string{
	"A",
	"K",
	"Q",
	"T",
	"9",
	"8",
	"7",
	"6",
	"5",
	"4",
	"3",
	"2",
	"J"}

func countHand(hand string) []hc {
  hcr:=[]hc{}
  hasj:=false
  for s:=0;s<len(hand);s++{
    if "J" == string(hand[s]){
      hasj=true
    }
    found:=false
    for h:=0;h<len(hcr);h++{
      if hcr[h].label==hand[s] {
        hcr[h].count+=1
	found=true
	break
      }
    }
    if !found {
      hcn:=hc{label:hand[s],count:1}
      hcr=append(hcr, hcn)
    }
  }
  sort.Slice(hcr, func(i, j int) bool{
    return hcr[i].count > hcr[j].count
  })
  if hasj {//find J and move points to first
    for s:=0;s<len(hcr);s++{
      if string(hcr[s].label) == "J"{
	if s==0{
	  if len(hcr) ==1{
            //add biggy label
	    fmt.Println(" ho ",hand," ",hcr)
	    //for
	  }else{
	    hcr[s+1].count+=hcr[s].count
	    hcr[s].count =0
          }
        }else{
	  hcr[0].count+=hcr[s].count
	  hcr[s].count =0
	}
      }
    }
    sort.Slice(hcr, func(i, j int) bool{
      return hcr[i].count > hcr[j].count
    })
  }
  return hcr
}

func identifyHand(hcr []hc) int {
  cmp:=[]int{}
  for h:=0;h<len(hcr);h++{
    cmp=append(cmp,hcr[h].count)
  }
  good:=false
  for i:=0;i<len(htv);i++{
    good=true
    for j:=0;j<len(cmp);j++{
      if htv[i][j] != cmp[j]{
	good=false
        break
      }
    }
    if good{
      return i
    }
  }
  return -1
}

func rankMyHand(hando hrht, mando hrht) bool{
  if hando.handtype == mando.handtype {
    for h:=0;h<len(hando.hand);h++{
      if hando.hand[h] == mando.hand[h]{
        continue
      }
      lid:=getLid(hando.hand[h])
      lidl:=getLid(mando.hand[h])
      return lid < lidl
    }
  }
  return hando.handtype < mando.handtype
}

func getLid(b byte) int {
  for l:=0;l<len(hlabel);l++{
    if hlabel[l] == string(b) { return l }
  }
  return -1
}

func collect(collector []hrht, hando hrht) ([]hrht,int) {
    into:=0
    for h:=0;h<len(collector);h++{

      //find next to insert
      if rankMyHand(collector[h],hando){
	into=h+1
        continue
      }
      into=h
      break
    }
    collector = slices.Insert(collector, into, hando)

    wienings:=0
    for h:=0;h<len(collector);h++{
      rank:=len(collector)-h
      if collector[h].rank != rank{
	collector[h].rank =rank
      }
      wienings+=rank*collector[h].bid
    }
    return collector, wienings
}
///////
///////
func main() {
	data, err := os.ReadFile("input7.txt")
	//data, err := os.ReadFile("input7_small.txt")
	if err != nil {
		log.Fatal(err)
	}
	str:=string(data);
	l_arr:=strings.Split(str, "\n")
	l:= len(l_arr);
	/////
	collector:=[]hrht{}
	wienings:=0
	for li:=0;li<l;li++{
	  if l_arr[li] == "" { break }

	  lp:=strings.Split(l_arr[li], " ")
	  hand:=lp[0]
	  bid,_:=strconv.Atoi(lp[1])
	  ch:=countHand(hand)
	  idh:=identifyHand(ch)

	  hando:=hrht{hand:hand,handtype:idh,bid:bid}
	  collector,wienings=collect(collector, hando)
	  fmt.Println(hand," ", bid," ",ch," idh=",idh)



	}

	fmt.Println("collector")
	fmt.Println(collector)
	fmt.Println("wienings")
	fmt.Println(wienings)
	/////
	//fmt.Println(l, l_arr,htypes,hlabel,htv)

        //os.Stdout.Write(l);
	//os.Stdout.Write(data)

}
