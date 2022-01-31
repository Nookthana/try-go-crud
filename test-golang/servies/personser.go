package personser

import (
	"encoding/json"
	c "golang/controller"
	m "golang/models"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetDataPersonAll(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	var err error
	vars := mux.Vars(req)
	id := vars["id"]
	id_str, _ := strconv.Atoi(id)
	//w.Write([]byte(`"Message": "OK"`))
	var infomation = map[string]string{"Message": "OK"}

	if err != nil {
		json.NewEncoder(w).Encode("Data Not Fount")
		panic(err.Error())

	} else {
		res := c.GetDataId(id_str)

		json.NewEncoder(w).Encode(res)
		json.NewEncoder(w).Encode(infomation)

	}

}

func GetDataPersonID(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	res := c.GetData()
	if len(res) != 0 {
		json.NewEncoder(w).Encode(res)
	} else {
		json.NewEncoder(w).Encode("Data Not Fount")
	}
}

func CreatePerson(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	reqBody, _ := ioutil.ReadAll(req.Body)
	var person m.Person
	json.Unmarshal(reqBody, &person)
	f_name := person.Firstname
	l_name := person.Firstname

	if f_name != "" && l_name != "" {
		json.NewEncoder(w).Encode("data inserted")
		res := c.InsertData(f_name, l_name)
		json.NewEncoder(w).Encode(res)

	} else {
		json.NewEncoder(w).Encode("Cannot insert data value is null")
	}

}

func DeletePersonID(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var err error
	vars := mux.Vars(req)
	id := vars["id"]
	id_str, _ := strconv.Atoi(id)

	if err != nil {
		json.NewEncoder(w).Encode("Data Not in database")
		panic(err.Error())

	} else {
		res := c.DeleteDataId(id_str)
		json.NewEncoder(w).Encode(res)
		json.NewEncoder(w).Encode("Data Deleteted")
	}
}

func UpDatePersonID(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//var err error
	vars := mux.Vars(req)
	id := vars["id"]
	id_str, _ := strconv.Atoi(id)
	reqBody, _ := ioutil.ReadAll(req.Body)
	var person m.Person
	json.Unmarshal(reqBody, &person)
	f_name := person.Firstname
	l_name := person.Firstname

	if f_name != "" && l_name != "" {
		res := c.UpDataId(id_str, f_name, l_name)
		json.NewEncoder(w).Encode(res)
		json.NewEncoder(w).Encode("data has updated successfully")
	} else {
		json.NewEncoder(w).Encode("cannot update data value is null")

	}

}
