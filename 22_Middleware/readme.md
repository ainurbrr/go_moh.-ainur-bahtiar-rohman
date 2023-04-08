## (22) Middleware

1. Middleware adalah fungsi yang digunakan untuk memodifikasi request dan response dalam proses pengiriman data antara klien dan server. Middleware dapat digunakan untuk melakukan berbagai tugas, seperti melakukan logging, melakukan autentikasi, menambahkan header pada request dan response, dan banyak lagi.

2. Log middleware di Golang adalah salah satu jenis middleware yang berfungsi untuk mencatat setiap permintaan yang masuk ke dalam server ke dalam file log. Log ini dapat digunakan untuk keperluan debugging dan pemantauan performa server dll.

3. Auth middleware di Golang adalah salah satu jenis middleware yang berfungsi untuk melakukan autentikasi pengguna sebelum mengakses resource tertentu di dalam server. Middleware ini biasanya digunakan untuk melindungi resource yang sensitif seperti halaman profil pengguna atau data pribadi contohnya:
    - Basic Authentication
    - Authentication menggunakan JWT(JSON Web Token)