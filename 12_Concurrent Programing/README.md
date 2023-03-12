# Pengenalan Concurrent Program merujuk pada kemampuan untuk menjalankan beberapa tugas atau proses secara simultan atau bersamaan dalam sebuah program. Golang didesain dengan built-in support untuk concurrent programming dengan menggunakan goroutine, channel, dan mutex.

## Goroutine adalah unit dasar untuk concurrent programming di Golang, yang merupakan sebuah fungsi yang dapat dijalankan secara bersamaan dengan fungsi lainnya di dalam program. Goroutine dijalankan secara asinkron dan memiliki kemampuan untuk mengkomunikasikan data dengan menggunakan channel.

## Channel adalah sebuah struktur data yang digunakan untuk mengirim dan menerima data antar goroutine. Channel memungkinkan goroutine untuk berkomunikasi dengan cara yang aman dan terkontrol, sehingga menghindari adanya masalah seperti race condition dan deadlock.

## Mutex adalah sebuah mekanisme yang digunakan untuk mengatur akses ke sumber daya bersama seperti variabel atau file. Mutex dapat membantu menghindari adanya masalah seperti race condition dan memastikan bahwa hanya satu goroutine yang dapat mengakses sumber daya bersama pada suatu waktu.

### Dengan dukungan untuk goroutine, channel, dan mutex, Golang memungkinkan para pengembang untuk membuat program yang efisien, scalable, dan mudah untuk dikelola dalam lingkungan concurrent yang kompleks.