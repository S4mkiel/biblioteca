package Biblioteca

import "fmt"

type biblioteca struct{
	Nome 	string
	Assunto string
	Autor 	string
	ISBN 	string
}


func Biblioteca(b biblioteca) string{
	fmt.Printf("Informe os dados do livro:\nNome:%s\nAssunto:%s\nAutor:%s\nISBN:%s", b.Nome, b.Assunto, b.Autor, b.ISBN)
	//for count := 1; count < 5; count++{
		//if (count < 5){
		//	fmt.Sprint("Informe os dados do livro:\nNome:%s\nAssunto:%s\nAutor:%s\nISBN:%s", b.Nome, b.Assunto, b.Autor, b.ISBN)
		//}else{
	//		fmt.Print("O número maximo de adição já foi batido")
	//	}
//	}
	return b.Nome

}