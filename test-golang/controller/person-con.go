package person_con

import (
	g "golang/connect"
	m "golang/models"

	_ "github.com/go-sql-driver/mysql"
)

var mem m.Person

//var id_response m.Person

var empData = make(map[string]m.Person)

//info := map[string]string{"message": "User Registered successfully"}

func GetData() (empData []m.Person) {
	con := g.ConnectDB()
	results, err := con.Query("SELECT * FROM `user`")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {

		err = results.Scan(&mem.ID, &mem.Firstname, &mem.Lastname)
		if err != nil {
			panic(err.Error())
		}
		//empData["message"] = append(empData, mem)
	}
	//fmt.Printf("%+v\n", empData)
	return empData

}

func GetDataId(id int) (empData []m.Person) {

	con := g.ConnectDB()
	results, err := con.Query("SELECT * FROM `user` WHERE `id` = ?", id)
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {

		err = results.Scan(&mem.ID, &mem.Firstname, &mem.Lastname)
		if err != nil {
			panic(err.Error())
		}
		empData = append(empData, mem)
	}
	return empData
}

func DeleteDataId(id int) m.Person {
	con := g.ConnectDB()
	results, err := con.Query("SELECT * FROM `user` WHERE `id` = ?", id)

	if err != nil {

		panic(err.Error())
	}
	for results.Next() {

		err = results.Scan(&mem.ID, &mem.Firstname, &mem.Lastname)
		if err != nil {
			panic(err.Error())
		}

	}

	con.Exec("DELETE FROM `user` WHERE id = ?", id)

	if err != nil {
		panic(err.Error())
	}

	res := &mem

	return *res

}

func UpDataId(id int, f_name string, l_name string) m.Person {

	con := g.ConnectDB()
	sql, err := con.Prepare("UPDATE `user` SET fname= ?,lname= ? where id=?")
	sql.Exec(f_name, l_name, id)
	if err != nil {
		panic(err.Error())
	}

	mem.ID = id
	mem.Firstname = f_name
	mem.Lastname = l_name
	res := &mem

	return *res
}

func InsertData(f_name string, l_name string) m.Person {
	var err error

	con := g.ConnectDB()

	sql := `INSERT INTO user( fname, lname) VALUES (?,?)`

	_, err = con.Exec(sql, f_name, l_name)
	if err != nil {
		panic(err.Error())
	}
	results, err := con.Query("SELECT * FROM user ORDER BY ID DESC LIMIT 1")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {

		err = results.Scan(&mem.ID, &mem.Firstname, &mem.Lastname)
		if err != nil {
			panic(err.Error())
		}

	}
	res := &mem
	return *res
}
