package main
import (
	"fmt"
	//"math/rand"
	//"math"
	"math/cmplx"
	"math"
	"runtime"
	"time"
	"strings"
	"io"
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
	//fmt.Println(sqrt(2), sqrt(-4))
	//testPow()
	//testPow1()
	//switchTest()
	//switchTest1()
	//switchTest2()
	//deferTest()
	//deferTest1()
	//pointerTest()
	//structTest()
	//structTest1()
	//structTest2()
	//arrayTest()
	//sliceTest()
	//sliceTest2()
	//sliceTest3()
	//sliceTest4()
	//sliceTest5()

	//methodsTest()
	//methodsTest1()
	//methodsTest2()
	//stringTest()
	//errorTest()
	readerTest()
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

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}
func testPow()  {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

func pow1(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// 这里开始就不能使用 v 了
	return lim
}

func testPow1()  {
	fmt.Println(
		pow1(3, 2, 10),
		pow1(3, 3, 20),
	)
}

func switchTest()  {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
}

func switchTest1()  {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {//此处为常量，换成Monday则进第一个分支
	case today + 0://case分支中的条件为变量
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

func switchTest2()  {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func deferTest() {
	defer fmt.Println("world")
	fmt.Println("hello")
}

func deferTest1()  {
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}

func pointerTest()  {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}

type Vertex struct {
	X int
	Y int
}
func structTest() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}
func structTest1()  {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}
func structTest2()  {
	var (
		v1 = Vertex{1, 2}  // 类型为 Vertex
		v2 = Vertex{X: 1}  // Y:0 被省略
		v3 = Vertex{}      // X:0 和 Y:0
		p  = &Vertex{1, 2} // 类型为 *Vertex
	)
	fmt.Println(v1, p, v2, v3)
}

func arrayTest()  {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
}

func sliceTest()  {
	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("p ==", p)

	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d] == %d\n", i, p[i])
	}
}

func sliceTest2()  {
	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("p ==", p)
	fmt.Println("p[1:4] ==", p[1:4])//4代表从头开始4个元素

	// 省略下标代表从 0 开始
	fmt.Println("p[:3] ==", p[:3])

	// 省略上标代表到 len(s) 结束
	fmt.Println("p[4:] ==", p[4:])

	fmt.Println("p[2:3] ==", p[2:3])//[5]
	fmt.Println("p[2:3] ==", p[2:2])//[]
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
func sliceTest3()  {
	a := make([]int, 5)//初始元素为5个，初始值都为0,cap为5
	printSlice("a", a)
	b := make([]int, 0, 5)//初始元素为0个，cap为5
	printSlice("b", b)
	c := b[:2]
	printSlice("c", c)
	d := c[2:5]
	printSlice("d", d)
}

func sliceTest4()  {
	var z []int
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil!")
	}
}

func sliceTest5()  {
	var a []int
	printSlice("a", a)

	// append works on nil slices.
	a = append(a, 0)
	printSlice("a", a)

	// the slice grows as needed.
	a = append(a, 1)
	printSlice("a", a)

	// we can add more than one element at a time.
	a = append(a, 2, 3, 4)
	printSlice("a", a)
}

func rangeTest()  {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

func rangeTest1()  {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i)
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

func mapsTest()  {
	type Vertex struct {
		Lat, Long float64
	}
	var m map[string]Vertex
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}

func mapsTest1()  {
	type Vertex struct {
		Lat, Long float64
	}
	var m = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}
	fmt.Println(m)
}

func mapsTest2()  {
	type Vertex struct {
		Lat, Long float64
	}
	var m = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(m)
}

func mapsTest3()  {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	//delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

func functionValues()  {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(3, 4))
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
func functionClosures()  {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

type Vertex1 struct {
	X, Y float64
}
func (v *Vertex1) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func methodsTest()  {
	v := &Vertex1{3, 4}
	fmt.Println(v.Abs())
}

type MyFloat float64
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
func methodsTest1()  {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}

func (v *Vertex1) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
func (v Vertex1) Scale1(f float64) {
	//当 v 是一个值（非指针），方法看到的是 Vertex 的副本，并且无法修改原始值
	v.X = v.X * f
	v.Y = v.Y * f
}
func (v *Vertex1) Abs1() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func methodsTest2() {
	v := &Vertex1{3, 4}
	//v.Scale(5)
	v.Scale1(5)
	fmt.Println(v, v.Abs1())
}


type Person struct {
	Name string
	Age  int
}
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}
func stringTest()  {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}


type MyError struct {
	When time.Time
	What string
}
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}
func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}
func errorTest()  {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

func readerTest()  {
	r := strings.NewReader("Hello, Reader!")
	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}