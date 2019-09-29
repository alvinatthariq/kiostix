package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func returnAllBook(w http.ResponseWriter, r *http.Request) {
	var buku Book
	var arrBook []Book
	var response Response

	db := connect()
	defer db.Close()

	rows, err := db.Query("Select id,judul_buku,penulis,kategori from book")
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&buku.ID, &buku.JudulBuku, &buku.Penulis, &buku.Kategori); err != nil {
			log.Fatal(err.Error())

		} else {
			arrBook = append(arrBook, buku)
		}
	}

	response.Status = 1
	response.Message = "Success"
	response.Data = arrBook

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func returnBookPenulis(w http.ResponseWriter, r *http.Request) {
	var buku Book
	var arrBook []Book
	var response Response

	db := connect()
	defer db.Close()
	penulis := r.FormValue("penulis")
	rows, err := db.Query("Select id,judul_buku,penulis,kategori from book where penulis = ?", penulis)
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&buku.ID, &buku.JudulBuku, &buku.Penulis, &buku.Kategori); err != nil {
			log.Fatal(err.Error())

		} else {
			arrBook = append(arrBook, buku)
		}
	}

	response.Status = 1
	response.Message = "Success"
	response.Data = arrBook

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func returnBookJudul(w http.ResponseWriter, r *http.Request) {
	var buku Book
	var arrBook []Book
	var response Response

	db := connect()
	defer db.Close()
	judul_buku := r.FormValue("judul_buku")
	rows, err := db.Query("Select id,judul_buku,penulis,kategori from book where judul_buku = ?", judul_buku)
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&buku.ID, &buku.JudulBuku, &buku.Penulis, &buku.Kategori); err != nil {
			log.Fatal(err.Error())

		} else {
			arrBook = append(arrBook, buku)
		}
	}

	response.Status = 1
	response.Message = "Success"
	response.Data = arrBook

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func returnBookKategori(w http.ResponseWriter, r *http.Request) {
	var buku Book
	var arrBook []Book
	var response Response

	db := connect()
	defer db.Close()

	kategori := r.FormValue("kategori")

	rows, err := db.Query("Select id,judul_buku,penulis,kategori from book where kategori = ?", kategori)
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&buku.ID, &buku.JudulBuku, &buku.Penulis, &buku.Kategori); err != nil {
			log.Fatal(err.Error())

		} else {
			arrBook = append(arrBook, buku)
		}
	}

	response.Status = 1
	response.Message = "Success"
	response.Data = arrBook

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func insertBook(w http.ResponseWriter, r *http.Request) {
	var response Response

	db := connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	judul_buku := r.FormValue("judul_buku")
	penulis := r.FormValue("penulis")
	kategori := r.FormValue("kategori")

	_, err = db.Exec("INSERT INTO book (judul_buku, penulis, kategori) values (?,?,?)",
		judul_buku,
		penulis,
		kategori,
	)
	if err != nil {
		log.Print(err)
	}

	response.Status = 1
	response.Message = "Success Add"
	log.Print("Insert data to database")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/books", returnAllBook).Methods("GET")
	router.HandleFunc("/books/penulis", returnBookPenulis).Methods("GET")
	router.HandleFunc("/books/judul", returnBookJudul).Methods("GET")
	router.HandleFunc("/books/kategori", returnBookKategori).Methods("GET")

	router.HandleFunc("/books", insertBook).Methods("POST")
	http.Handle("/", router)
	fmt.Println("Connected to port 1234")
	log.Fatal(http.ListenAndServe(":1234", router))

}
