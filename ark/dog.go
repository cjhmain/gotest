package ark

import (
	"fmt"
)

type Dog struct {
	tail 	int
	Animal
}

func (d Dog)Move(){
	msg := fmt.Sprintf("%s animal moving", d.name)
	fmt.Println(msg)
}

func (d Dog)Eat(){
	msg := fmt.Sprintf("%s animal eating", d.name)
	fmt.Println(msg)
}

