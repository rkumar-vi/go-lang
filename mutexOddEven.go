package main

import (
	"fmt"
	"sync"
)

var x int = 0

func increment(wg *sync.WaitGroup, mutex *sync.Mutex) {
	mutex.Lock()
	x = x + 1
	if x%2 == 0 {
		fmt.Println("even ", x)
	} else {

		fmt.Print("odd ", x)
	}
	mutex.Unlock()
	wg.Done()
}

func main() {

	var waitGroup sync.WaitGroup
	var mutex sync.Mutex
	for i := 0; i < 1000; i++ {
		waitGroup.Add(1)
		go increment(&waitGroup, &mutex)
	}
	waitGroup.Wait()

	fmt.Println("sum of 1 thousand times = ", x)

}

// output with line 11 and 19 with comment
/* odd 993even  994
odd 995even  996
odd 997odd 981odd 613even  998
odd 619odd 437even  438
even  544
odd 639odd 815odd 911odd 459odd 869sum of 1 thousand times =  999
rupeshkumar@ENCOPDBANLT0254 go-lang %
*/
//output with line 11 and 19 without comment
/* odd 987even  988
odd 989even  990
odd 991even  992
odd 993even  994
odd 995even  996
odd 997even  998
odd 999even  1000
sum of 1 thousand times =  1000
rupeshkumar@ENCOPDBANLT0254 go-lang %  */
