package Biblioteca

import (
	"fmt"
	
)

type biblioteca struct{
	Nome 	string
	Assunto string
	Autor 	string
	ISBN 	string
}


func Bd(b biblioteca) string{
	fmt.Printf("Informe os dados do livro:\nNome:%s\nAssunto:%s\nAutor:%s\nISBN:%s", b.Nome, b.Assunto, b.Autor, b.ISBN)
	return b.Nome
}

func internal() bool{
	return true
}