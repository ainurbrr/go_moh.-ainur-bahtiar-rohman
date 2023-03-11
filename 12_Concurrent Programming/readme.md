## (12) Concurrent Programming

1. Sequential Programming vs Parallel Programming vs Concurrent Programming 
- pada Sequential programming sebelum program mengerjakan task baru, task sebelumnya harus finish terlebih dahulu
- pada Parallel programming bisa mengerjakan beberapa task bersamaan dan serentak tetapi membutuhkan multi-core machines
- pada Concurrent programming dapat mengeksekusi banyak task secara independen dan mungkin saja selesainya pun secara serentak

2. Goroutines
Goroutines adalah unit konkuren dalam bahasa pemrograman Go. Goroutine merupakan representasi dari tugas-tugas yang berjalan secara asinkronus, artinya goroutine dapat dijalankan secara bersamaan tanpa harus menunggu goroutine lainnya selesai terlebih dahulu. Goroutine memungkinkan untuk menjalankan banyak tugas secara bersamaan dalam sebuah program tanpa memakan banyak sumber daya komputer.

Dalam Go, sebuah goroutine dapat dijalankan dengan menggunakan keyword "go". Ketika sebuah goroutine dijalankan dengan "go", tugas tersebut akan berjalan secara independen dari goroutine utama dan dapat melakukan pekerjaan di latar belakang tanpa mempengaruhi jalannya program.

3. Channel
Channel adalah salah satu fitur yang penting dalam bahasa pemrograman Go. Channel merupakan sebuah mekanisme yang digunakan untuk melakukan komunikasi antar goroutine dalam sebuah program. Dengan menggunakan channel, goroutine dapat saling berkomunikasi dan bertukar data dengan aman dan terkoordinasi. Channel dapat dibuat dengan menggunakan fungsi built-in make(). Secara umum, sebuah channel memiliki tiga properti utama yaitu tipe data yang akan ditransferkan, buffer size, dan status operasi (closed atau open).
Channel bisa digunakan untuk mengirim dan menerima data. Untuk mengirim data ke sebuah channel, digunakan operator <- dengan nilai yang akan dikirim. Sedangkan untuk menerima data dari sebuah channel, operator <- digunakan tanpa nilai.

4. Select
Select merupakan salah satu fitur penting dalam bahasa pemrograman Go yang digunakan untuk melakukan pemilihan antara beberapa operasi yang berjalan secara konkuren. Dengan menggunakan select, kita dapat melakukan pemilihan operasi yang siap untuk dieksekusi di antara beberapa channel yang aktif. Select digunakan untuk menangani kasus dimana ada beberapa operasi yang harus dijalankan secara konkuren, tetapi kita tidak tahu operasi mana yang akan selesai terlebih dahulu.

Select terdiri dari beberapa klausa yang berisi operasi-operasi yang akan dieksekusi. Setiap klausa dipisahkan oleh operator case dan diakhiri dengan operator :. Operasi-operasi yang ditempatkan di dalam klausa dapat berupa pengiriman atau penerimaan data pada channel.

5. Race condition
adalah kondisi dimana dua atau lebih goroutine dalam sebuah program Go mencoba untuk mengakses dan memodifikasi sebuah sumber daya bersama seperti variabel, file, atau resource lainnya pada saat yang sama tanpa ada koordinasi yang memadai. Akibatnya, perilaku program menjadi tidak dapat diprediksi dan sumber daya bersama tersebut dapat mengalami data race atau race condition.

Go merupakan bahasa pemrograman yang mendukung konkurensi, sehingga race condition dapat terjadi pada aplikasi Go yang menggunakan goroutine. Untuk menghindari race condition pada Go, kita dapat menggunakan beberapa teknik seperti menggunakan mutex, channel, atau atomic operations. Selain itu, Go juga memiliki tools bawaan seperti race detector yang dapat membantu mengidentifikasi dan menangani race condition pada aplikasi Go.

Untuk menghindari race condition dalam program Go, ada beberapa teknik yang dapat digunakan seperti:

Menggunakan mutex: Mutex atau mutual exclusion adalah teknik yang digunakan untuk mengatur akses ke sumber daya bersama dalam lingkungan konkuren. Dengan menggunakan mutex, hanya satu goroutine yang dapat mengakses sumber daya pada satu waktu tertentu.
Menggunakan channel: Channel dapat digunakan untuk melakukan komunikasi dan koordinasi antar goroutine dalam program Go. Dengan menggunakan channel, goroutine dapat saling berkomunikasi dan bertukar data dengan aman dan terkoordinasi.
Menggunakan atomic operations: Atomic operations adalah operasi-operasi yang dapat dilakukan secara atomic atau seutuhnya, tanpa ada kemungkinan terjadinya race condition. Operasi-operasi ini sering digunakan untuk melakukan operasi penghitungan atau penambahan pada variabel yang digunakan bersama oleh beberapa goroutine.
Menggunakan package sync: Package sync adalah package standar dalam Go yang menyediakan beberapa fitur untuk mengatur akses dan sinkronisasi dalam lingkungan konkuren seperti WaitGroup dan Once.