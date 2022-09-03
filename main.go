package main

import (
	"fmt"
	"os"
	"strconv"
)

type Classmate struct {
	Absen     int
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

var classmate []Classmate

func init() {
	classmate = []Classmate{
		{Absen: 1, Nama: "Budi", Alamat: "abc", Pekerjaan: "Golang Programmer", Alasan: "Ingin menambah wawasan"},
		{Absen: 2, Nama: "Dinda", Alamat: "def", Pekerjaan: "Backend Engineer", Alasan: "Ingin menambah ilmu"},
		{Absen: 3, Nama: "Upin", Alamat: "ghi", Pekerjaan: "Frontend Engineer", Alasan: "Ingin menjadi orang sukses"},
		{Absen: 4, Nama: "Ipin", Alamat: "jkl", Pekerjaan: "UI/UX Engineer", Alasan: "Ingin membahagiakan orangtua"},
		{Absen: 5, Nama: "Indomie", Alamat: "mno", Pekerjaan: "Data Analyst", Alasan: "Ingin menjadi orang teladan"},
		{Absen: 6, Nama: "Marimas", Alamat: "pqr", Pekerjaan: "Marketing", Alasan: "Ingin membangun website dengan Go"},
		{Absen: 7, Nama: "Pop Ice", Alamat: "stu", Pekerjaan: "Business Management", Alasan: "Ingin belajar Go lebih dalam"},
		{Absen: 8, Nama: "Nutri Sari", Alamat: "vwx", Pekerjaan: "Ruby Programmer", Alasan: "Ingin menjadi programmer Go"},
		{Absen: 9, Nama: "Kopiko", Alamat: "yz", Pekerjaan: "Senior HMTL Programmer", Alasan: "Ingin menjadi programmer handal"},
		{Absen: 10, Nama: "Milkita", Alamat: "abc", Pekerjaan: "Junior Backend Engineer", Alasan: "Ingin menjadi orang yang keren"},
	}
}
func main() {
	allArgs := os.Args[1:]
	if len(os.Args) > 1 {
		for i := 0; i < len(allArgs); i++ {
			absen, err := strconv.Atoi(allArgs[i])
			if err != nil {
				fmt.Println("Error. Mohon di cek kembali")
			} else {
				search := 0
				for i := 0; i < len(classmate); i++ {
					if classmate[i].Absen == absen {
						ShowData(classmate[i])
						search++
						break
					}
				}
				if search == 0 {
					fmt.Println("Data Tidak Ditemukan")
				}
			}
		}
	} else {
		fmt.Println("Mohon masukan nomor absen mahasiswa yang sesuai")
	}
}

func ShowData(classmate Classmate) {
	fmt.Println("Nama :", classmate.Nama)
	fmt.Println("Alamat :", classmate.Alamat)
	fmt.Println("Pekerjaan :", classmate.Pekerjaan)
	fmt.Println("Alasan :", classmate.Alasan)
}
