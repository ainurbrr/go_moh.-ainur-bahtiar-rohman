## (21) ORM and Code Structure (MVC)

1. ORM adalah Object Relational Mapping (ORM) adalah teknik pemetaan data antara basis data relasional dan representasi objek dalam bahasa pemrograman. Di Golang, terdapat beberapa ORM populer yang dapat digunakan, seperti GORM, XORM, dan beego ORM.

GORM adalah ORM yang paling populer di Golang. GORM dapat digunakan untuk mengakses basis data relasional seperti MySQL, PostgreSQL, dan SQLite. Dengan GORM, pengguna dapat membuat struktur model untuk tabel basis data relasional dan melakukan operasi CRUD (Create, Read, Update, Delete) pada database. GORM juga menyediakan fitur seperti migrasi basis data, validasi data, dan asosiasi antara tabel.

2. Advantages and Disadvantages using ORM
    - ORM Advantages Less repetitive query, Automatically fetch data into ready to use object, Simple way if you want to screening data before store it in database, Some have feature cache query
    - ORM Disadvantages Add a layer in code and cost the overhead process, Load un-necessary relationship data, Complex raw query can be longy to write with ORM (>10 table joins), Specific SQL function related to one vendor may not supported or no specific function (MySQL : FORCE INDEX)

3. Code Structure & MVC (Model-View-Controller)
    - MVC is short for Model, View, and Controller. MVC is a popular way of organizing your code.
    The big idea behind MVC is that each section of your code has a purpose, and those purposes are different.
    - Why Need Structure ? To achieve to modular application. Implement separation of concerns. Less conflict on versioning.
