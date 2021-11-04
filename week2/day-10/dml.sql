SELECT id FROM transcation WHERE id_pelanggan=1 or id_pelanggan=2

SELECT SUM(product.price) as total_price from (transcation JOIN product ON transcation.id_product=product.id) WHERE transcation.id_pelanggan=1

SELECT SUM(price) from product where product_type='adidas'







