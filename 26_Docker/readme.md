## (26) Docker

1. Docker adalah platform perangkat lunak open-source yang digunakan untuk membangun, mengemas, dan menjalankan aplikasi di lingkungan yang terisolasi yang disebut container. Dalam lingkungan container, aplikasi dan dependensinya dikemas bersama dengan konfigurasi dan sumber daya yang diperlukan untuk menjalankannya, sehingga memungkinkan aplikasi berjalan secara konsisten di berbagai lingkungan yang berbeda. Docker memanfaatkan teknologi container untuk mengisolasi aplikasi dan memastikan bahwa mereka berjalan dengan aman dan konsisten, serta memudahkan pengembangan, pengujian, dan pengiriman aplikasi secara efisien. Docker juga menyediakan berbagai layanan dan alat untuk mengelola dan memonitor lingkungan container, seperti Docker Hub untuk menyimpan dan berbagi image Docker, dan Docker Swarm untuk mengelola cluster container yang besar.

2. Infrastruktur Docker juga terdiri dari beberapa komponen lainnya, seperti Docker Engine, Docker Compose, dan Docker Swarm. Docker Engine adalah mesin runtime yang menjalankan container, sedangkan Docker Compose adalah alat yang digunakan untuk mengelola aplikasi multi-container. Docker Swarm adalah alat pengelolaan cluster yang digunakan untuk mengelola dan memonitor banyak container dalam lingkungan yang luas. Dalam hal ini, Docker menyediakan infrastruktur yang mudah digunakan dan efisien untuk pengembangan, pengujian, dan pengiriman aplikasi.

3. Docker Volume Docker Volume bisa dianggap sebagai storage atau tempat penyimpanan data di container. Tentunya saat kita membuat container kita tidak ingin ketika container kita mati data yang ada pada container ikut terhapus juga. Untuk itu kita dapat memanfaatkan Volume pada docker. 

4. Docker Network Defaultnya container pada docker akan saling terisolasi satu sama lain. Kita tidak dapat melakukan request api (misal) dari container satu ke container lain. Untuk itu kita harus membuat dan mendaftarkan container pada network yang sama.

