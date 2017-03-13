package main
import (
	"fmt"
	//"math/rand"
	//"math"
	"math/cmplx"
	"math"
)

func main() {
	//fmt.Println("hello go")
	//fmt.Println("My favorite number is",rand.Intn(10))
	//fmt.Printf("Now you have %g problems.", math.Nextafter(2, 3))
	//fmt.Println(math.Pi)
	//fmt.Println(add(2,3))
	//a, b := swap("hello", "word")
	//fmt.Println(a, b)
	//fmt.Println(split(17))
	//variable()
	//variableInit()
	//variableDeclare()
	//baseTypes()
	//zeroVal()
	//typeConvert()
	//valConst()
	//forTest()
	fmt.Println(sqrt(2), sqrt(-4))
}

func add(x,y int) int {
	return x + y
}

func swap(x,y string)(string,string) {
	return y, x
}

func split(sum int)(x,y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func variable() {
	var c, python, java bool
	var i int
	fmt.Println(i, c, python, java)
}

func variableInit() {
	var i, j = 1, 2
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}

func variableDeclare() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"
	fmt.Println(i, j, k, c, python, java)
}

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)
func baseTypes() {
	const f = "%T(%v)\n"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)
}

func zeroVal(){
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

func typeConvert()  {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z int = int(f)
	fmt.Println(x, y, z)
}

const (
	Big   = 1 << 100
	Small = Big >> 99 //Small实际等于1左移1位
)
func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}
func valConst()  {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

func forTest() {
	//sum := 0
	//for i := 0; i < 10; i++ {
	//	sum += 1
	//}
	//fmt.Println(sum)

	sum1 := 1
	//count := 0
	for sum1 < 1000 {
		sum1 += sum1
		//count += 1
		fmt.Println(sum1)
	}
	fmt.Println(sum1)
	//fmt.Println(count)
}

func sqrt(x float64) string {
	if x < 0{
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}