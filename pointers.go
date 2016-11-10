// takes two pointer, change their value with each other
func swap(x *int, y *int) {
  temp := *x
  *x = *y
  *y = temp
  
  fmt.Println(*x, *y)
}

func main() {
  x := int(1)
  y := int(2)
  // addresses are passing to the swap function
  swap(&x, &y)
}
