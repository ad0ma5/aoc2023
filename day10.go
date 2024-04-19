package main

import (
	"log"
	"os"
	"fmt"
	"strings"
	//"sort"
	//"slices"
	//"unicode"
	"strconv"
)

///////
/*
The pipes are arranged in a two-dimensional grid of tiles:

| is a vertical pipe connecting north and south.
- is a horizontal pipe connecting east and west.
L is a 90-degree bend connecting north and east.
J is a 90-degree bend connecting north and west.
7 is a 90-degree bend connecting south and west.
F is a 90-degree bend connecting south and east.
. is ground; there is no pipe in this tile.
S is the starting position of the animal; there is a pipe on this tile, but your sketch doesn't show what shape the pipe has.
*/
///////
var pipes = []string{
"|",
"-",
"L",
"J",
"7",
"F"}

var np = []string{
".",
"S"}

var cup = []string{
"|",
"7",
"F"}
var cdn = []string{
"|",
"L",
"J"}
var clt = []string{
"-",
"L",
"F"}
var crt = []string{
"-",
"J",
"7",
}

type dot struct {
  x int
  y int
  c int
}

func inarr(they []string, it string) bool{
    for c:=0;c<len(they);c++{
      if it==they[c]{
	//it is next
	return true
      }
    }
  return false
}

func diffdot(a dot, b dot) (int, int){
  x:=a.x-b.x
  y:=a.y-b.y
  return x,y
}

func nextpb(last dot, cur dot, board [][]string, nb [][]string) dot {
  out:=cur
  cit:= board[cur.y][cur.x]

  //fmt.Println("------------------")
  //fmt.Println("cit")
  //fmt.Println(cit)
  //fmt.Println(cur)

  x,y:=diffdot(last, cur)
  //fmt.Println(x,y)
  ngo:=""
  if y == 1 {// it went up
    //fmt.Println("from up to ")
    if cit == "|"{
	ngo="up"
    }
    if cit == "7"{
	ngo="lt"
    }
    if cit == "F"{
	ngo="rt"
    }
  }
  if y == -1 {// it went down
    //fmt.Println("from down to ")
    if cit == "|" {
      ngo="dn"
    }
    if cit == "J" {
      ngo="lt"
    }
    if cit == "L" {
      ngo="rt"
    }
  }
  if x == 1 {// it went left
    //fmt.Println("from left to ")
    if cit == "-" {
      ngo="lt"
    }
    if cit == "L" {
      ngo="up"
    }
    if cit == "F" {
      ngo="dn"
    }

  }
  if x == -1 {// it went rt
    //fmt.Println("from rt to ")
    if cit == "-" {
      ngo="rt"
    }
    if cit == "J" {
      ngo="up"
    }
    if cit == "7" {
      ngo="dn"
    }

  }
  //fmt.Println("ngo")
  //fmt.Println(ngo)

  if ngo == "lt" {
    out.x=cur.x-1
    out.c++
    return out
  }
  if ngo == "rt" {
    out.x=cur.x+1
    out.c++
    return out
  }
  if ngo == "up" {
    out.y=cur.y-1
    out.c++
    return out
  }
  if ngo == "dn" {
    out.y=cur.y+1
    out.c++
    return out
  }

  fmt.Println("out")
  fmt.Println(out)

  nit:= board[out.y][out.x]
  fmt.Println(nit)
  os.Exit(3)
  return cur
}
func nextp(last dot, cur dot, board [][]string, nb [][]string) dot {
  out:=cur

  //up
  if cur.y > 0  {
    if last.x != cur.x || last.y != cur.y-1 {
      it:= board[cur.y-1][cur.x]
      if inarr(cup, it){
      fmt.Println(it)
        out.y=cur.y-1
        out.c++
        return out
      }
    }
  }
  //dn
  if cur.y < len(board)-1 {
    if last.x != cur.x || last.y != cur.y+1 {
      it:= board[cur.y+1][cur.x]
      if inarr(cdn, it){
      fmt.Println(it)
        out.y=cur.y+1
        out.c++
        return out
      }
    }
  }
  //lt
  if cur.x > 0 {
    if last.x != cur.x-1 || last.y != cur.y {
      it:= board[cur.y][cur.x-1]
      if inarr(clt, it){
      fmt.Println(it)
        out.x=cur.x-1
        out.c++
        return out
      }
    }
  }
  //rt
  if cur.x < len(board[0])-1 {

    if last.x != cur.x+1 || last.y != cur.y {
      it:= board[cur.y][cur.x+1]
      fmt.Println(it)
      if inarr(crt, it){
        out.x=cur.x+1
        out.c++
        return out
      }
    }
  }
      fmt.Println("o ",out)
      fmt.Println("olast ",last)
      fmt.Println("ocur ",cur)
      pbo(board)
      fmt.Println("oooooooooooooooooooooooooooooo ")
      pbo(nb)
      os.Exit(3)
  return out
}

