## (8) Data Structure

1. Array
- array adalah suatu struktur data yang berisi elemen-elemen, bisa berisi satu type dari variabel dengan ukuran yang sudah ditetapkan. array bisa berupa tipe data numeric, string, dan boolean
- deklarasi array 
var arr [100]tipedatanya
arr := [100]tipedatanya{..,..,..}
- akses array
arr[0] = "array"

2. Slice
- Slice adalah struktur data yang berisi tentang elemen-elemen sama seperti array tetapi Slice memiliki ukuran yang dapat berubah dinamis jika penuh maka ukuran akan bertambah 2 kali lipat.
- deklarasi slice
var slice []tipeDatanya
slice bisa dideklarasikan tanpa ukuran
-ada fungsi-fungsi yang bisa dilakukan dengan slice salah satunya make :
make([]tipedata, len, cap) len adalah panjang isi dan cap adalah kapasitas slice

3. Map
- Map adalah struktur data yang disimpan dalam form yang memiliki key and value yang mana setiap key adalah unik tidak sama
-deklarasi map
var nama= map[tipedata]tipedata{}
var nama= map[tipedata]tipedata{"nama1":1992,}
-akses map menggunakan key
nama["nama2"] = 2002

4. Function
- Function adalah bagian dari code yang dapat dipanggil dengan nama, manfaat Function adalah cara mudah untuk membagi kode Anda menjadi blok-blok yang bisa dipanggil ketika diperlukan, sehingga code menjadi rapi dan modular.
-  Deklarasi function
func <name_function> (<parameter>) <type_return> { <statements> }

