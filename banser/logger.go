package banser

import (
	"log"
	"os"
	"time"

	"github.com/bintang/go-jomb/ringin"
)

var (
	// DefaultLogger adalah pemantau sistem bawaan yang mengarahkan pesan proses ke terminal.
	DefaultLogger = log.New(os.Stdout, "[BANSER SEC] ", log.LstdFlags|log.Lmsgprefix)
)

// Logf mencetak format pesan eksekusi ke saluran keluaran layar dengan standar yang rapi.
func Logf(format string, v ...any) {
	DefaultLogger.Printf(format, v...)
}

// === IMPLEMENTASI MIDDLEWARE BAWAAN GO-JOMB ===

// Logger adalah Penjaga bawaan framework yang bertugas mencetak kecepatan (latency) eksekusi rute API.
// Cocok dipasang secara Global.
func Logger() ringin.Kyai {
	return func(c *ringin.Konteks) {
		waktuTiba := time.Now()

		// Izinkan gerbong urutan berikutnya (Handler aslinya) tereksekusi.
		c.Next()

		// Setelah proses di dalam handler beres total, kita inspeksi durasinya.
		durasi := time.Since(waktuTiba)
		Logf("Rute Terakses | %s %s - %v", c.Request.Method, c.Request.URL.Path, durasi)
	}
}
