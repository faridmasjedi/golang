package main
import "fmt"

func main(){
	i := 1
	for i<=3 {
		fmt.Println(i)
		i++
	}

	fmt.Println("\n---\n")

	for j:=2; j<6; j++ {
		fmt.Println(j)
	}

	fmt.Println("\n---\n")

	for {
		fmt.Println("loop??!!!")
		break
	}

	fmt.Println("\n---\n")

	for n:=1; n <= 10; n++ {
		if n % 2 == 0 {
			fmt.Println(n)
		}
	}
}