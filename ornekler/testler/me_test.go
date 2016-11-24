// Go Test
//
// Go ile yazilan kodlari test etmek icin:
//   
//   - test kodlarini icerecek dosyanin ismi `_test.go` ile bitmeli ve `testing` paketi import edilmeli.
//   - test fonksiyonlarin ismi `Test` ismi ile baslamali ve bu fonksiyonlar parametre olarak
//     `*testing.Test` tipinde bir parametre almalidir.
//   - hata mesajlari icin fonksiyona verilen parametrenin metodu olan `Error()` fonksiyonu kullanilabilir.
//
// Test'i baslatmak icin: $ go test
// Birden fazla test dosyasi kullanacaksaniz her dosyadaki paket ismi ayni olmalidir. Diger turlu hata alirsiniz.

package me

import "math"
import "testing"

func TestMath(t *testing.T) {
	if 4 != math.Pow(2, 2) {
		t.Error("Not equals!")
	}
	if add(3) > 1 {
		t.Error("Threshold violation")
	}
}

func add(a int) int {
	return a
}
