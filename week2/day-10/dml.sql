//nomor1
SELECT * FROM transcation WHERE id_pelanggan=1 UNION SELECT * FROM transcation WHERE id_pelanggan=2

//nomor2
SELECT SUM(product.price) as total_price from (transcation JOIN product ON transcation.id_product=product.id) WHERE transcation.id_pelanggan=1

//nomor3
SELECT COUNT(product.price) as total_price from (transcation JOIN product ON transcation.id_product=product.id) WHERE product.product_type='adidas'

//nomor4
SELECT * from transcation INNER JOIN product ON transcation.id_product=product.id

//nomor5
SELECT * from product INNER JOIN transcation INNER JOIN pelanggan

//nomor6
DELIMITER %%
CREATE TRIGGER hapus_semua_data_transaksi
AFTER DELETE ON transaction FOR EACH ROW
BEGIN
DECLARE d_user_id INT;
DELETE FROM transaction_detail WHERE order_id=d_user_id;
END %%

//nomor7
DELIMITER &&
CREATE TRIGGER hapus_data_detail_transaksi
BEFORE DELETE ON transaction_detail FOR EACH ROW
BEGIN
DECLARE d_user_id INT;
DECLARE d_qty INT;
SET d_user_id=transaction_detail.order_id;
set d_qty =transaction_detail.product_id;
UPDATE transaction_detail SET
total_qty=totalqty-d_qty WHERE id=d_user_id;
END &&

//nomor8
SELECT * FROM product WHERE product.id_product NOT IN(SELECT transaction_detail.product_id FROM transaction_detail)