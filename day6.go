package main

import (
	"log"
	"os"
	"fmt"
	"strings"
	//"unicode"
	"strconv"
)

func main() {
	data, err := os.ReadFile("input6.txt")
	//data, err := os.ReadFile("input6_small.txt")
	if err != nil {
		log.Fatal(err)
	}
	str:=string(data);
	l_arr:=strings.Split(str, "\n")
	l:= len(l_arr);
	time:=[]string{}
	distance:=[]string{}
        speed:=1//mm/ms
	bigw:=1
	//extract from string
	for li:=0;li<l;li++{
	  if l_arr[li] ==""{ break }
	  ts_:=strings.Split(l_arr[li],":")
	  ts:=ts_[1]
	  ts=strings.TrimSpace(ts)
	  tsa:=strings.Split(ts,"   ")
          if li==0 { //time
	    time=tsa
    	  }else{
	    distance=tsa
          }
	}
	//do job
	for ti:=0;ti<len(time);ti++{
	  tt,_:=strconv.Atoi(strings.TrimSpace(time[ti]))
	  td,_:=strconv.Atoi(distance[ti])
	  fmt.Println("tt",tt," td",td)
	  w:=0
	  for i:=0;i<tt;i++{
	    speed=i
	    went:=(tt-speed)*speed
	    win:=went-td
	    if win > 0 {w+=1}
	    fmt.Println(">",i," w ",went, " w=",win," wc=",w)
	  }
	  bigw *=w
	}
	fmt.Println(l,l_arr, len(time), time, distance)
	fmt.Println("bw=",bigw)

        //os.Stdout.Write(l);
	//os.Stdout.Write(data)

}
