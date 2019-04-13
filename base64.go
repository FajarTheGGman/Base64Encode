package main

import "fmt"
import "encoding/base64"

func main(){
fmt.Printf("Base64 Encode By Fajar Firdaus\n")
fmt.Printf(" \n")
fmt.Printf("    _______         [======================]\n")
fmt.Printf("|==|_______D         Coder : Fajar Firdaus\n")
fmt.Printf("       |             Fb : Fajar Firdaus\n")
fmt.Printf("       |             IG : fajar_firdaus_7\n")
fmt.Printf("       |            [======================]\n")
	var in string
	fmt.Printf("Masukan Text : ")
	fmt.Scan(&in)
	var s = base64.StdEncoding.EncodeToString([]byte(in))
	fmt.Println("Hasil : " + s);
}