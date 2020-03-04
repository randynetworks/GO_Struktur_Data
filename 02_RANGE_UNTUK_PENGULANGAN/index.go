package main

import "fmt"

// array

func main() {

	days := [7]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}

	// pengulangan for
	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// dengan range
	for index, day := range days {
		// Jika ditampilkan keduanya (index dan day)
		fmt.Printf("hari ke-%d = %s\n", index+1, day)
		// jika salahsatu yang ditampilkan, maka pendeklarasian variable yang tidak digunakan, diganti dengan ( _ )
	}

}
