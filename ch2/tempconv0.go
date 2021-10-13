// tempconv パッケージは摂氏と華氏の温度計算を行う
package main

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.5
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CToF(c Celsius) Fahrenheit  { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius  { return Celsius((f - 32) * 5 / 9) }
func (c Celsius) String() string { return fmt.Sprintf("%gC", c) }

func main() {
	fmt.Printf("%g\n", BoilingC-FreezingC) // 100 C
	boilingF := CToF(BoilingC)
	fmt.Printf("%g\n", boilingF-CToF(FreezingC)) //180 F
	//	fmt.Printf("g\n", boilingF-FreezingC) // コンパイルエラー

	//実装したメソッドの動作確認
	c := FToC(212.0)
	fmt.Println(c.String()) //100C
	fmt.Printf("%v\n", c)   //100C 明示的に呼び出す必要はない
	fmt.Printf("%s\n", c)   //100C
	fmt.Println(c)
	fmt.Printf("%g\n", c)   //100; String() を呼び出さない
	fmt.Println(float64(c)) //100; String() を呼び出さない
}
