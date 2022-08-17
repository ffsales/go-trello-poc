===================================================================================================================================================
Comando para subir uma instância do MySQL

docker run -d -p 3306:3306 -e MYSQL_ROOT_PASSWORD=12345 -e MYSQL_DATABASE=go-db -e MYSQL_USER=root-go -e MYSQL_PASSWORD=12345-go mysql/mysql-server

===================================================================================================================================================
Criação do banco de dados

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

===================================================================================================================================================
Importar o arquivo "collection.json" no Insomnia para utilizar as Apis do projeto