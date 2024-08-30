package main

import (
	//"log"
	"os"
	"fmt"
	"strings"
	"strconv"
	//"unicode"
)


func main() {
    file:="input12s.txt"
	data, err := os.ReadFile(file)
	if err != nil {
		//log.Fatal(err)
	fmt.Println("not ok file", file)

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
      riddle:=larry[0]
	  for t:=0;t<len(templ);t++{
	    tint,_:=strconv.Atoi(templ[t])
	    for ti:=0;ti<tint;ti++ {
	      expected+="#"
	    }
        if t < len(templ)-1 {
	      expected+="."
        }
	  }

	  fmt.Println(riddle, "expected", expected, templ)
      if len(expected) == len(riddle) {
          fmt.Println("EQUAL: ", expected)
          continue
      }
      ec:=0
      for c:=0;c<len(riddle); c++ {
          if larr[i][c] != expected[ec]{
              fmt.Println(c, " not equal" ,larry[i][c], expected[c])
          }else{
              fmt.Println(c, " equal" ,larry[i][c], expected[c])
          }
      }

    }
	fmt.Println("ok")
        //os.Stdout.Write(l);
	//os.Stdout.Write(data)

}
