## (19) intro RESTful API

1. Apa itu API? API adalah kumpulan fungsi dan prosedur yang digunakan untuk pembuatan aplikasi yang dapat mengakses data dari sistem.
cara kerja API : ketika client membirkan request ke sistem atau server kemudian server akan memberikan respon kepada client.
API (Application Programming Interface) adalah kumpulan protokol, aturan, dan alat yang memungkinkan aplikasi perangkat lunak untuk berkomunikasi dengan satu sama lain. API sering digunakan dalam pengembangan perangkat lunak dan aplikasi web untuk memungkinkan interaksi antara berbagai sistem atau komponen.

2. REST (Representational State Transfer) API adalah jenis API yang menggunakan protokol HTTP untuk mengirim dan menerima data antara aplikasi web dan server. REST API biasanya digunakan dalam pengembangan web untuk memungkinkan komunikasi antara berbagai sistem yang berbeda.
REST API menggunakan konsep-konsep seperti resource (sumber daya), URI (Uniform Resource Identifier), method (metode), dan status code (kode status) untuk memungkinkan akses ke data dan fungsionalitas dari aplikasi web.
Dalam REST API, setiap resource direpresentasikan sebagai URI yang unik, dan setiap metode (GET, POST, PUT, DELETE, dll.) digunakan untuk melakukan operasi pada resource tersebut. Status code digunakan untuk memberi tahu klien tentang keberhasilan atau kegagalan permintaan.

3. JSON (JavaScript Object Notation) adalah format pertukaran data ringan yang digunakan dalam pengembangan web dan aplikasi. JSON merupakan sebuah format data yang berisi koleksi pasangan key-value yang mudah dibaca dan ditulis oleh manusia dan juga dapat dengan mudah diproses oleh komputer.
JSON biasanya digunakan sebagai format data untuk mentransmisikan data antara server dan aplikasi web. Dalam JSON, data diatur dalam objek atau array, dan setiap objek atau array terdiri dari pasangan key-value.

4. HTTP response code adalah kode status yang diberikan oleh server HTTP sebagai respon terhadap permintaan yang diberikan oleh klien HTTP (biasanya sebuah web browser atau aplikasi web). HTTP response code digunakan untuk memberikan informasi tentang keberhasilan atau kegagalan permintaan, dan seringkali digunakan dalam debugging atau pemecahan masalah masalah terkait jaringan atau aplikasi web.

Berikut adalah beberapa HTTP response code yang umum digunakan:

200 OK: Permintaan berhasil dilakukan dan server memberikan respons yang sesuai.
201 Created: Permintaan berhasil dilakukan dan server membuat objek baru sesuai permintaan.
204 No Content: Permintaan berhasil dilakukan tetapi tidak ada konten yang diberikan dalam respons.
400 Bad Request: Permintaan tidak dapat diproses karena kesalahan sintaks atau data yang tidak valid.
401 Unauthorized: Permintaan tidak dapat dilakukan karena klien belum terotentikasi atau tidak memiliki akses yang diperlukan.
403 Forbidden: Permintaan tidak dapat dilakukan karena klien tidak diizinkan mengakses sumber daya yang diminta.
404 Not Found: Sumber daya yang diminta tidak ditemukan di server.
500 Internal Server Error: Server mengalami kesalahan internal saat memproses permintaan.