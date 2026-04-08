package ringin

import (
	"strings"
)

// Kyai mendefinisikan standar tipe fungsi Pengendali atau Pasukan Penjaga Logika di framework ini.
type Kyai func(*Konteks)

// rutePola menyimpan detail URL pendaftaran termasuk susunan pecahannya.
type rutePola struct {
	method   string
	path     string
	segmen   []string
	handlers []Kyai // Slice Kyai, artinya bisa dikelilingi pasukan Middleware penjaga sebelum Handler inti.
}

// RinginContong adalah modul router yang mensupport Rantai Dinamis Penjagaan.
type RinginContong struct {
	rute []rutePola
}

// NewRingin mempersiapkan subsistem router yang baru.
func NewRingin() *RinginContong {
	return &RinginContong{
		rute: make([]rutePola, 0),
	}
}

func pecahJalan(path string) []string {
	parts := strings.Split(path, "/")
	hasil := []string{}
	for _, p := range parts {
		if p != "" {
			hasil = append(hasil, p)
		}
	}
	return hasil
}

// Add mendaftarkan tumpukan protokol rute baru dari yang paling luar (Middleware) sampai dalam (Handler).
func (rc *RinginContong) Add(method, path string, handlers []Kyai) {
	rc.rute = append(rc.rute, rutePola{
		method:   method,
		path:     path,
		segmen:   pecahJalan(path),
		handlers: handlers,
	})
}

// Match memproses URL request saat ini, dan mengekstrak array penjaganya sekalian.
func (rc *RinginContong) Match(method, path string) ([]Kyai, map[string]string) {
	segmenMasuk := pecahJalan(path)

	for _, pola := range rc.rute {
		if pola.method != method {
			continue
		}
		if len(pola.segmen) != len(segmenMasuk) {
			continue
		}

		cocok := true
		params := make(map[string]string)

		for i, seg := range pola.segmen {
			if strings.HasPrefix(seg, ":") {
				key := seg[1:]
				params[key] = segmenMasuk[i]
			} else if seg != segmenMasuk[i] {
				cocok = false
				break
			}
		}

		if cocok {
			return pola.handlers, params // Kalau cocok rutenya, tarik seluruh tim pasukannya
		}
	}

	return nil, nil // Tersesat (404)
}
