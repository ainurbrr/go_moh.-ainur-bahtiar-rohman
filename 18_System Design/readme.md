## (18) System Design

1. Diagram design meliputi pembuatan diagram yang digunakan untuk merancang atau menganalisis arsitektur, desain, dan implementasi sistem perangkat lunak. Diagram ini dapat membantu pengembang dalam memahami hubungan antara komponen perangkat lunak, mengidentifikasi kebutuhan fungsional dan non-fungsional, dan menggambarkan alur logika program dll.

Beberapa tools yang dapat digunakan untuk membuat diagram:
- SmartDraw
- Lucidchart
- Draw.io
- Visio
- Whimsical

2. Horizontal Scaling vs Vertical Scaling
- Horizontal Scaling: Menambah jumlah instance komputer server pada suatu aplikasi dengan tujuan membagi beban kerja
- Vertical Scaling: Menambah kapasitas pada suatu instance, misalnya dengan menambah kekuatan komputer servernya

3. Job/Work Queue
Job Queue atau Work Queue adalah mekanisme untuk menyelesaikan tugas secara terstruktur. Tugas-tugas tersebut ditempatkan pada antrian, dan instance yang tersedia dapat mengambil tugas dari antrian tersebut untuk dieksekusi.

4. Load Balancing
Load Balancing adalah mekanisme untuk membagi beban kerja pada beberapa instance agar tidak terjadi overload pada suatu instance. Load Balancer akan memantau beban kerja pada setiap instance dan memutuskan instance mana yang akan menerima permintaan selanjutnya.

5. Monolitik vs Microservice
- Monolitik: Aplikasi yang dibangun sebagai satu kesatuan yang besar, biasanya sulit untuk melakukan scaling secara horizontal
- Microservice: Aplikasi yang dibangun sebagai kumpulan service yang terpisah, memungkinkan scaling secara horizontal

6. SQL vs NoSQL
- SQL: Database yang memanfaatkan Structured Query Language sebagai bahasa query, cocok untuk data terstruktur
- NoSQL: Database yang tidak memanfaatkan Structured Query Language, cocok untuk data yang tidak terstruktur

7. Caching
Caching adalah mekanisme untuk menyimpan data sementara pada suatu instance, sehingga permintaan selanjutnya tidak perlu memuat data dari sumber asli. Hal ini dapat meningkatkan performa dan mengurangi beban kerja pada sumber asli.

8. Database Indexing
Indexing adalah mekanisme untuk mempercepat pencarian data pada suatu database dengan membuat indeks pada kolom-kolom yang sering digunakan

9. Database Replication
Database Replication adalah mekanisme untuk membuat salinan data dari suatu database pada instance lain, sehingga terdapat cadangan data jika terjadi kegagalan pada suatu instance.