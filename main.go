package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// Init users var as a slice User Struct
var customers []customer

type customer struct {
	ID string `json:"id"`
	FullName string `json:"full_name"`
	NIC string `json:"nic"`
	ContactNumber int `json:"contact_number"`
	Address string `json:"address"`
	ImgFilePath string `json:"img_file_path"`
}


// enable cross Origin policy
func setupCorsResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
}


func getCustomers(w http.ResponseWriter,r *http.Request){
	var	customers2 []customer
	setupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	w.Header().Set("Content-Type","application/json")

	db, _ := sql.Open("mysql", "root:ijse@tcp(127.0.0.1:3306)/custRegFormDB")
	row, err := db.Query("select * from customer")
	if err != nil {
		panic(err.Error())
	}else{
		for row.Next(){
			var cusid string
			var fullname string
			var nic string
			var contact int
			var address string
			var filepath string

			err2 := row.Scan(&cusid	, &fullname, &nic, &contact, &address,&filepath)
			row.Columns()
			if err2 != nil{
				panic(err2.Error())
			}else{
				customer2 := customer{
					ID:cusid,
					FullName:fullname,
					NIC: nic,
					ContactNumber: contact,
					Address:address,
					ImgFilePath:filepath,
				}
				customers2 = append(customers2, customer2)
			}
		}
	}
	defer row.Close()
	json.NewEncoder(w).Encode(customers2)
}

func addCustomer(w http.ResponseWriter, r *http.Request)  {

	setupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	// Maximum upload of 10 MB files
	r.ParseMultipartForm(1 << 2)

	// Get handler for filename, size and headers
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	r.MultipartReader()




	cusid := r.FormValue("cus_id")
	cusname :=r.FormValue("cus_name")
	cusnic :=r.FormValue("cus_nic")
	cuscontact :=r.FormValue("cus_contact_number")
	cusadress :=r.FormValue("cus_address")
	filepathmain:="\\CustRegForm_V01\\assets\\img-dir\\"+handler.Filename

	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Create file
	//dst, err := os.Create(handler.Filename)
	dst, err := os.Create(filepath.Join("assets/img-dir", filepath.Base(handler.Filename)))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer dst.Close()
	if _, err = io.Copy(dst, file); err != nil {
		fmt.Println(err)
		return
	}

	// Copy the uploaded file to the created file on the filesystem
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	db, _ := sql.Open("mysql", "root:ijse@tcp(127.0.0.1:3306)/custRegFormDB")
	insert, err := db.Query("INSERT INTO customer VALUES (?, ?, ?, ?, ?, ?)",cusid,cusname, cusnic,cuscontact, cusadress,filepathmain)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()

}

//get Customer
func getCustomer(w http.ResponseWriter, r *http.Request) {
	var cust customer

	setupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)

	db, _ := sql.Open("mysql", "root:ijse@tcp(127.0.0.1:3306)/custRegFormDB")
	row, err := db.Query("select * from customer where cusid=?", params["id"])
	if err != nil {
		panic(err.Error())
	}else{
		for row.Next(){
			var cusid string
			var fullname string
			var nic string
			var contact int
			var address string
			var filepath string
			err2 := row.Scan(&cusid, &fullname, &nic, &contact, &address,  &filepath)
			row.Columns()
			if err2 != nil{
				panic(err2.Error())
			}else{
				customer := customer{
					ID:      cusid,
					FullName:    fullname,
					NIC:     nic,
					ContactNumber:contact,
					Address: address,
					ImgFilePath:filepath,
				}
				cust = customer
			}
		}
	}
	defer row.Close()
	json.NewEncoder(w).Encode(cust)
}

func updateCustomer(w http.ResponseWriter, r *http.Request)  {
	setupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	// Maximum upload of 10 MB files
	r.ParseMultipartForm(1 << 2)

	// Get handler for filename, size and headers
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	r.MultipartReader()



	params := mux.Vars(r)
	//cusid := r.FormValue("cus_id")
	cusname :=r.FormValue("cus_name")
	cusnic :=r.FormValue("cus_nic")
	cuscontact :=r.FormValue("cus_contact_number")
	cusadress :=r.FormValue("cus_address")
	filepathmain:="\\CustRegForm_V01\\assets\\img-dir\\"+handler.Filename

	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Create file
	//dst, err := os.Create(handler.Filename)
	dst, err := os.Create(filepath.Join("assets/img-dir", filepath.Base(handler.Filename)))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer dst.Close()
	if _, err = io.Copy(dst, file); err != nil {
		fmt.Println(err)
		return
	}

	// Copy the uploaded file to the created file on the filesystem
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	db, _ := sql.Open("mysql", "root:ijse@tcp(127.0.0.1:3306)/custRegFormDB")
	update, err := db.Query("update customer set fullname=? , nic=? , contact=? , address=?, imgpath=?   where cusid= ?", cusname, cusnic, cuscontact,cusadress,filepathmain, params["id"])
	if err != nil {
		panic(err.Error())
	}
	defer update.Close()

}

func deleteCustomer(w http.ResponseWriter, r *http.Request)  {
	setupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	db, _ := sql.Open("mysql", "root:ijse@tcp(127.0.0.1:3306)/custRegFormDB")
	delete, err := db.Query("delete from customer where cusid=?", params["id"])

	if err != nil {
		panic(err.Error())
	}
	defer delete.Close()
	json.NewEncoder(w).Encode(customers)
}



func main() {
	r := mux.NewRouter()
	fmt.Println("Server Running...")
	r.HandleFunc("/api/customer", getCustomers).Methods("GET","OPTIONS")
	r.HandleFunc("/api/customer/{id}", getCustomer).Methods("GET","OPTIONS")
	r.HandleFunc("/api/customer", addCustomer).Methods("POST","OPTIONS")
	r.HandleFunc("/api/customer/{id}", updateCustomer).Methods("PUT","OPTIONS")
	r.HandleFunc("/api/customer/{id}", deleteCustomer).Methods("DELETE","OPTIONS")

	log.Fatal(http.ListenAndServe(":8000", r))
}