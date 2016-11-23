//
// var ile degisken tanimlama:
//   var degisken_adi tipi
//   var ile birden fazla degisken tek seferde tanimlanabilir.
//
// := ile degisken tanimlama:
// degisken tanimlamanin kisa yoludur. Degiskene deger atanmasi gerekir.
// tip belirtme yoktur ancak soyle bir kullanimda var:
//	a := int(5) a degerine 5 degerini atar.
// 

package main

import "fmt"

func main() {

  var msg string = "Merhaba, nasilsin?"
  fmt.Println(msg)
  var a, b int = 1, 2
  fmt.Println(a, b)
  var kontrol = true
  fmt.Println(kontrol)
  var bos int
  fmt.Println(bos)
  kisa := 45
  fmt.Println(kisa)
  kisa_float := float64(0)
  fmt.Println(kisa_float)

}

/* CIKTI:
$ go run variables.go
Merhaba, nasilsin?
1 2
true
0
45
0
*/
