package person

type Person struct {
	ID        int    `json:"id"`
	Firstname string `json:"fname"`
	Lastname  string `json:"lname"`
}
