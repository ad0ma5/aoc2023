package main

import (
	"log"
	"os"
	"fmt"
	"strings"
	//"unicode"
	"strconv"
)

type srange struct{
	source int
	dest int
	length int
}
type smap struct{
	title string
	data []srange
}

func main() {
	data, err := os.ReadFile("input5.txt")
	//data, err := os.ReadFile("input5_small.txt")
	if err != nil {
		log.Fatal(err)
	}
	str:=string(data);
	l_arr:=strings.Split(str, "\n\n")
	l:= len(l_arr);
	//seeds:=[]int{}
	seedsf:=[]int{}
	seedst:=[]int{}
	smaps:=[]smap{}
	fmt.Println("cats", l )
	fmt.Println(str )

	minloc:=0

	for b:=0; b<l;b++{
	  if b==0{//seeds
	    l_p:=strings.Split(l_arr[b],": ")
	    seeds_:=strings.Split(l_p[1]," ")
	    for ss:=0;ss<len(seeds_)/2;ss++{
	      from,_:=strconv.Atoi(seeds_[ss*2])
	      offset,_:=strconv.Atoi(seeds_[ss*2+1])
	      to:=from+offset
	fmt.Println("from ",from," to ",to )
	/*
	      for ff:=from;ff<to;ff++{
		seeds=append(seeds,ff)
		fmt.Print(ff," ")
	      }
	      */

		seedsf=append(seedsf,from)
		seedst=append(seedst,to)
	    }

	    continue
          }

	    l_p:=strings.Split(l_arr[b],":\n")
	    ltit:=l_p[0]
	    nmap:=[]srange{}

	    lmap:=strings.Split(l_p[1],"\n")
	    for r:=0;r<len(lmap);r++{

	      lnmap:=strings.Split(lmap[r]," ")
	      fmt.Println(lnmap, len(lnmap) )
	      if len(lnmap) < 3 { continue }
	      dest,_:=strconv.Atoi(lnmap[0])
	      src,_:=strconv.Atoi(lnmap[1])
	      length,_:=strconv.Atoi(lnmap[2])
	      rel:=srange{source:src,dest:dest,length:length}
              nmap = append(nmap, rel)
            }
	    sm :=smap{ title: ltit, data: nmap }
	    smaps = append(smaps, sm)
        }

	fmt.Println(seedsf , smaps)
	for s:=0;s<len(seedsf);s++{
	  fmt.Print("seed f",seedsf[s])
	for z:=seedsf[s];z<seedst[s];z++{

	  //fmt.Println()
          //cs,_:=strconv.Atoi(seeds[s])
	  cs:=z
	  re:=cs
	  for m:=0;m<len(smaps);m++{
	    md:=smaps[m].data
	    mapped:=false

	    //tit:=smaps[m].title
	    //fmt.Println(tit)
	    ore:=re
	    _ = ore
            for d:=0;d<len(md);d++{
	      bsrc:=md[d].source
	      esrc:=md[d].source + md[d].length
	      dest:=md[d].dest
	      sddt:=bsrc-dest
              if re >= bsrc && re < esrc{

		re=re-sddt
	        mapped=true
	        //fmt.Println(m," ",d," yes in range ",bsrc,"<=",ore,"<",esrc," new ",re)
		break
	      }
	    }
	    if !mapped{
	      //fmt.Println(m," not in range ", re)

	    }
	  }
	  if minloc == 0 { minloc = re }
	  if minloc > re { minloc = re }
        }
        }
	fmt.Println(" 21054567 minloc:", minloc)
        //os.Stdout.Write(l);
	//os.Stdout.Write(data)

}
