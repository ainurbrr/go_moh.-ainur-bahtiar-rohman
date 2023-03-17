## (14) Database - DDL - DML

1. Database adalah kumpulan data yang terorganisir secara sistematis dan terstruktur, yang dirancang untuk memudahkan penyimpanan, pengambilan, dan pengolahan data. 

2. Database Relationship
    - One-to-one relationship adalah jenis relasi antara dua tabel dalam sebuah database di mana setiap baris atau entitas dalam satu tabel hanya terhubung dengan satu baris atau entitas di tabel lainnya. Artinya, setiap nilai pada kolom kunci di satu tabel hanya memiliki pasangan nilai unique di kolom kunci pada tabel lainnya.
    - One-to-Many relationship adalah jenis relasi antara dua tabel dalam sebuah database di mana setiap baris atau entitas dalam satu tabel dapat terhubung dengan satu atau lebih baris atau entitas di tabel lainnya. Artinya, setiap nilai pada kolom kunci di satu tabel dapat memiliki beberapa nilai yang sesuai di kolom kunci pada tabel lainnya. 
    - Many-to-many relationship adalah jenis relasi antara dua tabel dalam sebuah database di mana setiap baris atau entitas dalam satu tabel dapat terhubung dengan banyak baris atau entitas di tabel lainnya dan sebaliknya. Artinya, setiap nilai pada kolom kunci di satu tabel dapat memiliki banyak nilai yang sesuai di kolom kunci pada tabel lainnya, dan sebaliknya.

3. DDL (Data Definition Language) dan DML (Data Manipulation Language)
    *DDL STATEMENT
     - CREATE DATABASE database_name
     - USE database_name
     - CREATE TABLE ...
     - DROP TABLE ...
     - RENAME TABLE ...
     - ALTER TABLE ...
    *DML adalah perintah yang difunakan untuk memanipulasi data dalam tabel dari suatu database
     - mamasukkan data pada tabel INSERT INTO ...
     - menampilkan semua data pada tabel SELECT ...
     - mengubah data pada table UPDATE ...
     - mengahpus data pada tabel DELETE ...
    *DML Statement
     - LIKE / BETWEEN
     - AND / OR
     - ORDER BY
     - LIMIT