func pbo(bo [][]string) {
  for i, value := range bo {

      fmt.Println(i," ",value)
  }

}
//6733
func main() {
	data, err := os.ReadFile("input10.txt")
	//data, err := os.ReadFile("input10_mid.txt")
	//data, err := os.ReadFile("input10_mid1.txt")
	//data, err := os.ReadFile("input10_small.txt")
	//data, err := os.ReadFile("input10_small1.txt")
	if err != nil {
	  log.Fatal(err)
	}
	str:=string(data);
	l_arr:=strings.Split(str, "\n")
	l:= len(l_arr);
	sdot:=dot{}
	board:=[][]string{}
	nb:=[][]string{}
	///// Extract dat
	for li:=0;li<l;li++{
	  if l_arr[li] == "" { continue }


	  //seq:=[][]int{}
	  no:=strings.Split(l_arr[li], "")
	  noc:=strings.Split(l_arr[li], "")
	  for ni:=0;ni<len(no);ni++{
	    if no[ni] == "S"{//start
	      sdot.x=ni
	      sdot.y=li
	      //fmt.Println(ni,li)
	      //break
	    }

	  }
	  board=append(board, no)
	  nb=append(nb, noc)
	  //fmt.Println(no)
	  fmt.Println(noc)
	  //fmt.Println(len(no))

	}
	//fmt.Println(board)
	//DO
	//Action
	//done:=false
	crawl1:=sdot
	crawl2:=sdot
	l1:=sdot
	l2:=sdot
	count:=0
	//for !done {

	  if count==0{
	    fmt.Println("S",sdot)
	    c1:=nextp(sdot, sdot, board, nb)
	    nb[c1.y][c1.x] = strconv.Itoa(c1.c)
	    crawl1=c1
	    fmt.Println("S",sdot)
	    c2:=nextp(crawl1,sdot, board, nb)
	    nb[c2.y][c2.x] = strconv.Itoa(c2.c)
	    crawl2=c2
          }
	  leave := false
	  for !leave {
	    //fmt.Println("==========")
	    //fmt.Println("c1 before",crawl1)
	    c1:=nextpb(l1, crawl1, board, nb)
	    //nb[c1.y][c1.x] = strconv.Itoa(c1.c)
	    nb[c1.y][c1.x] = strconv.Itoa(0)
	    l1=crawl1
	    crawl1=c1
	    //fmt.Println("c1aafter",crawl1)
	    //fmt.Println("==========")

	    //fmt.Println("c2 before",crawl2)
	    c2:=nextpb(l2, crawl2, board, nb)
	    nb[c2.y][c2.x] = strconv.Itoa(9)
	    l2=crawl2
	    crawl2=c2
	    //fmt.Println("c2 after",crawl2)
	    //fmt.Println("==========")

	    if crawl1.x == crawl2.x && crawl1.y == crawl2.y{
	      fmt.Println("EQUAL")
	      leave = true
	    }

	  }
          //done=true
	//}

	for i:=0;i<len(nb);i++{
		fmt.Println(i,":",nb[i])
        }


	fmt.Println(crawl1,crawl2)


	/////
	//fmt.Println(l, l_arr)

        //os.Stdout.Write(l);
	//os.Stdout.Write(data)

}
