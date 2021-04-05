package main

//func main() {
//	c := Calculator{2, 2}
//	print(c.addition())
//}

type Calculator struct {
	v1, v2 int
}

func (c *Calculator) division() int {
	return c.v1 / c.v2
}

func (c *Calculator) multiplication() int {
	return c.v1 * c.v2
}

func (c *Calculator) addition() int {
	return c.v1 + c.v2
}

func (c *Calculator) subtraction() int {
	return c.v1 - c.v2
}
