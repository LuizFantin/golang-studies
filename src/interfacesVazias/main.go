package interfacesVazias

import (
	"fmt"
	"reflect"
)

func Test() {

	var genericList []interface{}

	genericList = append(genericList, 10)
	genericList = append(genericList, 7.5)
	genericList = append(genericList, true)
	genericList = append(genericList, "Teste")

	for _, valor := range genericList {
		fmt.Println("Valor: ", valor)
		fmt.Println("Valor: ", reflect.TypeOf(valor))
	}

}
