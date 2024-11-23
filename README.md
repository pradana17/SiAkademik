# SiAkademik

# Sistem Informasi Akademik

Aplikasi ini adalah sistem manajemen akademik yang dirancang untuk memudahkan pengelolaan data pengguna, mata kuliah, peran, dan pendaftaran mahasiswa. Sistem ini menggunakan framework **Gin** di Go (Golang) untuk menangani permintaan HTTP dan API. Aplikasi ini juga dilengkapi dengan fitur autentikasi menggunakan **JWT (JSON Web Token)** untuk memastikan keamanan setiap permintaan yang diterima oleh server.

Aplikasi ini memiliki berbagai macam endpoint yang dibagi berdasarkan peran pengguna (role-based access control), seperti **Admin**, **Dosen**, dan **Mahasiswa**. Setiap peran memiliki akses yang berbeda, sesuai dengan fungsinya.

### Fitur Utama:
1. **Autentikasi Pengguna**: Pengguna dapat login untuk mendapatkan token JWT yang diperlukan untuk mengakses endpoint yang dilindungi.
2. **Manajemen Pengguna dan Peran**: Admin dapat mengelola pengguna dan peran di sistem.
3. **MataKuliah dan Semester**: Admin dapat membuat MataKuliah dan semester baru.
4. **Pendaftaran Mahasiswa**: Mahasiswa dapat mendaftar ke mata kuliah dan melihat nilai serta IPK mereka.
5. **Manajemen Nilai**: Dosen dapat memberikan nilai untuk mahasiswa.

### Teknologi yang Digunakan:
- **Go (Golang)**: Bahasa pemrograman untuk backend.
- **Gin**: Web framework untuk Golang yang digunakan untuk menangani permintaan HTTP.
- **JWT**: Digunakan untuk autentikasi dan otorisasi.
- **Swagger**: Digunakan untuk dokumentasi API.

## Endpoints API
Di bawah ini adalah daftar lengkap endpoint yang tersedia dalam aplikasi ini, yang terbagi berdasarkan peran pengguna.
| **Endpoint**                        | **Method**  | **Description**                                  | **Authentication**      | **Roles**        |
|-------------------------------------|-------------|--------------------------------------------------|-------------------------|------------------|
| `/auth/login`                       | POST        | Login untuk mendapatkan JWT token                | Tidak diperlukan         | Semua pengguna   |
| `/userprofile`                      | GET         | Mendapatkan profil pengguna                      | Diperlukan Authorization | Semua pengguna   |
| `/user`                             | PUT         | Memperbarui informasi pengguna                   | Diperlukan Authorization | Semua pengguna   |
| `/userprofile`                      | PUT         | Memperbarui profil pengguna                      | Diperlukan Authorization | Semua pengguna   |
| `/admin/user`                       | POST        | Membuat pengguna baru                            | Diperlukan Authorization | Admin            |
| `/admin/role`                       | POST        | Membuat peran baru                               | Diperlukan Authorization | Admin            |
| `/admin/role/{id}`                  | DELETE      | Menghapus peran berdasarkan ID                    | Diperlukan Authorization | Admin            |
| `/admin/course`                     | POST        | Membuat kursus baru                              | Diperlukan Authorization | Admin            |
| `/admin/semesters`                  | POST        | Membuat semester baru                            | Diperlukan Authorization | Admin            |
| `/dosen/course`                     | GET         | Mendapatkan kursus yang diajarkan oleh dosen     | Diperlukan Authorization | Dosen            |
| `/dosen/grade`                      | POST        | Membuat nilai untuk mahasiswa                    | Diperlukan Authorization | Dosen            |
| `/mahasiswa/enrollment`             | POST        | Mendaftar mata kuliah untuk mahasiswa            | Diperlukan Authorization | Mahasiswa        |
| `/mahasiswa/gpa`                    | GET         | Mendapatkan GPA mahasiswa                        | Diperlukan Authorization | Mahasiswa        |
| `/mahasiswa/course`                 | GET         | Mendapatkan daftar mata kuliah yang diambil mahasiswa | Diperlukan Authorization | Mahasiswa        |
