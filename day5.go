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
	seeds:=[]string{}
	smaps:=[]smap{}
	fmt.Println("cats", l )
	fmt.Println(str )

	minloc:=0

	for b:=0; b<l;b++{
	  if b==0{//seeds
	    l_p:=strings.Split(l_arr[b],": ")
	    seeds=strings.Split(l_p[1]," ")
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

	fmt.Println(seeds , smaps)
	for s:=0;s<len(seeds);s++{
	  fmt.Print("seed ",seeds[s])
	  fmt.Println()
          cs,_:=strconv.Atoi(seeds[s])
	  re:=cs
	  for m:=0;m<len(smaps);m++{
	    md:=smaps[m].data
	    mapped:=false

	    tit:=smaps[m].title
	    fmt.Println(tit)
	    ore:=re
            for d:=0;d<len(md);d++{
	      bsrc:=md[d].source
	      esrc:=md[d].source + md[d].length
	      dest:=md[d].dest
	      sddt:=bsrc-dest
              if re >= bsrc && re < esrc{

		re=re-sddt
	        mapped=true
	        fmt.Println(m," ",d," yes in range ",bsrc,"<=",ore,"<",esrc," new ",re)
		break
	      }
	    }
	    if !mapped{
	      fmt.Println(m," not in range ", re)

	    }
	  }
	  if minloc == 0 { minloc = re }
	  if minloc > re { minloc = re }
        }
	fmt.Println("minloc:", minloc)
        //os.Stdout.Write(l);
	//os.Stdout.Write(data)

}
