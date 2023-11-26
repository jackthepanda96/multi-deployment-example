## Alur setup project saat ini

Penjelasan dibawah adalah alur `GITHUB ACTION` yang kita buat


1. Github action akan melakukan build image dari project kita berdasarkan `dockerfile`
2. Setelah image sukses dibuat, akan di push ke `dockerhub`
3. Setelah proses push selesai, github action akan login ke server kita melalui SSH dengan kredensial yang sudah kita simpan di bagian secrets
4. Setelah berhasil login, github action akan menjalankan perintah pada bagian `scripts` sesuai dengan apa yang kita tuliskan pada bagian tersebut. Environment variable untuk container kita tuliskan lengkap pada perintah `docker run`.