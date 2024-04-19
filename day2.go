package main

import (
	"log"
	"os"
	"fmt"
	"strings"
	//"unicode"
	"strconv"
)

func isThisPossible(color string, count int) (bool){
  red:=12
  green:=13
  blue:=14
  if color == "red" && count <= red {
    return true
  }
  if color == "green" && count <= green {
    return true
  }
  if color == "blue" && count <= blue {
    return true
  }

  return false
}
//Determine which games would have been possible if the bag had been loaded with only 12 red cubes, 13 green cubes, and 14 blue cubes. What is the sum of the IDs of those games?
func main() {
	data, err := os.ReadFile("input2.txt")
	if err != nil {
		log.Fatal(err)
	}
	str:=string(data);
	str_l:= strings.Split(str, "\n");

	sumpow:=0
	countid:=0
	for i:=0;i<len(str_l);i++{
	  line_arr:=strings.Split(str_l[i],":")
	  if len(line_arr) > 1{
	    fmt.Println("l: ", i, line_arr[1]   )
	    mred:=0
	    mgreen:=0
	    mblue:=0

	    isok:=true
	    r_take := strings.Split(line_arr[1],";")
	    for j:=0;j<len(r_take);j++{
	      s_color:=strings.Split(r_take[j],",")
	      for k:=0;k<len(s_color);k++{
	        //fmt.Println("c", s_color[k],  j )
		s:=strings.TrimSpace(s_color[k])
		single:=strings.Split(s," ")
		color:=single[1]
		count, _ :=strconv.Atoi(single[0])
	        fmt.Println("", color,  count )
	  if color == "red" && mred < count {
	    mred = count
	  }
	  if color == "green" && mgreen < count{
	    mgreen = count
	  }
	  if color == "blue" && count > mblue {
	    mblue = count
	  }
		isP:=isThisPossible(color,count)

	        //fmt.Println("p: ", i, isP)

		if isP {
	        }else{
		  isok =false
	        }
	      }

            }//end take

	    if isok {

		countid+=i+1
	        //fmt.Println("add: ", countid, i+1, isok)
	    }
	    pow:=mred*mgreen*mblue
	    sumpow+=pow
fmt.Println(": ", mred,mgreen,mblue,pow)
	        //fmt.Println("ll: ", i, isok)
          }//end line
	}
	//total:=len(str_l);

	fmt.Println("countid", countid )
	fmt.Println("sumpow", sumpow )
        //os.Stdout.Write(l);
	//os.Stdout.Write(data)

}
