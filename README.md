# 🚀 Realtime Chat App Backend (Golang)

Aplikasi Backend Chat Realtime berperforma tinggi yang dibangun menggunakan **Go (Golang)**, **PostgreSQL**, dan **Docker**. Aplikasi ini dirancang untuk skalabilitas dan kemudahan deployment.
![GitHub Logo](https://cdn.prod.website-files.com/6100d0111a4ed76bc1b9fd54/62217e885f52b860da9f00cc_Apa%20Itu%20Golang%3F%20Apa%20Saja%20Fungsi%20Dan%20Keunggulannya%20-%20Binar%20Academy.jpeg)

## ✨ Fitur Utama
* **User Authentication**: Register & Login menggunakan JWT (JSON Web Token).
* **Realtime Communication**: Pengiriman pesan instan secara dua arah menggunakan WebSockets (Gorilla WebSocket).
* **Secure**: Enkripsi password menggunakan Bcrypt.
* **Database Migration**: Skema database otomatis dikelola saat aplikasi pertama kali dijalankan di Docker.
* **Dockerized**: Siap dijalankan di server mana pun hanya dengan satu perintah.

## 🛠️ Teknologi yang Digunakan
* **Language**: Go 1.24+
* **Framework**: Gin Gonic
* **Database**: PostgreSQL 14
* **Container**: Docker & Docker Compose

## 🚀 Cara Menjalankan (Quick Start)

Pastikan kamu sudah menginstall **Docker** dan **Docker Compose** di laptopmu.

1. **Clone Repositori**
   ```bash
   git clone [https://github.com/AgustiBayu/realtime-chatapp.git](https://github.com/AgustiBayu/realtime-chatapp.git)
   cd realtime-chatapp
2. **Jalankan dengan Docker**
Cukup jalankan perintah berikut, Docker akan mengatur Database PostgreSQL dan Backend secara otomatis:
   ```bash
   docker-compose up --build
3. **Akses Aplikasi** akan berjalan pada: http://localhost:8080

## 🧪 Cara Pengujian (Testing)
1. **Registrasi User Baru**
Sebelum bisa chatting, kamu harus membuat akun terlebih dahulu.
* **Method:** POST
* **Url:**
   ``` bash 
   http://localhost:8080/v1/api/auth/register
* **Json Body:**
   ``` bash 
   {
    "name":"",
    "email":"",
    "password":""
   }
2. **Login (Mendapatkan Token)**
Login digunakan untuk mendapatkan JWT Token yang diperlukan untuk mengakses fitur chat.
* **Method:** POST
* **Url:**
   ``` bash 
   http://localhost:8080/v1/api/auth/login
* **Json Body:**
   ``` bash 
   {    		
    "email":"",
    "password":""
   }
* **Output:** Copy nilai token yang muncul di response untuk digunakan pada langkah berikutnya.
3. **Mengambil Riwayat Chat**
* **Method:** GET
* **Url:**
   ```bash 
   http://localhost:8080/messages/history?receiver_id=1
* **Authorization:** Pilih tab Auth, pilih Bearer Token, lalu paste token hasil login tadi.
4. **Mencoba WebSocket (Realtime)**
* **Type:** WebSocket Request
* **Url:**
   ``` bash
   ws://localhost:8080/messages/ws?token=PASTE_TOKEN_DISINI
* **Message (JSON):** 
   ``` bash
   {
    "receiver_id": 2,
    "content": "Halo, ini pesan realtime via Channels!"
   }