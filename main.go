/// struct, interface tanımlamaları:
///	her ikiside type anahtar kelimesi, isim ve sırasıyla struct, interface alarak tanımlanırlar
///	oluşturulan structlara interface uygulanması için interface de tanımlanan metot, her struct için
///	bir metot olmalıdır. Yani tanımlama şu şekilde olmalıdır: 
///		func (degisken_adi Struct_Adi) interfacedeki_metot_ismi() varsa_geri_donus_degeri {}
///	Not: Bu şekilde tanımlanan her metot ilgili struct'ın belirtilen interface'yi uyguladığı anlamına gelir. Önemli olan
///	bu metot isminin aynı olmasıdır.
///	Bu metotlar pointer struct olarak da belirtilebilirler:
///		func (degisken_adi *Struct_Adi) interfacedeki_metot_ismi() varsa_geri_donus_degeri {}

package main

import (
	"fmt"
	"math"
)

type Daire struct {
	a float64
	b float64
	r float64
}
type Dikdortgen struct {
	genislik float64
	yukseklik float64
}
func (d Dikdortgen) alan() float64 {
	return d.genislik * d.yukseklik
}
func (d Daire) alan() float64 {
	return math.Pi * d.r * d.r
}

type Sekil interface {
	 alan() float64
}
type CokSekil struct {
	sekiller []Sekil
}	

func main() {
	cokSekil := CokSekil {
		sekiller: []Sekil{
			Daire { 0, 0, 5},
			Dikdortgen{ 10, 10 },
		},
	}
	for num, sek := range cokSekil.sekiller {
		fmt.Println(name, sek.alan())
	}
}
