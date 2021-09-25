//WAP to check if a number is prime or not
package main

import "fmt"

func main() {
	var number int
	fmt.Println("Enter no to check")
	fmt.Scanln(&number)
	checkPrime(number)
}

func checkPrime(number int) {
	isPrime := true
	if number == 0 || number == 1 {
		fmt.Printf("%d no is no prime", number)
	} else {
		for i := 2; i <= number/2; i++ {
			if number%1 == 0 {
				fmt.Printf("%dno is no prime", number)
				isPrime = false
				break
			}
		}
		if isPrime == true {
			fmt.Printf("%d no is prime", number)
		}
	}

}
