# 🏛️ Go-jomb Framework

**Go-jomb** adalah *High-Performance Micro-Framework* untuk Golang yang terinspirasi dari arsitektur web framework kelas enterprise internasional, namun tetap setia membawakan nuansa kearifan lokal "Kota Santri" Jombang. 

Framework ini dirancang memiliki performa secepat **Gin**, kemudahan instalasi **Fiber**, dengan manajemen memori kelas atas berkat implementasi fitur *Sync Pool Zero-Allocation*. Sangat ideal bagi kalangan developer Golang di Nusantara yang memprioritaskan produktivitas dan kualitas *Developer Experience (DX)* tingkat tinggi.

---

## 🚀 Fitur Unggulan Go-jomb

1. **Pendopo Engine (`gojomb`)** 🏭
   - Mesin server inti berbasis `http.Handler` standar murni Golang. Tahan banting dan dibekali manajemen alokasi Pool untuk meringankan beban **Garbage Collector (GC)**.

2. **Router Ringin Contong (`ringin`)** 🚥
   - Pengatur rute otomatis (*Router System/Mux*) yang fleksibel. Secara otomatis mengekstrak *Dynamic Parameters Extractor* untuk path param seperti (`/api/data/:id`).

3. **Konteks JSON Dinamis (`ringin.H{}`)** 📦
   - Selamat tinggal konfigurasi header membosankan! Menyusun respons JSON yang dinamis, me-render String dan HTML sekarang semudah menjentikkan jari.
   - Punya alat pembedah *Body* langsung ke Struct Golang via `TangkapJSON()`.

4. **Sistem Pengamanan Middleware Berlapis (Banser)** 🛡️
   - Terintegrasi penuh dengan metode pendelegasian rantai *Middlewares*. Pasang penjagaan seperti pencegat *Autentikasi (JWT/Token)* dan *Logger server* di level Global maupun tingkat Khusus di satu fungsi eksekusi rute tanpa masalah!

---

## 📦 Pemasangan (Installation)

Siapkan terminal ruang kerja Anda, jalankan *go-getter* sakti berikut:
```bash
go get -u github.com/alfalaq12/Go-Jomb
```

---

## 🗂️ Kamus Padanan Istilah Lokal
| Kosakata Go-jomb | Makna Global Konvensional |
| --- | --- |
| **`Pendopo`** | Server Engine / Instance Web Utama (`app`) |
| **`RinginContong`** | Komponen Router Inti Pembelah URL (`mux`) |
| **`Kyai`** | Eksekutor Logika HTTP / `HandlerFunc`  |
| **`Banser`** | *Middlewares*, Logger dan Keamanan (*Defense*) |
| **`GassPoll()`** | Pemacu Server Listen Port (`http.ListenAndServe`) |

---

## 💻 Penggunaan Utama (Quick Start)

Lihat seberapa cantik kodenya saat mengeksekusi program REST API sederhana.

```go
package main

import (
	"fmt"
	"net/http"

	gojomb "github.com/bintang/go-jomb"
	"github.com/bintang/go-jomb/banser"
	"github.com/bintang/go-jomb/ringin"
)

func main() {
	// 1. Inisiasi Mesin Utama
	app := gojomb.New()

	// 2. Pasang Penjaga Banser Logger Global untuk mencatat kecepatan respons di terminal
	app.Pake(banser.Logger())

	// 3. Tarik Tali Rute Dinamis dengan ekstraksi ":id"
	app.GET("/api/jamaah/:id", func(c *ringin.Konteks) {
		idTangkapan := c.Param("id") 
		
        // 4. Balas seketika lemparkan Output JSON yang manis
		c.JSON(http.StatusOK, ringin.H{
			"status":  "sukses",
			"pesan":   fmt.Sprintf("Mencari informasi jamaah ID %s", idTangkapan),
			"id_user": idTangkapan,
		})
	})

	// 5. GassPoll Server
	if err := app.GassPoll(":8080"); err != nil {
		banser.Logf("Mesin trouble: %v", err)
	}
}
```

## 🤝 Turut Berkontribusi (Soko Guru)
Bangga rilis murni arsitektur anak bangsa! Sangat terbuka bebas (*Open-Source*) bagi rekan-rekan Developer Nusantara maupun internasional yang memiliki ide perbaikan (Misal implementasi ORM Database *Brantas*, WebSocket terintegrasi, hingga optimasi radix tree murni). **Silakan layangkan *Pull-Request* Anda!**

---
🚀 *Crafted with ❤️ and Code in Jombang.*
