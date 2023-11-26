## Mekanisme setup yang baru

1. Lakukan instalasi docker compose
2. Buat docker-compose.yaml dengan acuan yang tertera pada repository
3. Lakukan modifikasi pada action `appleboy/ssh-action@v1.0.0`
4. Tambahkan bagian `env` setelah bagian `uses`, dengan format sebagai berikut sesuai dengan kebutuhan. Contoh:
    ```yaml
        name: login server with ssh
        uses: appleboy/ssh-action@v1.0.0
        env:
            var1: ${{ secrets.nama_parameter_github_secret1 }}
            var2: ${{ secrets.nama_parameter_github_secret2 }}
    ```
5. Pada segmen `with` tambahkan `envs` yang berisi variable apa saja yang akan diteruskan ke server. Contoh
    ```yaml
        with:
            envs: var1, var2, var3, var4
    ```
6. Pada bagian awal perintah `script` pastikan untuk melakukan set environment variable dilanjutkan dengan perintah yang diperlukan. Contoh:
    ```bash
        export var1=$var1
        export var2=$var2
        
        #lanjutkan perintah sesuai kebutuhan
    ```
7. Untuk menggunakan perintah `docker compose` kita wajib memiliki `docker-compose.yaml` file yang berisi konfigurasi setup container yang kita perlukan. Untuk melakukan hal ini, kita bisa melakukan clone project / download file `docker-compose.yaml` dari github kita.
8. Jalankan `docker compose up` untuk melakukan instalasi seluruh keperluan.
9. Jalankan `docker compose down` untuk melakukan teardown seluruh instalasi