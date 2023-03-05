package main

import (
	"fmt"
	"strconv"
)

func munculsekali(angka string) []int{
count := make (map[int]map[int])
for _,

func main() {

	fmt.Println(munculsekali ("1234123"))  //[4]
	fmt.Println(munculsekali("76523752")) //[6,3]
	fmt.Println(munculsekali("12345"))     //[12345]
	fmt.Println(munculsekali("1122334455")) //[]
	fmt.Println(munculsekali("0872504"))     //[87254]
}
