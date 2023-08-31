package models

import (
	"golang-web/db"
)

type Produto struct {
	Id              int
	Nome, Descricao string
	Preco           float64
	Quantidade      int
}

func CadastraProduto() {
	//
}

func GetAllProdutos() []Produto {
	db := db.ConectaComBancoDeDados()
	resultado, err := db.Query("select * from produtos")

	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for resultado.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = resultado.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)

	}

	defer db.Close()
	return produtos
}
