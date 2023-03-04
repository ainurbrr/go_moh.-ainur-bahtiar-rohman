## (9) String - Advance Function - Pointer - Method - Struct and Interface

1. String
Go menyediakan package strings, isinya banyak fungsi untuk keperluan pengolahan data string. antara lain : len(string) yaitu menhitung panjang string, compare string untuk membandingkan dua buah string, strings.Contains untuk melihat apakah substring bagian dari string, mengambil substring, strings.Replace untuk mengganti bagian dari string, dan Insert pada String

2. More Function
-Variadic Function
    Konsep Variadic Function atau pembuatan fungsi dengan parameter sejenis yang tak terbatas. Maksud tak terbatas di sini adalah jumlah parameter yang disisipkan ketika pemanggilan fungsi bisa berapa saja. Parameter pada variadic function memiliki sifat yang mirip dengan slice. Nilai dari parameter-parameter yang disisipkan bertipe data sama, dan ditampung oleh sebuah variabel saja. Cara pengaksesan tiap datanya juga sama, dengan menggunakan index
-Anonymous Function
    Anonymous function adalah fungsi yang tidak mengandung nama apa pun. Ini berguna saat Anda ingin membuat fungsi sebaris.
-Closure Function
    adalah sebuah fungsi yang bisa disimpan dalam variabel. Dengan menerapkan konsep tersebut, kita bisa membuat fungsi di dalam fungsi, atau bahkan membuat fungsi yang mengembalikan fungsi.
-Defer Function
    Defer function adalah fungsi yang ditangguhkan atau dijalankan diakhir setelah fungsi lainnya selesai. defer function bisa digunakan lebih dari sekali dan mereka dijalankan sebagai tumpukan atau stack, atau last in first out

3. Pointer
Pointer adalah reference atau alamat memori. Variabel pointer berarti variabel yang berisi alamat memori suatu nilai.
cara mendeklarasikan pointer yaitu menggunakan tanda * misalnya var *pointer dan untuk menambahkan isinya  var *pointer = &namaVariable

4. Struct 
Struct adalah kumpulan definisi variabel (atau property) dan atau fungsi (atau method), yang dibungkus sebagai tipe data baru dengan nama tertentu. Property dalam struct, tipe datanya bisa bervariasi. Mirip seperti map, hanya saja key-nya sudah didefinisikan di awal, dan tipe data tiap itemnya bisa berbeda.
Dari sebuah struct, kita bisa buat variabel baru, yang memiliki atribut sesuai skema struct tersebut. Kita sepakati dalam hal ini, variabel tersebut dipanggil dengan istilah object atau object struct.
Konsep struct di golang mirip dengan konsep class pada OOP

5. Method
Method adalah fungsi yang menempel pada type (bisa struct atau tipe data lainnya). Method bisa diakses lewat variabel objek.

6. Interface 
Interface adalah kumpulan definisi method yang tidak memiliki isi (hanya definisi saja), yang dibungkus dengan nama tertentu.
Interface merupakan tipe data. Nilai objek bertipe interface zero value-nya adalah nil. Interface mulai bisa digunakan jika sudah ada isinya, yaitu objek konkret yang memiliki definisi method minimal sama dengan yang ada di interface-nya.

7. Package
Paket adalah kumpulan fungsi dan data. yang bisa dipanggil tetapi harus di import terlebih dahulu

8. Error Handling (Panic & Recover)
Panic
    Saat runtime Go mendeteksi error ini, ia akan mengembalikan nilai panic.
Recover
    Untuk menambahkan kemampuan untuk memulihkan dari error panic, tambahkan fungsi anonim atau tentukan fungsi khusus dan panggil dengan kata kunci 'defer' dari dalam method, di mana panic mungkin terjadi dari panggilan internal lainnya