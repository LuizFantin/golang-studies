package panicDeferRecover

import (
	"fmt"
	"os"
)

func Test() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Teste recover")
		}
	}()

	file, err := os.Create("./index.txt")
	//? Garante que vai dar o close como ultima ação da função. Roda depois do return.
	defer file.Close()

	if err != nil {
		panic(err)
	}

	_, err = file.WriteString("Teste")

	if err != nil {
		panic(err)
	}
}
