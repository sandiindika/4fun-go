package basics

import "fmt"

func Fundamentals() {
	fmt.Println("Hello, World!")

	//	todo: integer(1) declaration

	// int8: -128 to 127
	// int16: -32,768 to 32,767
	// int32: -2,147,483,648 to 2,147,483,647
	// int64: -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807

	var i int8 = 127
	fmt.Printf("The value of i is %d\n", i)

	// todo: integer(2) declaration

	// uint8: 0 to 255
	// uint16: 0 to 65,535
	// uint32: 0 to 4,294,967,295
	// uint64: 0 to 18,446,744,073,709,551,615

	var u uint8 = 255
	fmt.Printf("The value of u is %d\n", u)

	// todo: float declaration

	// float32: 1.18e-38 to 3.4e+38
	// float64: 2.23e-308 to 1.80e+308

	var f float32 = 3.14159
	fmt.Printf("The value of f is %.2f\n", f)

	// todo: complex declaration

	// complex64: complex numbers with float32 real and imaginary parts
	// complex128: complex numbers with float64 real and imaginary parts

	var c complex64 = 2 + 3i
	fmt.Printf("The value of c is %.2f + %.2fi\n", real(c), imag(c))

	// todo: alias

	// byte is an alias for uint8
	// rune is an alias for int32
	// int is an alias for minimal int32
	// uint is an alias for minimal uint32

	// todo: boolean declaration

	var b bool = true
	fmt.Printf("The value of b is %t\n", b)

	// todo: string declaration

	var s string = "Hello, World!"
	fmt.Printf("The value of s is %s\n", s)
	fmt.Println("Length of s:", len(s))
	fmt.Println("First character of s:", s[0])

	// todo: keyword var

	// var san string = "Sandi Indika"
	var san = "Sandi Indika"
	fmt.Println("The value of san is:", san)

	di := "Saputra"
	fmt.Println("The value of di is:", di)

	// todo: variable shadowing

	var (
		x float64 = 3.14
	)

	fmt.Printf("The value of x is %f and the type is %T\n", x, x)
}
