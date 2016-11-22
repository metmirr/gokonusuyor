// Go'da bir string sadece okunabilir(read-only) bytlardan oluşmuş bir parçadır(slice).
// Ve string Unicode text, utf-8 text veya herhangi bir ön tanımlı formata sahip değildir.
package main

import "fmt"

func main() {
	msg := "Merhaba"
	for n, s := range msg {
		fmt.Println(n, "->", s)
	}
}
/*  ÇIKTI:
  0 -> 77
  1 -> 101
  2 -> 114
  3 -> 104
  4 -> 97
  5 -> 98
  6 -> 97
*/
