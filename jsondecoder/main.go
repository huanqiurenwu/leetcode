package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	//"unsafe"
)

type people struct {
	name string
	age int
}

func (peo *people)getName(){
	fmt.Println(peo.name, " ", peo)
}
func (peo people)getAge(){
	fmt.Println(peo.age, " ", &peo)
}

//func main() {
//	numbers := []int{12, 7, 21, 12, 12, 26, 25, 21, 30}
//	fmt.Println("过滤之前：\n\t", numbers)
//	// 三个非接口值被包裹在一个Filter切片
//	// 的三个接口元素中。
//	filters := []Filter{
//		0:UniqueFilter{},
//		1:MultiFilter(2),
//		2:MultiFilter(3),
//	}
//	// 每个切片元素将被赋值给类型为Filter的
//	// 循环变量fltr。每个元素中的动态值也将
//	// 被同时复制并被包裹在循环变量fltr中。
//	for _, fltr := range filters {
//		numbers = filterAndPrint(fltr, numbers)
//	}
//}

type Person struct {
	Name string
	Age int
}
func (p Person) PrintName() {
	fmt.Println("Name:", p.Name)
}
func (p *Person) SetAge(age int) {
	p.Age = age
}

type Singer struct {
	Person // 通过内嵌Person类型来扩展之
	works []string
}
//
//func main() {
//	t := reflect.TypeOf(Singer{}) // the Singer type
//	fmt.Println(t, "has", t.NumField(), "fields:")
//	for i := 0; i < t.NumField(); i++ {
//		fmt.Print(" field#", i, ": ", t.Field(i).Name, "\n")
//	}
//	fmt.Println(t, "has", t.NumMethod(), "methods:")
//	for i := 0; i < t.NumMethod(); i++ {
//		fmt.Print(" method#", i, ": ", t.Method(i).Name, "\n")
//	}
//	pt := reflect.TypeOf(&Singer{}) // the *Singer type
//	fmt.Println(pt, "has", pt.NumMethod(), "methods:")
//	for i := 0; i < pt.NumMethod(); i++ {
//		fmt.Print(" method#", i, ": ", pt.Method(i).Name, "\n")
//	}
//}

type I interface {
	m()
}

type T struct{
	I
}

//func main(){
//	var t T
//	var i = &t
//	t.I = i
//	i.m()
//}

//func main(){
//	var x float64 = 341434.444443243
//	f := math.Float64bits(x)
//	ft := uint64(f)
//	fmt.Println(f)
//	fmt.Println(ft)
//	type MyString string
//	ms := []MyString{"C", "C++", "Go"}
//	fmt.Printf("%s\n", uintptr(unsafe.Pointer(&ms)))
//	ss := *(*[]string)(unsafe.Pointer(&ms))
//	ss[1] = "Rust"
//	fmt.Printf("%sv\n", uintptr(unsafe.Pointer(&ss)))
//}

//func main(){
//	a := [...]byte{'G', 'o', 'l', 'a', 'n', 'g'}
//	s := "Java"
//	hdr := (*reflect.StringHeader)(unsafe.Pointer(&s))
//	hdr.Data = uintptr(unsafe.Pointer(&a))
//	hdr.Len = len(a)
//	fmt.Println(s)
//	a[2], a[3], a[4], a[5] = 'o', 'g', 'l', 'e'
//	fmt.Println(s)
//}

//func main(){
//	a := [...]byte{'G', 'o', 'l', 'a', 'n', 'g'}
//	bs := []byte("Go101")
//	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&bs))
//	hdr.Data = uintptr(unsafe.Pointer(&a))
//	hdr.Len = len(a)
//	hdr.Cap = len(a)
//	fmt.Println("%s\n", bs)
//}

type F func(string, int) bool
func (f F) m(s string) bool{
	return f(s, 32)
}
func (f F)M(){}
type II interface{m(s string) bool; M()}

//func main() {
//	var x struct {
//		F F
//		i II
//	}
//	tx := reflect.TypeOf(x)
//	fmt.Println(tx.Kind())
//	fmt.Println(tx.NumField())
//	fmt.Println(tx.Field(1).Name)
//	fmt.Println(tx.Field(0).PkgPath)
//	fmt.Println(tx.Field(1).PkgPath)
//
//	tf, ti := tx.Field(0).Type, tx.Field(1).Type
//	fmt.Println("tf::::::", tf.Kind())
//
//	fmt.Println(tf.IsVariadic())         // false
//	fmt.Println(tf.NumIn(), tf.NumOut()) // 2 1
//	t0, t1, t2 := tf.In(0), tf.In(1), tf.Out(0)
//	// 下一行打印出：string int bool
//	fmt.Println(t0.Kind(), t1.Kind(), t2.Kind())
//	fmt.Println(tf.NumMethod(), ti.NumMethod()) // 1 2
//	fmt.Print(tf.Method(0).Name + " ")              // M
//	fmt.Println(ti.Method(0).Name)              // m
//	_, ok1 := tf.MethodByName("m")
//	_, ok2 := ti.MethodByName("m")
//	fmt.Println(ok1, ok2) // false true
//}

type TAG struct {
	X int `max:"99" min:"0" default:"0"`
	Y, Z bool `optional:"yes"`
}

/*func main(){
	t := reflect.TypeOf(TAG{})
	x := t.Field(0).Tag
	y := t.Field(1).Tag
	z := t.Field(2).Tag
	fmt.Println(reflect.TypeOf(x))
	v, present := x.Lookup("max")
	fmt.Println(len(v), present)
	fmt.Println(x.Get("max"))
	fmt.Println(x.Lookup("optional"))
	fmt.Println(y.Lookup("optional"))
	fmt.Println(z.Lookup("optional"))
	var sl []int
	sl = append(sl, 2)
	fmt.Println(sl)
}*/

