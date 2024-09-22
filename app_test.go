package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

var a App

func Testmain(m *testing.M) {
	err := a.Initailize(DbUser, DbPassword, DbName)
	if err != nil {
		log.Fatal("Error Occured while Initialising the database")
	}
	m.Run()
}

func createTable() {
	createTableQuery := `Create table If not exists products(
    id int NOT NULL AUTO_INCREMENT,
    name varchar(255) NOT NULL,
    quantity int,
    price float(10,7),
    PRIMARY KEY(id)
    );`

	_, err := a.Db.Exec(createTableQuery)
	if err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	a.Db.Exec("Delete from products")
}

func addProduct(name string, quantity int, price float64) {
	query := fmt.Sprintf("Insert Into products(name,quantity,price) values('%v','%v','%v')", name, quantity, price)
	a.Db.Exec(query)
}

func TestGetProduct(t *testing.T) {
	clearTable()
	addProduct("keyboard", 100, 5000)
	request, _ := http.NewRequest("GET", "/product/1", nil)
	response := sendRequest(request)
	checkStatusCode(t, http.StatusOK, response.Code)

}

func checkStatusCode(t *testing.T, expectedStatusCode int, actualStatusCode int) {
	if expectedStatusCode != actualStatusCode {
		t.Errorf("Expected status: %v, Received: %v", expectedStatusCode, actualStatusCode)
	}
}

func sendRequest(request *http.Request) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()
	a.Router.ServeHTTP(recorder, request)
	return recorder
}
