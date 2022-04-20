CREATE DATABASE alta_online_shop;
USE alta_online_shop;
CREATE TABLE user (
	user_id int  NOT NULL,
    	username varchar(50) NOT NULL unique,
    	PRIMARY KEY (user_id)
);

CREATE TABLE product_type (
	type_id int NOT NULL,
    	type_name varchar(50) NOT NULL,
    	PRIMARY KEY (type_id)
);
CREATE TABLE operator(
	operator_id int NOT NULL,
    	operator_name varchar(50) NOT NULL,
    	PRIMARY KEY (operator_id)
);
CREATE TABLE product(
	product_id int NOT NULL,
    	operator_id int NOT NULL,
    	product_type int NOT NULL,
    	product_name varchar(50) NOT NULL,
    	product_descripion LONGTEXT NOT NULL,
    	PRIMARY KEY (product_id),
    	FOREIGN KEY (operator_id) REFERENCES operator(operator_id),
    	FOREIGN KEY (product_type) REFERENCES product_type(type_id)
);
CREATE TABLE payment_method(
	payment_id int NOT NULL,
    	payment_name varchar(50) NOT NULL,
    	PRIMARY KEY (payment_id)
);
CREATE TABLE transactions(
	transactions_id int NOT NULL,
    	user_id int NOT NULL,
    	payment_id int NOT NULL,
    	totalAmount int NOT NULL,
    	PRIMARY KEY (transactions_id),
    	FOREIGN KEY (user_id) REFERENCES user(user_id),
    	FOREIGN KEY (payment_id) REFERENCES payment_method(payment_id)
);
CREATE TABLE transaction_details(
	transaction_details_id int NOT NULL,
    	product_id int NOT NULL,
    	transactions_id int NOT NULL,
    	quantity int NOT NULL,
    	subtotal int NOT NULL,
    	PRIMARY KEY (transaction_details_id),
    	FOREIGN KEY (product_id) REFERENCES product(product_id),
    	FOREIGN KEY (transactions_id) REFERENCES transactions(transactions_id)
);
CREATE TABLE kurir(
	kurir_id int NOT NULL,
    	name varchar(50) NOT NULL,
    	create_at datetime NOT NULL,
    	update_at datetime NOT NULL
);
ALTER TABLE kurir
ADD ongkos_dasar int NOT NULL;

ALTER TABLE kurir 
RENAME shipping;

DROP TABLE shipping;

/*
Silahkan menambahkan entity baru dengan relation 1-to-1, 1-to-many, many-to-many.

Sudah dilakukan pada saat pembuatan table
*/