# kudemo
Ini merupakan aplikasi demo untuk buku "Kubernetes untuk Pemula" yang ditulis oleh Giri Kuncoro, Aris Risdianto, and Onno W. Purbo, diterbitkan oleh Penerbit Andi. Terdiri dari 2 aplikasi: halo, sebuah halaman Vue sederhana, dan todo, aplikasi pencatatan tugas sederhana.

# Requirements
Instalasi [Docker](https://www.docker.com/) di komputer.

# Quickstart
## Halo
Untuk menjalankan aplikasi halo menggunakan Docker, bisa menggunakan langkah-langkah berikut:

Clone repositori ini, lalu masuk ke direktori `halo`.
```
cd halo
```
Lalu build kontainer image di lokal.
```
docker build --tag kudemo-halo:1.0 .
```
Jalankan kontainer image di lokal.
```
docker run --publish 8080:80 --detach --name kudemo-halo kudemo-halo:1.0
```
Kunjungi aplikasi pada kontainer dengan browser dengan alamat http://localhost:8080. Akan terlihat aplikasi kudemo-halo aktif dan berjalan.

## Todo
Aplikasi todo terdiri dari 3 komponen: ui, api, dan postgres, bisa dijalankan dengan mudah menggunakan [docker compose](https://docs.docker.com/compose/) di lokal.

Clone repositori ini, lalu masuk ke direktori `todo`.
```
cd todo
```
Uncomment endpoint localhost di [file ini](https://github.com/ngoprek-kubernetes/kudemo/blob/master/todo/ui/src/libs/axios.js#L5).

Lalu jalankan 3 kontainer menggunakan docker compose.
```
docker compose up
```
Kunjungin aplikasi pada browser favorit kamu dengan alamat http://localhost:3000.
