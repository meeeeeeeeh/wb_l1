//Разработать программу, которая перемножает, делит,
//складывает, вычитает две числовых переменных a,b, значение которых > 2^20 (???)

package main

import (
	"errors"
	"fmt"
	"math/big"
)

func sum(num1 *big.Int, num2 *big.Int) *big.Int {
	res := new(big.Int)
	res.Add(num1, num2)

	return res
}

func sub(num1 *big.Int, num2 *big.Int) *big.Int {
	res := new(big.Int)
	res.Sub(num1, num2)

	return res
}

func mult(num1 *big.Int, num2 *big.Int) *big.Int {
	res := new(big.Int)
	res.Mul(num1, num2)

	return res
}

func div(num1 *big.Int, num2 *big.Int) (*big.Int, error) {
	// конвертация в int64
	if num2.Int64() == 0 {
		err := errors.New("zero division")
		return nil, err
	}

	res := new(big.Int)
	res.Div(num1, num2)

	return res, nil
}

// для 2^20 ьожно использовать int64
// но если имелось в виду 10^22 то произойдет переполнение int64
// тогда надо исползовать пакет math/big
func main() {
	// принимает либо int64 либо строку и систему в которую надо перевести
	// здесь будет 10-ричная система
	num1 := new(big.Int)
	num1.SetString("24000000000000000000", 10)

	num2 := new(big.Int)
	num2.SetString("6000000000000000000", 10)

	fmt.Println(sum(num1, num2))

	fmt.Println(sub(num1, num2))
	fmt.Println(mult(num1, num2))
	fmt.Println(div(num1, num2))

	num2.SetString("0", 10)
	fmt.Println(div(num1, num2))

}
