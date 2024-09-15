package src

import (
	"log"
	"os"
)

const fileName = "cotacao.txt"

func CreateFile() {
	f, e := os.Create(fileName)
	if e != nil {
		panic(e)
	}
	defer f.Close()

	txt := "Lorem ipsum dolor sit amet"
	size, e := f.WriteString(txt)
	if e != nil {
		panic(e)
	}
	log.Printf("Arquivo %s criado com sucesso. Tamanho %d bytes.\n", fileName, size)
}
