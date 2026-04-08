package gojomb

import (
	"net/http"
	"sync"

	"github.com/bintang/go-jomb/banser"
	"github.com/bintang/go-jomb/ringin"
)

// Pendopo berperan sebagai mesin aplikasi (core engine) untuk framework go-jomb.
type Pendopo struct {
	router      *ringin.RinginContong
	middlewares []ringin.Kyai // Wadah buat penjaga Banser level Global Murni (Applies to All).
	pool        sync.Pool
}

// New menginisialisasi instansi mandiri dari mesin Pendopo.
func New() *Pendopo {
	p := &Pendopo{
		router:      ringin.NewRingin(),
		middlewares: make([]ringin.Kyai, 0),
	}
	p.pool.New = func() any {
		return &ringin.Konteks{}
	}
	return p
}

// Pake (Use) mendaftarkan Middleware Banser tingkat dewa yang selalu hidup di tiap request apapun.
func (p *Pendopo) Pake(middlewares ...ringin.Kyai) {
	p.middlewares = append(p.middlewares, middlewares...)
}

// GET menyusupkan Rute HTTP GET sambil mengikat berlapis-lapis penjaga bersama handler utama terujung.
func (p *Pendopo) GET(path string, handlers ...ringin.Kyai) {
	// Menumpuk penjaga global terlebih dahulu, lalu diiringi penjaga khusus (jika ada), ditutup dengan Handler Pokok.
	gabunganSusunan := append(p.middlewares, handlers...)
	p.router.Add(http.MethodGet, path, gabunganSusunan)
}

// POST mendaftarkan susunan gerbong HTTP POST persis seperti GET.
func (p *Pendopo) POST(path string, handlers ...ringin.Kyai) {
	gabunganSusunan := append(p.middlewares, handlers...)
	p.router.Add(http.MethodPost, path, gabunganSusunan)
}

// ServeHTTP diakui oleh Golang asli sebagai perawakan server web sesungguhnya.
func (p *Pendopo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := p.pool.Get().(*ringin.Konteks)
	ctx.Reset(w, r)

	handlers, tangkapanSistem := p.router.Match(r.Method, r.URL.Path)
	if handlers != nil {
		ctx.Params = tangkapanSistem
		ctx.SetHandlers(handlers) // Susun panitia gerbong kereta
		ctx.Next()                // Beri instruksi ke pimpinan regu (masinis) untuk memacu barisan
	} else {
		ctx.JSON(http.StatusNotFound, ringin.H{
			"error": "Ngapunten, rute niki mboten wonten (404 Not Found).",
		})
	}

	p.pool.Put(ctx)
}

// GassPoll menghidupkan lalu lintas pengawasan terminalnya dengan cepat.
func (p *Pendopo) GassPoll(addr string) error {
	banser.Logf("Mesin Pendopo wis mbedhal nang http://localhost%s 🚀", addr)
	return http.ListenAndServe(addr, p)
}
