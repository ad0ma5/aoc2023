package main

import (
	"log"
	"os"
	"fmt"
	"unicode"
	"strconv"
)

func compareTextNum(bytes [5]byte) (int){
numbers:=[9]string{"one","two","three","four","five","six","seven","eight","nine"}
  for i:=0; i<len(numbers); i++{
    for j:=0; j<len(numbers[i]);j++{

	//fmt.Println("fors", i,  j , string(bytes[j]), bytes)

      if string(numbers[i][j]) == string(bytes[j]){
        if len(numbers[i]) == j+1{
          return i+1
        }
      }else{
        break
      }
    }
  }
  return 0
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	str:=string(data);
	l:= len(str);
    first:=string("x")
    last:=string("x")
    both:=string("")
	_ = first
	_ = last
	_ = both
	total:=0
	both_int:=0
	for i:= 0;i<l;i++{
	  b:=str[i];
	  fmt.Printf("b=%c\n", b)
	  isn:=unicode.IsNumber(rune(b))
	  iss:=string(b)==string("\n")
	  if isn {
	    fmt.Printf("%q looks like a number.\n", b)
	    last=string(b);
	    if first == string("x"){ first=string(b) }
	  }else{
	    if i+4 < l{
	      tt:=[5]byte{str[i],str[i+1],str[i+2],str[i+3],str[i+4]}
	//fmt.Println("sub", tt)
	    ist:=compareTextNum(tt)
		  if ist>0 {

	    last=strconv.Itoa(ist);
	    if first == string("x"){ first=strconv.Itoa(ist) }
		    fmt.Printf("%d looks like a text number.\n", ist)
		  }
	    }
	  }
	  if iss {
	    both=first+last
	    fmt.Printf("%q looks like a \\n.\n", b)
	    fmt.Printf("%q looks like first \n", first)
	    fmt.Printf("%q looks like last \n", last)
	    fmt.Printf("%q looks like both \n", both)
	    both_int, _ = strconv.Atoi(both)
	    total += both_int
	    first=string("x")
	    last=string("x")

	  }
        }
	fmt.Println("range", l,  total )
        //os.Stdout.Write(l);
	//os.Stdout.Write(data)

}
