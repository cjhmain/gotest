package ark

import (
	"fmt"
)

type Stringer interface{
	Move()
	Eat()
}

type Animal struct{
	name string
}

func (an Animal)Laugh(){
	fmt.Println(an.name, "is laughing")
}

func Test(){
	fmt.Println("-------------------- submodule ark --------------------")

	var s1, s2 Stringer
	s1 = Dog{
		tail: 10,
		Animal: Animal{
			name:"dog HEI",
		},
	}
	s2 = Horse{
		tail: 100,
		Animal :Animal{
			name: "horse WU",
		},
	}
	
	// 多态
	s1.Move()
	s2.Eat()

	s3 := Dog{
		tail: 10,
		Animal: Animal{
			name:"dog MINI HEI",
		},
	}
	// 继承
	s3.Laugh()
}
