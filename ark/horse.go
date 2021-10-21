package ark

import (
	"fmt"
)

type Horse struct{
	tail 	int
	Animal
}

func (h Horse)Move(){
	msg := fmt.Sprintf("%s animal moving", h.name)
	fmt.Println(msg)
}

func (h Horse)Eat(){
	msg := fmt.Sprintf("%s animal eating", h.name)
	fmt.Println(msg)
}
