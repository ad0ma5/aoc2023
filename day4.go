package main

import (
	"log"
	"os"
	"fmt"
	"strings"
	//"unicode"
	//"strconv"
)



type gear struct {
  i int
  j int
  no1 int
  no2 int
}

func main() {
	//data, err := os.ReadFile("input4_small.txt")
	data, err := os.ReadFile("input4.txt")
	if err != nil {
	  log.Fatal(err)
	}
	str:=string(data);
	line_arr:= strings.Split(str, "\n");

	gran_sum:=0
	for l:=0;l<len(line_arr);l++{
	  line_w := 0
	  if len(line_arr[l]) == 0 { break }
	  l_part:= strings.Split(line_arr[l], " | ");
	  w_part:= strings.Split(l_part[0], ": ");
	  try_numbers_str:=strings.TrimSpace(l_part[1])
	  try_arr:= strings.Split(try_numbers_str, " ");
	  win_numbers_str:=strings.TrimSpace(w_part[1])
	  win_arr:= strings.Split(win_numbers_str, " ");
	  for t:=0;t<len(try_arr);t++{
	    for w:=0;w<len(win_arr);w++{
	      //fmt.Print(":" , try_arr[t],"=", win_arr[w])
	      if try_arr[t] == win_arr[w] && try_arr[t] != ""{
	        fmt.Print(":" , try_arr[t],"=", win_arr[w])
		if line_w == 0 { line_w = 1 } else { line_w += line_w }
	        fmt.Println( "win[ ", line_w, " ]")
		break
	      }
            }
          }
	  gran_sum += line_w
	  fmt.Println("l w:" , line_w)

	}//for line_arr
	fmt.Println(" BW: ", gran_sum, " = 101660 Bad ")
	//fmt.Println( str)
	//total:=len(str_l);

        //os.Stdout.Write(l);
	//os.Stdout.Write(data)

}
