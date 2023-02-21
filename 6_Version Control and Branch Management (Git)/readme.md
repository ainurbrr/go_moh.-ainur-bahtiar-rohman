## (6) Versioning Control and Branch Management (Git)

1. Version Control System dan git 
Git adalah satu version control system populer yang digunakan para developer untuk mengembanagkan software secara bersama sam
git dapat mencatat perubahan yang ada pada repository
ada beberapa tahap untuk mengupload file di local ke server git yaitu staging area yaitu 
- git add ... "untuk menambahkan file ke staging area"
- git commit ... "untuk menghubungkan file di staging area ke repository github"
- git push ... "untuk mengupload file setelah dicommit ke repository github"
.gitignore untuk mendeklarasikan file-file yang tidak diupload ke git repository

2. Branching & Pull request
Branch yaitu cabang repository yang digunakan untuk melakukan perubahan" ataupun menambahkan fitur sebagai tempat testing ataupun development sebelum di upload ke master
Pull request yaitu cara untuk meminta pembaharuan dari suatu branch ke branch lain

3. Workflow collaboration
cara melakukan workflow yang baik yaitu :
- Jadikan branch master jarang dirubah ataupun di ganggu
- Hindari mengedit langsung di branch utama ataupun bersama
- Aplikasikan fitur baru hanya ke branch development untuk di test bersama
- Aplikasikan branch development ke branch master jika semua sudah selesai