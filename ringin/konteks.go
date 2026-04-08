package ringin

import (
	"encoding/json"
	"math"
	"net/http"
)

// H adalah pintasan untuk map[string]any, terinspirasi dari gin.H.
type H map[string]any

// indexBerhenti menandakan angka ekstrem biar perulangan middleware distop paksa.
const indexBerhenti = int8(math.MaxInt8 / 2)

// Konteks mewakili urat nadi kehidupan dari HTTP request yang sedang berjalan.
type Konteks struct {
	Request *http.Request
	Writer  http.ResponseWriter
	Params  map[string]string // Tangkapan variabel dari rute URL dinamis

	// Fitur Middleware Rantai Pengawal
	handlers []Kyai
	index    int8
}

// Reset digunakan untuk menginisiasi ulang Konteks saat ditarik dari memori daur ulang.
func (c *Konteks) Reset(w http.ResponseWriter, r *http.Request) {
	c.Request = r
	c.Writer = w
	c.Params = make(map[string]string)
	c.index = -1 // Mulai rute rantai dari titik nol
	c.handlers = nil
}

// === MESIN RANTAI MIDDLEWARE ===

// SetHandlers menaruh rantai pasukan penjaga beserta handler tujuan akhir ke dalam konteks ini.
func (c *Konteks) SetHandlers(handlers []Kyai) {
	c.handlers = handlers
}

// Next mengizinkan middleware di lapisan ini untuk mengeksekusi lapisan fungsi selanjutnya.
// Ini adalah kunci agar middleware bisa menghitung durasi "sebelum" vs "sesudah" dieksekusi.
func (c *Konteks) Next() {
	c.index++
	for c.index < int8(len(c.handlers)) {
		c.handlers[c.index](c)
		c.index++
	}
}

// Berhenti memaksa potong rantai eksekusi (cocok dipakai jika token ditolak).
// Mirip fungsi Abort() di framework Gin.
func (c *Konteks) Berhenti(kode int) {
	c.index = indexBerhenti
	c.Writer.WriteHeader(kode)
}

// === ALAT TANGKAP INPUTAN ====

func (c *Konteks) Param(kunci string) string {
	return c.Params[kunci]
}

func (c *Konteks) Query(kunci string) string {
	return c.Request.URL.Query().Get(kunci)
}

func (c *Konteks) TangkapJSON(obj any) error {
	defer c.Request.Body.Close()
	return json.NewDecoder(c.Request.Body).Decode(obj)
}

// === ALAT PENGELUARAN (OUTPUT) ====

func (c *Konteks) JSON(code int, obj any) {
	c.Writer.Header().Set("Content-Type", "application/json")
	// Kita pastikan header sukses belum tertimpa instruksi penghentian paksa
	if c.index < indexBerhenti {
		c.Writer.WriteHeader(code)
	}
	if err := json.NewEncoder(c.Writer).Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}
}

func (c *Konteks) HTML(code int, html string) {
	c.Writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	if c.index < indexBerhenti {
		c.Writer.WriteHeader(code)
	}
	c.Writer.Write([]byte(html))
}
