package main

import (
	"log"
	"os"
	"fmt"
	"strings"
	"unicode"
	"strconv"
)


//Of course, the actual engine schematic is much larger. What is the sum of all of the part numbers in the engine schematic?
func isd(c byte) bool {
  return string(c)==string(".")
}

func isn(c byte) bool {
  return unicode.IsNumber(rune(c))
}

func isp(c byte) bool {
  return !isd(c) //&& !isn(c)
}

func ispn(c byte) bool {
  return !isd(c) && !isn(c)
}

func iss(c byte) bool {
  return string(c) == string("*")
}

type gear struct {
  i int
  j int
  no1 int
  no2 int
}

func main() {
	data, err := os.ReadFile("input3.txt")
	//data, err := os.ReadFile("day3_small.txt")
	if err != nil {
		log.Fatal(err)
	}
	str:=string(data);
	str_l:= strings.Split(str, "\n");

	sum:=0

	  number_add :=false;
	  number := string("")
	  the_gear := gear{}
	  has_gear:=false
	  all_gear:= []gear{}
	for i:=0;i<len(str_l);i++{
	  fmt.Println()
	  fmt.Println("l:", str_l[i], i )
          for j:=0;j<len(str_l[i]);j++{
	    c:=str_l[i][j];
	    if ( !isn(c)  || j == 0) && number != string("")  {
	      fmt.Print(":", number  )
	      if number_add {
                var n, _ =strconv.Atoi(number)
		sum += n
	        fmt.Print(" :sum:", sum, " ")
		if has_gear {
		  for g:=0;g<len(all_gear);g++{
		  if all_gear[g].i == the_gear.i && all_gear[g].j == the_gear.j {
		    all_gear[g].no2 = n
		    has_gear = false
		    the_gear = gear{}
		    break
	          }
	          }
		  if has_gear {//not found above so still has
		    the_gear.no1 = n
		    all_gear = append(all_gear, the_gear)
		    has_gear = false
		    the_gear = gear{}
	          }
	        }
	      }

	      number = string("")
	      number_add =false;
	      //continue;
	    }

	    if isn(c) { number +=string(c) }
	    //does it touch
	    if isn(c) && i > 0 && j > 0 {
              if isp(str_l[i-1][j-1]) {
                //it touch that
		number_add = true
	      }
              if iss(str_l[i-1][j-1]) {
		the_gear= gear{i:i-1,j:j-1}
	        has_gear = true
	      }
	    }
	    if isn(c) && i > 0 {
              if isp(str_l[i-1][j]) {
                //it touch that
		number_add = true
	      }
              if iss(str_l[i-1][j]) {
		the_gear= gear{i:i-1,j:j}
	        has_gear = true
	      }
	    }
	    if isn(c) && i > 0 && j < len(str_l[i])-2{
              if isp(str_l[i-1][j+1]) {
                //it touch that
		number_add = true
	      }
              if iss(str_l[i-1][j+1]) {
		the_gear= gear{i:i-1,j:j+1}
	        has_gear = true
	      }
	    }
	    if isn(c) && i < len(str_l)-2 && j > 0 {
              if isp(str_l[i+1][j-1]) {
                //it touch that
		number_add = true
	      }
              if iss(str_l[i+1][j-1]) {
		the_gear= gear{i:i+1,j:j-1}
	        has_gear = true
	      }
	    }
	    if isn(c) && i < len(str_l)-2 {
              if isp(str_l[i+1][j]) {
                //it touch that
		number_add = true
	      }
              if iss(str_l[i+1][j]) {
		the_gear= gear{i:i+1,j:j}
	        has_gear = true
	      }
	    }
	    if isn(c) && i  < len(str_l)-2  && j < len(str_l[i])-2{
              if isp(str_l[i+1][j+1]) {
                //it touch that
		number_add = true
	      }
              if iss(str_l[i+1][j+1]) {
		the_gear= gear{i:i+1,j:j+1}
	        has_gear = true
	      }
	    }

	    if isn(c) && j > 0 {
              if ispn(str_l[i][j-1]) {
                //it touch that
		number_add = true
	      }
              if iss(str_l[i][j-1]) {
		the_gear= gear{i:i,j:j-1}
	        has_gear = true
	      }
            }
	    if isn(c) && j < len(str_l[i])-2 {
              if ispn(str_l[i][j+1]) {
                //it touch that
		number_add = true
	      }
              if iss(str_l[i][j+1]) {
		the_gear= gear{i:i,j:j+1}
	        has_gear = true
	      }
            }
            //done above :D
	    //if number_add { number += string(c) }
	    //fmt.Println("c:", string(c),isn(c) , isd(c), number_add, number)

	  }
	}
	all_gsum:=0
	for ind:=0; ind<len(all_gear);ind++{
	  all_gsum += all_gear[ind].no1 * all_gear[ind].no2
        }
	    fmt.Println("sum:", sum, " gears", len(all_gear) , all_gsum)
	//total:=len(str_l);

        //os.Stdout.Write(l);
	//os.Stdout.Write(data)

}
