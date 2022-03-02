packe main 

import "fmt"

func main (){
	var angka int

	fmt.Println("masukan angka :")
	fmt.Scanln(&angka)

	fmt.Println("faktor dari angka ",angka, "yaitu:")
	for i :=1 ; i <= angka; i++{
		if angka & i ==0{
			fmt,Println(i)
		}
	}

}