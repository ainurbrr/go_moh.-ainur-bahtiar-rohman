-- 1. Gabungkan data transaksi dari user id 1 dan user id 2.

SELECT * FROM transactions t INNER JOIN users u ON t.user_id = u.id WHERE t.user_id = 1 OR t.user_id = 2;

-- 2. Tampilkan jumlah harga transaksi user id 1.

SELECT SUM(total_price) FROM transactions WHERE user_id = 1;

-- 3. Tampilkan total transaksi dengan product type 2.

SELECT SUM(total_price) as total_price FROM transactions WHERE user_id = 2 group by user_id;

-- 4. Tampilkan semua field table product dan field name table products_type yang saling berhubungan.

SELECT products.*, product_types.name FROM products 
JOIN product_types ON products.product_type_id = product_types.id;

-- 5. Tampilkan semua field table transaction, field name table product dan field name table user.

SELECT t.*, p.name, u.name as users FROM transactions_details td 
JOIN transactions t ON t.id = td.transaction_id
JOIN users u ON u.id = t.user_id
JOIN products p ON p.id = td.product_id

 -- 6. Buat function setelah data transaksi dihapus maka transaction detail terhapus juga dengan transaction id yang dimaksud.

delimiter $$
CREATE trigger delete_transactions_details
BEFORE DELETE ON transactions for each row
BEGIN
    DECLARE trans_id INT;
    set trans_id = old.id;
    DELETE FROM transactions_details WHERE transaction_id = trans_id
END $$;

DELETE FROM transactions WHERE id = 4; 

SELECT * FROM transactions_details;

-- 7. Buat function setelah data transaksi detail dihapus maka data total_qty terupdate berdasarkan qty data transaction id yang dihapus.

delimiter $$
CREATE trigger update_transaksi_id 
AFTER DELETE ON transactions_details for each row
BEGIN
    DECLARE trans_id INT;
    set trans_id = old.transaction_id;
    UPDATE transactions set total_qty = (SELECT  SUM (qty) FROM transactions_details WHERE transaction_id = trans_id) 
    WHERE id = trans_id;
END $$ 

SELECT * FROM transactions_details;

SELECT * FROM transactions_details WHERE transaction_id = 2;

DELETE FROM transactions_details WHERE transaction_id = 2 AND product_id = 9;

-- 8. Tampilkan data products yang tidak pernah ada di tabel transaction_details dengan sub-query.

SELECT * FROM products WHERE id NOT IN (SELECT product_id FROM transactions_details);

