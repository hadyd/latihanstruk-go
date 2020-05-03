package main

import "fmt"

type Sepeda struct {
	warnaSepeda string
	jumlahBan   int
	jumlahGigi  int
}

func (s *Sepeda) waktuTempuh(jarak float32) float32 {

	waktutempuh := jarak * 2.5

	return waktutempuh
}

func main() {
	dataSepeda := make([]Sepeda, 5, 5)
	dataSepeda[0] = Sepeda{warnaSepeda: "Merah", jumlahBan: 2, jumlahGigi: 6}
	dataSepeda[1] = Sepeda{warnaSepeda: "Kuning", jumlahBan: 2, jumlahGigi: 5}
	dataSepeda[2] = Sepeda{warnaSepeda: "Biru", jumlahBan: 2, jumlahGigi: 3}
	dataSepeda[3] = Sepeda{warnaSepeda: "Hijau", jumlahBan: 2, jumlahGigi: 2}
	dataSepeda[4] = Sepeda{warnaSepeda: "Putih", jumlahBan: 2, jumlahGigi: 1}
	fmt.Println("Berikut Data Sepeda :\n")

	for _, value := range dataSepeda {

		fmt.Println("Warna Sepeda 		= " + value.warnaSepeda)
		fmt.Println("Jumlah Ban   		= ", value.jumlahBan)
		fmt.Println("Jumlah Gigi  		= ", value.jumlahGigi)
		fmt.Println("Waktu Tempuh 20KM 	= ", value.waktuTempuh(20), "Menit\n")
	}
}
