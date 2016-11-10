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

type Circle struct {
	x float64
	y float64
	r float64
}
type Rectangle struct {
	width float64
	height float64
}
func (r Rectangle) area() float64 {
	return r.width * r.height
}
func (c Circle) area() float64 {
	return math.Pi * c.r * c.r
}

type Shape interface {
	 area() float64
}
type MultipleShape struct {
	shapes []Shape
}	

func main() {
	multiShape := MultipleShape {
		shapes: []Shape{
			Circle { 0, 0, 5},
			Rectangle{ 10, 10 },
		},
	}
	for name, v := range multiShape.shapes {
		fmt.Println(name, v.area())
	}
	fmt.Println()
}
