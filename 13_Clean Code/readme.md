## (13) Clean Code

1. Clean code adalah istilah untuk kode yang mudah dibaca, difahami dan diubah oleh programmernya sendiri ataupun programmer lain.

2. Kenapa Clean code?
- Mempermudah kolaborasi antar programmer
- Mempermudah dalam mengerjakan feature development 
- Mempercepat development

3. Karakteristik Clean Code :
    1. Penamaan mudah difahami
    2. Mudah dieja dan dicari
    3. Singkat namun mendeskripsikan konteks
    4. Konsisten dalam penuisan
    5. Hindari penambahan yang tidak perlu
    6. Memberikan komentar sesuai kebutuhan
    7. Membuat fungsi yang mudah dipahami dan tidak terlalu panjang
    8. Gunakan konvensi style guide
    9. Menggunakan formatting atau code prettier, dekatkan code yang berhubungan secara berdekatan

4. Clean code principle 
    - Keep It So Simple (KISS):
        Fungsi atau class harus kecil
        Fungsi dibuat untuk melakukan satu tugas saja
        Jangan gunakan terlalu banyak argumen pada fungsi
        Harus diperhatikan untuk mencapai kondisi yang seimbang
        kecil dan jumlahnya minimal
    - Don't Repeat Yourself (DRY):
        Duplikasi code terjadi karena sering copy paste. Untuk menghindari duplikasi code buatlah fungsi yang dapat digunakan secara berulang-ulang.
    - Refactoring 
        Refactoring adalah proses restrukturisasi kode yang dibuat, dengan cara mengubah struktur internal tanpa mengubah perilaku eksternal. Prinsip KISS dan DRY bisa dicapai dengan cara refactor.
    - Teknik Refactoring:
        Membuat sebuah abstraksi.
        Memecah kode dengan fungsi/class.
        Perbaiki penamaan dan lokasi kode.
        Deteksi kode yang memiliki duplikasi.