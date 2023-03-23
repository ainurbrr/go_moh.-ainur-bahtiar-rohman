-- QUERY INSERT

-- 1. Insert 5 operators pada table operators.

INSERT INTO operators(name) VALUES ('Toko Muji Jaya'), 
('Toko Sinar Abadi'), 
('Toko Sumber Rejeki'), 
('Toko Berkah Jaya'), 
('Toko OKE');

-- 2. Insert 3 product type.

INSERT INTO product_types(name) VALUES
('Makanan'),
('Sabun/Deterjen'),
('Elektronik');

-- 3. Insert 2 product dengan product type id = 1, dan operators id = 3.

INSERT INTO products(product_type_id, operator_id, code, name, status) VALUES
(1, 3, 'MNM001', 'RedBull', 0),
(1, 3, 'MNM004', 'Kratingdeng', 1);

-- 4. Insert 3 product dengan product type id = 2, dan operators id = 1.

INSERT INTO products(product_type_id, operator_id, code, name, status) VALUES
(2, 1, 'ELK001', 'Kulkas', 1),
(2, 1, 'ELK003', 'Televisi', 1),
(2, 1, 'ELK005', 'Lampu Taman', 1);

-- 5. Insert 3 product dengan product type id = 3, dan operators id = 4.

INSERT INTO products(product_type_id, operator_id, code, name, status) VALUES
(3, 4, 'MKN004', 'Indomie Goreng', 1),
(3, 4, 'MKN005', 'Chocolate SilverQueen', 1),
(3, 4, 'MKN006', 'Biskuat', 1);

-- 6. Insert product description pada setiap product.

INSERT INTO product_descriptions(product_id, description) VALUES
(2, 'Minuman berenergi'),
(1, 'Minuman Gambar banteng'),
(5, 'lampu taman untuk menerangi taman'),
(4, 'televisi untuk menonton'),
(3, 'Kulkas untuk mendinginkan sesuatu'),
(8, 'Biskuat untuk cemilan'),
(6, 'Indomie Goreng untuk makan'),
(7, 'Chocolate Lezat');

-- 7. Insert 3 payment methods.

INSERT INTO payment_methods(name, status) VALUES
('Transfer Bank', 1),
('Gopay', 1),
('Dana', 1);

-- 8. Insert 5 user pada tabel user.

INSERT INTO users(name, status, dob, gender) VALUES
('MOMO', 1, '1998-04-19', 'L'),
('Jamal', 1, '1992-02-02', 'L'),
('Lika', 1, '1999-05-03', 'P'),
('Sri', 1, '1997-04-04', 'P'),
('Dani', 1, '2001-04-05', 'L');

-- 9. Insert 3 transaksi di masing-masing user. (soal berlanjut ke soal 1.j)

INSERT INTO transactions(user_id, payment_method_id, status, total_qty, total_price) VALUES
(1, 1, 'lunas', 2, 20000),
(2, 2, 'belum', 3, 30000),
(3, 3, 'lunas', 4, 40000);

-- 10. Insert 3 product di masing-masing transaksi.

UPDATE transactions
SET total_qty = 5, total_price= 50000
WHERE user_id = 1;

UPDATE transactions
SET total_qty = 6, total_price= 60000
WHERE user_id = 2;

UPDATE transactions
SET total_qty = 7, total_price= 70000
WHERE user_id = 3;

-- QUERY SELECT

-- a. Tampilkan nama user / pelanggan dengan gender Laki-laki / M.

SELECT name FROM users WHERE gender = 'L' or gender = 'M';

-- b. Tampilkan product dengan id = 3.

SELECT * FROM products WHERE id = 3;

-- c. Tampilkan data pelanggan yang created_at dalam range 7 hari kebelakang dan mempunyai nama mengandung kata ‘a’.

SELECT * FROM users WHERE created_at BETWEEN DATE_SUB(NOW(), INTERVAL 7 DAY) AND NOW() AND name LIKE '%a%';

-- d. Hitung jumlah user / pelanggan dengan status gender Perempuan.

SELECT COUNT(*) FROM users WHERE gender = 'P';

-- e. Tampilkan data pelanggan dengan urutan sesuai nama abjad

SELECT * FROM users ORDER BY name ASC;

-- f. Tampilkan 5 data pada data product

SELECT * FROM products LIMIT 5;

-- QUERY UPDATE

-- a. Ubah data product id 1 dengan nama ‘product dummy’.

UPDATE products SET name = 'product dummy' WHERE id = 1;

-- b. Update qty = 3 pada transaction detail dengan product id = 1.

UPDATE transactions_details SET qty = 3 WHERE product_id = 1;

-- QUERY DELETE

-- a. Delete data pada tabel product dengan id = 1.

DELETE FROM products WHERE id = 1;

-- b. Delete pada pada tabel product dengan product type id 1.

DELETE FROM products WHERE product_type_id = 1;


