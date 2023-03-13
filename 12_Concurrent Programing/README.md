# Concurrent programing
===============================

ketika kita membuat sebuah pncarian website itu ada 3 hasil :
  - hasil websitenya
  - video
  - image

  cara kerjanya salah satunya sequential , paralel, concurrent

* seqeuntial program 
Berurutan , dalam program seqeuntial, sebelum tugas baru dimulai, tugas sebelumnya harus selesai

* parallel programs
dalam program paralel, banyak tugas dapat dijalankan secara bersamaan

* concureent programs
dalam program bersamaan, banyak tugas dapat dijalankan secara independen dan mungkin muncul secara bersamaan
  

  + Concureewncy allow paralleslism 
  Go's concureency (gorountines) mempermudah pembuatan program paralel yang memanfaatkan mesin dengan banyak prosesor (sebagian besar mesin saat ini)


+ feature go
 - ekskusi serentak
 - sinkronisasi dan pengiriman pesan
 - kontrol bersamaan multi-arah
  

  * Goroutine
  adalah fungsi atau metode yang dijalankan secara bersamaan (indenpenden ) dengan fungsi atau metode lain.


  * Gomaxprocs
  digunakan untuk mengontrol jumlah utas sistem operasi yang tersedia untuk eksekusi goroutine ke program go tertentu

  + chanel & select
  chanel adalah objek komunikasi yang digunakan goroutine untuk berkomunikasi satu sama lain

  + select 
  pilih membuatnya lebih mudah untuk mengontrol komunikasi data melalui satu atau banyak saluran

  * Race condition
  
kondisi balapan adalah saat 2 utas mengakses memori pada saat yang sama, salah satunya adalah penulisan. Kondisi balapan terjadi karena akses yang tidak disinkronkan ke memori bersama


* memperbaiki ras data
1# waitgroups
   blocking with waitgroups
   cara paling mudah untuk menyelesaikan data race, adalah memblokir akses baca hingga operasi tulis selesai

2# CHANEL BLOCKING
ini memungkinkan goroutine kita untuk melakukan sinkronisasi satu sama lain untuk sesaat

3# Mutex
di mana kami tidak peduli dengan urutan baca dan tulis, kami hanya mensyaratkan bahwa itu tidak terjadi secara bersamaan