//func main(){
//	matchString()
//	input := []int{1,2,3,4}
//	fmt.Println(subsets1(input))
//	fmt.Println(combine(4,2))
//	f1()
//}

func say(s string) {

	for i := 0; i < 2; i++ {

		//runtime.Gosched()

		fmt.Println(s)

	}

}

var p0, p1, p2, p3, p4, p5 *int
var x = 9999 // x#0
func main() {
	bo, err := strconv.ParseBool("true")
	// 9999 888 77 6 3 -3
	fmt.Println(bo, " ", err)
}

//

func modify(input []int){
	input[1] = 1
	input[2] = 2
	input[3] = 3
}

func subsets1(in []int) [][]int{
	if len(in) == 0{
		return [][]int{[]int{}}
	}
	n := in[0]
	res := subsets1(in[1:])
	for _, v := range res{
		res = append(res, v)
		res[len(res) - 1] = append(res[len(res) - 1],n)
	}
	return res
}

func combine(n int, k int) [][]int {
	result := make([][]int, 0, n*(n-1)/k) //总数n*(n-1)/k
	combineBacktrace([]int{}, k, 1, n, &result)
	return result
}

func combineBacktrace(path []int, k int, firstRemain int, lastRemain int, result *[][]int){
	if len(path) == k{
		*result = append(*result, path)
		return
	}
	for i:= firstRemain; i<= lastRemain; i++ {
		newPath := make([]int, 0, k)
		newPath = append(newPath, path...)
		newPath = append(newPath, i)
		combineBacktrace(newPath, k, i+1, lastRemain, result)
	}
}
//func isPalindrome(input int) bool{
//	str := strconv.Itoa(input)
//	var sl []int
//	strHdr := (*reflect.StringHeader)(unsafe.Pointer(&str))
//	slHdr := (*reflect.SliceHeader)(unsafe.Pointer(&sl))
//	slHdr.Data = strHdr.Data
//	slHdr.Len = strHdr.Len
//	slHdr.Cap =  strHdr.Len
//	for i,j := 0, len(sl)-1; i < j; i,j = i+1, j-1{
//		if sl[i] != sl[j]{
//			return false
//		}
//	}
//	return true
//}


func matchString(){
	fmt.Println(strings.Contains("network is unreachable haha", ""))
}

func subsets(in []int)[][]int {
	result := make([][]int,0, uint64(math.Sqrt(float64(len(in)))))
	result = append(result, []int{})
	zuhe([]int{}, in, &result)
	return result
}

func zuhe(baseHead, remain []int, result *[][]int){
	if len(remain) > 0 {
		for i,_ := range remain {
			addedHead := append([]int{}, baseHead...)
			addedHead = append(addedHead, remain[i])
			*result = append(*result, addedHead)
			zuhe(addedHead, remain[i+1:], result)
		}
	}
}

func lengthOfLongestSubstring(s string)(result int) {
	return
}

//func str2byte(in string) []byte{
//	unsafe.Pointer(&in)
//}

func f0() int {
	var x = 1
	defer fmt.Println("正常退出：", x)
	x++
	return x
}

func f1() {
	var x = 1
	defer fmt.Println("正常退出：", x)
	x++
}
//func isPalindrome(input int) bool {
//
//}
	//c := make(chan *[16]byte)
	//go func() {
	//	// 使用两个数组以避免数据竞争。
	//	var dataA, dataB = new([16]byte), new([16]byte)
	//	for {
	//		_, err := rand.Read(dataA[:])
	//		if err != nil {
	//			close(c)
	//		} else {
	//			c <- dataA
	//			dataA, dataB = dataB, dataA
	//		}
	//	}
	//}()
	//for data := range c {
	//	fmt.Println(data[:])
	//	time.Sleep(time.Second / 2)
	//}


//func subs(input []int)[]int {
//	result := make([][]int, 0, int64(math.Sqrt(float64(len(input)))))
//	result = append(result, []int{}, []int{input[len(input)-1]})
//	for keyOutside := 0; keyOutside < len(input)-1; keyOutside++ {
//		initEle := []int{input[keyOutside]}
//		result = append(result, initEle)
//		for keyInside, valueInside := range input[keyOutside+1:]{
//
//			result = append(result, )
//		}
//	}
//}

func q1_1(){
	for i:=0; i< 10; i++ {
		fmt.Println(i)
	}
}

func q1_2(arr []int){
	i:=0
	Loop:
		{
			if i<len(arr){
				fmt.Println(arr[i])
				i++
				goto Loop
			}
		}
}

//接口，多态示例：
type Filter interface {
	About() string
	Process([]int) []int
}

type UniqueFilter struct{
}
func (UniqueFilter) About() string{
	return "删除重复的数字"
}

func (UniqueFilter) Process(inputs []int)[]int{
	outs := make([]int, 0, len(inputs))
	pusheds := make(map[int]bool)
	for _,n := range inputs{
		pushed, ok := pusheds[n]
		if !pushed || !ok {
			pusheds[n] = true
			outs = append(outs, n)
		}
	}
	return outs
}

type MultiFilter int

func (mf MultiFilter)About() string{
	return fmt.Sprintf("保留%v的倍数", mf)
}

func (mf MultiFilter)Process(inputs []int)[]int{
	var outs = make([]int, 0, len(inputs))
	for _, n := range inputs{
		if n % int(mf) == 0{
			outs = append(outs, n)
		}
	}
	return outs
}

func filterAndPrint(fltr Filter, unfiltered []int) []int {
	filtered := fltr.Process(unfiltered)
	fmt.Println(fltr.About() + ":\n\t", filtered)
	return filtered
}


