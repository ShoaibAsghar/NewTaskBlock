import "fmt"

type Person struct {
	Name   string
	age    int
	job    string
	salary int
}

//function that takes struct type as argument

func argu_StructFunc(obj Person) {
	fmt.Println("Name:   ", obj.Name)
	fmt.Println("Age:    ", obj.age)
	fmt.Println("Job:    ", obj.job)
	fmt.Println("Salary: ", obj.salary)

}
