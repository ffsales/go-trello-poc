<h1 align="center">Project Go Trello Poc</h1>

<div align="center">

[![Status](https://img.shields.io/badge/status-active-success.svg)]()
[![GitHub Issues](https://img.shields.io/github/issues/ffsales/go-trello-poc.svg)](https://github.com/ffsales/go-trello-poc/issues)
[![GitHub Pull Requests](https://img.shields.io/github/issues-pr/ffsales/go-trello-poc.svg)](https://github.com/ffsales/go-trello-poc/pulls)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](/LICENSE)

</div>

---


## 📝 Conteúdo

- [Sobre](#about)
- [Iniciando](#getting_started)
- [Deploy](#deployment)
- [Uso](#usage)
- [Tecnologias Usadas](#built_using)
- [TODO](./TODO.md)
- [Autor](#authors)
- [Reconhecimento](#acknowledgement)

## 🧐 Sobre <a name = "about"></a>

O projeto tem como objetivo fazer uma demonstração de uma api rest na linguagem GO simulando a aplicação Trello.

## 🏁 Iniciando <a name = "getting_started"></a>

Abaixo estão as instruções:

Importar o arquivo "collection.json" no Insomnia para utilizar as Apis do projeto

### Pré-requisitos

```
Esta aplicação foi desenvolvida usando:
- [Go] 1.18
```

### Instalando

- Primeiro execute o comando docker abaixo
```
docker run -d -p 3306:3306 -e MYSQL_ROOT_PASSWORD=12345 -e MYSQL_DATABASE=go-db -e MYSQL_USER=root-go -e MYSQL_PASSWORD=12345-go mysql/mysql-server

Este comando irá inciar uma instância do MySQL na porta 3306, com a senha root 12345, uma database com o nome go-db, usuário root igual a root-go e senha root igual a 12345-go
```
- Criando tabelas

```
CREATE DATABASE trello-go-db

CREATE TABLE board (
	id int not null auto_increment, 
	name varchar(255), 
	description varchar(255), 
	primary key (id)
);

CREATE TABLE list (
	id int not null auto_increment, 
	name varchar(255),
	pos int,
	id_board int, 
	primary key (id),
	FOREIGN KEY(id_board) REFERENCES board(id)
);

CREATE TABLE card (
	id int not null auto_increment, 
	name varchar(255), 
	finished bool,
    id_list int, 
	primary key (id),
    FOREIGN KEY(id_list) REFERENCES list(id)
);
```
- Executando programa

```
go run *.go
```


## 🔧 Rodando testes <a name = "tests"></a>

Em construção.


## 🎈 Uso <a name="usage"></a>

Importe o arquivo collection.json no aplicativo Insomnia e execute as requisições

## 🚀 Deploy <a name = "deployment"></a>

Em construção.

## ⛏️ Built Using <a name = "built_using"></a>

- [Go](https://go.dev/) - Language
- [Chi](https://go-chi.io/#/) - Server Framework
- [MySQL](https://www.mysql.com/) - Database
- [Docker](https://www.docker.com/) - Container

## ✍️ Autor <a name = "authors"></a>

- [@ffsales](https://github.com/ffsales) - Planejamento e execução


## 🎉 Reconhecimento <a name = "acknowledgement"></a>
Em construção

