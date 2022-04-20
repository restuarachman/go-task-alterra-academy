INSERT INTO operators(id, name) VALUES 
(1, "TELKOMSEL"),
(2, "XL"),
(3, "IM3"),
(4, "By-U"),
(5, "MENTARI");

INSERT INTO product_types(id, name) VALUES
(1, "Pulsa"),
(2, "Paket Data"),
(3, "Voucher Game");

INSERT INTO products(id, product_type_id, operator_id, code, name, status) VALUES
(1, 1, 3, "113", "Pulsa IM3 Darurat", 1),
(2, 1, 3, "213", "Pulsa IM3 Eksklusif", 1),
(3, 2, 1, "221", "Paket Data Telkomsel Harian", 1),
(4, 2, 1, "421", "Paket Data Telkomsel Mingguan", 1),
(5, 2, 1, "521", "Paket Data Telkomsel Bulanan", 1),
(6, 3, 4, "634", "Voucher Game By-U PUBG", 1),
(7, 3, 4, "734", "Voucher Game By-U Valorant", 1),
(8, 3, 4, "834", "Voucher Game By-U Mobile Legend", 1);

INSERT INTO product_descriptions(id, description) VALUES
(1, "description 1"),
(2, "description 2"),
(3, "description 3"),
(4, "description 4"),
(5, "description 5"),
(6, "description 6"),
(7, "description 7"),
(8, "description 8");

INSERT INTO  payment_methods(id, name, status) VALUES
(1, "E-Wallet", 1),
(2, "ATM", 1),
(3, "Tunai", 1);

INSERT INTO users(id, status, dob, gender, name) VALUES
(1, 1, '2001-06-01', 'M', "Ramadito"),
(2, 1, '2002-06-01', 'M', "Ferdian"),
(3, 1, '2001-02-01', 'M', "Assa"),
(4, 1, '2002-01-01', 'F', "Haviza"),
(5, 1, '2002-10-01', 'F', "Aufa");

INSERT INTO transactions(id, user_id, payment_method_id, total_price, total_qty, status) VALUES
(1, 1, 1, 20000, 2, 1),
(2, 1, 1, 20000, 2, 1),
(3, 1, 1, 20000, 2, 1),

(4, 2, 2, 20000, 2, 1),
(5, 2, 2, 20000, 2, 1),
(6, 2, 2, 20000, 2, 1),

(7, 3, 3, 20000, 2, 1),
(8, 3, 3, 20000, 2, 1),
(9, 3, 3, 20000, 2, 1),

(10, 4, 1, 20000, 2, 1),
(11, 4, 1, 20000, 2, 1),
(12, 4, 1, 20000, 2, 1),

(13, 5, 1, 20000, 2, 1),
(14, 5, 1, 20000, 2, 1),
(15, 5, 1, 20000, 2, 1);

INSERT INTO `alta_online_shop`.`transaction_details` (`transaction_id`, `product_id`, `status`, `qty`, `price`) VALUES 
('1', '1', '1', '2', '2000'),
('1', '2', '1', '2', '2000'),
('1', '3', '1', '2', '2000'),
('2', '4', '1', '2', '2000'),
('2', '5', '1', '2', '2000'),
('2', '6', '1', '2', '2000'),
('3', '7', '1', '2', '2000'),
('3', '8', '1', '2', '2000'),
('4', '6', '1', '2', '2000'),
('4', '5', '1', '2', '2000'),
('4', '4', '1', '2', '2000'),
('5', '6', '1', '2', '2000'),
('5', '4', '1', '2', '2000'),
('5', '7', '1', '2', '2000'),
('6', '3', '1', '2', '2000'),
('6', '2', '1', '2', '2000'),
('6', '7', '1', '2', '2000'),
('7', '4', '1', '2', '2000'),
('7', '8', '1', '2', '2000'),
('7', '2', '1', '2', '2000'),
('8', '1', '1', '2', '2000'),
('8', '4', '1', '2', '2000'),
('8', '7', '1', '2', '2000'),
('9', '3', '1', '2', '2000'),
('9', '8', '1', '2', '2000'),
('9', '4', '1', '2', '2000'),
('10', '3', '1', '2', '2000'),
('10', '5', '1', '2', '2000'),
('10', '8', '1', '2', '2000'),
('11', '8', '1', '2', '2000'),
('11', '4', '1', '2', '2000'),
('11', '2', '1', '2', '2000'),
('12', '2', '1', '2', '2000'),
('12', '4', '1', '2', '2000'),
('12', '5', '1', '2', '2000'),
('13', '7', '1', '2', '2000'),
('13', '5', '1', '2', '2000'),
('13', '8', '1', '2', '2000'),
('14', '3', '1', '2', '2000'),
('14', '7', '1', '2', '2000'),
('14', '2', '1', '2', '2000'),
('15', '4', '1', '2', '2000'),
('15', '6', '1', '2', '2000'),
('15', '7', '1', '2', '2000');

SELECT name FROM users WHERE gender='M';

SELECT * FROM product WHERE id=3;

SELECT * FROM users WHERE 
date(created_at) >= CURDATE() - interval 7 day AND 
name LIKE '%a%';

SELECT COUNT(gender) FROM users WHERE 
gender = 'F';

SELECT * FROM users 
ORDER BY name

SELECT * FROM products LIMIT 5;

UPDATE products
SET name = 'product dummy'
WHERE id = 1;

UPDATE transaction_details
SET qty = 3
WHERE product_id = 1;

DELETE FROM products where id=1;

SELECT * FROM transactions WHERE user_id = 1 or user_id = 2;

SELECT SUM(total_price) FROM transactions WHERE user_id = 1;

SELECT COUNT(user_id) FROM transactions WHERE user_id = 2;

SELECT p.*, pt.name FROM products as p
INNER JOIN product_types as pt
ON product_type_id=product_types.id;

SELECT t.*, p.name, u.name FROM products as p
INNER JOIN  transaction_details as td
ON td.product_id=p.id
INNER JOIN transactions as t 
ON t.id = td.transaction_id
INNER JOIN users as u
ON u.id = t.user_id;

DELIMITER $$
CREATE TRIGGER delete_transactions
BEFORE DELETE ON transactions FOR EACH ROW
BEGIN
DECLARE v_transaction_id INT;
SET v_transaction_id=OLD.id;
DELETE FROM transactoin_details WHERE transaction_id = v_transaction_id;
END$$
DELIMITER ;

DELIMITER $$
CREATE TRIGGER delete_transaction_details
BEFORE DELETE ON transaction_details FOR EACH ROW
BEGIN
DECLARE v_total_qty INT;
DECLARE v_id INT;
SET v_id = OLD.transaction_id;
SELECT SUM(qty) INTO v_total_qty FROM transactoin_details WHERE transaction_id = v_id;
UPDATE transactions
SET total_qty = v_total_qty
WHERE transactions.id = v_id;
END$$
DELIMITER ;

SELECT * FROM products WHERE id NOT IN
(SELECT product_id FROM transaction_details);