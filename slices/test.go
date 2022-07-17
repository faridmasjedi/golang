package main
import "fmt"

func main(){ 
	var ss [3]string
	s := make([]string , 3)

	s[0] = "s1"
	s[1] = "s2"
	s[2] = "s3"
	ss[0] = "ss0"
	ss[1] = "ss1"
	ss[2] = "ss2"
	fmt.Println("ss: ", ss)
	fmt.Println("s: ", s)

	s = append(s, "s4")
	fmt.Println("s: ",s)
	// ss has been defined as [3]string, it can not use append as it is not slice
	// ss = append(ss, "ss3")
	// ss = append(ss, "ss4", "ss5")
	// fmt.Println(ss)
	s = append(s, "s5", "s6")
	fmt.Println("s: ", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("c: ", c)

	twoD := make([][]int, 3)
	for i := 0; i<len(twoD); i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j:=0; j<innerLen; j++ {
			twoD[i][j] = i+j
		}
	}
	fmt.Println("2d: ", twoD)


	l := make([][]int, 5)
	for i := 5; i > 0; i-- {
		var innerLen int = i
		l[5-i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			l[5-i][j] = i - j
		}
	} 
	fmt.Println("l: ", l)
}