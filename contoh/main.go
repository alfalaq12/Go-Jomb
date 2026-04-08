package main

import (
	"net/http"

	gojomb "github.com/bintang/go-jomb"
	"github.com/bintang/go-jomb/banser"
	"github.com/bintang/go-jomb/ringin"
)

// FormulirPendaftaran dicetak untuk menangkap struktur JSON dari payload
type FormulirPendaftaran struct {
	Nama   string `json:"nama"`
	Asal   string `json:"asal"`
	Status string `json:"status"`
}

// JagaSantri adalah CONTOH custom Middleware buatan pengguna (Kalian bisa bikin logika Auth sesuka hati).
// Ini contoh cara menggunakan fitur c.Next() dan c.Berhenti() untuk keamanan JWT dsb.
func JagaSantri() ringin.Kyai {
	return func(c *ringin.Konteks) {
		tokenKTP := c.Request.Header.Get("KTP-Santri")

		if tokenKTP != "ijo-royoroyo" {
			banser.Logf("🚨 Akses ditolak! Parameter sekuritas KTP Santri tidak valid.")
			c.Berhenti(http.StatusUnauthorized)
			c.JSON(http.StatusUnauthorized, ringin.H{
				"error": "Akses Ditolak. Harus verifikasi Token. (401 Unauthorized)",
			})
			return
		}

		// Lolos inspeksi keamanan, lanjutkan proses
		c.Next()
	}
}

func main() {
	app := gojomb.New()

	// 1. Integrasi Middleware Bawaan Go-jomb (Penjaga Logger Performa Global)
	app.Pake(banser.Logger())

	// 2. Rute Publik
	app.GET("/", func(c *ringin.Konteks) {
		c.HTML(http.StatusOK, "<h1>Sugeng Rawuh di Go-jomb!</h1><p>Memacu Aplikasi Kencang Ala Jombang.</p>")
	})

	app.GET("/api/santri/:id", func(c *ringin.Konteks) {
		idTangkapan := c.Param("id")
		c.JSON(http.StatusOK, ringin.H{
			"status":  "sukses",
			"pesan":   "Pencarian profil santri area publik",
			"id_user": idTangkapan,
		})
	})

	// 3. Rute Rahasia (Menggunakan Custom Middleware Authentication Sendiri)
	app.GET("/api/rahasia-dokumen", JagaSantri(), func(c *ringin.Konteks) {
		c.JSON(http.StatusOK, ringin.H{
			"status": "Lolos Verifikasi",
			"data":   "Kitab Babad Tanah Jawi Asli 📜",
		})
	})

	// 4. Test Endpoint POST Input Data Terstruktur
	app.POST("/api/tambah_data", func(c *ringin.Konteks) {
		var inputan FormulirPendaftaran
		if err := c.TangkapJSON(&inputan); err != nil {
			c.JSON(http.StatusBadRequest, ringin.H{"error": "Format JSON tidak valid atau struktur rusak"})
			return
		}
		c.JSON(http.StatusOK, ringin.H{
			"status":    "Penyimpanan Sukses",
			"data_baru": inputan,
		})
	})

	// 5. Eksekusi Server
	if err := app.GassPoll(":8080"); err != nil {
		banser.Logf("Mesin trouble mas: %v", err)
	}
}
