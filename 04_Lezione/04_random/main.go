//genera K numeri casuali in [0,10] e stampali

package main

import ("fmt"
		"math/rand"
		"time"
)

func main() {
	const K int = 10 

	rand.Seed(time.Now().UnixNano()) // secondi da the Epoch

	 

	for i := 1; i <= K; i++ {

		fmt.Println(rand.Intn(11))
	
	}
}