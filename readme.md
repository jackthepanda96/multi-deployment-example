## Container setup using docker-compose

### Instalasi docker-compose
pastikan sudah melakukan instalasi docker
```bash
    sudo apt update
    mkdir -p ~/.docker/cli-plugins/
    curl -SL https://github.com/docker/compose/releases/download/v2.23.3/docker-compose-linux-x86_64 -o ~/.docker/cli-plugins/docker-compose
    chmod +x ~/.docker/cli-plugins/docker-compose
```

jalankan perintah ini untuk memastikan instalasi berhasil
```bash
    docker compose version
```
**jika sukses maka akan tampil versi docker compose yang telah diinstall**

Perintah yang docker compose yang digunakan:

- `docker compose up` - digunakan untuk melakukan setup kontainer yang dituliskan pada `docker-compose.yaml`
- `docker compose down` - digunakan untuk mencopot container yang disebutkan dalam `docker-compose.yaml`

untuk parameter yang digunakan, silahakan baca pada `docker compose <up|down> --help`. Silahkan sesuaikan parameter dengan keperluan masing-masing.