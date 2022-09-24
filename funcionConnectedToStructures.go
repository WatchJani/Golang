//VIDEO 27

package main

import "fmt"

type Person struct {
	ID string
}

func (p *Person) SectorName() string {
	number := p.ID[:1]

	return fmt.Sprintf("first letter: %v", number)

}

//==================================================================================================================

func SetupConnections(m []string, f []int32) (map[string]int32, string) {
	if len(m) != len(f) {
		return nil, "Slice nema isti broj elemenata :D"
	}

	my_map := map[string]int32{}

	for index, value := range m {
		my_map[value] = f[index]
	}

	return my_map, "Your list is created"

}

func main() {
	l1 := []string{"Pero", "Petar"}
	l2 := []int32{321, 321}

	my_new_list, message := SetupConnections(l1, l2)

	fmt.Printf("%v, %s\n", my_new_list, message)

	p := &Person{"sda"}
	p2 := Person{"+654"}

	fmt.Println(p.SectorName())
	fmt.Println(p2.SectorName())
}
