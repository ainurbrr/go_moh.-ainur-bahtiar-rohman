## (11) Problem Solving Paradigm - Brute Force Greedy and Dynamic Programming

1. Problem Solving Paradigm adalah cara untuk memnyelesaikan problem
- Brute force(complete search) adalah sebuah metode untuk memecahkan masalah dengan mencari ke semua bagian yang memungkinkan untuk mendapatkan solusi yang diinginkan brute force merupakan algoritma yang mudah dipahami tetapi kompleksitas waktunya kurang efisien contohnya untuk mencari nilai min & max
- Divide and conquer (D&C) adalah paradigma pemecahan masalah dimana suatu masalah dibuat lebih sederhana dengan membagi menjadi bagian-bagian menjadi lebih kecil kemudian mencari di bagian yang lebih kecil tsb. contohnya binary search
- Greedy algorithm dapat membuat lebih optimal pada setiap langkahnya dengan harapan pada akhirnya mencapai mencapai solusi yang optimal. dalam beberapa kasus greesy bekerja dan menghasilkan solusi yang singkat dan berjalan efisien contohnya djikstra, huffman coding dll.

2. Dynamic Programming adalah algoritma untuk memecahkan masalah secara optimal  dengan cara memecahnya menjadi submaskah yang sederhana dan memanfaatkan hasilnya menjadi solusi yang optimal keseluruhan

3. Dynamic Programming method
- Top-down with Memoization dalam pendekatan  ini kita mencoba memecahkan masalah yang lebih besar dengan mencari solusi secara rekursif untuk sub-masalahnya dan memecahkannya pada tiap sub masalah dan menyimpannya dalam cache contohnya memecahkan fibonacci secara rekursi
- Bottom-up with Tabulation adalah kebalikan dari top-down yaitu pemecahan masalah dengan cara menghindari menggunakan rekursi yaitu dengan cara menyelsaikan sub masalah terlebih dahulu dan menyimpannya dalam tabulasi.