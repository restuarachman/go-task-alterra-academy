package main

type Kendaraan struct {
	totalRoda int
	kecepatanPerJam int
}

// menambahkan kecepatanPerJam sebanyak 10
func (kendaraan *Kendaraan) berjalan() {
	kendaraan.tambahKecepatan(10)
}

// Menambah kecepatanPerJam dengan kecepatanBaru
func (kendaraan *Kendaraan) tambahKecepatan(kecepatanBaru int) {
	kendaraan.kecepatanPerJam = kendaraan.kecepatanPerJam + kecepatanBaru
}

func main() {
	mobilCepat := Kendaraan{}
	mobilCepat.berjalan()
	mobilCepat.berjalan()
	mobilCepat.berjalan()

	mobilLamban := Kendaraan{}
	mobilLamban.berjalan()
}