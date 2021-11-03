create database alta_online_shop;

use alta_online_shop

create table product(product varchar(44) not null primary key,product_type varchar(44),product_description varchar(44),operator varchar(44),payment_methods varchar(44));

create table transcation(transaction int not null auto_increment primary key,transaction_detail varchar(44) not null);

create table kurir(id int not null auto_increment primary key, name varchar(44), created_at date, updated_at date);

alter table kurir add ongkos_dasar int(24);

rename table kurir to shipping;

create table userdetail (id int primary key,
nik int,
email varchar(50),
usia int,
foreign key (id) references users(user_id));

alter table transaction
add column user_id int;

alter table transaction add constraint foreign key (user_id) references users(user_id);

create table transaction_products (
transaction_id int,
product_id int,
foreign key (transaction_id) references transaction(transaction_id) on delete restrict on update cascade,
foreign key (product_id) references itemproduct(product_id) on delete restrict on update cascade,
primary key (transaction_id, product_id));