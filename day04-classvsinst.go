#Author : iarvi

package main
import "fmt"

type person struct {
	age int
}

func (p person) NewPerson(initialAge int) person {
	//Add some more code to run some checks on initialAge
    //fmt.Println(p.age,"received NP")
    if initialAge <= 0{
        p.age = 0
        fmt.Println("Age is not valid, setting age to 0.")
    } else {
        p.age = initialAge
    }

	return p
}

func (p person) amIOld() {
	//Do some computation in here and print out the correct statement to the console
    //fmt.Println(p.age,"received amiold")
    if p.age < 13 {
        fmt.Println(p.age,"You are young")
    } else if p.age >= 13 && p.age < 18{
        fmt.Println(p.age,"You are teenager")
    }else if p.age >= 18{
        fmt.Println(p.age,"You are old")
    }
}

func (p person) yearPasses() person {
	//Increment the age of the person in here
    //fmt.Println(p.age,"received yearpasses")
    if p.age >= 18{
    p.age = p.age + 1
    
    }
	return p
}

func main() {
