//nomor1
SELECT id FROM transcation WHERE id_pelanggan=1 or id_pelanggan=2

//nomor2
SELECT SUM(product.price) as total_price from (transcation JOIN product ON transcation.id_product=product.id) WHERE transcation.id_pelanggan=1

//nomor3
SELECT SUM(price) from product where product_type='adidas'

//nomor4
SELECT * from product INNER JOIN transcation

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


//nomor8



