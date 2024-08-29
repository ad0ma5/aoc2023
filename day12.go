package main

import (
	"log"
	"os"
	"fmt"
	"strings"
	"strconv"
	//"unicode"
)


func main() {
	data, err := os.ReadFile("input12s.txt")
	if err != nil {
		//log.Fatal(err)
	fmt.Println("not ok file")

	}
	str:=string(data);
	larr:=strings.Split(str,"\n")
	l:= len(larr);
	for i:= 0;i<l;i++{
	  fmt.Println("oki:" , larr[i])
	  if larr[i]=="" { continue }
	  larry:=strings.Split(larr[i], " ")
	  templ:=strings.Split(larry[1], ",")
	  expected:=""
	  for t:=0;t<len(templ);t++{
	    tint,_:=strconv.Atoi(templ[t])
	    for ti:=0;ti<tint;ti++ {
	      expected+="#"
	    }
	    expected+=" "
	  }
	  fmt.Println(larr[i], expected)

        }
	fmt.Println("ok")
        //os.Stdout.Write(l);
	//os.Stdout.Write(data)

}
