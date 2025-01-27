# Struktur Modular Domain Project Golang (Core and Modules)
## _By: Isra Andika Bakhri_

[![N|Solid](https://cldup.com/dTxpPi9lDf.thumb.png)](https://nodesource.com/products/nsolid)

Seperti yang kita ketahui bahwa Golang memilki keunggulan dalam berbagai macam solusi dalam persoalan di dunia backend seperti ``Goroutine`` dan ``Concurrency``, ``Microservices``, ``gRPC``, ``Rest APi`` serta ``Api gateway``. Namun dari beberapa keunggulan Golang tadi kita perlu merencanakan struktur kode yang sangat ideal bukan cuma dari aspek keterbacaan ataupun dari sisi ``maintable`` tetapi juga tetap memaksimalkan potensi Golang secara powerfull. Sungguh ironis jika seandanya kode yang kita buat sangat bisa dibaca dan ``maintable`` namun di sisi lain justru menggangu performa Golang itu sendiri.

Setidaknya ada 4 indikator yang perlu kita garis bawahi dalam menulis atau menyusun struktur kode Proyek Golang yakni:
- **Scalability** : Setiap domain dapat diatur secara independen, sehingga aplikasi lebih mudah berkembang (Modules)
- **Maintainability** : Kode lebih terorganisir dan mudah dipahami.
- **Reusability** : Repositori, service, dan validator dapat dengan mudah digunakan kembali di berbagai tempat (Core)
- **Team Collaboration** : Mudah untuk mendistribusikan pekerjaan berdasarkan domain.

Ada banyak teknik dalam penulisan struktur kode dan yang paling unggul yakni sistem ``Modular Domain``

Struktur ``Modular Domain`` menolak adanya centralisasi pengelolaan module yang mewakilai tiap table di database. Jika di ibaratkan pemerintah ``Modular Domain`` ini adalah **Pemerintah Otonom** (_disebut_ ``Modules``) seperti **Provinsi/Kabupaten/kota** yang dimana mereka mengelola urusannya sendiri namun tetap dalam backup pusat (_Disebut_ ``Core``)

## Struktur Kode yang Di Usulkan


```sh
project/
├── cmd/
│   └── main.go           // Entry point aplikasi
├── config/
│   └── config.go         // Konfigurasi aplikasi (env, database, dll.)
├── core/
│   ├── handlers/
│   │   ├── auth_handler.go // Shared handler untuk autentikasi
│   │   ├── upload_handler.go // Shared handler untuk upload
│   │   ├── email_handler.go  // Shared handler untuk email
│   │   ├── api_handler.go    // Shared handler untuk API konsumsi
│   ├── services/
│   │   ├── auth_service.go   // Shared service untuk autentikasi
│   │   ├── upload_service.go // Shared service untuk upload
│   │   ├── email_service.go  // Shared service untuk email
│   ├── middlewares/
│   │   ├── cors.go           // Middleware untuk CORS
│   │   ├── jwt.go            // Middleware untuk JWT dan role-based auth
│   │   ├── response.go       // Middleware response JSON
│   │   ├── error.go          // Middleware error JSON
│   ├── utils/
│   │   ├── uuid.go           // UUID generator
│   │   ├── date.go           // Format tanggal
│   │   ├── currency.go       // Format uang
│   │   ├── encrypt.go        // Enkripsi/dekripsi
│   │   ├── file.go           // Helper untuk file upload
├── modules/
│   ├── siswa/
│   │   ├── handlers/
│   │   │   ├── siswa_handler.go    // CRUD handler siswa
│   │   ├── services/
│   │   │   ├── siswa_service.go    // Logika bisnis siswa
│   │   ├── repositories/
│   │   │   ├── siswa_repo.go       // Query database siswa
│   │   ├── models/
│   │   │   ├── siswa.go            // Model siswa
│   │   ├── validators/
│   │   │   ├── siswa_validator.go  // Validator siswa
│   │   ├── siswa_routes.go         // Routing siswa
│   ├── guru/
│   │   ├── handlers/
│   │   │   ├── guru_handler.go
│   │   ├── services/
│   │   │   ├── guru_service.go
│   │   ├── repositories/
│   │   │   ├── guru_repo.go
│   │   ├── models/
│   │   │   ├── guru.go
│   │   ├── validators/
│   │   │   ├── guru_validator.go
│   │   ├── guru_routes.go
├── routes/
│   └── routes.go           // Definisi routing utama
├── templates/
│   └── email.html          // Template email
└── go.mod

```


## Deskripsi Singkat

### Folder core/ untuk Shared Functionality:
  
> Semua fungsi umum seperti ``auth_handler``, ``upload_handler``, dan ``email_handler`` dikelompokkan di folder ``core/handlers``.

> Shared services seperti ``auth_service`` juga disimpan di ``core/services`` untuk menghindari duplikasi kode.

> Dengan cara ini, modul domain ``(siswa, guru, sekolah)`` dapat memanfaatkan layanan ini tanpa menduplikasi logika.

### Modularisasi Domain Tetap Dipertahankan:

> Folder ``modules/siswa`` hanya berisi elemen spesifik untuk domain siswa, sehingga modularitas tetap terjaga.

### Middleware dan Utils Terorganisir:

> Folder middlewares dan utils berisi semua elemen teknis umum seperti **JWT**, **response handler**, dan **utility functions**.

### Shared Upload dan Email Services:

> Shared functionality seperti upload dokumen atau pengiriman email dipusatkan di ``core/services``, sehingga lebih mudah di-maintain dan digunakan kembali.

### Routing Tetap Terpisah:

> File ``routes.go`` hanya bertugas meregistrasi semua routing dari **modul** dan **core.**

### Dampak yang Diharapkan

- **Highly Reusable**: Shared functionality di folder ``core/`` dapat digunakan oleh semua modul, mengurangi duplikasi kode.
- **Maintainability**: Setiap modul domain fokus pada tugasnya tanpa harus mengurusi logika umum.
- **Scalability**: Mudah menambah domain atau fitur baru tanpa mengubah struktur besar aplikasi.
- **Clean Code**: Dengan pemisahan yang jelas, kode lebih mudah dibaca, diubah, dan diuji.


## Menjawab Tantangan Penerapan gRPC

Pada folder ``core/services``, Anda dapat membuat implementasi service **gRPC** yang mengikuti proto file. Setiap service **gRPC** dapat memanfaatkan logika bisnis yang sama dengan HTTP handler

Contoh: ``gRPC SiswaService`` dapat diintegrasikan dengan logika di ``siswa_service.go.``

```sh
service SiswaService {
    rpc GetSiswa (SiswaRequest) returns (SiswaResponse);
}
```

Implementasi di Go:

```sh
type SiswaServiceServer struct {
    siswaService *services.SiswaService // Integrasi ke siswa_service.go
}

func (s *SiswaServiceServer) GetSiswa(ctx context.Context, req *pb.SiswaRequest) (*pb.SiswaResponse, error) {
    siswa, err := s.siswaService.GetSiswaByID(req.Id)
    if err != nil {
        return nil, status.Errorf(codes.NotFound, "Siswa not found")
    }
    return &pb.SiswaResponse{
        Id:   siswa.ID,
        Name: siswa.Name,
    }, nil
}
```

### Proto File Shared Across Modules

Letakkan file ``.proto`` di folder terpisah seperti proto/ agar mudah diakses oleh semua domain atau bahkan proyek lain.

```sh
project/
├── proto/
│   ├── siswa.proto
│   ├── guru.proto
│   ├── sekolah.proto
```


### Generate gRPC Stubs Secara Modular:

File stub hasil dari ``protoc`` bisa di-generate dan disimpan di folder domain terkait:

```sh
modules/
├── siswa/
│   ├── proto/
│   │   ├── siswa.pb.go
├── guru/
│   ├── proto/
│   │   ├── guru.pb.go
```

## Menjawab Tantangan Penerapan Go Rountine

Dengan memisahkan fungsi-fungsi ke dalam domain ```(modules/siswa, modules/guru, dll.)```, setiap domain dapat menjalankan tugas secara independen. Ini berarti Anda bisa menjalankan goroutines di setiap handler atau service untuk memproses permintaan secara paralel.

Pada handler yang membutuhkan data dari beberapa layanan, Anda dapat menggunakan ```goroutines``` dan **WaitGroup** untuk mempercepat eksekusi:

```sh
var wg sync.WaitGroup
wg.Add(2)

go func() {
    defer wg.Done()
    siswaService.GetSiswaData()
}()

go func() {
    defer wg.Done()
    guruService.GetGuruData()
}()

wg.Wait()
```

Dalam layanan yang memerlukan komunikasi antar-proses, seperti pengelolaan batch data _(e.g., batch upload siswa)_, channel bisa digunakan:


```sh
func ProcessBatchUpload(files []string) {
    ch := make(chan string)
    for _, file := range files {
        go func(f string) {
            // Simulate processing
            time.Sleep(2 * time.Second)
            ch <- f + " processed"
        }(file)
    }
    for range files {
        fmt.Println(<-ch)
    }
}
```



## Menjawab Tantangan Penerapan Microservice

Jadikan setiap modul (siswa, guru, sekolah) sebagai layanan independen dengan struktur seperti berikut:

Usulan struktur kode yang merupakan versi minify dari struktur lengkap:
```sh
siswa-service/
├── cmd/
│   └── main.go
├── config/
│   └── config.go
├── handlers/
│   └── siswa_handler.go
├── services/
│   └── siswa_service.go
├── repositories/
│   └── siswa_repo.go
├── models/
│   └── siswa.go
├── validators/
│   └── siswa_validator.go
├── routes/
│   └── siswa_routes.go
└── go.mod
```

Pastikan setiap layanan memiliki database terpisah. Sebagai contoh:
**siswa-service** → D**atabase siswa_db.**
**guru-service** → D**atabase guru_db.**

Setiap layanan memiliki file ```go.mod``` sendiri. Jangan menggabungkan semua dependensi dalam satu file untuk seluruh proyek.


## Menjawab Tantangan Penerapan Api Gateway

Usulan struktur kode khusus implementasi Api Gateway:

```sh
api-gateway/
├── cmd/
│   └── main.go                 // Entry point aplikasi
├── config/
│   ├── config.go               // Konfigurasi aplikasi (env, service URLs, dll.)
├── core/
│   ├── middlewares/
│   │   ├── cors.go             // Middleware untuk CORS
│   │   ├── jwt.go              // Middleware untuk autentikasi JWT
│   │   ├── rate_limit.go       // Middleware untuk rate limiting
│   ├── utils/
│   │   ├── logger.go           // Utility untuk logging
│   │   ├── http_client.go      // Utility untuk HTTP client (proxy request)
├── routes/
│   ├── api_routes.go           // Definisi routing utama
├── handlers/
│   ├── proxy_handler.go        // Handler untuk meneruskan permintaan ke microservice
│   ├── health_handler.go       // Handler untuk health check API Gateway
├── services/
│   ├── auth_service.go         // Service untuk validasi token atau autentikasi
│   ├── cache_service.go        // Service untuk caching (opsional, Redis)
├── docs/
│   └── api_docs.yaml           // Dokumentasi API (Swagger/OpenAPI)
├── internal/
│   ├── load_balancer/
│   │   ├── round_robin.go      // Implementasi load balancer round-robin
│   │   ├── weighted.go         // Implementasi weighted load balancer
│   ├── registry/
│       ├── service_registry.go // Registry untuk mendata microservice (Consul, etcd, hardcoded)
├── templates/
│   └── error_template.html     // Template error untuk respon HTML (jika diperlukan)
├── go.mod
└── go.sum
```

### Komponen Penting dalam Api Gateway
1. ```cmd/main.go```
```sh
func main() {
    // Load configuration
    config := config.LoadConfig()

    // Initialize router
    router := routes.SetupRoutes()

    // Start server
    log.Printf("Starting API Gateway at port %s", config.Port)
    log.Fatal(http.ListenAndServe(":"+config.Port, router))
}
```

2. ```routes/api_routes.go```
```sh
func SetupRoutes() *mux.Router {
    router := mux.NewRouter()

    // Health check
    router.HandleFunc("/health", handlers.HealthHandler).Methods("GET")

    // Proxy route
    router.PathPrefix("/api").HandlerFunc(handlers.ProxyHandler)

    return router
}
```

3. ```handlers/proxy_handler.go```
```sh
func ProxyHandler(w http.ResponseWriter, r *http.Request) {
    targetService := determineService(r.URL.Path)
    proxyURL := targetService + r.URL.Path

    // Forward request to target service
    resp, err := utils.ForwardRequest(proxyURL, r)
    if err != nil {
        http.Error(w, "Service unavailable", http.StatusServiceUnavailable)
        return
    }
    defer resp.Body.Close()

    // Write response back to client
    w.WriteHeader(resp.StatusCode)
    io.Copy(w, resp.Body)
}
```

4. ```middlewares/rate_limit.go```
```sh
func RateLimitMiddleware(next http.Handler) http.Handler {
    limiter := rate.NewLimiter(1, 5) // 1 permintaan/detik, burst 5
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if !limiter.Allow() {
            http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
            return
        }
        next.ServeHTTP(w, r)
    })
}
```

5. ```internal/registry/service_registry.go```
```sh
var serviceRegistry = map[string]string{
    "auth": "http://auth-service:8080",
    "user": "http://user-service:8081",
    "order": "http://order-service:8082",
}

func determineService(path string) string {
    if strings.HasPrefix(path, "/api/auth") {
        return serviceRegistry["auth"]
    } else if strings.HasPrefix(path, "/api/user") {
        return serviceRegistry["user"]
    } else if strings.HasPrefix(path, "/api/order") {
        return serviceRegistry["order"]
    }
    return ""
}
```

6. ```services/auth_service.go```
```sh
func ValidateToken(token string) (bool, error) {
    // Validasi JWT token di sini
    _, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
        return []byte("secret-key"), nil
    })
    if err != nil {
        return false, err
    }
    return true, nil
}
```

## Hasil yang diharapkan Golang sebagai API Gateway
- **Kinerja Tinggi**: Golang mendukung concurrency yang efisien untuk menangani permintaan dalam jumlah besar.
- **Efisiensi Memori**: Golang memiliki footprint memori rendah dibandingkan dengan solusi lain seperti Node.js.
- **Scalability**: Mudah diskalakan dengan menambahkan lebih banyak instance API Gateway.
- **Ecosystem yang Luas**: Banyak library untuk load balancing, rate limiting, logging, dan monitoring.


Jadi, mungkin itu saja konsep yang saya paparkan dalam tulisan ini semoga mendapatkan masukan yang baik untuk menjadi bahan evaluasi kedepannya. Terima Kasih.




