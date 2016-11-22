// iki tane pointer alır, değerlerini birbiri ile değiştirir
func swap(x *int, y *int) {
  temp := *x
  *x = *y
  *y = temp
  // pointerların değerlerini ekrana yazdırır
  fmt.Println(*x, *y)
}

func main() {
  x := int(1)
  y := int(2)
  // swap fonksiyonuna iki tane adres veriliyor
  swap(&x, &y)
}
