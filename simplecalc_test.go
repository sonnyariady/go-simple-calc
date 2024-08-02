package mysimplecalc

import "testing"

func TestJumlah(t *testing.T) {
	hasil := Jumlah(4, 2)
	if hasil != 6 {
		panic("Hasilnya salah")
	}
}

func TestKali(t *testing.T) {
	hasil := Kali(13, 3)
	if hasil != 39 {
		panic("Hasilnya salah")
	}
}

func TestPitagoras(t *testing.T) {
	hasil := Pitagoras(3, 4)
	if hasil != 5 {
		panic("Hasilnya salah")
	}
}

func TestJumlahKuadrat(t *testing.T) {
	hasil := JumlahKuadrat(3, 4, 5)
	if hasil != 50 {
		panic("Hasilnya salah")
	}
}
